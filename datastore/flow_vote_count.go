package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/environment"
)

var (
	envVoteHost             = environment.NewVariable("VOTE_HOST", "localhost", "Host of the vote-service.")
	envVotePort             = environment.NewVariable("VOTE_PORT", "9013", "Port of the vote-service.")
	envVoteProtocol         = environment.NewVariable("VOTE_PROTOCOL", "http", "Protocol of the vote-service.")
	envDebugHasVotedUserIDs = environment.NewVariable("DEBUG_HAS_VOTED_USER_IDS", "false", "Enable Debug message for an error from May 2025.")
)

const allVotedIdsPath = "/internal/vote/all_voted_ids"

// FlowVoteCount is a datastore flow for the poll/vote_count value.
type FlowVoteCount struct {
	voteServiceURL string
	client         *http.Client
	id             uint64

	mu            sync.Mutex
	pollToUserIDs map[int][]int
	update        chan map[int][]int
	ready         chan struct{}

	debugHasVotedUserIDs bool
}

// NewFlowVoteCount initializes the object.
func NewFlowVoteCount(lookup environment.Environmenter) *FlowVoteCount {
	url := fmt.Sprintf(
		"%s://%s:%s",
		envVoteProtocol.Value(lookup),
		envVoteHost.Value(lookup),
		envVotePort.Value(lookup),
	)

	debug2025, _ := strconv.ParseBool(envDebugHasVotedUserIDs.Value(lookup))

	flow := FlowVoteCount{
		voteServiceURL: url,
		client:         &http.Client{},
		update:         make(chan map[int][]int, 1),
		pollToUserIDs:  make(map[int][]int),
		ready:          make(chan struct{}),

		debugHasVotedUserIDs: debug2025,
	}

	return &flow
}

func (s *FlowVoteCount) printDebugHasVotedUserIDs(format string, a ...any) {
	if !s.debugHasVotedUserIDs {
		return
	}
	fmt.Printf("DEBUG: HAS VOTED USER IDs: %s\n", fmt.Sprintf(format, a...))
}

// Connect creates a connection to the vote service and makes sure, it stays
// open.
//
// eventProvider is a function that returns a channel. If the connection fails,
// this function fetches such a channel and waits for a signal before it tries
// to open a new connection.
func (s *FlowVoteCount) Connect(ctx context.Context, eventProvider func() (<-chan time.Time, func() bool), errHandler func(error)) {
	for ctx.Err() == nil {
		s.printDebugHasVotedUserIDs("Create connection to vote service")
		if err := s.connect(ctx); err != nil {
			s.printDebugHasVotedUserIDs("Error with vote service connection: %v", err)
			errHandler(fmt.Errorf("connecting to vote service: %w", err))
		}

		// TODO: When the connection is closed, s.ready has to be resetted. But
		// this would be a race condition.

		s.printDebugHasVotedUserIDs("Waiting for reconnect")
		s.wait(ctx, eventProvider)
	}

	s.printDebugHasVotedUserIDs("Stop trying to connect to vote service: %v", ctx.Err())
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
	s.pollToUserIDs = make(map[int][]int)
	s.mu.Unlock()

	req, err := http.NewRequestWithContext(ctx, "GET", s.voteServiceURL+allVotedIdsPath, nil)
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
		var fromVoteService map[int][]int
		if err := decoder.Decode(&fromVoteService); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("decoding poll data: %w", err)
		}

		s.mu.Lock()
		for pollID, userIDs := range fromVoteService {
			if userIDs == nil {
				delete(s.pollToUserIDs, pollID)
				continue
			}
			s.pollToUserIDs[pollID] = append(s.pollToUserIDs[pollID], userIDs...)
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
			s.printDebugHasVotedUserIDs("Skripping message from vote service: %v", fromVoteService)
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
			if key.Field() != "has_voted_user_ids" {
				continue
			}

			userIDs, ok := s.pollToUserIDs[key.ID()]
			if !ok {
				continue
			}

			bytes, err := json.Marshal(userIDs)
			if err != nil {
				return nil, fmt.Errorf("converting user_ids to json: %w", err)
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
		var fromVoteService map[int][]int
		select {
		case <-ctx.Done():
			s.printDebugHasVotedUserIDs("Update exists after context is done: %v", ctx.Err())
			return // TODO: Should the error be returned?

		case fromVoteService = <-s.update:
		}

		var keys []dskey.Key
		for pollID := range fromVoteService {
			pollKey, err := dskey.FromParts("poll", pollID, "has_voted_user_ids")
			if err != nil {
				updateFn(nil, err)
				s.printDebugHasVotedUserIDs("Update exists on error: %v", err)
				return
			}
			keys = append(keys, pollKey)
		}

		updateFn(s.Get(ctx, keys...))
	}
}
