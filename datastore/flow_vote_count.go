package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/environment"
)

var (
	envVoteHost       = environment.NewVariable("VOTE_HOST", "localhost", "Host of the vote-service.")
	envVotePort       = environment.NewVariable("VOTE_PORT", "9013", "Port of the vote-service.")
	envVoteProtocol   = environment.NewVariable("VOTE_PROTOCOL", "http", "Protocol of the vote-service.")
	envDebugLiveVotes = environment.NewVariable("DEBUG_HAS_VOTED_USER_IDS", "false", "Enable Debug message for an error from May 2025.")
)

const liveVotesPath = "/internal/vote/live_votes"

// FlowVoteCount is a datastore flow for the poll/vote_count value.
type FlowVoteCount struct {
	voteServiceURL string
	client         *http.Client
	id             uint64

	mu            sync.Mutex
	pollLiveVotes map[int]map[int]*string
	update        chan map[int]map[int]*string
	ready         chan struct{}

	debugLiveVotes bool
}

// NewFlowVoteCount initializes the object.
func NewFlowVoteCount(lookup environment.Environmenter) *FlowVoteCount {
	url := fmt.Sprintf(
		"%s://%s:%s",
		envVoteProtocol.Value(lookup),
		envVoteHost.Value(lookup),
		envVotePort.Value(lookup),
	)

	debug2025, _ := strconv.ParseBool(envDebugLiveVotes.Value(lookup))

	flow := FlowVoteCount{
		voteServiceURL: url,
		client:         &http.Client{},
		update:         make(chan map[int]map[int]*string, 1),
		pollLiveVotes:  make(map[int]map[int]*string),
		ready:          make(chan struct{}),

		debugLiveVotes: debug2025,
	}

	return &flow
}

func (s *FlowVoteCount) printDebugLiveVotes(format string, a ...any) {
	if !s.debugLiveVotes {
		return
	}
	fmt.Printf("DEBUG: Live Votes: %s\n", fmt.Sprintf(format, a...))
}

// Connect creates a connection to the vote service and makes sure, it stays
// open.
//
// eventProvider is a function that returns a channel. If the connection fails,
// this function fetches such a channel and waits for a signal before it tries
// to open a new connection.
func (s *FlowVoteCount) Connect(ctx context.Context, eventProvider func() (<-chan time.Time, func() bool), errHandler func(error)) {
	for ctx.Err() == nil {
		s.printDebugLiveVotes("Create connection to vote service")
		if err := s.connect(ctx); err != nil {
			s.printDebugLiveVotes("Error with vote service connection: %v", err)
			errHandler(fmt.Errorf("connecting to vote service: %w", err))
		}

		// TODO: When the connection is closed, s.ready has to be resetted. But
		// this would be a race condition.

		s.printDebugLiveVotes("Waiting for reconnect")
		s.wait(ctx, eventProvider)
	}

	s.printDebugLiveVotes("Stop trying to connect to vote service: %v", ctx.Err())
}

// wait waits for an event in s.eventProvider.
func (s *FlowVoteCount) wait(ctx context.Context, eventProvider func() (<-chan time.Time, func() bool)) {
	event, close := eventProvider()
	defer close()

	select {
	case <-ctx.Done():
	case <-event:
	}
}

func (s *FlowVoteCount) connect(ctx context.Context) error {
	s.mu.Lock()
	s.pollLiveVotes = make(map[int]map[int]*string)
	s.mu.Unlock()

	req, err := http.NewRequestWithContext(ctx, "GET", s.voteServiceURL+liveVotesPath, nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		// TODO External Error
		return fmt.Errorf("sending request to vote service: %w", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var fromVoteService map[int]map[int]*string
		if err := decoder.Decode(&fromVoteService); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("decoding poll data: %w", err)
		}

		s.mu.Lock()
		for pollID, userID2Vote := range fromVoteService {
			if userID2Vote == nil {
				// The userID2Vote map is nil, if a poll was removed.
				delete(s.pollLiveVotes, pollID)
				continue
			}
			if s.pollLiveVotes[pollID] == nil {
				s.pollLiveVotes[pollID] = make(map[int]*string)
			}
			maps.Copy(s.pollLiveVotes[pollID], userID2Vote)
		}
		s.mu.Unlock()

		select {
		case <-s.ready:
		default:
			close(s.ready)
		}

		select {
		case s.update <- fromVoteService:
		default:
			s.printDebugLiveVotes("Skripping message from vote service: %v", fromVoteService)
		}
	}
}

// Get is called when a key is not in the cache.
func (s *FlowVoteCount) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	select {
	case <-s.ready:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	out := make(map[dskey.Key][]byte, len(keys))
	for _, key := range keys {
		out[key] = nil

		switch key.Collection() {
		case "poll":
			if key.Field() != "live_votes" {
				continue
			}

			userID2Vote, ok := s.pollLiveVotes[key.ID()]
			if !ok {
				continue
			}

			bytes, err := json.Marshal(userID2Vote)
			if err != nil {
				return nil, fmt.Errorf("converting userID2Vote to json: %w", err)
			}
			out[key] = bytes

		default:
			continue
		}
	}
	return out, nil
}

// Update has to be called frequently. It blocks, until there is new data.
func (s *FlowVoteCount) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	for {
		var fromVoteService map[int]map[int]*string
		select {
		case <-ctx.Done():
			s.printDebugLiveVotes("Update exists after context is done: %v", ctx.Err())
			return // TODO: Should the error be returned?

		case fromVoteService = <-s.update:
		}

		var keys []dskey.Key
		for pollID := range fromVoteService {
			pollKey, err := dskey.FromParts("poll", pollID, "live_votes")
			if err != nil {
				updateFn(nil, err)
				s.printDebugLiveVotes("Update exists on error: %v", err)
				return
			}
			keys = append(keys, pollKey)
		}

		updateFn(s.Get(ctx, keys...))
	}
}
