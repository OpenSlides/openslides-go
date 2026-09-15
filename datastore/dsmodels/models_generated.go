// Code generated from meta collections. DO NOT EDIT.
package dsmodels

import (
	"encoding/json"

	"github.com/OpenSlides/openslides-go/datastore/dstypes"
	"github.com/OpenSlides/openslides-go/maybe"
	"github.com/shopspring/decimal"
)

// ActionWorker has all fields from action_worker.
type ActionWorker struct {
	Created   int
	ID        int
	Name      string
	Result    maybe.Maybe[json.RawMessage]
	State     dstypes.ActionWorker_State
	Timestamp int
	UserID    int
}

type actionWorkerBuilder struct {
	builder[actionWorkerBuilder, *actionWorkerBuilder, ActionWorker]
}

func (b *actionWorkerBuilder) lazy(ds *Fetch, id int) *ActionWorker {
	c := ActionWorker{}
	ds.ActionWorker_Created(id).Lazy(&c.Created)
	ds.ActionWorker_ID(id).Lazy(&c.ID)
	ds.ActionWorker_Name(id).Lazy(&c.Name)
	ds.ActionWorker_Result(id).Lazy(&c.Result)
	ds.ActionWorker_State(id).Lazy(&c.State)
	ds.ActionWorker_Timestamp(id).Lazy(&c.Timestamp)
	ds.ActionWorker_UserID(id).Lazy(&c.UserID)
	return &c
}

func (b *actionWorkerBuilder) Preload(rel builderWrapperI) *actionWorkerBuilder {
	b.builder.Preload(rel)
	return b
}

func (r *Fetch) ActionWorker(ids ...int) *actionWorkerBuilder {
	return &actionWorkerBuilder{
		builder: builder[actionWorkerBuilder, *actionWorkerBuilder, ActionWorker]{
			ids:   ids,
			fetch: r,
		},
	}
}

// AgendaItem has all fields from agenda_item.
type AgendaItem struct {
	ChildIDs        maybe.Maybe[[]int]
	Closed          maybe.Maybe[bool]
	Comment         maybe.Maybe[string]
	ContentObjectID string
	Duration        maybe.Maybe[int]
	ID              int
	IsHidden        maybe.Maybe[bool]
	IsInternal      maybe.Maybe[bool]
	ItemNumber      maybe.Maybe[string]
	Level           maybe.Maybe[int]
	MeetingID       int
	ParentID        maybe.Maybe[int]
	ProjectionIDs   maybe.Maybe[[]int]
	TagIDs          maybe.Maybe[[]int]
	Type            maybe.Maybe[dstypes.AgendaItem_Type]
	Weight          maybe.Maybe[int]
	ChildList       []AgendaItem
	Meeting         *Meeting
	Parent          *maybe.Maybe[AgendaItem]
	ProjectionList  []Projection
	TagList         []Tag
}

type agendaItemBuilder struct {
	builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]
}

func (b *agendaItemBuilder) lazy(ds *Fetch, id int) *AgendaItem {
	c := AgendaItem{}
	ds.AgendaItem_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.AgendaItem_Closed(id).Lazy(&c.Closed)
	ds.AgendaItem_Comment(id).Lazy(&c.Comment)
	ds.AgendaItem_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.AgendaItem_Duration(id).Lazy(&c.Duration)
	ds.AgendaItem_ID(id).Lazy(&c.ID)
	ds.AgendaItem_IsHidden(id).Lazy(&c.IsHidden)
	ds.AgendaItem_IsInternal(id).Lazy(&c.IsInternal)
	ds.AgendaItem_ItemNumber(id).Lazy(&c.ItemNumber)
	ds.AgendaItem_Level(id).Lazy(&c.Level)
	ds.AgendaItem_MeetingID(id).Lazy(&c.MeetingID)
	ds.AgendaItem_ParentID(id).Lazy(&c.ParentID)
	ds.AgendaItem_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.AgendaItem_TagIDs(id).Lazy(&c.TagIDs)
	ds.AgendaItem_Type(id).Lazy(&c.Type)
	ds.AgendaItem_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *agendaItemBuilder) Preload(rel builderWrapperI) *agendaItemBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *agendaItemBuilder) ChildList() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChildIDs",
			relField: "ChildList",
			many:     true,
		},
	}
}

func (b *agendaItemBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *agendaItemBuilder) Parent() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ParentID",
			relField: "Parent",
		},
	}
}

func (b *agendaItemBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *agendaItemBuilder) TagList() *tagBuilder {
	return &tagBuilder{
		builder: builder[tagBuilder, *tagBuilder, Tag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TagIDs",
			relField: "TagList",
			many:     true,
		},
	}
}

func (r *Fetch) AgendaItem(ids ...int) *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Assignment has all fields from assignment.
type Assignment struct {
	AgendaItemID                   maybe.Maybe[int]
	AttachmentMeetingMediafileIDs  maybe.Maybe[[]int]
	CandidateIDs                   maybe.Maybe[[]int]
	DefaultPollDescription         maybe.Maybe[string]
	Description                    maybe.Maybe[string]
	HistoryEntryIDs                maybe.Maybe[[]int]
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	NumberPollCandidates           maybe.Maybe[bool]
	OpenPosts                      maybe.Maybe[int]
	Phase                          maybe.Maybe[dstypes.Assignment_Phase]
	PollIDs                        maybe.Maybe[[]int]
	ProjectionIDs                  maybe.Maybe[[]int]
	SequentialNumber               int
	TagIDs                         maybe.Maybe[[]int]
	Title                          string
	AgendaItem                     *maybe.Maybe[AgendaItem]
	AttachmentMeetingMediafileList []MeetingMediafile
	CandidateList                  []AssignmentCandidate
	HistoryEntryList               []HistoryEntry
	ListOfSpeakers                 *ListOfSpeakers
	Meeting                        *Meeting
	PollList                       []Poll
	ProjectionList                 []Projection
	TagList                        []Tag
}

type assignmentBuilder struct {
	builder[assignmentBuilder, *assignmentBuilder, Assignment]
}

func (b *assignmentBuilder) lazy(ds *Fetch, id int) *Assignment {
	c := Assignment{}
	ds.Assignment_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Assignment_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Assignment_CandidateIDs(id).Lazy(&c.CandidateIDs)
	ds.Assignment_DefaultPollDescription(id).Lazy(&c.DefaultPollDescription)
	ds.Assignment_Description(id).Lazy(&c.Description)
	ds.Assignment_HistoryEntryIDs(id).Lazy(&c.HistoryEntryIDs)
	ds.Assignment_ID(id).Lazy(&c.ID)
	ds.Assignment_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Assignment_MeetingID(id).Lazy(&c.MeetingID)
	ds.Assignment_NumberPollCandidates(id).Lazy(&c.NumberPollCandidates)
	ds.Assignment_OpenPosts(id).Lazy(&c.OpenPosts)
	ds.Assignment_Phase(id).Lazy(&c.Phase)
	ds.Assignment_PollIDs(id).Lazy(&c.PollIDs)
	ds.Assignment_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Assignment_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Assignment_TagIDs(id).Lazy(&c.TagIDs)
	ds.Assignment_Title(id).Lazy(&c.Title)
	return &c
}

func (b *assignmentBuilder) Preload(rel builderWrapperI) *assignmentBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *assignmentBuilder) AgendaItem() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AgendaItemID",
			relField: "AgendaItem",
		},
	}
}

func (b *assignmentBuilder) AttachmentMeetingMediafileList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AttachmentMeetingMediafileIDs",
			relField: "AttachmentMeetingMediafileList",
			many:     true,
		},
	}
}

func (b *assignmentBuilder) CandidateList() *assignmentCandidateBuilder {
	return &assignmentCandidateBuilder{
		builder: builder[assignmentCandidateBuilder, *assignmentCandidateBuilder, AssignmentCandidate]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CandidateIDs",
			relField: "CandidateList",
			many:     true,
		},
	}
}

func (b *assignmentBuilder) HistoryEntryList() *historyEntryBuilder {
	return &historyEntryBuilder{
		builder: builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HistoryEntryIDs",
			relField: "HistoryEntryList",
			many:     true,
		},
	}
}

func (b *assignmentBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *assignmentBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *assignmentBuilder) PollList() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollIDs",
			relField: "PollList",
			many:     true,
		},
	}
}

func (b *assignmentBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *assignmentBuilder) TagList() *tagBuilder {
	return &tagBuilder{
		builder: builder[tagBuilder, *tagBuilder, Tag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TagIDs",
			relField: "TagList",
			many:     true,
		},
	}
}

func (r *Fetch) Assignment(ids ...int) *assignmentBuilder {
	return &assignmentBuilder{
		builder: builder[assignmentBuilder, *assignmentBuilder, Assignment]{
			ids:   ids,
			fetch: r,
		},
	}
}

// AssignmentCandidate has all fields from assignment_candidate.
type AssignmentCandidate struct {
	AssignmentID  int
	ID            int
	MeetingID     int
	MeetingUserID maybe.Maybe[int]
	Weight        maybe.Maybe[int]
	Assignment    *Assignment
	Meeting       *Meeting
	MeetingUser   *maybe.Maybe[MeetingUser]
}

type assignmentCandidateBuilder struct {
	builder[assignmentCandidateBuilder, *assignmentCandidateBuilder, AssignmentCandidate]
}

func (b *assignmentCandidateBuilder) lazy(ds *Fetch, id int) *AssignmentCandidate {
	c := AssignmentCandidate{}
	ds.AssignmentCandidate_AssignmentID(id).Lazy(&c.AssignmentID)
	ds.AssignmentCandidate_ID(id).Lazy(&c.ID)
	ds.AssignmentCandidate_MeetingID(id).Lazy(&c.MeetingID)
	ds.AssignmentCandidate_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.AssignmentCandidate_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *assignmentCandidateBuilder) Preload(rel builderWrapperI) *assignmentCandidateBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *assignmentCandidateBuilder) Assignment() *assignmentBuilder {
	return &assignmentBuilder{
		builder: builder[assignmentBuilder, *assignmentBuilder, Assignment]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AssignmentID",
			relField: "Assignment",
		},
	}
}

func (b *assignmentCandidateBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *assignmentCandidateBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (r *Fetch) AssignmentCandidate(ids ...int) *assignmentCandidateBuilder {
	return &assignmentCandidateBuilder{
		builder: builder[assignmentCandidateBuilder, *assignmentCandidateBuilder, AssignmentCandidate]{
			ids:   ids,
			fetch: r,
		},
	}
}

// ChatGroup has all fields from chat_group.
type ChatGroup struct {
	ChatMessageIDs  maybe.Maybe[[]int]
	ID              int
	MeetingID       int
	Name            string
	ReadGroupIDs    maybe.Maybe[[]int]
	Weight          maybe.Maybe[int]
	WriteGroupIDs   maybe.Maybe[[]int]
	ChatMessageList []ChatMessage
	Meeting         *Meeting
	ReadGroupList   []Group
	WriteGroupList  []Group
}

type chatGroupBuilder struct {
	builder[chatGroupBuilder, *chatGroupBuilder, ChatGroup]
}

func (b *chatGroupBuilder) lazy(ds *Fetch, id int) *ChatGroup {
	c := ChatGroup{}
	ds.ChatGroup_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.ChatGroup_ID(id).Lazy(&c.ID)
	ds.ChatGroup_MeetingID(id).Lazy(&c.MeetingID)
	ds.ChatGroup_Name(id).Lazy(&c.Name)
	ds.ChatGroup_ReadGroupIDs(id).Lazy(&c.ReadGroupIDs)
	ds.ChatGroup_Weight(id).Lazy(&c.Weight)
	ds.ChatGroup_WriteGroupIDs(id).Lazy(&c.WriteGroupIDs)
	return &c
}

func (b *chatGroupBuilder) Preload(rel builderWrapperI) *chatGroupBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *chatGroupBuilder) ChatMessageList() *chatMessageBuilder {
	return &chatMessageBuilder{
		builder: builder[chatMessageBuilder, *chatMessageBuilder, ChatMessage]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChatMessageIDs",
			relField: "ChatMessageList",
			many:     true,
		},
	}
}

func (b *chatGroupBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *chatGroupBuilder) ReadGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReadGroupIDs",
			relField: "ReadGroupList",
			many:     true,
		},
	}
}

func (b *chatGroupBuilder) WriteGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "WriteGroupIDs",
			relField: "WriteGroupList",
			many:     true,
		},
	}
}

func (r *Fetch) ChatGroup(ids ...int) *chatGroupBuilder {
	return &chatGroupBuilder{
		builder: builder[chatGroupBuilder, *chatGroupBuilder, ChatGroup]{
			ids:   ids,
			fetch: r,
		},
	}
}

// ChatMessage has all fields from chat_message.
type ChatMessage struct {
	ChatGroupID   int
	Content       string
	Created       int
	ID            int
	MeetingID     int
	MeetingUserID maybe.Maybe[int]
	ChatGroup     *ChatGroup
	Meeting       *Meeting
	MeetingUser   *maybe.Maybe[MeetingUser]
}

type chatMessageBuilder struct {
	builder[chatMessageBuilder, *chatMessageBuilder, ChatMessage]
}

func (b *chatMessageBuilder) lazy(ds *Fetch, id int) *ChatMessage {
	c := ChatMessage{}
	ds.ChatMessage_ChatGroupID(id).Lazy(&c.ChatGroupID)
	ds.ChatMessage_Content(id).Lazy(&c.Content)
	ds.ChatMessage_Created(id).Lazy(&c.Created)
	ds.ChatMessage_ID(id).Lazy(&c.ID)
	ds.ChatMessage_MeetingID(id).Lazy(&c.MeetingID)
	ds.ChatMessage_MeetingUserID(id).Lazy(&c.MeetingUserID)
	return &c
}

func (b *chatMessageBuilder) Preload(rel builderWrapperI) *chatMessageBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *chatMessageBuilder) ChatGroup() *chatGroupBuilder {
	return &chatGroupBuilder{
		builder: builder[chatGroupBuilder, *chatGroupBuilder, ChatGroup]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChatGroupID",
			relField: "ChatGroup",
		},
	}
}

func (b *chatMessageBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *chatMessageBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (r *Fetch) ChatMessage(ids ...int) *chatMessageBuilder {
	return &chatMessageBuilder{
		builder: builder[chatMessageBuilder, *chatMessageBuilder, ChatMessage]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Committee has all fields from committee.
type Committee struct {
	AllChildIDs                         maybe.Maybe[[]int]
	AllParentIDs                        maybe.Maybe[[]int]
	ChildIDs                            maybe.Maybe[[]int]
	DefaultMeetingID                    maybe.Maybe[int]
	Description                         maybe.Maybe[string]
	ExternalID                          maybe.Maybe[string]
	ForwardToCommitteeIDs               maybe.Maybe[[]int]
	ID                                  int
	ManagerIDs                          maybe.Maybe[[]int]
	MeetingIDs                          maybe.Maybe[[]int]
	Name                                string
	NativeUserIDs                       maybe.Maybe[[]int]
	OrganizationID                      int
	OrganizationTagIDs                  maybe.Maybe[[]int]
	ParentID                            maybe.Maybe[int]
	ReceiveForwardingsFromCommitteeIDs  maybe.Maybe[[]int]
	UserIDs                             maybe.Maybe[[]int]
	AllChildList                        []Committee
	AllParentList                       []Committee
	ChildList                           []Committee
	DefaultMeeting                      *maybe.Maybe[Meeting]
	ForwardToCommitteeList              []Committee
	ManagerList                         []User
	MeetingList                         []Meeting
	NativeUserList                      []User
	Organization                        *Organization
	OrganizationTagList                 []OrganizationTag
	Parent                              *maybe.Maybe[Committee]
	ReceiveForwardingsFromCommitteeList []Committee
	UserList                            []User
}

type committeeBuilder struct {
	builder[committeeBuilder, *committeeBuilder, Committee]
}

func (b *committeeBuilder) lazy(ds *Fetch, id int) *Committee {
	c := Committee{}
	ds.Committee_AllChildIDs(id).Lazy(&c.AllChildIDs)
	ds.Committee_AllParentIDs(id).Lazy(&c.AllParentIDs)
	ds.Committee_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.Committee_DefaultMeetingID(id).Lazy(&c.DefaultMeetingID)
	ds.Committee_Description(id).Lazy(&c.Description)
	ds.Committee_ExternalID(id).Lazy(&c.ExternalID)
	ds.Committee_ForwardToCommitteeIDs(id).Lazy(&c.ForwardToCommitteeIDs)
	ds.Committee_ID(id).Lazy(&c.ID)
	ds.Committee_ManagerIDs(id).Lazy(&c.ManagerIDs)
	ds.Committee_MeetingIDs(id).Lazy(&c.MeetingIDs)
	ds.Committee_Name(id).Lazy(&c.Name)
	ds.Committee_NativeUserIDs(id).Lazy(&c.NativeUserIDs)
	ds.Committee_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Committee_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Committee_ParentID(id).Lazy(&c.ParentID)
	ds.Committee_ReceiveForwardingsFromCommitteeIDs(id).Lazy(&c.ReceiveForwardingsFromCommitteeIDs)
	ds.Committee_UserIDs(id).Lazy(&c.UserIDs)
	return &c
}

func (b *committeeBuilder) Preload(rel builderWrapperI) *committeeBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *committeeBuilder) AllChildList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AllChildIDs",
			relField: "AllChildList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) AllParentList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AllParentIDs",
			relField: "AllParentList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) ChildList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChildIDs",
			relField: "ChildList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) DefaultMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultMeetingID",
			relField: "DefaultMeeting",
		},
	}
}

func (b *committeeBuilder) ForwardToCommitteeList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ForwardToCommitteeIDs",
			relField: "ForwardToCommitteeList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) ManagerList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ManagerIDs",
			relField: "ManagerList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) MeetingList() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingIDs",
			relField: "MeetingList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) NativeUserList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "NativeUserIDs",
			relField: "NativeUserList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) Organization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationID",
			relField: "Organization",
		},
	}
}

func (b *committeeBuilder) OrganizationTagList() *organizationTagBuilder {
	return &organizationTagBuilder{
		builder: builder[organizationTagBuilder, *organizationTagBuilder, OrganizationTag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationTagIDs",
			relField: "OrganizationTagList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) Parent() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ParentID",
			relField: "Parent",
		},
	}
}

func (b *committeeBuilder) ReceiveForwardingsFromCommitteeList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReceiveForwardingsFromCommitteeIDs",
			relField: "ReceiveForwardingsFromCommitteeList",
			many:     true,
		},
	}
}

func (b *committeeBuilder) UserList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserIDs",
			relField: "UserList",
			many:     true,
		},
	}
}

func (r *Fetch) Committee(ids ...int) *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Gender has all fields from gender.
type Gender struct {
	ID             int
	Name           string
	OrganizationID int
	UserIDs        maybe.Maybe[[]int]
	Organization   *Organization
	UserList       []User
}

type genderBuilder struct {
	builder[genderBuilder, *genderBuilder, Gender]
}

func (b *genderBuilder) lazy(ds *Fetch, id int) *Gender {
	c := Gender{}
	ds.Gender_ID(id).Lazy(&c.ID)
	ds.Gender_Name(id).Lazy(&c.Name)
	ds.Gender_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Gender_UserIDs(id).Lazy(&c.UserIDs)
	return &c
}

func (b *genderBuilder) Preload(rel builderWrapperI) *genderBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *genderBuilder) Organization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationID",
			relField: "Organization",
		},
	}
}

func (b *genderBuilder) UserList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserIDs",
			relField: "UserList",
			many:     true,
		},
	}
}

func (r *Fetch) Gender(ids ...int) *genderBuilder {
	return &genderBuilder{
		builder: builder[genderBuilder, *genderBuilder, Gender]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Group has all fields from group.
type Group struct {
	AdminGroupForMeetingID                   maybe.Maybe[int]
	AnonymousGroupForMeetingID               maybe.Maybe[int]
	DefaultGroupForMeetingID                 maybe.Maybe[int]
	ExternalID                               maybe.Maybe[string]
	ID                                       int
	MeetingID                                int
	MeetingMediafileAccessGroupIDs           maybe.Maybe[[]int]
	MeetingMediafileInheritedAccessGroupIDs  maybe.Maybe[[]int]
	MeetingUserIDs                           maybe.Maybe[[]int]
	Name                                     string
	Permissions                              maybe.Maybe[[]string]
	PollIDs                                  maybe.Maybe[[]int]
	ReadChatGroupIDs                         maybe.Maybe[[]int]
	ReadCommentSectionIDs                    maybe.Maybe[[]int]
	UsedAsAssignmentPollDefaultID            maybe.Maybe[int]
	UsedAsMotionPollDefaultID                maybe.Maybe[int]
	UsedAsPollDefaultID                      maybe.Maybe[int]
	UsedAsTopicPollDefaultID                 maybe.Maybe[int]
	Weight                                   maybe.Maybe[int]
	WriteChatGroupIDs                        maybe.Maybe[[]int]
	WriteCommentSectionIDs                   maybe.Maybe[[]int]
	AdminGroupForMeeting                     *maybe.Maybe[Meeting]
	AnonymousGroupForMeeting                 *maybe.Maybe[Meeting]
	DefaultGroupForMeeting                   *maybe.Maybe[Meeting]
	Meeting                                  *Meeting
	MeetingMediafileAccessGroupList          []MeetingMediafile
	MeetingMediafileInheritedAccessGroupList []MeetingMediafile
	MeetingUserList                          []MeetingUser
	PollList                                 []Poll
	ReadChatGroupList                        []ChatGroup
	ReadCommentSectionList                   []MotionCommentSection
	UsedAsAssignmentPollDefault              *maybe.Maybe[Meeting]
	UsedAsMotionPollDefault                  *maybe.Maybe[Meeting]
	UsedAsPollDefault                        *maybe.Maybe[Meeting]
	UsedAsTopicPollDefault                   *maybe.Maybe[Meeting]
	WriteChatGroupList                       []ChatGroup
	WriteCommentSectionList                  []MotionCommentSection
}

type groupBuilder struct {
	builder[groupBuilder, *groupBuilder, Group]
}

func (b *groupBuilder) lazy(ds *Fetch, id int) *Group {
	c := Group{}
	ds.Group_AdminGroupForMeetingID(id).Lazy(&c.AdminGroupForMeetingID)
	ds.Group_AnonymousGroupForMeetingID(id).Lazy(&c.AnonymousGroupForMeetingID)
	ds.Group_DefaultGroupForMeetingID(id).Lazy(&c.DefaultGroupForMeetingID)
	ds.Group_ExternalID(id).Lazy(&c.ExternalID)
	ds.Group_ID(id).Lazy(&c.ID)
	ds.Group_MeetingID(id).Lazy(&c.MeetingID)
	ds.Group_MeetingMediafileAccessGroupIDs(id).Lazy(&c.MeetingMediafileAccessGroupIDs)
	ds.Group_MeetingMediafileInheritedAccessGroupIDs(id).Lazy(&c.MeetingMediafileInheritedAccessGroupIDs)
	ds.Group_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.Group_Name(id).Lazy(&c.Name)
	ds.Group_Permissions(id).Lazy(&c.Permissions)
	ds.Group_PollIDs(id).Lazy(&c.PollIDs)
	ds.Group_ReadChatGroupIDs(id).Lazy(&c.ReadChatGroupIDs)
	ds.Group_ReadCommentSectionIDs(id).Lazy(&c.ReadCommentSectionIDs)
	ds.Group_UsedAsAssignmentPollDefaultID(id).Lazy(&c.UsedAsAssignmentPollDefaultID)
	ds.Group_UsedAsMotionPollDefaultID(id).Lazy(&c.UsedAsMotionPollDefaultID)
	ds.Group_UsedAsPollDefaultID(id).Lazy(&c.UsedAsPollDefaultID)
	ds.Group_UsedAsTopicPollDefaultID(id).Lazy(&c.UsedAsTopicPollDefaultID)
	ds.Group_Weight(id).Lazy(&c.Weight)
	ds.Group_WriteChatGroupIDs(id).Lazy(&c.WriteChatGroupIDs)
	ds.Group_WriteCommentSectionIDs(id).Lazy(&c.WriteCommentSectionIDs)
	return &c
}

func (b *groupBuilder) Preload(rel builderWrapperI) *groupBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *groupBuilder) AdminGroupForMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AdminGroupForMeetingID",
			relField: "AdminGroupForMeeting",
		},
	}
}

func (b *groupBuilder) AnonymousGroupForMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AnonymousGroupForMeetingID",
			relField: "AnonymousGroupForMeeting",
		},
	}
}

func (b *groupBuilder) DefaultGroupForMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultGroupForMeetingID",
			relField: "DefaultGroupForMeeting",
		},
	}
}

func (b *groupBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *groupBuilder) MeetingMediafileAccessGroupList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingMediafileAccessGroupIDs",
			relField: "MeetingMediafileAccessGroupList",
			many:     true,
		},
	}
}

func (b *groupBuilder) MeetingMediafileInheritedAccessGroupList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingMediafileInheritedAccessGroupIDs",
			relField: "MeetingMediafileInheritedAccessGroupList",
			many:     true,
		},
	}
}

func (b *groupBuilder) MeetingUserList() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserIDs",
			relField: "MeetingUserList",
			many:     true,
		},
	}
}

func (b *groupBuilder) PollList() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollIDs",
			relField: "PollList",
			many:     true,
		},
	}
}

func (b *groupBuilder) ReadChatGroupList() *chatGroupBuilder {
	return &chatGroupBuilder{
		builder: builder[chatGroupBuilder, *chatGroupBuilder, ChatGroup]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReadChatGroupIDs",
			relField: "ReadChatGroupList",
			many:     true,
		},
	}
}

func (b *groupBuilder) ReadCommentSectionList() *motionCommentSectionBuilder {
	return &motionCommentSectionBuilder{
		builder: builder[motionCommentSectionBuilder, *motionCommentSectionBuilder, MotionCommentSection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReadCommentSectionIDs",
			relField: "ReadCommentSectionList",
			many:     true,
		},
	}
}

func (b *groupBuilder) UsedAsAssignmentPollDefault() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsAssignmentPollDefaultID",
			relField: "UsedAsAssignmentPollDefault",
		},
	}
}

func (b *groupBuilder) UsedAsMotionPollDefault() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsMotionPollDefaultID",
			relField: "UsedAsMotionPollDefault",
		},
	}
}

func (b *groupBuilder) UsedAsPollDefault() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsPollDefaultID",
			relField: "UsedAsPollDefault",
		},
	}
}

func (b *groupBuilder) UsedAsTopicPollDefault() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsTopicPollDefaultID",
			relField: "UsedAsTopicPollDefault",
		},
	}
}

func (b *groupBuilder) WriteChatGroupList() *chatGroupBuilder {
	return &chatGroupBuilder{
		builder: builder[chatGroupBuilder, *chatGroupBuilder, ChatGroup]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "WriteChatGroupIDs",
			relField: "WriteChatGroupList",
			many:     true,
		},
	}
}

func (b *groupBuilder) WriteCommentSectionList() *motionCommentSectionBuilder {
	return &motionCommentSectionBuilder{
		builder: builder[motionCommentSectionBuilder, *motionCommentSectionBuilder, MotionCommentSection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "WriteCommentSectionIDs",
			relField: "WriteCommentSectionList",
			many:     true,
		},
	}
}

func (r *Fetch) Group(ids ...int) *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			ids:   ids,
			fetch: r,
		},
	}
}

// HistoryEntry has all fields from history_entry.
type HistoryEntry struct {
	Entries         maybe.Maybe[[]string]
	ID              int
	MeetingID       maybe.Maybe[int]
	ModelID         maybe.Maybe[string]
	OriginalModelID maybe.Maybe[string]
	PositionID      int
	Meeting         *maybe.Maybe[Meeting]
	Position        *HistoryPosition
}

type historyEntryBuilder struct {
	builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]
}

func (b *historyEntryBuilder) lazy(ds *Fetch, id int) *HistoryEntry {
	c := HistoryEntry{}
	ds.HistoryEntry_Entries(id).Lazy(&c.Entries)
	ds.HistoryEntry_ID(id).Lazy(&c.ID)
	ds.HistoryEntry_MeetingID(id).Lazy(&c.MeetingID)
	ds.HistoryEntry_ModelID(id).Lazy(&c.ModelID)
	ds.HistoryEntry_OriginalModelID(id).Lazy(&c.OriginalModelID)
	ds.HistoryEntry_PositionID(id).Lazy(&c.PositionID)
	return &c
}

func (b *historyEntryBuilder) Preload(rel builderWrapperI) *historyEntryBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *historyEntryBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *historyEntryBuilder) Position() *historyPositionBuilder {
	return &historyPositionBuilder{
		builder: builder[historyPositionBuilder, *historyPositionBuilder, HistoryPosition]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PositionID",
			relField: "Position",
		},
	}
}

func (r *Fetch) HistoryEntry(ids ...int) *historyEntryBuilder {
	return &historyEntryBuilder{
		builder: builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]{
			ids:   ids,
			fetch: r,
		},
	}
}

// HistoryPosition has all fields from history_position.
type HistoryPosition struct {
	EntryIDs       maybe.Maybe[[]int]
	ID             int
	OriginalUserID maybe.Maybe[int]
	Timestamp      maybe.Maybe[int]
	UserID         maybe.Maybe[int]
	EntryList      []HistoryEntry
	User           *maybe.Maybe[User]
}

type historyPositionBuilder struct {
	builder[historyPositionBuilder, *historyPositionBuilder, HistoryPosition]
}

func (b *historyPositionBuilder) lazy(ds *Fetch, id int) *HistoryPosition {
	c := HistoryPosition{}
	ds.HistoryPosition_EntryIDs(id).Lazy(&c.EntryIDs)
	ds.HistoryPosition_ID(id).Lazy(&c.ID)
	ds.HistoryPosition_OriginalUserID(id).Lazy(&c.OriginalUserID)
	ds.HistoryPosition_Timestamp(id).Lazy(&c.Timestamp)
	ds.HistoryPosition_UserID(id).Lazy(&c.UserID)
	return &c
}

func (b *historyPositionBuilder) Preload(rel builderWrapperI) *historyPositionBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *historyPositionBuilder) EntryList() *historyEntryBuilder {
	return &historyEntryBuilder{
		builder: builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "EntryIDs",
			relField: "EntryList",
			many:     true,
		},
	}
}

func (b *historyPositionBuilder) User() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserID",
			relField: "User",
		},
	}
}

func (r *Fetch) HistoryPosition(ids ...int) *historyPositionBuilder {
	return &historyPositionBuilder{
		builder: builder[historyPositionBuilder, *historyPositionBuilder, HistoryPosition]{
			ids:   ids,
			fetch: r,
		},
	}
}

// ImportPreview has all fields from import_preview.
type ImportPreview struct {
	Created int
	ID      int
	Name    dstypes.ImportPreview_Name
	Result  maybe.Maybe[json.RawMessage]
	State   dstypes.ImportPreview_State
}

type importPreviewBuilder struct {
	builder[importPreviewBuilder, *importPreviewBuilder, ImportPreview]
}

func (b *importPreviewBuilder) lazy(ds *Fetch, id int) *ImportPreview {
	c := ImportPreview{}
	ds.ImportPreview_Created(id).Lazy(&c.Created)
	ds.ImportPreview_ID(id).Lazy(&c.ID)
	ds.ImportPreview_Name(id).Lazy(&c.Name)
	ds.ImportPreview_Result(id).Lazy(&c.Result)
	ds.ImportPreview_State(id).Lazy(&c.State)
	return &c
}

func (b *importPreviewBuilder) Preload(rel builderWrapperI) *importPreviewBuilder {
	b.builder.Preload(rel)
	return b
}

func (r *Fetch) ImportPreview(ids ...int) *importPreviewBuilder {
	return &importPreviewBuilder{
		builder: builder[importPreviewBuilder, *importPreviewBuilder, ImportPreview]{
			ids:   ids,
			fetch: r,
		},
	}
}

// ListOfSpeakers has all fields from list_of_speakers.
type ListOfSpeakers struct {
	Closed                           maybe.Maybe[bool]
	ContentObjectID                  string
	ID                               int
	MeetingID                        int
	ModeratorNotes                   maybe.Maybe[string]
	ProjectionIDs                    maybe.Maybe[[]int]
	SequentialNumber                 int
	SpeakerIDs                       maybe.Maybe[[]int]
	StructureLevelListOfSpeakersIDs  maybe.Maybe[[]int]
	Meeting                          *Meeting
	ProjectionList                   []Projection
	SpeakerList                      []Speaker
	StructureLevelListOfSpeakersList []StructureLevelListOfSpeakers
}

type listOfSpeakersBuilder struct {
	builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]
}

func (b *listOfSpeakersBuilder) lazy(ds *Fetch, id int) *ListOfSpeakers {
	c := ListOfSpeakers{}
	ds.ListOfSpeakers_Closed(id).Lazy(&c.Closed)
	ds.ListOfSpeakers_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.ListOfSpeakers_ID(id).Lazy(&c.ID)
	ds.ListOfSpeakers_MeetingID(id).Lazy(&c.MeetingID)
	ds.ListOfSpeakers_ModeratorNotes(id).Lazy(&c.ModeratorNotes)
	ds.ListOfSpeakers_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.ListOfSpeakers_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.ListOfSpeakers_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.ListOfSpeakers_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
	return &c
}

func (b *listOfSpeakersBuilder) Preload(rel builderWrapperI) *listOfSpeakersBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *listOfSpeakersBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *listOfSpeakersBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *listOfSpeakersBuilder) SpeakerList() *speakerBuilder {
	return &speakerBuilder{
		builder: builder[speakerBuilder, *speakerBuilder, Speaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SpeakerIDs",
			relField: "SpeakerList",
			many:     true,
		},
	}
}

func (b *listOfSpeakersBuilder) StructureLevelListOfSpeakersList() *structureLevelListOfSpeakersBuilder {
	return &structureLevelListOfSpeakersBuilder{
		builder: builder[structureLevelListOfSpeakersBuilder, *structureLevelListOfSpeakersBuilder, StructureLevelListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelListOfSpeakersIDs",
			relField: "StructureLevelListOfSpeakersList",
			many:     true,
		},
	}
}

func (r *Fetch) ListOfSpeakers(ids ...int) *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Mediafile has all fields from mediafile.
type Mediafile struct {
	ChildIDs                            maybe.Maybe[[]int]
	CreateTimestamp                     maybe.Maybe[int]
	Filename                            maybe.Maybe[string]
	Filesize                            maybe.Maybe[int]
	ID                                  int
	IsDirectory                         maybe.Maybe[bool]
	MeetingMediafileIDs                 maybe.Maybe[[]int]
	Mimetype                            maybe.Maybe[string]
	OwnerID                             string
	ParentID                            maybe.Maybe[int]
	PdfInformation                      maybe.Maybe[json.RawMessage]
	PublishedToMeetingsInOrganizationID maybe.Maybe[int]
	Title                               string
	Token                               maybe.Maybe[string]
	ChildList                           []Mediafile
	MeetingMediafileList                []MeetingMediafile
	Parent                              *maybe.Maybe[Mediafile]
	PublishedToMeetingsInOrganization   *maybe.Maybe[Organization]
}

type mediafileBuilder struct {
	builder[mediafileBuilder, *mediafileBuilder, Mediafile]
}

func (b *mediafileBuilder) lazy(ds *Fetch, id int) *Mediafile {
	c := Mediafile{}
	ds.Mediafile_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.Mediafile_CreateTimestamp(id).Lazy(&c.CreateTimestamp)
	ds.Mediafile_Filename(id).Lazy(&c.Filename)
	ds.Mediafile_Filesize(id).Lazy(&c.Filesize)
	ds.Mediafile_ID(id).Lazy(&c.ID)
	ds.Mediafile_IsDirectory(id).Lazy(&c.IsDirectory)
	ds.Mediafile_MeetingMediafileIDs(id).Lazy(&c.MeetingMediafileIDs)
	ds.Mediafile_Mimetype(id).Lazy(&c.Mimetype)
	ds.Mediafile_OwnerID(id).Lazy(&c.OwnerID)
	ds.Mediafile_ParentID(id).Lazy(&c.ParentID)
	ds.Mediafile_PdfInformation(id).Lazy(&c.PdfInformation)
	ds.Mediafile_PublishedToMeetingsInOrganizationID(id).Lazy(&c.PublishedToMeetingsInOrganizationID)
	ds.Mediafile_Title(id).Lazy(&c.Title)
	ds.Mediafile_Token(id).Lazy(&c.Token)
	return &c
}

func (b *mediafileBuilder) Preload(rel builderWrapperI) *mediafileBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *mediafileBuilder) ChildList() *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChildIDs",
			relField: "ChildList",
			many:     true,
		},
	}
}

func (b *mediafileBuilder) MeetingMediafileList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingMediafileIDs",
			relField: "MeetingMediafileList",
			many:     true,
		},
	}
}

func (b *mediafileBuilder) Parent() *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ParentID",
			relField: "Parent",
		},
	}
}

func (b *mediafileBuilder) PublishedToMeetingsInOrganization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PublishedToMeetingsInOrganizationID",
			relField: "PublishedToMeetingsInOrganization",
		},
	}
}

func (r *Fetch) Mediafile(ids ...int) *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Meeting has all fields from meeting.
type Meeting struct {
	AdminGroupID                                 maybe.Maybe[int]
	AgendaEnableNumbering                        maybe.Maybe[bool]
	AgendaItemCreation                           maybe.Maybe[dstypes.Meeting_AgendaItemCreation]
	AgendaItemIDs                                maybe.Maybe[[]int]
	AgendaNewItemsDefaultVisibility              maybe.Maybe[dstypes.Meeting_AgendaNewItemsDefaultVisibility]
	AgendaNumberPrefix                           maybe.Maybe[string]
	AgendaNumeralSystem                          maybe.Maybe[dstypes.Meeting_AgendaNumeralSystem]
	AgendaShowInternalItemsOnProjector           maybe.Maybe[bool]
	AgendaShowSubtitles                          maybe.Maybe[bool]
	AgendaShowTopicNavigationOnDetailView        maybe.Maybe[bool]
	AllProjectionIDs                             maybe.Maybe[[]int]
	AnonymousGroupID                             maybe.Maybe[int]
	ApplauseEnable                               maybe.Maybe[bool]
	ApplauseMaxAmount                            maybe.Maybe[int]
	ApplauseMinAmount                            maybe.Maybe[int]
	ApplauseParticleImageUrl                     maybe.Maybe[string]
	ApplauseShowLevel                            maybe.Maybe[bool]
	ApplauseTimeout                              maybe.Maybe[int]
	ApplauseType                                 maybe.Maybe[dstypes.Meeting_ApplauseType]
	AssignmentCandidateIDs                       maybe.Maybe[[]int]
	AssignmentIDs                                maybe.Maybe[[]int]
	AssignmentPollAddCandidatesToListOfSpeakers  maybe.Maybe[bool]
	AssignmentPollBallotPaperNumber              maybe.Maybe[int]
	AssignmentPollBallotPaperSelection           maybe.Maybe[dstypes.BallotPaperSelection]
	AssignmentPollDefaultBackend                 maybe.Maybe[dstypes.PollBackends]
	AssignmentPollDefaultGroupIDs                maybe.Maybe[[]int]
	AssignmentPollDefaultMethod                  maybe.Maybe[string]
	AssignmentPollDefaultOnehundredPercentBase   maybe.Maybe[dstypes.OnehundredPercentBases]
	AssignmentPollDefaultType                    maybe.Maybe[string]
	AssignmentPollEnableMaxVotesPerOption        maybe.Maybe[bool]
	AssignmentPollSortPollResultByVotes          maybe.Maybe[bool]
	AssignmentsExportPreamble                    maybe.Maybe[string]
	AssignmentsExportTitle                       maybe.Maybe[string]
	ChatGroupIDs                                 maybe.Maybe[[]int]
	ChatMessageIDs                               maybe.Maybe[[]int]
	CommitteeID                                  int
	ConferenceAutoConnect                        maybe.Maybe[bool]
	ConferenceAutoConnectNextSpeakers            maybe.Maybe[int]
	ConferenceEnableHelpdesk                     maybe.Maybe[bool]
	ConferenceLosRestriction                     maybe.Maybe[bool]
	ConferenceOpenMicrophone                     maybe.Maybe[bool]
	ConferenceOpenVideo                          maybe.Maybe[bool]
	ConferenceShow                               maybe.Maybe[bool]
	ConferenceStreamPosterUrl                    maybe.Maybe[string]
	ConferenceStreamUrl                          maybe.Maybe[string]
	CustomTranslations                           maybe.Maybe[json.RawMessage]
	DefaultGroupID                               int
	DefaultMeetingForCommitteeID                 maybe.Maybe[int]
	DefaultProjectorAgendaItemListIDs            []int
	DefaultProjectorAmendmentIDs                 []int
	DefaultProjectorAssignmentIDs                []int
	DefaultProjectorAssignmentPollIDs            []int
	DefaultProjectorCountdownIDs                 []int
	DefaultProjectorCurrentLosIDs                []int
	DefaultProjectorListOfSpeakersIDs            []int
	DefaultProjectorMediafileIDs                 []int
	DefaultProjectorMessageIDs                   []int
	DefaultProjectorMotionBlockIDs               []int
	DefaultProjectorMotionIDs                    []int
	DefaultProjectorMotionPollIDs                []int
	DefaultProjectorPollIDs                      []int
	DefaultProjectorTopicIDs                     []int
	Description                                  maybe.Maybe[string]
	EnableAnonymous                              maybe.Maybe[bool]
	EndTime                                      maybe.Maybe[int]
	ExportCsvEncoding                            maybe.Maybe[dstypes.Meeting_ExportCsvEncoding]
	ExportCsvSeparator                           maybe.Maybe[string]
	ExportPdfFontsize                            maybe.Maybe[int]
	ExportPdfLineHeight                          maybe.Maybe[float64]
	ExportPdfPageMarginBottom                    maybe.Maybe[int]
	ExportPdfPageMarginLeft                      maybe.Maybe[int]
	ExportPdfPageMarginRight                     maybe.Maybe[int]
	ExportPdfPageMarginTop                       maybe.Maybe[int]
	ExportPdfPagenumberAlignment                 maybe.Maybe[dstypes.Meeting_ExportPdfPagenumberAlignment]
	ExportPdfPagesize                            maybe.Maybe[dstypes.Meeting_ExportPdfPagesize]
	ExternalID                                   maybe.Maybe[string]
	FontBoldID                                   maybe.Maybe[int]
	FontBoldItalicID                             maybe.Maybe[int]
	FontChyronSpeakerNameID                      maybe.Maybe[int]
	FontItalicID                                 maybe.Maybe[int]
	FontMonospaceID                              maybe.Maybe[int]
	FontProjectorH1ID                            maybe.Maybe[int]
	FontProjectorH2ID                            maybe.Maybe[int]
	FontRegularID                                maybe.Maybe[int]
	ForwardedMotionIDs                           maybe.Maybe[[]int]
	GroupIDs                                     maybe.Maybe[[]int]
	ID                                           int
	ImportedAt                                   maybe.Maybe[int]
	IsActiveInOrganizationID                     maybe.Maybe[int]
	IsArchivedInOrganizationID                   maybe.Maybe[int]
	JitsiDomain                                  maybe.Maybe[string]
	JitsiRoomName                                maybe.Maybe[string]
	JitsiRoomPassword                            maybe.Maybe[string]
	Language                                     maybe.Maybe[dstypes.Languages]
	ListOfSpeakersAllowMultipleSpeakers          maybe.Maybe[bool]
	ListOfSpeakersAmountLastOnProjector          maybe.Maybe[int]
	ListOfSpeakersAmountNextOnProjector          maybe.Maybe[int]
	ListOfSpeakersCanCreatePointOfOrderForOthers maybe.Maybe[bool]
	ListOfSpeakersCanSetContributionSelf         maybe.Maybe[bool]
	ListOfSpeakersClosingDisablesPointOfOrder    maybe.Maybe[bool]
	ListOfSpeakersCountdownID                    maybe.Maybe[int]
	ListOfSpeakersCoupleCountdown                maybe.Maybe[bool]
	ListOfSpeakersDefaultStructureLevelTime      maybe.Maybe[int]
	ListOfSpeakersEnableInterposedQuestion       maybe.Maybe[bool]
	ListOfSpeakersEnablePointOfOrderCategories   maybe.Maybe[bool]
	ListOfSpeakersEnablePointOfOrderSpeakers     maybe.Maybe[bool]
	ListOfSpeakersEnableProContraSpeech          maybe.Maybe[bool]
	ListOfSpeakersHideContributionCount          maybe.Maybe[bool]
	ListOfSpeakersIDs                            maybe.Maybe[[]int]
	ListOfSpeakersInitiallyClosed                maybe.Maybe[bool]
	ListOfSpeakersInterventionTime               maybe.Maybe[int]
	ListOfSpeakersPresentUsersOnly               maybe.Maybe[bool]
	ListOfSpeakersShowAmountOfSpeakersOnSlide    maybe.Maybe[bool]
	ListOfSpeakersShowFirstContribution          maybe.Maybe[bool]
	ListOfSpeakersSpeakerNoteForEveryone         maybe.Maybe[bool]
	Location                                     maybe.Maybe[string]
	LockedFromInside                             maybe.Maybe[bool]
	LogoPdfBallotPaperID                         maybe.Maybe[int]
	LogoPdfFooterLID                             maybe.Maybe[int]
	LogoPdfFooterRID                             maybe.Maybe[int]
	LogoPdfHeaderLID                             maybe.Maybe[int]
	LogoPdfHeaderRID                             maybe.Maybe[int]
	LogoProjectorHeaderID                        maybe.Maybe[int]
	LogoProjectorMainID                          maybe.Maybe[int]
	LogoWebHeaderID                              maybe.Maybe[int]
	MediafileIDs                                 maybe.Maybe[[]int]
	MeetingMediafileIDs                          maybe.Maybe[[]int]
	MeetingUserIDs                               maybe.Maybe[[]int]
	MotionBlockIDs                               maybe.Maybe[[]int]
	MotionCategoryIDs                            maybe.Maybe[[]int]
	MotionChangeRecommendationIDs                maybe.Maybe[[]int]
	MotionCommentIDs                             maybe.Maybe[[]int]
	MotionCommentSectionIDs                      maybe.Maybe[[]int]
	MotionEditorIDs                              maybe.Maybe[[]int]
	MotionIDs                                    maybe.Maybe[[]int]
	MotionPollBallotPaperNumber                  maybe.Maybe[int]
	MotionPollBallotPaperSelection               maybe.Maybe[dstypes.BallotPaperSelection]
	MotionPollDefaultBackend                     maybe.Maybe[dstypes.PollBackends]
	MotionPollDefaultGroupIDs                    maybe.Maybe[[]int]
	MotionPollDefaultMethod                      maybe.Maybe[string]
	MotionPollDefaultOnehundredPercentBase       maybe.Maybe[dstypes.OnehundredPercentBases]
	MotionPollDefaultType                        maybe.Maybe[string]
	MotionPollProjectionMaxColumns               int
	MotionPollProjectionNameOrderFirst           dstypes.Meeting_MotionPollProjectionNameOrderFirst
	MotionStateIDs                               maybe.Maybe[[]int]
	MotionSubmitterIDs                           maybe.Maybe[[]int]
	MotionSupporterIDs                           maybe.Maybe[[]int]
	MotionWorkflowIDs                            maybe.Maybe[[]int]
	MotionWorkingGroupSpeakerIDs                 maybe.Maybe[[]int]
	MotionsAmendmentsEnabled                     maybe.Maybe[bool]
	MotionsAmendmentsInMainList                  maybe.Maybe[bool]
	MotionsAmendmentsMultipleParagraphs          maybe.Maybe[bool]
	MotionsAmendmentsOfAmendments                maybe.Maybe[bool]
	MotionsAmendmentsPrefix                      maybe.Maybe[string]
	MotionsAmendmentsTextMode                    maybe.Maybe[dstypes.Meeting_MotionsAmendmentsTextMode]
	MotionsBlockSlideColumns                     maybe.Maybe[int]
	MotionsCreateEnableAdditionalSubmitterText   maybe.Maybe[bool]
	MotionsDefaultAmendmentWorkflowID            int
	MotionsDefaultLineNumbering                  maybe.Maybe[dstypes.Meeting_MotionsDefaultLineNumbering]
	MotionsDefaultSorting                        maybe.Maybe[dstypes.Meeting_MotionsDefaultSorting]
	MotionsDefaultWorkflowID                     int
	MotionsEnableEditor                          maybe.Maybe[bool]
	MotionsEnableOriginMotionDisplay             maybe.Maybe[bool]
	MotionsEnableReasonOnProjector               maybe.Maybe[bool]
	MotionsEnableRecommendationOnProjector       maybe.Maybe[bool]
	MotionsEnableRestrictedEditorForManager      maybe.Maybe[bool]
	MotionsEnableRestrictedEditorForNonManager   maybe.Maybe[bool]
	MotionsEnableSideboxOnProjector              maybe.Maybe[bool]
	MotionsEnableTextOnProjector                 maybe.Maybe[bool]
	MotionsEnableWorkingGroupSpeaker             maybe.Maybe[bool]
	MotionsExportFollowRecommendation            maybe.Maybe[bool]
	MotionsExportPreamble                        maybe.Maybe[string]
	MotionsExportSubmitterRecommendation         maybe.Maybe[bool]
	MotionsExportTitle                           maybe.Maybe[string]
	MotionsHideMetadataBackground                maybe.Maybe[bool]
	MotionsLineLength                            maybe.Maybe[int]
	MotionsNumberMinDigits                       maybe.Maybe[int]
	MotionsNumberType                            maybe.Maybe[dstypes.Meeting_MotionsNumberType]
	MotionsNumberWithBlank                       maybe.Maybe[bool]
	MotionsOriginMotionToggleDefault             maybe.Maybe[bool]
	MotionsPreamble                              maybe.Maybe[string]
	MotionsReasonRequired                        maybe.Maybe[bool]
	MotionsRecommendationTextMode                maybe.Maybe[dstypes.Meeting_MotionsRecommendationTextMode]
	MotionsRecommendationsBy                     maybe.Maybe[string]
	MotionsShowReferringMotions                  maybe.Maybe[bool]
	MotionsShowSequentialNumber                  maybe.Maybe[bool]
	MotionsSupportersMinAmount                   maybe.Maybe[int]
	Name                                         string
	OptionIDs                                    maybe.Maybe[[]int]
	OrganizationTagIDs                           maybe.Maybe[[]int]
	PersonalNoteIDs                              maybe.Maybe[[]int]
	PointOfOrderCategoryIDs                      maybe.Maybe[[]int]
	PollBallotPaperNumber                        maybe.Maybe[int]
	PollBallotPaperSelection                     maybe.Maybe[dstypes.BallotPaperSelection]
	PollCandidateIDs                             maybe.Maybe[[]int]
	PollCandidateListIDs                         maybe.Maybe[[]int]
	PollCountdownID                              maybe.Maybe[int]
	PollCoupleCountdown                          maybe.Maybe[bool]
	PollDefaultBackend                           maybe.Maybe[dstypes.PollBackends]
	PollDefaultGroupIDs                          maybe.Maybe[[]int]
	PollDefaultLiveVotingEnabled                 maybe.Maybe[bool]
	PollDefaultMethod                            maybe.Maybe[string]
	PollDefaultOnehundredPercentBase             maybe.Maybe[dstypes.OnehundredPercentBases]
	PollDefaultType                              maybe.Maybe[string]
	PollIDs                                      maybe.Maybe[[]int]
	PollSortPollResultByVotes                    maybe.Maybe[bool]
	PresentUserIDs                               maybe.Maybe[[]int]
	ProjectionIDs                                maybe.Maybe[[]int]
	ProjectorCountdownDefaultTime                int
	ProjectorCountdownIDs                        maybe.Maybe[[]int]
	ProjectorCountdownWarningTime                int
	ProjectorIDs                                 maybe.Maybe[[]int]
	ProjectorMessageIDs                          maybe.Maybe[[]int]
	ReferenceProjectorID                         int
	RelevantHistoryEntryIDs                      maybe.Maybe[[]int]
	SpeakerIDs                                   maybe.Maybe[[]int]
	StartTime                                    maybe.Maybe[int]
	StructureLevelIDs                            maybe.Maybe[[]int]
	StructureLevelListOfSpeakersIDs              maybe.Maybe[[]int]
	TagIDs                                       maybe.Maybe[[]int]
	TemplateForOrganizationID                    maybe.Maybe[int]
	TimeZone                                     maybe.Maybe[string]
	TopicIDs                                     maybe.Maybe[[]int]
	TopicPollDefaultGroupIDs                     maybe.Maybe[[]int]
	UserIDs                                      maybe.Maybe[[]int]
	UsersAllowSelfSetPresent                     maybe.Maybe[bool]
	UsersEmailBody                               maybe.Maybe[string]
	UsersEmailReplyto                            maybe.Maybe[string]
	UsersEmailSender                             maybe.Maybe[string]
	UsersEmailSubject                            maybe.Maybe[string]
	UsersEnablePresenceView                      maybe.Maybe[bool]
	UsersEnableVoteDelegations                   maybe.Maybe[bool]
	UsersEnableVoteWeight                        maybe.Maybe[bool]
	UsersForbidDelegatorAsSubmitter              maybe.Maybe[bool]
	UsersForbidDelegatorAsSupporter              maybe.Maybe[bool]
	UsersForbidDelegatorInListOfSpeakers         maybe.Maybe[bool]
	UsersForbidDelegatorToVote                   maybe.Maybe[bool]
	UsersPdfWelcometext                          maybe.Maybe[string]
	UsersPdfWelcometitle                         maybe.Maybe[string]
	UsersPdfWlanEncryption                       maybe.Maybe[dstypes.Meeting_UsersPdfWlanEncryption]
	UsersPdfWlanPassword                         maybe.Maybe[string]
	UsersPdfWlanSsid                             maybe.Maybe[string]
	VoteIDs                                      maybe.Maybe[[]int]
	WelcomeText                                  maybe.Maybe[string]
	WelcomeTitle                                 maybe.Maybe[string]
	AdminGroup                                   *maybe.Maybe[Group]
	AgendaItemList                               []AgendaItem
	AllProjectionList                            []Projection
	AnonymousGroup                               *maybe.Maybe[Group]
	AssignmentCandidateList                      []AssignmentCandidate
	AssignmentList                               []Assignment
	AssignmentPollDefaultGroupList               []Group
	ChatGroupList                                []ChatGroup
	ChatMessageList                              []ChatMessage
	Committee                                    *Committee
	DefaultGroup                                 *Group
	DefaultMeetingForCommittee                   *maybe.Maybe[Committee]
	DefaultProjectorAgendaItemListList           []Projector
	DefaultProjectorAmendmentList                []Projector
	DefaultProjectorAssignmentList               []Projector
	DefaultProjectorAssignmentPollList           []Projector
	DefaultProjectorCountdownList                []Projector
	DefaultProjectorCurrentLosList               []Projector
	DefaultProjectorListOfSpeakersList           []Projector
	DefaultProjectorMediafileList                []Projector
	DefaultProjectorMessageList                  []Projector
	DefaultProjectorMotionBlockList              []Projector
	DefaultProjectorMotionList                   []Projector
	DefaultProjectorMotionPollList               []Projector
	DefaultProjectorPollList                     []Projector
	DefaultProjectorTopicList                    []Projector
	FontBold                                     *maybe.Maybe[MeetingMediafile]
	FontBoldItalic                               *maybe.Maybe[MeetingMediafile]
	FontChyronSpeakerName                        *maybe.Maybe[MeetingMediafile]
	FontItalic                                   *maybe.Maybe[MeetingMediafile]
	FontMonospace                                *maybe.Maybe[MeetingMediafile]
	FontProjectorH1                              *maybe.Maybe[MeetingMediafile]
	FontProjectorH2                              *maybe.Maybe[MeetingMediafile]
	FontRegular                                  *maybe.Maybe[MeetingMediafile]
	ForwardedMotionList                          []Motion
	GroupList                                    []Group
	IsActiveInOrganization                       *maybe.Maybe[Organization]
	IsArchivedInOrganization                     *maybe.Maybe[Organization]
	ListOfSpeakersCountdown                      *maybe.Maybe[ProjectorCountdown]
	ListOfSpeakersList                           []ListOfSpeakers
	LogoPdfBallotPaper                           *maybe.Maybe[MeetingMediafile]
	LogoPdfFooterL                               *maybe.Maybe[MeetingMediafile]
	LogoPdfFooterR                               *maybe.Maybe[MeetingMediafile]
	LogoPdfHeaderL                               *maybe.Maybe[MeetingMediafile]
	LogoPdfHeaderR                               *maybe.Maybe[MeetingMediafile]
	LogoProjectorHeader                          *maybe.Maybe[MeetingMediafile]
	LogoProjectorMain                            *maybe.Maybe[MeetingMediafile]
	LogoWebHeader                                *maybe.Maybe[MeetingMediafile]
	MediafileList                                []Mediafile
	MeetingMediafileList                         []MeetingMediafile
	MeetingUserList                              []MeetingUser
	MotionBlockList                              []MotionBlock
	MotionCategoryList                           []MotionCategory
	MotionChangeRecommendationList               []MotionChangeRecommendation
	MotionCommentList                            []MotionComment
	MotionCommentSectionList                     []MotionCommentSection
	MotionEditorList                             []MotionEditor
	MotionList                                   []Motion
	MotionPollDefaultGroupList                   []Group
	MotionStateList                              []MotionState
	MotionSubmitterList                          []MotionSubmitter
	MotionSupporterList                          []MotionSupporter
	MotionWorkflowList                           []MotionWorkflow
	MotionWorkingGroupSpeakerList                []MotionWorkingGroupSpeaker
	MotionsDefaultAmendmentWorkflow              *MotionWorkflow
	MotionsDefaultWorkflow                       *MotionWorkflow
	OptionList                                   []Option
	OrganizationTagList                          []OrganizationTag
	PersonalNoteList                             []PersonalNote
	PointOfOrderCategoryList                     []PointOfOrderCategory
	PollCandidateList                            []PollCandidate
	PollCandidateListList                        []PollCandidateList
	PollCountdown                                *maybe.Maybe[ProjectorCountdown]
	PollDefaultGroupList                         []Group
	PollList                                     []Poll
	PresentUserList                              []User
	ProjectionList                               []Projection
	ProjectorCountdownList                       []ProjectorCountdown
	ProjectorList                                []Projector
	ProjectorMessageList                         []ProjectorMessage
	ReferenceProjector                           *Projector
	RelevantHistoryEntryList                     []HistoryEntry
	SpeakerList                                  []Speaker
	StructureLevelList                           []StructureLevel
	StructureLevelListOfSpeakersList             []StructureLevelListOfSpeakers
	TagList                                      []Tag
	TemplateForOrganization                      *maybe.Maybe[Organization]
	TopicList                                    []Topic
	TopicPollDefaultGroupList                    []Group
	UserList                                     []User
	VoteList                                     []Vote
}

type meetingBuilder struct {
	builder[meetingBuilder, *meetingBuilder, Meeting]
}

func (b *meetingBuilder) lazy(ds *Fetch, id int) *Meeting {
	c := Meeting{}
	ds.Meeting_AdminGroupID(id).Lazy(&c.AdminGroupID)
	ds.Meeting_AgendaEnableNumbering(id).Lazy(&c.AgendaEnableNumbering)
	ds.Meeting_AgendaItemCreation(id).Lazy(&c.AgendaItemCreation)
	ds.Meeting_AgendaItemIDs(id).Lazy(&c.AgendaItemIDs)
	ds.Meeting_AgendaNewItemsDefaultVisibility(id).Lazy(&c.AgendaNewItemsDefaultVisibility)
	ds.Meeting_AgendaNumberPrefix(id).Lazy(&c.AgendaNumberPrefix)
	ds.Meeting_AgendaNumeralSystem(id).Lazy(&c.AgendaNumeralSystem)
	ds.Meeting_AgendaShowInternalItemsOnProjector(id).Lazy(&c.AgendaShowInternalItemsOnProjector)
	ds.Meeting_AgendaShowSubtitles(id).Lazy(&c.AgendaShowSubtitles)
	ds.Meeting_AgendaShowTopicNavigationOnDetailView(id).Lazy(&c.AgendaShowTopicNavigationOnDetailView)
	ds.Meeting_AllProjectionIDs(id).Lazy(&c.AllProjectionIDs)
	ds.Meeting_AnonymousGroupID(id).Lazy(&c.AnonymousGroupID)
	ds.Meeting_ApplauseEnable(id).Lazy(&c.ApplauseEnable)
	ds.Meeting_ApplauseMaxAmount(id).Lazy(&c.ApplauseMaxAmount)
	ds.Meeting_ApplauseMinAmount(id).Lazy(&c.ApplauseMinAmount)
	ds.Meeting_ApplauseParticleImageUrl(id).Lazy(&c.ApplauseParticleImageUrl)
	ds.Meeting_ApplauseShowLevel(id).Lazy(&c.ApplauseShowLevel)
	ds.Meeting_ApplauseTimeout(id).Lazy(&c.ApplauseTimeout)
	ds.Meeting_ApplauseType(id).Lazy(&c.ApplauseType)
	ds.Meeting_AssignmentCandidateIDs(id).Lazy(&c.AssignmentCandidateIDs)
	ds.Meeting_AssignmentIDs(id).Lazy(&c.AssignmentIDs)
	ds.Meeting_AssignmentPollAddCandidatesToListOfSpeakers(id).Lazy(&c.AssignmentPollAddCandidatesToListOfSpeakers)
	ds.Meeting_AssignmentPollBallotPaperNumber(id).Lazy(&c.AssignmentPollBallotPaperNumber)
	ds.Meeting_AssignmentPollBallotPaperSelection(id).Lazy(&c.AssignmentPollBallotPaperSelection)
	ds.Meeting_AssignmentPollDefaultBackend(id).Lazy(&c.AssignmentPollDefaultBackend)
	ds.Meeting_AssignmentPollDefaultGroupIDs(id).Lazy(&c.AssignmentPollDefaultGroupIDs)
	ds.Meeting_AssignmentPollDefaultMethod(id).Lazy(&c.AssignmentPollDefaultMethod)
	ds.Meeting_AssignmentPollDefaultOnehundredPercentBase(id).Lazy(&c.AssignmentPollDefaultOnehundredPercentBase)
	ds.Meeting_AssignmentPollDefaultType(id).Lazy(&c.AssignmentPollDefaultType)
	ds.Meeting_AssignmentPollEnableMaxVotesPerOption(id).Lazy(&c.AssignmentPollEnableMaxVotesPerOption)
	ds.Meeting_AssignmentPollSortPollResultByVotes(id).Lazy(&c.AssignmentPollSortPollResultByVotes)
	ds.Meeting_AssignmentsExportPreamble(id).Lazy(&c.AssignmentsExportPreamble)
	ds.Meeting_AssignmentsExportTitle(id).Lazy(&c.AssignmentsExportTitle)
	ds.Meeting_ChatGroupIDs(id).Lazy(&c.ChatGroupIDs)
	ds.Meeting_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.Meeting_CommitteeID(id).Lazy(&c.CommitteeID)
	ds.Meeting_ConferenceAutoConnect(id).Lazy(&c.ConferenceAutoConnect)
	ds.Meeting_ConferenceAutoConnectNextSpeakers(id).Lazy(&c.ConferenceAutoConnectNextSpeakers)
	ds.Meeting_ConferenceEnableHelpdesk(id).Lazy(&c.ConferenceEnableHelpdesk)
	ds.Meeting_ConferenceLosRestriction(id).Lazy(&c.ConferenceLosRestriction)
	ds.Meeting_ConferenceOpenMicrophone(id).Lazy(&c.ConferenceOpenMicrophone)
	ds.Meeting_ConferenceOpenVideo(id).Lazy(&c.ConferenceOpenVideo)
	ds.Meeting_ConferenceShow(id).Lazy(&c.ConferenceShow)
	ds.Meeting_ConferenceStreamPosterUrl(id).Lazy(&c.ConferenceStreamPosterUrl)
	ds.Meeting_ConferenceStreamUrl(id).Lazy(&c.ConferenceStreamUrl)
	ds.Meeting_CustomTranslations(id).Lazy(&c.CustomTranslations)
	ds.Meeting_DefaultGroupID(id).Lazy(&c.DefaultGroupID)
	ds.Meeting_DefaultMeetingForCommitteeID(id).Lazy(&c.DefaultMeetingForCommitteeID)
	ds.Meeting_DefaultProjectorAgendaItemListIDs(id).Lazy(&c.DefaultProjectorAgendaItemListIDs)
	ds.Meeting_DefaultProjectorAmendmentIDs(id).Lazy(&c.DefaultProjectorAmendmentIDs)
	ds.Meeting_DefaultProjectorAssignmentIDs(id).Lazy(&c.DefaultProjectorAssignmentIDs)
	ds.Meeting_DefaultProjectorAssignmentPollIDs(id).Lazy(&c.DefaultProjectorAssignmentPollIDs)
	ds.Meeting_DefaultProjectorCountdownIDs(id).Lazy(&c.DefaultProjectorCountdownIDs)
	ds.Meeting_DefaultProjectorCurrentLosIDs(id).Lazy(&c.DefaultProjectorCurrentLosIDs)
	ds.Meeting_DefaultProjectorListOfSpeakersIDs(id).Lazy(&c.DefaultProjectorListOfSpeakersIDs)
	ds.Meeting_DefaultProjectorMediafileIDs(id).Lazy(&c.DefaultProjectorMediafileIDs)
	ds.Meeting_DefaultProjectorMessageIDs(id).Lazy(&c.DefaultProjectorMessageIDs)
	ds.Meeting_DefaultProjectorMotionBlockIDs(id).Lazy(&c.DefaultProjectorMotionBlockIDs)
	ds.Meeting_DefaultProjectorMotionIDs(id).Lazy(&c.DefaultProjectorMotionIDs)
	ds.Meeting_DefaultProjectorMotionPollIDs(id).Lazy(&c.DefaultProjectorMotionPollIDs)
	ds.Meeting_DefaultProjectorPollIDs(id).Lazy(&c.DefaultProjectorPollIDs)
	ds.Meeting_DefaultProjectorTopicIDs(id).Lazy(&c.DefaultProjectorTopicIDs)
	ds.Meeting_Description(id).Lazy(&c.Description)
	ds.Meeting_EnableAnonymous(id).Lazy(&c.EnableAnonymous)
	ds.Meeting_EndTime(id).Lazy(&c.EndTime)
	ds.Meeting_ExportCsvEncoding(id).Lazy(&c.ExportCsvEncoding)
	ds.Meeting_ExportCsvSeparator(id).Lazy(&c.ExportCsvSeparator)
	ds.Meeting_ExportPdfFontsize(id).Lazy(&c.ExportPdfFontsize)
	ds.Meeting_ExportPdfLineHeight(id).Lazy(&c.ExportPdfLineHeight)
	ds.Meeting_ExportPdfPageMarginBottom(id).Lazy(&c.ExportPdfPageMarginBottom)
	ds.Meeting_ExportPdfPageMarginLeft(id).Lazy(&c.ExportPdfPageMarginLeft)
	ds.Meeting_ExportPdfPageMarginRight(id).Lazy(&c.ExportPdfPageMarginRight)
	ds.Meeting_ExportPdfPageMarginTop(id).Lazy(&c.ExportPdfPageMarginTop)
	ds.Meeting_ExportPdfPagenumberAlignment(id).Lazy(&c.ExportPdfPagenumberAlignment)
	ds.Meeting_ExportPdfPagesize(id).Lazy(&c.ExportPdfPagesize)
	ds.Meeting_ExternalID(id).Lazy(&c.ExternalID)
	ds.Meeting_FontBoldID(id).Lazy(&c.FontBoldID)
	ds.Meeting_FontBoldItalicID(id).Lazy(&c.FontBoldItalicID)
	ds.Meeting_FontChyronSpeakerNameID(id).Lazy(&c.FontChyronSpeakerNameID)
	ds.Meeting_FontItalicID(id).Lazy(&c.FontItalicID)
	ds.Meeting_FontMonospaceID(id).Lazy(&c.FontMonospaceID)
	ds.Meeting_FontProjectorH1ID(id).Lazy(&c.FontProjectorH1ID)
	ds.Meeting_FontProjectorH2ID(id).Lazy(&c.FontProjectorH2ID)
	ds.Meeting_FontRegularID(id).Lazy(&c.FontRegularID)
	ds.Meeting_ForwardedMotionIDs(id).Lazy(&c.ForwardedMotionIDs)
	ds.Meeting_GroupIDs(id).Lazy(&c.GroupIDs)
	ds.Meeting_ID(id).Lazy(&c.ID)
	ds.Meeting_ImportedAt(id).Lazy(&c.ImportedAt)
	ds.Meeting_IsActiveInOrganizationID(id).Lazy(&c.IsActiveInOrganizationID)
	ds.Meeting_IsArchivedInOrganizationID(id).Lazy(&c.IsArchivedInOrganizationID)
	ds.Meeting_JitsiDomain(id).Lazy(&c.JitsiDomain)
	ds.Meeting_JitsiRoomName(id).Lazy(&c.JitsiRoomName)
	ds.Meeting_JitsiRoomPassword(id).Lazy(&c.JitsiRoomPassword)
	ds.Meeting_Language(id).Lazy(&c.Language)
	ds.Meeting_ListOfSpeakersAllowMultipleSpeakers(id).Lazy(&c.ListOfSpeakersAllowMultipleSpeakers)
	ds.Meeting_ListOfSpeakersAmountLastOnProjector(id).Lazy(&c.ListOfSpeakersAmountLastOnProjector)
	ds.Meeting_ListOfSpeakersAmountNextOnProjector(id).Lazy(&c.ListOfSpeakersAmountNextOnProjector)
	ds.Meeting_ListOfSpeakersCanCreatePointOfOrderForOthers(id).Lazy(&c.ListOfSpeakersCanCreatePointOfOrderForOthers)
	ds.Meeting_ListOfSpeakersCanSetContributionSelf(id).Lazy(&c.ListOfSpeakersCanSetContributionSelf)
	ds.Meeting_ListOfSpeakersClosingDisablesPointOfOrder(id).Lazy(&c.ListOfSpeakersClosingDisablesPointOfOrder)
	ds.Meeting_ListOfSpeakersCountdownID(id).Lazy(&c.ListOfSpeakersCountdownID)
	ds.Meeting_ListOfSpeakersCoupleCountdown(id).Lazy(&c.ListOfSpeakersCoupleCountdown)
	ds.Meeting_ListOfSpeakersDefaultStructureLevelTime(id).Lazy(&c.ListOfSpeakersDefaultStructureLevelTime)
	ds.Meeting_ListOfSpeakersEnableInterposedQuestion(id).Lazy(&c.ListOfSpeakersEnableInterposedQuestion)
	ds.Meeting_ListOfSpeakersEnablePointOfOrderCategories(id).Lazy(&c.ListOfSpeakersEnablePointOfOrderCategories)
	ds.Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(id).Lazy(&c.ListOfSpeakersEnablePointOfOrderSpeakers)
	ds.Meeting_ListOfSpeakersEnableProContraSpeech(id).Lazy(&c.ListOfSpeakersEnableProContraSpeech)
	ds.Meeting_ListOfSpeakersHideContributionCount(id).Lazy(&c.ListOfSpeakersHideContributionCount)
	ds.Meeting_ListOfSpeakersIDs(id).Lazy(&c.ListOfSpeakersIDs)
	ds.Meeting_ListOfSpeakersInitiallyClosed(id).Lazy(&c.ListOfSpeakersInitiallyClosed)
	ds.Meeting_ListOfSpeakersInterventionTime(id).Lazy(&c.ListOfSpeakersInterventionTime)
	ds.Meeting_ListOfSpeakersPresentUsersOnly(id).Lazy(&c.ListOfSpeakersPresentUsersOnly)
	ds.Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(id).Lazy(&c.ListOfSpeakersShowAmountOfSpeakersOnSlide)
	ds.Meeting_ListOfSpeakersShowFirstContribution(id).Lazy(&c.ListOfSpeakersShowFirstContribution)
	ds.Meeting_ListOfSpeakersSpeakerNoteForEveryone(id).Lazy(&c.ListOfSpeakersSpeakerNoteForEveryone)
	ds.Meeting_Location(id).Lazy(&c.Location)
	ds.Meeting_LockedFromInside(id).Lazy(&c.LockedFromInside)
	ds.Meeting_LogoPdfBallotPaperID(id).Lazy(&c.LogoPdfBallotPaperID)
	ds.Meeting_LogoPdfFooterLID(id).Lazy(&c.LogoPdfFooterLID)
	ds.Meeting_LogoPdfFooterRID(id).Lazy(&c.LogoPdfFooterRID)
	ds.Meeting_LogoPdfHeaderLID(id).Lazy(&c.LogoPdfHeaderLID)
	ds.Meeting_LogoPdfHeaderRID(id).Lazy(&c.LogoPdfHeaderRID)
	ds.Meeting_LogoProjectorHeaderID(id).Lazy(&c.LogoProjectorHeaderID)
	ds.Meeting_LogoProjectorMainID(id).Lazy(&c.LogoProjectorMainID)
	ds.Meeting_LogoWebHeaderID(id).Lazy(&c.LogoWebHeaderID)
	ds.Meeting_MediafileIDs(id).Lazy(&c.MediafileIDs)
	ds.Meeting_MeetingMediafileIDs(id).Lazy(&c.MeetingMediafileIDs)
	ds.Meeting_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.Meeting_MotionBlockIDs(id).Lazy(&c.MotionBlockIDs)
	ds.Meeting_MotionCategoryIDs(id).Lazy(&c.MotionCategoryIDs)
	ds.Meeting_MotionChangeRecommendationIDs(id).Lazy(&c.MotionChangeRecommendationIDs)
	ds.Meeting_MotionCommentIDs(id).Lazy(&c.MotionCommentIDs)
	ds.Meeting_MotionCommentSectionIDs(id).Lazy(&c.MotionCommentSectionIDs)
	ds.Meeting_MotionEditorIDs(id).Lazy(&c.MotionEditorIDs)
	ds.Meeting_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.Meeting_MotionPollBallotPaperNumber(id).Lazy(&c.MotionPollBallotPaperNumber)
	ds.Meeting_MotionPollBallotPaperSelection(id).Lazy(&c.MotionPollBallotPaperSelection)
	ds.Meeting_MotionPollDefaultBackend(id).Lazy(&c.MotionPollDefaultBackend)
	ds.Meeting_MotionPollDefaultGroupIDs(id).Lazy(&c.MotionPollDefaultGroupIDs)
	ds.Meeting_MotionPollDefaultMethod(id).Lazy(&c.MotionPollDefaultMethod)
	ds.Meeting_MotionPollDefaultOnehundredPercentBase(id).Lazy(&c.MotionPollDefaultOnehundredPercentBase)
	ds.Meeting_MotionPollDefaultType(id).Lazy(&c.MotionPollDefaultType)
	ds.Meeting_MotionPollProjectionMaxColumns(id).Lazy(&c.MotionPollProjectionMaxColumns)
	ds.Meeting_MotionPollProjectionNameOrderFirst(id).Lazy(&c.MotionPollProjectionNameOrderFirst)
	ds.Meeting_MotionStateIDs(id).Lazy(&c.MotionStateIDs)
	ds.Meeting_MotionSubmitterIDs(id).Lazy(&c.MotionSubmitterIDs)
	ds.Meeting_MotionSupporterIDs(id).Lazy(&c.MotionSupporterIDs)
	ds.Meeting_MotionWorkflowIDs(id).Lazy(&c.MotionWorkflowIDs)
	ds.Meeting_MotionWorkingGroupSpeakerIDs(id).Lazy(&c.MotionWorkingGroupSpeakerIDs)
	ds.Meeting_MotionsAmendmentsEnabled(id).Lazy(&c.MotionsAmendmentsEnabled)
	ds.Meeting_MotionsAmendmentsInMainList(id).Lazy(&c.MotionsAmendmentsInMainList)
	ds.Meeting_MotionsAmendmentsMultipleParagraphs(id).Lazy(&c.MotionsAmendmentsMultipleParagraphs)
	ds.Meeting_MotionsAmendmentsOfAmendments(id).Lazy(&c.MotionsAmendmentsOfAmendments)
	ds.Meeting_MotionsAmendmentsPrefix(id).Lazy(&c.MotionsAmendmentsPrefix)
	ds.Meeting_MotionsAmendmentsTextMode(id).Lazy(&c.MotionsAmendmentsTextMode)
	ds.Meeting_MotionsBlockSlideColumns(id).Lazy(&c.MotionsBlockSlideColumns)
	ds.Meeting_MotionsCreateEnableAdditionalSubmitterText(id).Lazy(&c.MotionsCreateEnableAdditionalSubmitterText)
	ds.Meeting_MotionsDefaultAmendmentWorkflowID(id).Lazy(&c.MotionsDefaultAmendmentWorkflowID)
	ds.Meeting_MotionsDefaultLineNumbering(id).Lazy(&c.MotionsDefaultLineNumbering)
	ds.Meeting_MotionsDefaultSorting(id).Lazy(&c.MotionsDefaultSorting)
	ds.Meeting_MotionsDefaultWorkflowID(id).Lazy(&c.MotionsDefaultWorkflowID)
	ds.Meeting_MotionsEnableEditor(id).Lazy(&c.MotionsEnableEditor)
	ds.Meeting_MotionsEnableOriginMotionDisplay(id).Lazy(&c.MotionsEnableOriginMotionDisplay)
	ds.Meeting_MotionsEnableReasonOnProjector(id).Lazy(&c.MotionsEnableReasonOnProjector)
	ds.Meeting_MotionsEnableRecommendationOnProjector(id).Lazy(&c.MotionsEnableRecommendationOnProjector)
	ds.Meeting_MotionsEnableRestrictedEditorForManager(id).Lazy(&c.MotionsEnableRestrictedEditorForManager)
	ds.Meeting_MotionsEnableRestrictedEditorForNonManager(id).Lazy(&c.MotionsEnableRestrictedEditorForNonManager)
	ds.Meeting_MotionsEnableSideboxOnProjector(id).Lazy(&c.MotionsEnableSideboxOnProjector)
	ds.Meeting_MotionsEnableTextOnProjector(id).Lazy(&c.MotionsEnableTextOnProjector)
	ds.Meeting_MotionsEnableWorkingGroupSpeaker(id).Lazy(&c.MotionsEnableWorkingGroupSpeaker)
	ds.Meeting_MotionsExportFollowRecommendation(id).Lazy(&c.MotionsExportFollowRecommendation)
	ds.Meeting_MotionsExportPreamble(id).Lazy(&c.MotionsExportPreamble)
	ds.Meeting_MotionsExportSubmitterRecommendation(id).Lazy(&c.MotionsExportSubmitterRecommendation)
	ds.Meeting_MotionsExportTitle(id).Lazy(&c.MotionsExportTitle)
	ds.Meeting_MotionsHideMetadataBackground(id).Lazy(&c.MotionsHideMetadataBackground)
	ds.Meeting_MotionsLineLength(id).Lazy(&c.MotionsLineLength)
	ds.Meeting_MotionsNumberMinDigits(id).Lazy(&c.MotionsNumberMinDigits)
	ds.Meeting_MotionsNumberType(id).Lazy(&c.MotionsNumberType)
	ds.Meeting_MotionsNumberWithBlank(id).Lazy(&c.MotionsNumberWithBlank)
	ds.Meeting_MotionsOriginMotionToggleDefault(id).Lazy(&c.MotionsOriginMotionToggleDefault)
	ds.Meeting_MotionsPreamble(id).Lazy(&c.MotionsPreamble)
	ds.Meeting_MotionsReasonRequired(id).Lazy(&c.MotionsReasonRequired)
	ds.Meeting_MotionsRecommendationTextMode(id).Lazy(&c.MotionsRecommendationTextMode)
	ds.Meeting_MotionsRecommendationsBy(id).Lazy(&c.MotionsRecommendationsBy)
	ds.Meeting_MotionsShowReferringMotions(id).Lazy(&c.MotionsShowReferringMotions)
	ds.Meeting_MotionsShowSequentialNumber(id).Lazy(&c.MotionsShowSequentialNumber)
	ds.Meeting_MotionsSupportersMinAmount(id).Lazy(&c.MotionsSupportersMinAmount)
	ds.Meeting_Name(id).Lazy(&c.Name)
	ds.Meeting_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.Meeting_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Meeting_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.Meeting_PointOfOrderCategoryIDs(id).Lazy(&c.PointOfOrderCategoryIDs)
	ds.Meeting_PollBallotPaperNumber(id).Lazy(&c.PollBallotPaperNumber)
	ds.Meeting_PollBallotPaperSelection(id).Lazy(&c.PollBallotPaperSelection)
	ds.Meeting_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
	ds.Meeting_PollCandidateListIDs(id).Lazy(&c.PollCandidateListIDs)
	ds.Meeting_PollCountdownID(id).Lazy(&c.PollCountdownID)
	ds.Meeting_PollCoupleCountdown(id).Lazy(&c.PollCoupleCountdown)
	ds.Meeting_PollDefaultBackend(id).Lazy(&c.PollDefaultBackend)
	ds.Meeting_PollDefaultGroupIDs(id).Lazy(&c.PollDefaultGroupIDs)
	ds.Meeting_PollDefaultLiveVotingEnabled(id).Lazy(&c.PollDefaultLiveVotingEnabled)
	ds.Meeting_PollDefaultMethod(id).Lazy(&c.PollDefaultMethod)
	ds.Meeting_PollDefaultOnehundredPercentBase(id).Lazy(&c.PollDefaultOnehundredPercentBase)
	ds.Meeting_PollDefaultType(id).Lazy(&c.PollDefaultType)
	ds.Meeting_PollIDs(id).Lazy(&c.PollIDs)
	ds.Meeting_PollSortPollResultByVotes(id).Lazy(&c.PollSortPollResultByVotes)
	ds.Meeting_PresentUserIDs(id).Lazy(&c.PresentUserIDs)
	ds.Meeting_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Meeting_ProjectorCountdownDefaultTime(id).Lazy(&c.ProjectorCountdownDefaultTime)
	ds.Meeting_ProjectorCountdownIDs(id).Lazy(&c.ProjectorCountdownIDs)
	ds.Meeting_ProjectorCountdownWarningTime(id).Lazy(&c.ProjectorCountdownWarningTime)
	ds.Meeting_ProjectorIDs(id).Lazy(&c.ProjectorIDs)
	ds.Meeting_ProjectorMessageIDs(id).Lazy(&c.ProjectorMessageIDs)
	ds.Meeting_ReferenceProjectorID(id).Lazy(&c.ReferenceProjectorID)
	ds.Meeting_RelevantHistoryEntryIDs(id).Lazy(&c.RelevantHistoryEntryIDs)
	ds.Meeting_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.Meeting_StartTime(id).Lazy(&c.StartTime)
	ds.Meeting_StructureLevelIDs(id).Lazy(&c.StructureLevelIDs)
	ds.Meeting_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
	ds.Meeting_TagIDs(id).Lazy(&c.TagIDs)
	ds.Meeting_TemplateForOrganizationID(id).Lazy(&c.TemplateForOrganizationID)
	ds.Meeting_TimeZone(id).Lazy(&c.TimeZone)
	ds.Meeting_TopicIDs(id).Lazy(&c.TopicIDs)
	ds.Meeting_TopicPollDefaultGroupIDs(id).Lazy(&c.TopicPollDefaultGroupIDs)
	ds.Meeting_UserIDs(id).Lazy(&c.UserIDs)
	ds.Meeting_UsersAllowSelfSetPresent(id).Lazy(&c.UsersAllowSelfSetPresent)
	ds.Meeting_UsersEmailBody(id).Lazy(&c.UsersEmailBody)
	ds.Meeting_UsersEmailReplyto(id).Lazy(&c.UsersEmailReplyto)
	ds.Meeting_UsersEmailSender(id).Lazy(&c.UsersEmailSender)
	ds.Meeting_UsersEmailSubject(id).Lazy(&c.UsersEmailSubject)
	ds.Meeting_UsersEnablePresenceView(id).Lazy(&c.UsersEnablePresenceView)
	ds.Meeting_UsersEnableVoteDelegations(id).Lazy(&c.UsersEnableVoteDelegations)
	ds.Meeting_UsersEnableVoteWeight(id).Lazy(&c.UsersEnableVoteWeight)
	ds.Meeting_UsersForbidDelegatorAsSubmitter(id).Lazy(&c.UsersForbidDelegatorAsSubmitter)
	ds.Meeting_UsersForbidDelegatorAsSupporter(id).Lazy(&c.UsersForbidDelegatorAsSupporter)
	ds.Meeting_UsersForbidDelegatorInListOfSpeakers(id).Lazy(&c.UsersForbidDelegatorInListOfSpeakers)
	ds.Meeting_UsersForbidDelegatorToVote(id).Lazy(&c.UsersForbidDelegatorToVote)
	ds.Meeting_UsersPdfWelcometext(id).Lazy(&c.UsersPdfWelcometext)
	ds.Meeting_UsersPdfWelcometitle(id).Lazy(&c.UsersPdfWelcometitle)
	ds.Meeting_UsersPdfWlanEncryption(id).Lazy(&c.UsersPdfWlanEncryption)
	ds.Meeting_UsersPdfWlanPassword(id).Lazy(&c.UsersPdfWlanPassword)
	ds.Meeting_UsersPdfWlanSsid(id).Lazy(&c.UsersPdfWlanSsid)
	ds.Meeting_VoteIDs(id).Lazy(&c.VoteIDs)
	ds.Meeting_WelcomeText(id).Lazy(&c.WelcomeText)
	ds.Meeting_WelcomeTitle(id).Lazy(&c.WelcomeTitle)
	return &c
}

func (b *meetingBuilder) Preload(rel builderWrapperI) *meetingBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *meetingBuilder) AdminGroup() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AdminGroupID",
			relField: "AdminGroup",
		},
	}
}

func (b *meetingBuilder) AgendaItemList() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AgendaItemIDs",
			relField: "AgendaItemList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) AllProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AllProjectionIDs",
			relField: "AllProjectionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) AnonymousGroup() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AnonymousGroupID",
			relField: "AnonymousGroup",
		},
	}
}

func (b *meetingBuilder) AssignmentCandidateList() *assignmentCandidateBuilder {
	return &assignmentCandidateBuilder{
		builder: builder[assignmentCandidateBuilder, *assignmentCandidateBuilder, AssignmentCandidate]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AssignmentCandidateIDs",
			relField: "AssignmentCandidateList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) AssignmentList() *assignmentBuilder {
	return &assignmentBuilder{
		builder: builder[assignmentBuilder, *assignmentBuilder, Assignment]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AssignmentIDs",
			relField: "AssignmentList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) AssignmentPollDefaultGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AssignmentPollDefaultGroupIDs",
			relField: "AssignmentPollDefaultGroupList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ChatGroupList() *chatGroupBuilder {
	return &chatGroupBuilder{
		builder: builder[chatGroupBuilder, *chatGroupBuilder, ChatGroup]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChatGroupIDs",
			relField: "ChatGroupList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ChatMessageList() *chatMessageBuilder {
	return &chatMessageBuilder{
		builder: builder[chatMessageBuilder, *chatMessageBuilder, ChatMessage]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChatMessageIDs",
			relField: "ChatMessageList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) Committee() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CommitteeID",
			relField: "Committee",
		},
	}
}

func (b *meetingBuilder) DefaultGroup() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultGroupID",
			relField: "DefaultGroup",
		},
	}
}

func (b *meetingBuilder) DefaultMeetingForCommittee() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultMeetingForCommitteeID",
			relField: "DefaultMeetingForCommittee",
		},
	}
}

func (b *meetingBuilder) DefaultProjectorAgendaItemListList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorAgendaItemListIDs",
			relField: "DefaultProjectorAgendaItemListList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorAmendmentList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorAmendmentIDs",
			relField: "DefaultProjectorAmendmentList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorAssignmentList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorAssignmentIDs",
			relField: "DefaultProjectorAssignmentList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorAssignmentPollList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorAssignmentPollIDs",
			relField: "DefaultProjectorAssignmentPollList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorCountdownList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorCountdownIDs",
			relField: "DefaultProjectorCountdownList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorCurrentLosList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorCurrentLosIDs",
			relField: "DefaultProjectorCurrentLosList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorListOfSpeakersList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorListOfSpeakersIDs",
			relField: "DefaultProjectorListOfSpeakersList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorMediafileList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorMediafileIDs",
			relField: "DefaultProjectorMediafileList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorMessageList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorMessageIDs",
			relField: "DefaultProjectorMessageList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorMotionBlockList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorMotionBlockIDs",
			relField: "DefaultProjectorMotionBlockList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorMotionList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorMotionIDs",
			relField: "DefaultProjectorMotionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorMotionPollList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorMotionPollIDs",
			relField: "DefaultProjectorMotionPollList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorPollList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorPollIDs",
			relField: "DefaultProjectorPollList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) DefaultProjectorTopicList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultProjectorTopicIDs",
			relField: "DefaultProjectorTopicList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) FontBold() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontBoldID",
			relField: "FontBold",
		},
	}
}

func (b *meetingBuilder) FontBoldItalic() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontBoldItalicID",
			relField: "FontBoldItalic",
		},
	}
}

func (b *meetingBuilder) FontChyronSpeakerName() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontChyronSpeakerNameID",
			relField: "FontChyronSpeakerName",
		},
	}
}

func (b *meetingBuilder) FontItalic() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontItalicID",
			relField: "FontItalic",
		},
	}
}

func (b *meetingBuilder) FontMonospace() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontMonospaceID",
			relField: "FontMonospace",
		},
	}
}

func (b *meetingBuilder) FontProjectorH1() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontProjectorH1ID",
			relField: "FontProjectorH1",
		},
	}
}

func (b *meetingBuilder) FontProjectorH2() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontProjectorH2ID",
			relField: "FontProjectorH2",
		},
	}
}

func (b *meetingBuilder) FontRegular() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FontRegularID",
			relField: "FontRegular",
		},
	}
}

func (b *meetingBuilder) ForwardedMotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ForwardedMotionIDs",
			relField: "ForwardedMotionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) GroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "GroupIDs",
			relField: "GroupList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) IsActiveInOrganization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "IsActiveInOrganizationID",
			relField: "IsActiveInOrganization",
		},
	}
}

func (b *meetingBuilder) IsArchivedInOrganization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "IsArchivedInOrganizationID",
			relField: "IsArchivedInOrganization",
		},
	}
}

func (b *meetingBuilder) ListOfSpeakersCountdown() *projectorCountdownBuilder {
	return &projectorCountdownBuilder{
		builder: builder[projectorCountdownBuilder, *projectorCountdownBuilder, ProjectorCountdown]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersCountdownID",
			relField: "ListOfSpeakersCountdown",
		},
	}
}

func (b *meetingBuilder) ListOfSpeakersList() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersIDs",
			relField: "ListOfSpeakersList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) LogoPdfBallotPaper() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoPdfBallotPaperID",
			relField: "LogoPdfBallotPaper",
		},
	}
}

func (b *meetingBuilder) LogoPdfFooterL() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoPdfFooterLID",
			relField: "LogoPdfFooterL",
		},
	}
}

func (b *meetingBuilder) LogoPdfFooterR() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoPdfFooterRID",
			relField: "LogoPdfFooterR",
		},
	}
}

func (b *meetingBuilder) LogoPdfHeaderL() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoPdfHeaderLID",
			relField: "LogoPdfHeaderL",
		},
	}
}

func (b *meetingBuilder) LogoPdfHeaderR() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoPdfHeaderRID",
			relField: "LogoPdfHeaderR",
		},
	}
}

func (b *meetingBuilder) LogoProjectorHeader() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoProjectorHeaderID",
			relField: "LogoProjectorHeader",
		},
	}
}

func (b *meetingBuilder) LogoProjectorMain() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoProjectorMainID",
			relField: "LogoProjectorMain",
		},
	}
}

func (b *meetingBuilder) LogoWebHeader() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LogoWebHeaderID",
			relField: "LogoWebHeader",
		},
	}
}

func (b *meetingBuilder) MediafileList() *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MediafileIDs",
			relField: "MediafileList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MeetingMediafileList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingMediafileIDs",
			relField: "MeetingMediafileList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MeetingUserList() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserIDs",
			relField: "MeetingUserList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionBlockList() *motionBlockBuilder {
	return &motionBlockBuilder{
		builder: builder[motionBlockBuilder, *motionBlockBuilder, MotionBlock]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionBlockIDs",
			relField: "MotionBlockList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionCategoryList() *motionCategoryBuilder {
	return &motionCategoryBuilder{
		builder: builder[motionCategoryBuilder, *motionCategoryBuilder, MotionCategory]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionCategoryIDs",
			relField: "MotionCategoryList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionChangeRecommendationList() *motionChangeRecommendationBuilder {
	return &motionChangeRecommendationBuilder{
		builder: builder[motionChangeRecommendationBuilder, *motionChangeRecommendationBuilder, MotionChangeRecommendation]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionChangeRecommendationIDs",
			relField: "MotionChangeRecommendationList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionCommentList() *motionCommentBuilder {
	return &motionCommentBuilder{
		builder: builder[motionCommentBuilder, *motionCommentBuilder, MotionComment]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionCommentIDs",
			relField: "MotionCommentList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionCommentSectionList() *motionCommentSectionBuilder {
	return &motionCommentSectionBuilder{
		builder: builder[motionCommentSectionBuilder, *motionCommentSectionBuilder, MotionCommentSection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionCommentSectionIDs",
			relField: "MotionCommentSectionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionEditorList() *motionEditorBuilder {
	return &motionEditorBuilder{
		builder: builder[motionEditorBuilder, *motionEditorBuilder, MotionEditor]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionEditorIDs",
			relField: "MotionEditorList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionIDs",
			relField: "MotionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionPollDefaultGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionPollDefaultGroupIDs",
			relField: "MotionPollDefaultGroupList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionStateList() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionStateIDs",
			relField: "MotionStateList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionSubmitterList() *motionSubmitterBuilder {
	return &motionSubmitterBuilder{
		builder: builder[motionSubmitterBuilder, *motionSubmitterBuilder, MotionSubmitter]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionSubmitterIDs",
			relField: "MotionSubmitterList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionSupporterList() *motionSupporterBuilder {
	return &motionSupporterBuilder{
		builder: builder[motionSupporterBuilder, *motionSupporterBuilder, MotionSupporter]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionSupporterIDs",
			relField: "MotionSupporterList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionWorkflowList() *motionWorkflowBuilder {
	return &motionWorkflowBuilder{
		builder: builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionWorkflowIDs",
			relField: "MotionWorkflowList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionWorkingGroupSpeakerList() *motionWorkingGroupSpeakerBuilder {
	return &motionWorkingGroupSpeakerBuilder{
		builder: builder[motionWorkingGroupSpeakerBuilder, *motionWorkingGroupSpeakerBuilder, MotionWorkingGroupSpeaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionWorkingGroupSpeakerIDs",
			relField: "MotionWorkingGroupSpeakerList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) MotionsDefaultAmendmentWorkflow() *motionWorkflowBuilder {
	return &motionWorkflowBuilder{
		builder: builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionsDefaultAmendmentWorkflowID",
			relField: "MotionsDefaultAmendmentWorkflow",
		},
	}
}

func (b *meetingBuilder) MotionsDefaultWorkflow() *motionWorkflowBuilder {
	return &motionWorkflowBuilder{
		builder: builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionsDefaultWorkflowID",
			relField: "MotionsDefaultWorkflow",
		},
	}
}

func (b *meetingBuilder) OptionList() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OptionIDs",
			relField: "OptionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) OrganizationTagList() *organizationTagBuilder {
	return &organizationTagBuilder{
		builder: builder[organizationTagBuilder, *organizationTagBuilder, OrganizationTag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationTagIDs",
			relField: "OrganizationTagList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PersonalNoteList() *personalNoteBuilder {
	return &personalNoteBuilder{
		builder: builder[personalNoteBuilder, *personalNoteBuilder, PersonalNote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PersonalNoteIDs",
			relField: "PersonalNoteList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PointOfOrderCategoryList() *pointOfOrderCategoryBuilder {
	return &pointOfOrderCategoryBuilder{
		builder: builder[pointOfOrderCategoryBuilder, *pointOfOrderCategoryBuilder, PointOfOrderCategory]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PointOfOrderCategoryIDs",
			relField: "PointOfOrderCategoryList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PollCandidateList() *pollCandidateBuilder {
	return &pollCandidateBuilder{
		builder: builder[pollCandidateBuilder, *pollCandidateBuilder, PollCandidate]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollCandidateIDs",
			relField: "PollCandidateList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PollCandidateListList() *pollCandidateListBuilder {
	return &pollCandidateListBuilder{
		builder: builder[pollCandidateListBuilder, *pollCandidateListBuilder, PollCandidateList]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollCandidateListIDs",
			relField: "PollCandidateListList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PollCountdown() *projectorCountdownBuilder {
	return &projectorCountdownBuilder{
		builder: builder[projectorCountdownBuilder, *projectorCountdownBuilder, ProjectorCountdown]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollCountdownID",
			relField: "PollCountdown",
		},
	}
}

func (b *meetingBuilder) PollDefaultGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollDefaultGroupIDs",
			relField: "PollDefaultGroupList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PollList() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollIDs",
			relField: "PollList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) PresentUserList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PresentUserIDs",
			relField: "PresentUserList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ProjectorCountdownList() *projectorCountdownBuilder {
	return &projectorCountdownBuilder{
		builder: builder[projectorCountdownBuilder, *projectorCountdownBuilder, ProjectorCountdown]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectorCountdownIDs",
			relField: "ProjectorCountdownList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ProjectorList() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectorIDs",
			relField: "ProjectorList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ProjectorMessageList() *projectorMessageBuilder {
	return &projectorMessageBuilder{
		builder: builder[projectorMessageBuilder, *projectorMessageBuilder, ProjectorMessage]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectorMessageIDs",
			relField: "ProjectorMessageList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) ReferenceProjector() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReferenceProjectorID",
			relField: "ReferenceProjector",
		},
	}
}

func (b *meetingBuilder) RelevantHistoryEntryList() *historyEntryBuilder {
	return &historyEntryBuilder{
		builder: builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "RelevantHistoryEntryIDs",
			relField: "RelevantHistoryEntryList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) SpeakerList() *speakerBuilder {
	return &speakerBuilder{
		builder: builder[speakerBuilder, *speakerBuilder, Speaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SpeakerIDs",
			relField: "SpeakerList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) StructureLevelList() *structureLevelBuilder {
	return &structureLevelBuilder{
		builder: builder[structureLevelBuilder, *structureLevelBuilder, StructureLevel]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelIDs",
			relField: "StructureLevelList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) StructureLevelListOfSpeakersList() *structureLevelListOfSpeakersBuilder {
	return &structureLevelListOfSpeakersBuilder{
		builder: builder[structureLevelListOfSpeakersBuilder, *structureLevelListOfSpeakersBuilder, StructureLevelListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelListOfSpeakersIDs",
			relField: "StructureLevelListOfSpeakersList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) TagList() *tagBuilder {
	return &tagBuilder{
		builder: builder[tagBuilder, *tagBuilder, Tag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TagIDs",
			relField: "TagList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) TemplateForOrganization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TemplateForOrganizationID",
			relField: "TemplateForOrganization",
		},
	}
}

func (b *meetingBuilder) TopicList() *topicBuilder {
	return &topicBuilder{
		builder: builder[topicBuilder, *topicBuilder, Topic]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TopicIDs",
			relField: "TopicList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) TopicPollDefaultGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TopicPollDefaultGroupIDs",
			relField: "TopicPollDefaultGroupList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) UserList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserIDs",
			relField: "UserList",
			many:     true,
		},
	}
}

func (b *meetingBuilder) VoteList() *voteBuilder {
	return &voteBuilder{
		builder: builder[voteBuilder, *voteBuilder, Vote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "VoteIDs",
			relField: "VoteList",
			many:     true,
		},
	}
}

func (r *Fetch) Meeting(ids ...int) *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MeetingMediafile has all fields from meeting_mediafile.
type MeetingMediafile struct {
	AccessGroupIDs                         maybe.Maybe[[]int]
	AttachmentIDs                          maybe.Maybe[[]string]
	ID                                     int
	InheritedAccessGroupIDs                maybe.Maybe[[]int]
	IsPublic                               bool
	ListOfSpeakersID                       maybe.Maybe[int]
	MediafileID                            int
	MeetingID                              int
	ProjectionIDs                          maybe.Maybe[[]int]
	UsedAsFontBoldInMeetingID              maybe.Maybe[int]
	UsedAsFontBoldItalicInMeetingID        maybe.Maybe[int]
	UsedAsFontChyronSpeakerNameInMeetingID maybe.Maybe[int]
	UsedAsFontItalicInMeetingID            maybe.Maybe[int]
	UsedAsFontMonospaceInMeetingID         maybe.Maybe[int]
	UsedAsFontProjectorH1InMeetingID       maybe.Maybe[int]
	UsedAsFontProjectorH2InMeetingID       maybe.Maybe[int]
	UsedAsFontRegularInMeetingID           maybe.Maybe[int]
	UsedAsLogoPdfBallotPaperInMeetingID    maybe.Maybe[int]
	UsedAsLogoPdfFooterLInMeetingID        maybe.Maybe[int]
	UsedAsLogoPdfFooterRInMeetingID        maybe.Maybe[int]
	UsedAsLogoPdfHeaderLInMeetingID        maybe.Maybe[int]
	UsedAsLogoPdfHeaderRInMeetingID        maybe.Maybe[int]
	UsedAsLogoProjectorHeaderInMeetingID   maybe.Maybe[int]
	UsedAsLogoProjectorMainInMeetingID     maybe.Maybe[int]
	UsedAsLogoWebHeaderInMeetingID         maybe.Maybe[int]
	AccessGroupList                        []Group
	InheritedAccessGroupList               []Group
	ListOfSpeakers                         *maybe.Maybe[ListOfSpeakers]
	Mediafile                              *Mediafile
	Meeting                                *Meeting
	ProjectionList                         []Projection
	UsedAsFontBoldInMeeting                *maybe.Maybe[Meeting]
	UsedAsFontBoldItalicInMeeting          *maybe.Maybe[Meeting]
	UsedAsFontChyronSpeakerNameInMeeting   *maybe.Maybe[Meeting]
	UsedAsFontItalicInMeeting              *maybe.Maybe[Meeting]
	UsedAsFontMonospaceInMeeting           *maybe.Maybe[Meeting]
	UsedAsFontProjectorH1InMeeting         *maybe.Maybe[Meeting]
	UsedAsFontProjectorH2InMeeting         *maybe.Maybe[Meeting]
	UsedAsFontRegularInMeeting             *maybe.Maybe[Meeting]
	UsedAsLogoPdfBallotPaperInMeeting      *maybe.Maybe[Meeting]
	UsedAsLogoPdfFooterLInMeeting          *maybe.Maybe[Meeting]
	UsedAsLogoPdfFooterRInMeeting          *maybe.Maybe[Meeting]
	UsedAsLogoPdfHeaderLInMeeting          *maybe.Maybe[Meeting]
	UsedAsLogoPdfHeaderRInMeeting          *maybe.Maybe[Meeting]
	UsedAsLogoProjectorHeaderInMeeting     *maybe.Maybe[Meeting]
	UsedAsLogoProjectorMainInMeeting       *maybe.Maybe[Meeting]
	UsedAsLogoWebHeaderInMeeting           *maybe.Maybe[Meeting]
}

type meetingMediafileBuilder struct {
	builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]
}

func (b *meetingMediafileBuilder) lazy(ds *Fetch, id int) *MeetingMediafile {
	c := MeetingMediafile{}
	ds.MeetingMediafile_AccessGroupIDs(id).Lazy(&c.AccessGroupIDs)
	ds.MeetingMediafile_AttachmentIDs(id).Lazy(&c.AttachmentIDs)
	ds.MeetingMediafile_ID(id).Lazy(&c.ID)
	ds.MeetingMediafile_InheritedAccessGroupIDs(id).Lazy(&c.InheritedAccessGroupIDs)
	ds.MeetingMediafile_IsPublic(id).Lazy(&c.IsPublic)
	ds.MeetingMediafile_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.MeetingMediafile_MediafileID(id).Lazy(&c.MediafileID)
	ds.MeetingMediafile_MeetingID(id).Lazy(&c.MeetingID)
	ds.MeetingMediafile_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.MeetingMediafile_UsedAsFontBoldInMeetingID(id).Lazy(&c.UsedAsFontBoldInMeetingID)
	ds.MeetingMediafile_UsedAsFontBoldItalicInMeetingID(id).Lazy(&c.UsedAsFontBoldItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(id).Lazy(&c.UsedAsFontChyronSpeakerNameInMeetingID)
	ds.MeetingMediafile_UsedAsFontItalicInMeetingID(id).Lazy(&c.UsedAsFontItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontMonospaceInMeetingID(id).Lazy(&c.UsedAsFontMonospaceInMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH1InMeetingID(id).Lazy(&c.UsedAsFontProjectorH1InMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH2InMeetingID(id).Lazy(&c.UsedAsFontProjectorH2InMeetingID)
	ds.MeetingMediafile_UsedAsFontRegularInMeetingID(id).Lazy(&c.UsedAsFontRegularInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(id).Lazy(&c.UsedAsLogoPdfBallotPaperInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(id).Lazy(&c.UsedAsLogoPdfFooterLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(id).Lazy(&c.UsedAsLogoPdfFooterRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(id).Lazy(&c.UsedAsLogoPdfHeaderLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(id).Lazy(&c.UsedAsLogoPdfHeaderRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(id).Lazy(&c.UsedAsLogoProjectorHeaderInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(id).Lazy(&c.UsedAsLogoProjectorMainInMeetingID)
	ds.MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(id).Lazy(&c.UsedAsLogoWebHeaderInMeetingID)
	return &c
}

func (b *meetingMediafileBuilder) Preload(rel builderWrapperI) *meetingMediafileBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *meetingMediafileBuilder) AccessGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AccessGroupIDs",
			relField: "AccessGroupList",
			many:     true,
		},
	}
}

func (b *meetingMediafileBuilder) InheritedAccessGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "InheritedAccessGroupIDs",
			relField: "InheritedAccessGroupList",
			many:     true,
		},
	}
}

func (b *meetingMediafileBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *meetingMediafileBuilder) Mediafile() *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MediafileID",
			relField: "Mediafile",
		},
	}
}

func (b *meetingMediafileBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *meetingMediafileBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontBoldInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontBoldInMeetingID",
			relField: "UsedAsFontBoldInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontBoldItalicInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontBoldItalicInMeetingID",
			relField: "UsedAsFontBoldItalicInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontChyronSpeakerNameInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontChyronSpeakerNameInMeetingID",
			relField: "UsedAsFontChyronSpeakerNameInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontItalicInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontItalicInMeetingID",
			relField: "UsedAsFontItalicInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontMonospaceInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontMonospaceInMeetingID",
			relField: "UsedAsFontMonospaceInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontProjectorH1InMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontProjectorH1InMeetingID",
			relField: "UsedAsFontProjectorH1InMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontProjectorH2InMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontProjectorH2InMeetingID",
			relField: "UsedAsFontProjectorH2InMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsFontRegularInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsFontRegularInMeetingID",
			relField: "UsedAsFontRegularInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoPdfBallotPaperInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoPdfBallotPaperInMeetingID",
			relField: "UsedAsLogoPdfBallotPaperInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoPdfFooterLInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoPdfFooterLInMeetingID",
			relField: "UsedAsLogoPdfFooterLInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoPdfFooterRInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoPdfFooterRInMeetingID",
			relField: "UsedAsLogoPdfFooterRInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoPdfHeaderLInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoPdfHeaderLInMeetingID",
			relField: "UsedAsLogoPdfHeaderLInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoPdfHeaderRInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoPdfHeaderRInMeetingID",
			relField: "UsedAsLogoPdfHeaderRInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoProjectorHeaderInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoProjectorHeaderInMeetingID",
			relField: "UsedAsLogoProjectorHeaderInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoProjectorMainInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoProjectorMainInMeetingID",
			relField: "UsedAsLogoProjectorMainInMeeting",
		},
	}
}

func (b *meetingMediafileBuilder) UsedAsLogoWebHeaderInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsLogoWebHeaderInMeetingID",
			relField: "UsedAsLogoWebHeaderInMeeting",
		},
	}
}

func (r *Fetch) MeetingMediafile(ids ...int) *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MeetingUser has all fields from meeting_user.
type MeetingUser struct {
	AboutMe                       maybe.Maybe[string]
	AssignmentCandidateIDs        maybe.Maybe[[]int]
	ChatMessageIDs                maybe.Maybe[[]int]
	Comment                       maybe.Maybe[string]
	GroupIDs                      []int
	ID                            int
	LockedOut                     maybe.Maybe[bool]
	MeetingID                     int
	MotionEditorIDs               maybe.Maybe[[]int]
	MotionSubmitterIDs            maybe.Maybe[[]int]
	MotionSupporterIDs            maybe.Maybe[[]int]
	MotionWorkingGroupSpeakerIDs  maybe.Maybe[[]int]
	Number                        maybe.Maybe[string]
	PersonalNoteIDs               maybe.Maybe[[]int]
	SpeakerIDs                    maybe.Maybe[[]int]
	StructureLevelIDs             maybe.Maybe[[]int]
	UserID                        int
	VoteDelegatedToID             maybe.Maybe[int]
	VoteDelegationsFromIDs        maybe.Maybe[[]int]
	VoteWeight                    maybe.Maybe[decimal.Decimal]
	AssignmentCandidateList       []AssignmentCandidate
	ChatMessageList               []ChatMessage
	GroupList                     []Group
	Meeting                       *Meeting
	MotionEditorList              []MotionEditor
	MotionSubmitterList           []MotionSubmitter
	MotionSupporterList           []MotionSupporter
	MotionWorkingGroupSpeakerList []MotionWorkingGroupSpeaker
	PersonalNoteList              []PersonalNote
	SpeakerList                   []Speaker
	StructureLevelList            []StructureLevel
	User                          *User
	VoteDelegatedTo               *maybe.Maybe[MeetingUser]
	VoteDelegationsFromList       []MeetingUser
}

type meetingUserBuilder struct {
	builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]
}

func (b *meetingUserBuilder) lazy(ds *Fetch, id int) *MeetingUser {
	c := MeetingUser{}
	ds.MeetingUser_AboutMe(id).Lazy(&c.AboutMe)
	ds.MeetingUser_AssignmentCandidateIDs(id).Lazy(&c.AssignmentCandidateIDs)
	ds.MeetingUser_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.MeetingUser_Comment(id).Lazy(&c.Comment)
	ds.MeetingUser_GroupIDs(id).Lazy(&c.GroupIDs)
	ds.MeetingUser_ID(id).Lazy(&c.ID)
	ds.MeetingUser_LockedOut(id).Lazy(&c.LockedOut)
	ds.MeetingUser_MeetingID(id).Lazy(&c.MeetingID)
	ds.MeetingUser_MotionEditorIDs(id).Lazy(&c.MotionEditorIDs)
	ds.MeetingUser_MotionSubmitterIDs(id).Lazy(&c.MotionSubmitterIDs)
	ds.MeetingUser_MotionSupporterIDs(id).Lazy(&c.MotionSupporterIDs)
	ds.MeetingUser_MotionWorkingGroupSpeakerIDs(id).Lazy(&c.MotionWorkingGroupSpeakerIDs)
	ds.MeetingUser_Number(id).Lazy(&c.Number)
	ds.MeetingUser_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.MeetingUser_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.MeetingUser_StructureLevelIDs(id).Lazy(&c.StructureLevelIDs)
	ds.MeetingUser_UserID(id).Lazy(&c.UserID)
	ds.MeetingUser_VoteDelegatedToID(id).Lazy(&c.VoteDelegatedToID)
	ds.MeetingUser_VoteDelegationsFromIDs(id).Lazy(&c.VoteDelegationsFromIDs)
	ds.MeetingUser_VoteWeight(id).Lazy(&c.VoteWeight)
	return &c
}

func (b *meetingUserBuilder) Preload(rel builderWrapperI) *meetingUserBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *meetingUserBuilder) AssignmentCandidateList() *assignmentCandidateBuilder {
	return &assignmentCandidateBuilder{
		builder: builder[assignmentCandidateBuilder, *assignmentCandidateBuilder, AssignmentCandidate]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AssignmentCandidateIDs",
			relField: "AssignmentCandidateList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) ChatMessageList() *chatMessageBuilder {
	return &chatMessageBuilder{
		builder: builder[chatMessageBuilder, *chatMessageBuilder, ChatMessage]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChatMessageIDs",
			relField: "ChatMessageList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) GroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "GroupIDs",
			relField: "GroupList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *meetingUserBuilder) MotionEditorList() *motionEditorBuilder {
	return &motionEditorBuilder{
		builder: builder[motionEditorBuilder, *motionEditorBuilder, MotionEditor]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionEditorIDs",
			relField: "MotionEditorList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) MotionSubmitterList() *motionSubmitterBuilder {
	return &motionSubmitterBuilder{
		builder: builder[motionSubmitterBuilder, *motionSubmitterBuilder, MotionSubmitter]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionSubmitterIDs",
			relField: "MotionSubmitterList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) MotionSupporterList() *motionSupporterBuilder {
	return &motionSupporterBuilder{
		builder: builder[motionSupporterBuilder, *motionSupporterBuilder, MotionSupporter]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionSupporterIDs",
			relField: "MotionSupporterList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) MotionWorkingGroupSpeakerList() *motionWorkingGroupSpeakerBuilder {
	return &motionWorkingGroupSpeakerBuilder{
		builder: builder[motionWorkingGroupSpeakerBuilder, *motionWorkingGroupSpeakerBuilder, MotionWorkingGroupSpeaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionWorkingGroupSpeakerIDs",
			relField: "MotionWorkingGroupSpeakerList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) PersonalNoteList() *personalNoteBuilder {
	return &personalNoteBuilder{
		builder: builder[personalNoteBuilder, *personalNoteBuilder, PersonalNote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PersonalNoteIDs",
			relField: "PersonalNoteList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) SpeakerList() *speakerBuilder {
	return &speakerBuilder{
		builder: builder[speakerBuilder, *speakerBuilder, Speaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SpeakerIDs",
			relField: "SpeakerList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) StructureLevelList() *structureLevelBuilder {
	return &structureLevelBuilder{
		builder: builder[structureLevelBuilder, *structureLevelBuilder, StructureLevel]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelIDs",
			relField: "StructureLevelList",
			many:     true,
		},
	}
}

func (b *meetingUserBuilder) User() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserID",
			relField: "User",
		},
	}
}

func (b *meetingUserBuilder) VoteDelegatedTo() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "VoteDelegatedToID",
			relField: "VoteDelegatedTo",
		},
	}
}

func (b *meetingUserBuilder) VoteDelegationsFromList() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "VoteDelegationsFromIDs",
			relField: "VoteDelegationsFromList",
			many:     true,
		},
	}
}

func (r *Fetch) MeetingUser(ids ...int) *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Motion has all fields from motion.
type Motion struct {
	AdditionalSubmitter                           maybe.Maybe[string]
	AgendaItemID                                  maybe.Maybe[int]
	AllDerivedMotionIDs                           maybe.Maybe[[]int]
	AllOriginIDs                                  maybe.Maybe[[]int]
	AmendmentIDs                                  maybe.Maybe[[]int]
	AmendmentParagraphs                           maybe.Maybe[json.RawMessage]
	AttachmentMeetingMediafileIDs                 maybe.Maybe[[]int]
	BlockID                                       maybe.Maybe[int]
	CategoryID                                    maybe.Maybe[int]
	CategoryWeight                                maybe.Maybe[int]
	ChangeRecommendationIDs                       maybe.Maybe[[]int]
	CommentIDs                                    maybe.Maybe[[]int]
	Created                                       maybe.Maybe[int]
	DerivedMotionIDs                              maybe.Maybe[[]int]
	DiffVersion                                   maybe.Maybe[string]
	EditorIDs                                     maybe.Maybe[[]int]
	Forwarded                                     maybe.Maybe[int]
	HistoryEntryIDs                               maybe.Maybe[[]int]
	ID                                            int
	IDenticalMotionIDs                            maybe.Maybe[[]int]
	LastModified                                  maybe.Maybe[int]
	LeadMotionID                                  maybe.Maybe[int]
	ListOfSpeakersID                              int
	MarkedForwarded                               maybe.Maybe[bool]
	MeetingID                                     int
	ModifiedFinalVersion                          maybe.Maybe[string]
	Number                                        maybe.Maybe[string]
	NumberValue                                   maybe.Maybe[int]
	OptionIDs                                     maybe.Maybe[[]int]
	OriginID                                      maybe.Maybe[int]
	OriginMeetingID                               maybe.Maybe[int]
	PersonalNoteIDs                               maybe.Maybe[[]int]
	PollIDs                                       maybe.Maybe[[]int]
	ProjectionIDs                                 maybe.Maybe[[]int]
	Reason                                        maybe.Maybe[string]
	RecommendationExtension                       maybe.Maybe[string]
	RecommendationExtensionReferenceIDs           maybe.Maybe[[]string]
	RecommendationID                              maybe.Maybe[int]
	ReferencedInMotionRecommendationExtensionIDs  maybe.Maybe[[]int]
	ReferencedInMotionStateExtensionIDs           maybe.Maybe[[]int]
	SequentialNumber                              int
	SortChildIDs                                  maybe.Maybe[[]int]
	SortParentID                                  maybe.Maybe[int]
	SortWeight                                    maybe.Maybe[int]
	StartLineNumber                               maybe.Maybe[int]
	StateExtension                                maybe.Maybe[string]
	StateExtensionReferenceIDs                    maybe.Maybe[[]string]
	StateID                                       int
	SubmitterIDs                                  maybe.Maybe[[]int]
	SupporterIDs                                  maybe.Maybe[[]int]
	TagIDs                                        maybe.Maybe[[]int]
	Text                                          maybe.Maybe[string]
	TextHash                                      maybe.Maybe[string]
	Title                                         string
	WorkflowTimestamp                             maybe.Maybe[int]
	WorkingGroupSpeakerIDs                        maybe.Maybe[[]int]
	AgendaItem                                    *maybe.Maybe[AgendaItem]
	AllDerivedMotionList                          []Motion
	AllOriginList                                 []Motion
	AmendmentList                                 []Motion
	AttachmentMeetingMediafileList                []MeetingMediafile
	Block                                         *maybe.Maybe[MotionBlock]
	Category                                      *maybe.Maybe[MotionCategory]
	ChangeRecommendationList                      []MotionChangeRecommendation
	CommentList                                   []MotionComment
	DerivedMotionList                             []Motion
	EditorList                                    []MotionEditor
	HistoryEntryList                              []HistoryEntry
	IDenticalMotionList                           []Motion
	LeadMotion                                    *maybe.Maybe[Motion]
	ListOfSpeakers                                *ListOfSpeakers
	Meeting                                       *Meeting
	OptionList                                    []Option
	Origin                                        *maybe.Maybe[Motion]
	OriginMeeting                                 *maybe.Maybe[Meeting]
	PersonalNoteList                              []PersonalNote
	PollList                                      []Poll
	ProjectionList                                []Projection
	Recommendation                                *maybe.Maybe[MotionState]
	ReferencedInMotionRecommendationExtensionList []Motion
	ReferencedInMotionStateExtensionList          []Motion
	SortChildList                                 []Motion
	SortParent                                    *maybe.Maybe[Motion]
	State                                         *MotionState
	SubmitterList                                 []MotionSubmitter
	SupporterList                                 []MotionSupporter
	TagList                                       []Tag
	WorkingGroupSpeakerList                       []MotionWorkingGroupSpeaker
}

type motionBuilder struct {
	builder[motionBuilder, *motionBuilder, Motion]
}

func (b *motionBuilder) lazy(ds *Fetch, id int) *Motion {
	c := Motion{}
	ds.Motion_AdditionalSubmitter(id).Lazy(&c.AdditionalSubmitter)
	ds.Motion_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Motion_AllDerivedMotionIDs(id).Lazy(&c.AllDerivedMotionIDs)
	ds.Motion_AllOriginIDs(id).Lazy(&c.AllOriginIDs)
	ds.Motion_AmendmentIDs(id).Lazy(&c.AmendmentIDs)
	ds.Motion_AmendmentParagraphs(id).Lazy(&c.AmendmentParagraphs)
	ds.Motion_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Motion_BlockID(id).Lazy(&c.BlockID)
	ds.Motion_CategoryID(id).Lazy(&c.CategoryID)
	ds.Motion_CategoryWeight(id).Lazy(&c.CategoryWeight)
	ds.Motion_ChangeRecommendationIDs(id).Lazy(&c.ChangeRecommendationIDs)
	ds.Motion_CommentIDs(id).Lazy(&c.CommentIDs)
	ds.Motion_Created(id).Lazy(&c.Created)
	ds.Motion_DerivedMotionIDs(id).Lazy(&c.DerivedMotionIDs)
	ds.Motion_DiffVersion(id).Lazy(&c.DiffVersion)
	ds.Motion_EditorIDs(id).Lazy(&c.EditorIDs)
	ds.Motion_Forwarded(id).Lazy(&c.Forwarded)
	ds.Motion_HistoryEntryIDs(id).Lazy(&c.HistoryEntryIDs)
	ds.Motion_ID(id).Lazy(&c.ID)
	ds.Motion_IDenticalMotionIDs(id).Lazy(&c.IDenticalMotionIDs)
	ds.Motion_LastModified(id).Lazy(&c.LastModified)
	ds.Motion_LeadMotionID(id).Lazy(&c.LeadMotionID)
	ds.Motion_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Motion_MarkedForwarded(id).Lazy(&c.MarkedForwarded)
	ds.Motion_MeetingID(id).Lazy(&c.MeetingID)
	ds.Motion_ModifiedFinalVersion(id).Lazy(&c.ModifiedFinalVersion)
	ds.Motion_Number(id).Lazy(&c.Number)
	ds.Motion_NumberValue(id).Lazy(&c.NumberValue)
	ds.Motion_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.Motion_OriginID(id).Lazy(&c.OriginID)
	ds.Motion_OriginMeetingID(id).Lazy(&c.OriginMeetingID)
	ds.Motion_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.Motion_PollIDs(id).Lazy(&c.PollIDs)
	ds.Motion_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Motion_Reason(id).Lazy(&c.Reason)
	ds.Motion_RecommendationExtension(id).Lazy(&c.RecommendationExtension)
	ds.Motion_RecommendationExtensionReferenceIDs(id).Lazy(&c.RecommendationExtensionReferenceIDs)
	ds.Motion_RecommendationID(id).Lazy(&c.RecommendationID)
	ds.Motion_ReferencedInMotionRecommendationExtensionIDs(id).Lazy(&c.ReferencedInMotionRecommendationExtensionIDs)
	ds.Motion_ReferencedInMotionStateExtensionIDs(id).Lazy(&c.ReferencedInMotionStateExtensionIDs)
	ds.Motion_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Motion_SortChildIDs(id).Lazy(&c.SortChildIDs)
	ds.Motion_SortParentID(id).Lazy(&c.SortParentID)
	ds.Motion_SortWeight(id).Lazy(&c.SortWeight)
	ds.Motion_StartLineNumber(id).Lazy(&c.StartLineNumber)
	ds.Motion_StateExtension(id).Lazy(&c.StateExtension)
	ds.Motion_StateExtensionReferenceIDs(id).Lazy(&c.StateExtensionReferenceIDs)
	ds.Motion_StateID(id).Lazy(&c.StateID)
	ds.Motion_SubmitterIDs(id).Lazy(&c.SubmitterIDs)
	ds.Motion_SupporterIDs(id).Lazy(&c.SupporterIDs)
	ds.Motion_TagIDs(id).Lazy(&c.TagIDs)
	ds.Motion_Text(id).Lazy(&c.Text)
	ds.Motion_TextHash(id).Lazy(&c.TextHash)
	ds.Motion_Title(id).Lazy(&c.Title)
	ds.Motion_WorkflowTimestamp(id).Lazy(&c.WorkflowTimestamp)
	ds.Motion_WorkingGroupSpeakerIDs(id).Lazy(&c.WorkingGroupSpeakerIDs)
	return &c
}

func (b *motionBuilder) Preload(rel builderWrapperI) *motionBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionBuilder) AgendaItem() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AgendaItemID",
			relField: "AgendaItem",
		},
	}
}

func (b *motionBuilder) AllDerivedMotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AllDerivedMotionIDs",
			relField: "AllDerivedMotionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) AllOriginList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AllOriginIDs",
			relField: "AllOriginList",
			many:     true,
		},
	}
}

func (b *motionBuilder) AmendmentList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AmendmentIDs",
			relField: "AmendmentList",
			many:     true,
		},
	}
}

func (b *motionBuilder) AttachmentMeetingMediafileList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AttachmentMeetingMediafileIDs",
			relField: "AttachmentMeetingMediafileList",
			many:     true,
		},
	}
}

func (b *motionBuilder) Block() *motionBlockBuilder {
	return &motionBlockBuilder{
		builder: builder[motionBlockBuilder, *motionBlockBuilder, MotionBlock]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "BlockID",
			relField: "Block",
		},
	}
}

func (b *motionBuilder) Category() *motionCategoryBuilder {
	return &motionCategoryBuilder{
		builder: builder[motionCategoryBuilder, *motionCategoryBuilder, MotionCategory]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CategoryID",
			relField: "Category",
		},
	}
}

func (b *motionBuilder) ChangeRecommendationList() *motionChangeRecommendationBuilder {
	return &motionChangeRecommendationBuilder{
		builder: builder[motionChangeRecommendationBuilder, *motionChangeRecommendationBuilder, MotionChangeRecommendation]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChangeRecommendationIDs",
			relField: "ChangeRecommendationList",
			many:     true,
		},
	}
}

func (b *motionBuilder) CommentList() *motionCommentBuilder {
	return &motionCommentBuilder{
		builder: builder[motionCommentBuilder, *motionCommentBuilder, MotionComment]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CommentIDs",
			relField: "CommentList",
			many:     true,
		},
	}
}

func (b *motionBuilder) DerivedMotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DerivedMotionIDs",
			relField: "DerivedMotionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) EditorList() *motionEditorBuilder {
	return &motionEditorBuilder{
		builder: builder[motionEditorBuilder, *motionEditorBuilder, MotionEditor]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "EditorIDs",
			relField: "EditorList",
			many:     true,
		},
	}
}

func (b *motionBuilder) HistoryEntryList() *historyEntryBuilder {
	return &historyEntryBuilder{
		builder: builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HistoryEntryIDs",
			relField: "HistoryEntryList",
			many:     true,
		},
	}
}

func (b *motionBuilder) IDenticalMotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "IDenticalMotionIDs",
			relField: "IDenticalMotionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) LeadMotion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "LeadMotionID",
			relField: "LeadMotion",
		},
	}
}

func (b *motionBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *motionBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionBuilder) OptionList() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OptionIDs",
			relField: "OptionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) Origin() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OriginID",
			relField: "Origin",
		},
	}
}

func (b *motionBuilder) OriginMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OriginMeetingID",
			relField: "OriginMeeting",
		},
	}
}

func (b *motionBuilder) PersonalNoteList() *personalNoteBuilder {
	return &personalNoteBuilder{
		builder: builder[personalNoteBuilder, *personalNoteBuilder, PersonalNote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PersonalNoteIDs",
			relField: "PersonalNoteList",
			many:     true,
		},
	}
}

func (b *motionBuilder) PollList() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollIDs",
			relField: "PollList",
			many:     true,
		},
	}
}

func (b *motionBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) Recommendation() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "RecommendationID",
			relField: "Recommendation",
		},
	}
}

func (b *motionBuilder) ReferencedInMotionRecommendationExtensionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReferencedInMotionRecommendationExtensionIDs",
			relField: "ReferencedInMotionRecommendationExtensionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) ReferencedInMotionStateExtensionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReferencedInMotionStateExtensionIDs",
			relField: "ReferencedInMotionStateExtensionList",
			many:     true,
		},
	}
}

func (b *motionBuilder) SortChildList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SortChildIDs",
			relField: "SortChildList",
			many:     true,
		},
	}
}

func (b *motionBuilder) SortParent() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SortParentID",
			relField: "SortParent",
		},
	}
}

func (b *motionBuilder) State() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StateID",
			relField: "State",
		},
	}
}

func (b *motionBuilder) SubmitterList() *motionSubmitterBuilder {
	return &motionSubmitterBuilder{
		builder: builder[motionSubmitterBuilder, *motionSubmitterBuilder, MotionSubmitter]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SubmitterIDs",
			relField: "SubmitterList",
			many:     true,
		},
	}
}

func (b *motionBuilder) SupporterList() *motionSupporterBuilder {
	return &motionSupporterBuilder{
		builder: builder[motionSupporterBuilder, *motionSupporterBuilder, MotionSupporter]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SupporterIDs",
			relField: "SupporterList",
			many:     true,
		},
	}
}

func (b *motionBuilder) TagList() *tagBuilder {
	return &tagBuilder{
		builder: builder[tagBuilder, *tagBuilder, Tag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TagIDs",
			relField: "TagList",
			many:     true,
		},
	}
}

func (b *motionBuilder) WorkingGroupSpeakerList() *motionWorkingGroupSpeakerBuilder {
	return &motionWorkingGroupSpeakerBuilder{
		builder: builder[motionWorkingGroupSpeakerBuilder, *motionWorkingGroupSpeakerBuilder, MotionWorkingGroupSpeaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "WorkingGroupSpeakerIDs",
			relField: "WorkingGroupSpeakerList",
			many:     true,
		},
	}
}

func (r *Fetch) Motion(ids ...int) *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionBlock has all fields from motion_block.
type MotionBlock struct {
	AgendaItemID     maybe.Maybe[int]
	ID               int
	Internal         maybe.Maybe[bool]
	ListOfSpeakersID int
	MeetingID        int
	MotionIDs        maybe.Maybe[[]int]
	ProjectionIDs    maybe.Maybe[[]int]
	SequentialNumber int
	Title            string
	AgendaItem       *maybe.Maybe[AgendaItem]
	ListOfSpeakers   *ListOfSpeakers
	Meeting          *Meeting
	MotionList       []Motion
	ProjectionList   []Projection
}

type motionBlockBuilder struct {
	builder[motionBlockBuilder, *motionBlockBuilder, MotionBlock]
}

func (b *motionBlockBuilder) lazy(ds *Fetch, id int) *MotionBlock {
	c := MotionBlock{}
	ds.MotionBlock_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.MotionBlock_ID(id).Lazy(&c.ID)
	ds.MotionBlock_Internal(id).Lazy(&c.Internal)
	ds.MotionBlock_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.MotionBlock_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionBlock_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionBlock_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.MotionBlock_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionBlock_Title(id).Lazy(&c.Title)
	return &c
}

func (b *motionBlockBuilder) Preload(rel builderWrapperI) *motionBlockBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionBlockBuilder) AgendaItem() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AgendaItemID",
			relField: "AgendaItem",
		},
	}
}

func (b *motionBlockBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *motionBlockBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionBlockBuilder) MotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionIDs",
			relField: "MotionList",
			many:     true,
		},
	}
}

func (b *motionBlockBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (r *Fetch) MotionBlock(ids ...int) *motionBlockBuilder {
	return &motionBlockBuilder{
		builder: builder[motionBlockBuilder, *motionBlockBuilder, MotionBlock]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionCategory has all fields from motion_category.
type MotionCategory struct {
	ChildIDs         maybe.Maybe[[]int]
	ID               int
	Level            maybe.Maybe[int]
	MeetingID        int
	MotionIDs        maybe.Maybe[[]int]
	Name             string
	ParentID         maybe.Maybe[int]
	Prefix           maybe.Maybe[string]
	SequentialNumber int
	Weight           maybe.Maybe[int]
	ChildList        []MotionCategory
	Meeting          *Meeting
	MotionList       []Motion
	Parent           *maybe.Maybe[MotionCategory]
}

type motionCategoryBuilder struct {
	builder[motionCategoryBuilder, *motionCategoryBuilder, MotionCategory]
}

func (b *motionCategoryBuilder) lazy(ds *Fetch, id int) *MotionCategory {
	c := MotionCategory{}
	ds.MotionCategory_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.MotionCategory_ID(id).Lazy(&c.ID)
	ds.MotionCategory_Level(id).Lazy(&c.Level)
	ds.MotionCategory_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionCategory_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionCategory_Name(id).Lazy(&c.Name)
	ds.MotionCategory_ParentID(id).Lazy(&c.ParentID)
	ds.MotionCategory_Prefix(id).Lazy(&c.Prefix)
	ds.MotionCategory_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionCategory_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *motionCategoryBuilder) Preload(rel builderWrapperI) *motionCategoryBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionCategoryBuilder) ChildList() *motionCategoryBuilder {
	return &motionCategoryBuilder{
		builder: builder[motionCategoryBuilder, *motionCategoryBuilder, MotionCategory]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ChildIDs",
			relField: "ChildList",
			many:     true,
		},
	}
}

func (b *motionCategoryBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionCategoryBuilder) MotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionIDs",
			relField: "MotionList",
			many:     true,
		},
	}
}

func (b *motionCategoryBuilder) Parent() *motionCategoryBuilder {
	return &motionCategoryBuilder{
		builder: builder[motionCategoryBuilder, *motionCategoryBuilder, MotionCategory]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ParentID",
			relField: "Parent",
		},
	}
}

func (r *Fetch) MotionCategory(ids ...int) *motionCategoryBuilder {
	return &motionCategoryBuilder{
		builder: builder[motionCategoryBuilder, *motionCategoryBuilder, MotionCategory]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionChangeRecommendation has all fields from motion_change_recommendation.
type MotionChangeRecommendation struct {
	CreationTime     maybe.Maybe[int]
	ID               int
	Internal         maybe.Maybe[bool]
	LineFrom         int
	LineTo           int
	MeetingID        int
	MotionID         int
	OtherDescription maybe.Maybe[string]
	Rejected         maybe.Maybe[bool]
	Text             maybe.Maybe[string]
	Type             maybe.Maybe[dstypes.MotionChangeRecommendation_Type]
	Meeting          *Meeting
	Motion           *Motion
}

type motionChangeRecommendationBuilder struct {
	builder[motionChangeRecommendationBuilder, *motionChangeRecommendationBuilder, MotionChangeRecommendation]
}

func (b *motionChangeRecommendationBuilder) lazy(ds *Fetch, id int) *MotionChangeRecommendation {
	c := MotionChangeRecommendation{}
	ds.MotionChangeRecommendation_CreationTime(id).Lazy(&c.CreationTime)
	ds.MotionChangeRecommendation_ID(id).Lazy(&c.ID)
	ds.MotionChangeRecommendation_Internal(id).Lazy(&c.Internal)
	ds.MotionChangeRecommendation_LineFrom(id).Lazy(&c.LineFrom)
	ds.MotionChangeRecommendation_LineTo(id).Lazy(&c.LineTo)
	ds.MotionChangeRecommendation_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionChangeRecommendation_MotionID(id).Lazy(&c.MotionID)
	ds.MotionChangeRecommendation_OtherDescription(id).Lazy(&c.OtherDescription)
	ds.MotionChangeRecommendation_Rejected(id).Lazy(&c.Rejected)
	ds.MotionChangeRecommendation_Text(id).Lazy(&c.Text)
	ds.MotionChangeRecommendation_Type(id).Lazy(&c.Type)
	return &c
}

func (b *motionChangeRecommendationBuilder) Preload(rel builderWrapperI) *motionChangeRecommendationBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionChangeRecommendationBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionChangeRecommendationBuilder) Motion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionID",
			relField: "Motion",
		},
	}
}

func (r *Fetch) MotionChangeRecommendation(ids ...int) *motionChangeRecommendationBuilder {
	return &motionChangeRecommendationBuilder{
		builder: builder[motionChangeRecommendationBuilder, *motionChangeRecommendationBuilder, MotionChangeRecommendation]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionComment has all fields from motion_comment.
type MotionComment struct {
	Comment   string
	ID        int
	MeetingID int
	MotionID  int
	SectionID int
	Meeting   *Meeting
	Motion    *Motion
	Section   *MotionCommentSection
}

type motionCommentBuilder struct {
	builder[motionCommentBuilder, *motionCommentBuilder, MotionComment]
}

func (b *motionCommentBuilder) lazy(ds *Fetch, id int) *MotionComment {
	c := MotionComment{}
	ds.MotionComment_Comment(id).Lazy(&c.Comment)
	ds.MotionComment_ID(id).Lazy(&c.ID)
	ds.MotionComment_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionComment_MotionID(id).Lazy(&c.MotionID)
	ds.MotionComment_SectionID(id).Lazy(&c.SectionID)
	return &c
}

func (b *motionCommentBuilder) Preload(rel builderWrapperI) *motionCommentBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionCommentBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionCommentBuilder) Motion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionID",
			relField: "Motion",
		},
	}
}

func (b *motionCommentBuilder) Section() *motionCommentSectionBuilder {
	return &motionCommentSectionBuilder{
		builder: builder[motionCommentSectionBuilder, *motionCommentSectionBuilder, MotionCommentSection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SectionID",
			relField: "Section",
		},
	}
}

func (r *Fetch) MotionComment(ids ...int) *motionCommentBuilder {
	return &motionCommentBuilder{
		builder: builder[motionCommentBuilder, *motionCommentBuilder, MotionComment]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionCommentSection has all fields from motion_comment_section.
type MotionCommentSection struct {
	CommentIDs        maybe.Maybe[[]int]
	ID                int
	MeetingID         int
	Name              string
	ReadGroupIDs      maybe.Maybe[[]int]
	SequentialNumber  int
	SubmitterCanWrite maybe.Maybe[bool]
	Weight            maybe.Maybe[int]
	WriteGroupIDs     maybe.Maybe[[]int]
	CommentList       []MotionComment
	Meeting           *Meeting
	ReadGroupList     []Group
	WriteGroupList    []Group
}

type motionCommentSectionBuilder struct {
	builder[motionCommentSectionBuilder, *motionCommentSectionBuilder, MotionCommentSection]
}

func (b *motionCommentSectionBuilder) lazy(ds *Fetch, id int) *MotionCommentSection {
	c := MotionCommentSection{}
	ds.MotionCommentSection_CommentIDs(id).Lazy(&c.CommentIDs)
	ds.MotionCommentSection_ID(id).Lazy(&c.ID)
	ds.MotionCommentSection_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionCommentSection_Name(id).Lazy(&c.Name)
	ds.MotionCommentSection_ReadGroupIDs(id).Lazy(&c.ReadGroupIDs)
	ds.MotionCommentSection_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionCommentSection_SubmitterCanWrite(id).Lazy(&c.SubmitterCanWrite)
	ds.MotionCommentSection_Weight(id).Lazy(&c.Weight)
	ds.MotionCommentSection_WriteGroupIDs(id).Lazy(&c.WriteGroupIDs)
	return &c
}

func (b *motionCommentSectionBuilder) Preload(rel builderWrapperI) *motionCommentSectionBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionCommentSectionBuilder) CommentList() *motionCommentBuilder {
	return &motionCommentBuilder{
		builder: builder[motionCommentBuilder, *motionCommentBuilder, MotionComment]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CommentIDs",
			relField: "CommentList",
			many:     true,
		},
	}
}

func (b *motionCommentSectionBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionCommentSectionBuilder) ReadGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ReadGroupIDs",
			relField: "ReadGroupList",
			many:     true,
		},
	}
}

func (b *motionCommentSectionBuilder) WriteGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "WriteGroupIDs",
			relField: "WriteGroupList",
			many:     true,
		},
	}
}

func (r *Fetch) MotionCommentSection(ids ...int) *motionCommentSectionBuilder {
	return &motionCommentSectionBuilder{
		builder: builder[motionCommentSectionBuilder, *motionCommentSectionBuilder, MotionCommentSection]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionEditor has all fields from motion_editor.
type MotionEditor struct {
	ID            int
	MeetingID     int
	MeetingUserID maybe.Maybe[int]
	MotionID      int
	Weight        maybe.Maybe[int]
	Meeting       *Meeting
	MeetingUser   *maybe.Maybe[MeetingUser]
	Motion        *Motion
}

type motionEditorBuilder struct {
	builder[motionEditorBuilder, *motionEditorBuilder, MotionEditor]
}

func (b *motionEditorBuilder) lazy(ds *Fetch, id int) *MotionEditor {
	c := MotionEditor{}
	ds.MotionEditor_ID(id).Lazy(&c.ID)
	ds.MotionEditor_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionEditor_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionEditor_MotionID(id).Lazy(&c.MotionID)
	ds.MotionEditor_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *motionEditorBuilder) Preload(rel builderWrapperI) *motionEditorBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionEditorBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionEditorBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (b *motionEditorBuilder) Motion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionID",
			relField: "Motion",
		},
	}
}

func (r *Fetch) MotionEditor(ids ...int) *motionEditorBuilder {
	return &motionEditorBuilder{
		builder: builder[motionEditorBuilder, *motionEditorBuilder, MotionEditor]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionState has all fields from motion_state.
type MotionState struct {
	AllowAmendmentForwarding         maybe.Maybe[bool]
	AllowCreatePoll                  maybe.Maybe[bool]
	AllowMotionForwarding            maybe.Maybe[bool]
	AllowSubmitterEdit               maybe.Maybe[bool]
	AllowSupport                     maybe.Maybe[bool]
	CssClass                         dstypes.MotionState_CssClass
	FirstStateOfWorkflowID           maybe.Maybe[int]
	ID                               int
	IsInternal                       maybe.Maybe[bool]
	MeetingID                        int
	MergeAmendmentIntoFinal          maybe.Maybe[dstypes.MotionState_MergeAmendmentIntoFinal]
	MotionIDs                        maybe.Maybe[[]int]
	MotionRecommendationIDs          maybe.Maybe[[]int]
	Name                             string
	NextStateIDs                     maybe.Maybe[[]int]
	PreviousStateIDs                 maybe.Maybe[[]int]
	RecommendationLabel              maybe.Maybe[string]
	Restrictions                     maybe.Maybe[[]string]
	SetNumber                        maybe.Maybe[bool]
	SetWorkflowTimestamp             maybe.Maybe[bool]
	ShowRecommendationExtensionField maybe.Maybe[bool]
	ShowStateExtensionField          maybe.Maybe[bool]
	StateButtonLabel                 maybe.Maybe[string]
	SubmitterWithdrawBackIDs         maybe.Maybe[[]int]
	SubmitterWithdrawStateID         maybe.Maybe[int]
	Weight                           int
	WorkflowID                       int
	FirstStateOfWorkflow             *maybe.Maybe[MotionWorkflow]
	Meeting                          *Meeting
	MotionList                       []Motion
	MotionRecommendationList         []Motion
	NextStateList                    []MotionState
	PreviousStateList                []MotionState
	SubmitterWithdrawBackList        []MotionState
	SubmitterWithdrawState           *maybe.Maybe[MotionState]
	Workflow                         *MotionWorkflow
}

type motionStateBuilder struct {
	builder[motionStateBuilder, *motionStateBuilder, MotionState]
}

func (b *motionStateBuilder) lazy(ds *Fetch, id int) *MotionState {
	c := MotionState{}
	ds.MotionState_AllowAmendmentForwarding(id).Lazy(&c.AllowAmendmentForwarding)
	ds.MotionState_AllowCreatePoll(id).Lazy(&c.AllowCreatePoll)
	ds.MotionState_AllowMotionForwarding(id).Lazy(&c.AllowMotionForwarding)
	ds.MotionState_AllowSubmitterEdit(id).Lazy(&c.AllowSubmitterEdit)
	ds.MotionState_AllowSupport(id).Lazy(&c.AllowSupport)
	ds.MotionState_CssClass(id).Lazy(&c.CssClass)
	ds.MotionState_FirstStateOfWorkflowID(id).Lazy(&c.FirstStateOfWorkflowID)
	ds.MotionState_ID(id).Lazy(&c.ID)
	ds.MotionState_IsInternal(id).Lazy(&c.IsInternal)
	ds.MotionState_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionState_MergeAmendmentIntoFinal(id).Lazy(&c.MergeAmendmentIntoFinal)
	ds.MotionState_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionState_MotionRecommendationIDs(id).Lazy(&c.MotionRecommendationIDs)
	ds.MotionState_Name(id).Lazy(&c.Name)
	ds.MotionState_NextStateIDs(id).Lazy(&c.NextStateIDs)
	ds.MotionState_PreviousStateIDs(id).Lazy(&c.PreviousStateIDs)
	ds.MotionState_RecommendationLabel(id).Lazy(&c.RecommendationLabel)
	ds.MotionState_Restrictions(id).Lazy(&c.Restrictions)
	ds.MotionState_SetNumber(id).Lazy(&c.SetNumber)
	ds.MotionState_SetWorkflowTimestamp(id).Lazy(&c.SetWorkflowTimestamp)
	ds.MotionState_ShowRecommendationExtensionField(id).Lazy(&c.ShowRecommendationExtensionField)
	ds.MotionState_ShowStateExtensionField(id).Lazy(&c.ShowStateExtensionField)
	ds.MotionState_StateButtonLabel(id).Lazy(&c.StateButtonLabel)
	ds.MotionState_SubmitterWithdrawBackIDs(id).Lazy(&c.SubmitterWithdrawBackIDs)
	ds.MotionState_SubmitterWithdrawStateID(id).Lazy(&c.SubmitterWithdrawStateID)
	ds.MotionState_Weight(id).Lazy(&c.Weight)
	ds.MotionState_WorkflowID(id).Lazy(&c.WorkflowID)
	return &c
}

func (b *motionStateBuilder) Preload(rel builderWrapperI) *motionStateBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionStateBuilder) FirstStateOfWorkflow() *motionWorkflowBuilder {
	return &motionWorkflowBuilder{
		builder: builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FirstStateOfWorkflowID",
			relField: "FirstStateOfWorkflow",
		},
	}
}

func (b *motionStateBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionStateBuilder) MotionList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionIDs",
			relField: "MotionList",
			many:     true,
		},
	}
}

func (b *motionStateBuilder) MotionRecommendationList() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionRecommendationIDs",
			relField: "MotionRecommendationList",
			many:     true,
		},
	}
}

func (b *motionStateBuilder) NextStateList() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "NextStateIDs",
			relField: "NextStateList",
			many:     true,
		},
	}
}

func (b *motionStateBuilder) PreviousStateList() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PreviousStateIDs",
			relField: "PreviousStateList",
			many:     true,
		},
	}
}

func (b *motionStateBuilder) SubmitterWithdrawBackList() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SubmitterWithdrawBackIDs",
			relField: "SubmitterWithdrawBackList",
			many:     true,
		},
	}
}

func (b *motionStateBuilder) SubmitterWithdrawState() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SubmitterWithdrawStateID",
			relField: "SubmitterWithdrawState",
		},
	}
}

func (b *motionStateBuilder) Workflow() *motionWorkflowBuilder {
	return &motionWorkflowBuilder{
		builder: builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "WorkflowID",
			relField: "Workflow",
		},
	}
}

func (r *Fetch) MotionState(ids ...int) *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionSubmitter has all fields from motion_submitter.
type MotionSubmitter struct {
	ID            int
	MeetingID     int
	MeetingUserID maybe.Maybe[int]
	MotionID      int
	Weight        maybe.Maybe[int]
	Meeting       *Meeting
	MeetingUser   *maybe.Maybe[MeetingUser]
	Motion        *Motion
}

type motionSubmitterBuilder struct {
	builder[motionSubmitterBuilder, *motionSubmitterBuilder, MotionSubmitter]
}

func (b *motionSubmitterBuilder) lazy(ds *Fetch, id int) *MotionSubmitter {
	c := MotionSubmitter{}
	ds.MotionSubmitter_ID(id).Lazy(&c.ID)
	ds.MotionSubmitter_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionSubmitter_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionSubmitter_MotionID(id).Lazy(&c.MotionID)
	ds.MotionSubmitter_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *motionSubmitterBuilder) Preload(rel builderWrapperI) *motionSubmitterBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionSubmitterBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionSubmitterBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (b *motionSubmitterBuilder) Motion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionID",
			relField: "Motion",
		},
	}
}

func (r *Fetch) MotionSubmitter(ids ...int) *motionSubmitterBuilder {
	return &motionSubmitterBuilder{
		builder: builder[motionSubmitterBuilder, *motionSubmitterBuilder, MotionSubmitter]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionSupporter has all fields from motion_supporter.
type MotionSupporter struct {
	ID            int
	MeetingID     int
	MeetingUserID maybe.Maybe[int]
	MotionID      int
	Meeting       *Meeting
	MeetingUser   *maybe.Maybe[MeetingUser]
	Motion        *Motion
}

type motionSupporterBuilder struct {
	builder[motionSupporterBuilder, *motionSupporterBuilder, MotionSupporter]
}

func (b *motionSupporterBuilder) lazy(ds *Fetch, id int) *MotionSupporter {
	c := MotionSupporter{}
	ds.MotionSupporter_ID(id).Lazy(&c.ID)
	ds.MotionSupporter_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionSupporter_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionSupporter_MotionID(id).Lazy(&c.MotionID)
	return &c
}

func (b *motionSupporterBuilder) Preload(rel builderWrapperI) *motionSupporterBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionSupporterBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionSupporterBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (b *motionSupporterBuilder) Motion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionID",
			relField: "Motion",
		},
	}
}

func (r *Fetch) MotionSupporter(ids ...int) *motionSupporterBuilder {
	return &motionSupporterBuilder{
		builder: builder[motionSupporterBuilder, *motionSupporterBuilder, MotionSupporter]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionWorkflow has all fields from motion_workflow.
type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID maybe.Maybe[int]
	DefaultWorkflowMeetingID          maybe.Maybe[int]
	FirstStateID                      int
	ID                                int
	MeetingID                         int
	Name                              string
	SequentialNumber                  int
	StateIDs                          maybe.Maybe[[]int]
	DefaultAmendmentWorkflowMeeting   *maybe.Maybe[Meeting]
	DefaultWorkflowMeeting            *maybe.Maybe[Meeting]
	FirstState                        *MotionState
	Meeting                           *Meeting
	StateList                         []MotionState
}

type motionWorkflowBuilder struct {
	builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]
}

func (b *motionWorkflowBuilder) lazy(ds *Fetch, id int) *MotionWorkflow {
	c := MotionWorkflow{}
	ds.MotionWorkflow_DefaultAmendmentWorkflowMeetingID(id).Lazy(&c.DefaultAmendmentWorkflowMeetingID)
	ds.MotionWorkflow_DefaultWorkflowMeetingID(id).Lazy(&c.DefaultWorkflowMeetingID)
	ds.MotionWorkflow_FirstStateID(id).Lazy(&c.FirstStateID)
	ds.MotionWorkflow_ID(id).Lazy(&c.ID)
	ds.MotionWorkflow_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionWorkflow_Name(id).Lazy(&c.Name)
	ds.MotionWorkflow_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionWorkflow_StateIDs(id).Lazy(&c.StateIDs)
	return &c
}

func (b *motionWorkflowBuilder) Preload(rel builderWrapperI) *motionWorkflowBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionWorkflowBuilder) DefaultAmendmentWorkflowMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultAmendmentWorkflowMeetingID",
			relField: "DefaultAmendmentWorkflowMeeting",
		},
	}
}

func (b *motionWorkflowBuilder) DefaultWorkflowMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DefaultWorkflowMeetingID",
			relField: "DefaultWorkflowMeeting",
		},
	}
}

func (b *motionWorkflowBuilder) FirstState() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "FirstStateID",
			relField: "FirstState",
		},
	}
}

func (b *motionWorkflowBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionWorkflowBuilder) StateList() *motionStateBuilder {
	return &motionStateBuilder{
		builder: builder[motionStateBuilder, *motionStateBuilder, MotionState]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StateIDs",
			relField: "StateList",
			many:     true,
		},
	}
}

func (r *Fetch) MotionWorkflow(ids ...int) *motionWorkflowBuilder {
	return &motionWorkflowBuilder{
		builder: builder[motionWorkflowBuilder, *motionWorkflowBuilder, MotionWorkflow]{
			ids:   ids,
			fetch: r,
		},
	}
}

// MotionWorkingGroupSpeaker has all fields from motion_working_group_speaker.
type MotionWorkingGroupSpeaker struct {
	ID            int
	MeetingID     int
	MeetingUserID maybe.Maybe[int]
	MotionID      int
	Weight        maybe.Maybe[int]
	Meeting       *Meeting
	MeetingUser   *maybe.Maybe[MeetingUser]
	Motion        *Motion
}

type motionWorkingGroupSpeakerBuilder struct {
	builder[motionWorkingGroupSpeakerBuilder, *motionWorkingGroupSpeakerBuilder, MotionWorkingGroupSpeaker]
}

func (b *motionWorkingGroupSpeakerBuilder) lazy(ds *Fetch, id int) *MotionWorkingGroupSpeaker {
	c := MotionWorkingGroupSpeaker{}
	ds.MotionWorkingGroupSpeaker_ID(id).Lazy(&c.ID)
	ds.MotionWorkingGroupSpeaker_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionWorkingGroupSpeaker_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionWorkingGroupSpeaker_MotionID(id).Lazy(&c.MotionID)
	ds.MotionWorkingGroupSpeaker_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *motionWorkingGroupSpeakerBuilder) Preload(rel builderWrapperI) *motionWorkingGroupSpeakerBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *motionWorkingGroupSpeakerBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *motionWorkingGroupSpeakerBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (b *motionWorkingGroupSpeakerBuilder) Motion() *motionBuilder {
	return &motionBuilder{
		builder: builder[motionBuilder, *motionBuilder, Motion]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MotionID",
			relField: "Motion",
		},
	}
}

func (r *Fetch) MotionWorkingGroupSpeaker(ids ...int) *motionWorkingGroupSpeakerBuilder {
	return &motionWorkingGroupSpeakerBuilder{
		builder: builder[motionWorkingGroupSpeakerBuilder, *motionWorkingGroupSpeakerBuilder, MotionWorkingGroupSpeaker]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Option has all fields from option.
type Option struct {
	Abstain                    maybe.Maybe[decimal.Decimal]
	ContentObjectID            maybe.Maybe[string]
	ID                         int
	MeetingID                  int
	No                         maybe.Maybe[decimal.Decimal]
	PollID                     maybe.Maybe[int]
	Text                       maybe.Maybe[string]
	UsedAsGlobalOptionInPollID maybe.Maybe[int]
	VoteIDs                    maybe.Maybe[[]int]
	Weight                     maybe.Maybe[int]
	Yes                        maybe.Maybe[decimal.Decimal]
	Meeting                    *Meeting
	Poll                       *maybe.Maybe[Poll]
	UsedAsGlobalOptionInPoll   *maybe.Maybe[Poll]
	VoteList                   []Vote
}

type optionBuilder struct {
	builder[optionBuilder, *optionBuilder, Option]
}

func (b *optionBuilder) lazy(ds *Fetch, id int) *Option {
	c := Option{}
	ds.Option_Abstain(id).Lazy(&c.Abstain)
	ds.Option_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Option_ID(id).Lazy(&c.ID)
	ds.Option_MeetingID(id).Lazy(&c.MeetingID)
	ds.Option_No(id).Lazy(&c.No)
	ds.Option_PollID(id).Lazy(&c.PollID)
	ds.Option_Text(id).Lazy(&c.Text)
	ds.Option_UsedAsGlobalOptionInPollID(id).Lazy(&c.UsedAsGlobalOptionInPollID)
	ds.Option_VoteIDs(id).Lazy(&c.VoteIDs)
	ds.Option_Weight(id).Lazy(&c.Weight)
	ds.Option_Yes(id).Lazy(&c.Yes)
	return &c
}

func (b *optionBuilder) Preload(rel builderWrapperI) *optionBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *optionBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *optionBuilder) Poll() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollID",
			relField: "Poll",
		},
	}
}

func (b *optionBuilder) UsedAsGlobalOptionInPoll() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsGlobalOptionInPollID",
			relField: "UsedAsGlobalOptionInPoll",
		},
	}
}

func (b *optionBuilder) VoteList() *voteBuilder {
	return &voteBuilder{
		builder: builder[voteBuilder, *voteBuilder, Vote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "VoteIDs",
			relField: "VoteList",
			many:     true,
		},
	}
}

func (r *Fetch) Option(ids ...int) *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Organization has all fields from organization.
type Organization struct {
	ActiveMeetingIDs                        maybe.Maybe[[]int]
	ArchivedMeetingIDs                      maybe.Maybe[[]int]
	CommitteeIDs                            maybe.Maybe[[]int]
	DefaultLanguage                         maybe.Maybe[dstypes.Languages]
	Description                             maybe.Maybe[string]
	DisableForwardWithAttachments           maybe.Maybe[bool]
	EnableAnonymous                         maybe.Maybe[bool]
	EnableChat                              maybe.Maybe[bool]
	EnableElectronicVoting                  maybe.Maybe[bool]
	GenderIDs                               maybe.Maybe[[]int]
	ID                                      int
	LegalNotice                             maybe.Maybe[string]
	LimitOfMeetings                         maybe.Maybe[int]
	LimitOfUsers                            maybe.Maybe[int]
	LoginText                               maybe.Maybe[string]
	MediafileIDs                            maybe.Maybe[[]int]
	Name                                    maybe.Maybe[string]
	OrganizationTagIDs                      maybe.Maybe[[]int]
	PrivacyPolicy                           maybe.Maybe[string]
	PublishedMediafileIDs                   maybe.Maybe[[]int]
	RequireDuplicateFrom                    maybe.Maybe[bool]
	ResetPasswordVerboseErrors              maybe.Maybe[bool]
	RestrictEditForwardCommittees           maybe.Maybe[bool]
	RestrictEditingSameLevelCommitteeAdmins maybe.Maybe[bool]
	SamlAttrMapping                         maybe.Maybe[json.RawMessage]
	SamlEnabled                             maybe.Maybe[bool]
	SamlLoginButtonText                     maybe.Maybe[string]
	SamlMetadataIDp                         maybe.Maybe[string]
	SamlMetadataSp                          maybe.Maybe[string]
	SamlPrivateKey                          maybe.Maybe[string]
	TemplateMeetingIDs                      maybe.Maybe[[]int]
	ThemeID                                 int
	ThemeIDs                                maybe.Maybe[[]int]
	TimeZone                                maybe.Maybe[string]
	Url                                     maybe.Maybe[string]
	UserIDs                                 maybe.Maybe[[]int]
	UsersEmailBody                          maybe.Maybe[string]
	UsersEmailReplyto                       maybe.Maybe[string]
	UsersEmailSender                        maybe.Maybe[string]
	UsersEmailSubject                       maybe.Maybe[string]
	ActiveMeetingList                       []Meeting
	ArchivedMeetingList                     []Meeting
	CommitteeList                           []Committee
	GenderList                              []Gender
	MediafileList                           []Mediafile
	OrganizationTagList                     []OrganizationTag
	PublishedMediafileList                  []Mediafile
	TemplateMeetingList                     []Meeting
	Theme                                   *Theme
	ThemeList                               []Theme
	UserList                                []User
}

type organizationBuilder struct {
	builder[organizationBuilder, *organizationBuilder, Organization]
}

func (b *organizationBuilder) lazy(ds *Fetch, id int) *Organization {
	c := Organization{}
	ds.Organization_ActiveMeetingIDs(id).Lazy(&c.ActiveMeetingIDs)
	ds.Organization_ArchivedMeetingIDs(id).Lazy(&c.ArchivedMeetingIDs)
	ds.Organization_CommitteeIDs(id).Lazy(&c.CommitteeIDs)
	ds.Organization_DefaultLanguage(id).Lazy(&c.DefaultLanguage)
	ds.Organization_Description(id).Lazy(&c.Description)
	ds.Organization_DisableForwardWithAttachments(id).Lazy(&c.DisableForwardWithAttachments)
	ds.Organization_EnableAnonymous(id).Lazy(&c.EnableAnonymous)
	ds.Organization_EnableChat(id).Lazy(&c.EnableChat)
	ds.Organization_EnableElectronicVoting(id).Lazy(&c.EnableElectronicVoting)
	ds.Organization_GenderIDs(id).Lazy(&c.GenderIDs)
	ds.Organization_ID(id).Lazy(&c.ID)
	ds.Organization_LegalNotice(id).Lazy(&c.LegalNotice)
	ds.Organization_LimitOfMeetings(id).Lazy(&c.LimitOfMeetings)
	ds.Organization_LimitOfUsers(id).Lazy(&c.LimitOfUsers)
	ds.Organization_LoginText(id).Lazy(&c.LoginText)
	ds.Organization_MediafileIDs(id).Lazy(&c.MediafileIDs)
	ds.Organization_Name(id).Lazy(&c.Name)
	ds.Organization_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Organization_PrivacyPolicy(id).Lazy(&c.PrivacyPolicy)
	ds.Organization_PublishedMediafileIDs(id).Lazy(&c.PublishedMediafileIDs)
	ds.Organization_RequireDuplicateFrom(id).Lazy(&c.RequireDuplicateFrom)
	ds.Organization_ResetPasswordVerboseErrors(id).Lazy(&c.ResetPasswordVerboseErrors)
	ds.Organization_RestrictEditForwardCommittees(id).Lazy(&c.RestrictEditForwardCommittees)
	ds.Organization_RestrictEditingSameLevelCommitteeAdmins(id).Lazy(&c.RestrictEditingSameLevelCommitteeAdmins)
	ds.Organization_SamlAttrMapping(id).Lazy(&c.SamlAttrMapping)
	ds.Organization_SamlEnabled(id).Lazy(&c.SamlEnabled)
	ds.Organization_SamlLoginButtonText(id).Lazy(&c.SamlLoginButtonText)
	ds.Organization_SamlMetadataIDp(id).Lazy(&c.SamlMetadataIDp)
	ds.Organization_SamlMetadataSp(id).Lazy(&c.SamlMetadataSp)
	ds.Organization_SamlPrivateKey(id).Lazy(&c.SamlPrivateKey)
	ds.Organization_TemplateMeetingIDs(id).Lazy(&c.TemplateMeetingIDs)
	ds.Organization_ThemeID(id).Lazy(&c.ThemeID)
	ds.Organization_ThemeIDs(id).Lazy(&c.ThemeIDs)
	ds.Organization_TimeZone(id).Lazy(&c.TimeZone)
	ds.Organization_Url(id).Lazy(&c.Url)
	ds.Organization_UserIDs(id).Lazy(&c.UserIDs)
	ds.Organization_UsersEmailBody(id).Lazy(&c.UsersEmailBody)
	ds.Organization_UsersEmailReplyto(id).Lazy(&c.UsersEmailReplyto)
	ds.Organization_UsersEmailSender(id).Lazy(&c.UsersEmailSender)
	ds.Organization_UsersEmailSubject(id).Lazy(&c.UsersEmailSubject)
	return &c
}

func (b *organizationBuilder) Preload(rel builderWrapperI) *organizationBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *organizationBuilder) ActiveMeetingList() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ActiveMeetingIDs",
			relField: "ActiveMeetingList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) ArchivedMeetingList() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ArchivedMeetingIDs",
			relField: "ArchivedMeetingList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) CommitteeList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CommitteeIDs",
			relField: "CommitteeList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) GenderList() *genderBuilder {
	return &genderBuilder{
		builder: builder[genderBuilder, *genderBuilder, Gender]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "GenderIDs",
			relField: "GenderList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) MediafileList() *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MediafileIDs",
			relField: "MediafileList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) OrganizationTagList() *organizationTagBuilder {
	return &organizationTagBuilder{
		builder: builder[organizationTagBuilder, *organizationTagBuilder, OrganizationTag]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationTagIDs",
			relField: "OrganizationTagList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) PublishedMediafileList() *mediafileBuilder {
	return &mediafileBuilder{
		builder: builder[mediafileBuilder, *mediafileBuilder, Mediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PublishedMediafileIDs",
			relField: "PublishedMediafileList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) TemplateMeetingList() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "TemplateMeetingIDs",
			relField: "TemplateMeetingList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) Theme() *themeBuilder {
	return &themeBuilder{
		builder: builder[themeBuilder, *themeBuilder, Theme]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ThemeID",
			relField: "Theme",
		},
	}
}

func (b *organizationBuilder) ThemeList() *themeBuilder {
	return &themeBuilder{
		builder: builder[themeBuilder, *themeBuilder, Theme]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ThemeIDs",
			relField: "ThemeList",
			many:     true,
		},
	}
}

func (b *organizationBuilder) UserList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserIDs",
			relField: "UserList",
			many:     true,
		},
	}
}

func (r *Fetch) Organization(ids ...int) *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			ids:   ids,
			fetch: r,
		},
	}
}

// OrganizationTag has all fields from organization_tag.
type OrganizationTag struct {
	Color          string
	ID             int
	Name           string
	OrganizationID int
	TaggedIDs      maybe.Maybe[[]string]
	Organization   *Organization
}

type organizationTagBuilder struct {
	builder[organizationTagBuilder, *organizationTagBuilder, OrganizationTag]
}

func (b *organizationTagBuilder) lazy(ds *Fetch, id int) *OrganizationTag {
	c := OrganizationTag{}
	ds.OrganizationTag_Color(id).Lazy(&c.Color)
	ds.OrganizationTag_ID(id).Lazy(&c.ID)
	ds.OrganizationTag_Name(id).Lazy(&c.Name)
	ds.OrganizationTag_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.OrganizationTag_TaggedIDs(id).Lazy(&c.TaggedIDs)
	return &c
}

func (b *organizationTagBuilder) Preload(rel builderWrapperI) *organizationTagBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *organizationTagBuilder) Organization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationID",
			relField: "Organization",
		},
	}
}

func (r *Fetch) OrganizationTag(ids ...int) *organizationTagBuilder {
	return &organizationTagBuilder{
		builder: builder[organizationTagBuilder, *organizationTagBuilder, OrganizationTag]{
			ids:   ids,
			fetch: r,
		},
	}
}

// PersonalNote has all fields from personal_note.
type PersonalNote struct {
	ContentObjectID string
	ID              int
	MeetingID       int
	MeetingUserID   int
	Note            maybe.Maybe[string]
	Star            maybe.Maybe[bool]
	Meeting         *Meeting
	MeetingUser     *MeetingUser
}

type personalNoteBuilder struct {
	builder[personalNoteBuilder, *personalNoteBuilder, PersonalNote]
}

func (b *personalNoteBuilder) lazy(ds *Fetch, id int) *PersonalNote {
	c := PersonalNote{}
	ds.PersonalNote_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.PersonalNote_ID(id).Lazy(&c.ID)
	ds.PersonalNote_MeetingID(id).Lazy(&c.MeetingID)
	ds.PersonalNote_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.PersonalNote_Note(id).Lazy(&c.Note)
	ds.PersonalNote_Star(id).Lazy(&c.Star)
	return &c
}

func (b *personalNoteBuilder) Preload(rel builderWrapperI) *personalNoteBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *personalNoteBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *personalNoteBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (r *Fetch) PersonalNote(ids ...int) *personalNoteBuilder {
	return &personalNoteBuilder{
		builder: builder[personalNoteBuilder, *personalNoteBuilder, PersonalNote]{
			ids:   ids,
			fetch: r,
		},
	}
}

// PointOfOrderCategory has all fields from point_of_order_category.
type PointOfOrderCategory struct {
	ID          int
	MeetingID   int
	Rank        int
	SpeakerIDs  maybe.Maybe[[]int]
	Text        string
	Meeting     *Meeting
	SpeakerList []Speaker
}

type pointOfOrderCategoryBuilder struct {
	builder[pointOfOrderCategoryBuilder, *pointOfOrderCategoryBuilder, PointOfOrderCategory]
}

func (b *pointOfOrderCategoryBuilder) lazy(ds *Fetch, id int) *PointOfOrderCategory {
	c := PointOfOrderCategory{}
	ds.PointOfOrderCategory_ID(id).Lazy(&c.ID)
	ds.PointOfOrderCategory_MeetingID(id).Lazy(&c.MeetingID)
	ds.PointOfOrderCategory_Rank(id).Lazy(&c.Rank)
	ds.PointOfOrderCategory_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.PointOfOrderCategory_Text(id).Lazy(&c.Text)
	return &c
}

func (b *pointOfOrderCategoryBuilder) Preload(rel builderWrapperI) *pointOfOrderCategoryBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *pointOfOrderCategoryBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *pointOfOrderCategoryBuilder) SpeakerList() *speakerBuilder {
	return &speakerBuilder{
		builder: builder[speakerBuilder, *speakerBuilder, Speaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SpeakerIDs",
			relField: "SpeakerList",
			many:     true,
		},
	}
}

func (r *Fetch) PointOfOrderCategory(ids ...int) *pointOfOrderCategoryBuilder {
	return &pointOfOrderCategoryBuilder{
		builder: builder[pointOfOrderCategoryBuilder, *pointOfOrderCategoryBuilder, PointOfOrderCategory]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Poll has all fields from poll.
type Poll struct {
	Backend               dstypes.PollBackends
	ContentObjectID       string
	Description           maybe.Maybe[string]
	EntitledGroupIDs      maybe.Maybe[[]int]
	EntitledUsersAtStop   maybe.Maybe[json.RawMessage]
	GlobalAbstain         maybe.Maybe[bool]
	GlobalNo              maybe.Maybe[bool]
	GlobalOptionID        maybe.Maybe[int]
	GlobalYes             maybe.Maybe[bool]
	ID                    int
	IsPseudoanonymized    maybe.Maybe[bool]
	LiveVotes             maybe.Maybe[json.RawMessage]
	LiveVotingEnabled     maybe.Maybe[bool]
	MaxVotesAmount        maybe.Maybe[int]
	MaxVotesPerOption     maybe.Maybe[int]
	MeetingID             int
	MinVotesAmount        maybe.Maybe[int]
	OnehundredPercentBase dstypes.OnehundredPercentBases
	OptionIDs             maybe.Maybe[[]int]
	Pollmethod            dstypes.Poll_Pollmethod
	ProjectionIDs         maybe.Maybe[[]int]
	SequentialNumber      int
	State                 maybe.Maybe[dstypes.Poll_State]
	Title                 string
	Type                  dstypes.Poll_Type
	VotedIDs              maybe.Maybe[[]int]
	Votescast             maybe.Maybe[decimal.Decimal]
	Votesinvalid          maybe.Maybe[decimal.Decimal]
	Votesvalid            maybe.Maybe[decimal.Decimal]
	EntitledGroupList     []Group
	GlobalOption          *maybe.Maybe[Option]
	Meeting               *Meeting
	OptionList            []Option
	ProjectionList        []Projection
	VotedList             []User
}

type pollBuilder struct {
	builder[pollBuilder, *pollBuilder, Poll]
}

func (b *pollBuilder) lazy(ds *Fetch, id int) *Poll {
	c := Poll{}
	ds.Poll_Backend(id).Lazy(&c.Backend)
	ds.Poll_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Poll_Description(id).Lazy(&c.Description)
	ds.Poll_EntitledGroupIDs(id).Lazy(&c.EntitledGroupIDs)
	ds.Poll_EntitledUsersAtStop(id).Lazy(&c.EntitledUsersAtStop)
	ds.Poll_GlobalAbstain(id).Lazy(&c.GlobalAbstain)
	ds.Poll_GlobalNo(id).Lazy(&c.GlobalNo)
	ds.Poll_GlobalOptionID(id).Lazy(&c.GlobalOptionID)
	ds.Poll_GlobalYes(id).Lazy(&c.GlobalYes)
	ds.Poll_ID(id).Lazy(&c.ID)
	ds.Poll_IsPseudoanonymized(id).Lazy(&c.IsPseudoanonymized)
	ds.Poll_LiveVotes(id).Lazy(&c.LiveVotes)
	ds.Poll_LiveVotingEnabled(id).Lazy(&c.LiveVotingEnabled)
	ds.Poll_MaxVotesAmount(id).Lazy(&c.MaxVotesAmount)
	ds.Poll_MaxVotesPerOption(id).Lazy(&c.MaxVotesPerOption)
	ds.Poll_MeetingID(id).Lazy(&c.MeetingID)
	ds.Poll_MinVotesAmount(id).Lazy(&c.MinVotesAmount)
	ds.Poll_OnehundredPercentBase(id).Lazy(&c.OnehundredPercentBase)
	ds.Poll_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.Poll_Pollmethod(id).Lazy(&c.Pollmethod)
	ds.Poll_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Poll_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Poll_State(id).Lazy(&c.State)
	ds.Poll_Title(id).Lazy(&c.Title)
	ds.Poll_Type(id).Lazy(&c.Type)
	ds.Poll_VotedIDs(id).Lazy(&c.VotedIDs)
	ds.Poll_Votescast(id).Lazy(&c.Votescast)
	ds.Poll_Votesinvalid(id).Lazy(&c.Votesinvalid)
	ds.Poll_Votesvalid(id).Lazy(&c.Votesvalid)
	return &c
}

func (b *pollBuilder) Preload(rel builderWrapperI) *pollBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *pollBuilder) EntitledGroupList() *groupBuilder {
	return &groupBuilder{
		builder: builder[groupBuilder, *groupBuilder, Group]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "EntitledGroupIDs",
			relField: "EntitledGroupList",
			many:     true,
		},
	}
}

func (b *pollBuilder) GlobalOption() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "GlobalOptionID",
			relField: "GlobalOption",
		},
	}
}

func (b *pollBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *pollBuilder) OptionList() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OptionIDs",
			relField: "OptionList",
			many:     true,
		},
	}
}

func (b *pollBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *pollBuilder) VotedList() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "VotedIDs",
			relField: "VotedList",
			many:     true,
		},
	}
}

func (r *Fetch) Poll(ids ...int) *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			ids:   ids,
			fetch: r,
		},
	}
}

// PollCandidate has all fields from poll_candidate.
type PollCandidate struct {
	ID                  int
	MeetingID           int
	PollCandidateListID int
	UserID              maybe.Maybe[int]
	Weight              int
	Meeting             *Meeting
	PollCandidateList   *PollCandidateList
	User                *maybe.Maybe[User]
}

type pollCandidateBuilder struct {
	builder[pollCandidateBuilder, *pollCandidateBuilder, PollCandidate]
}

func (b *pollCandidateBuilder) lazy(ds *Fetch, id int) *PollCandidate {
	c := PollCandidate{}
	ds.PollCandidate_ID(id).Lazy(&c.ID)
	ds.PollCandidate_MeetingID(id).Lazy(&c.MeetingID)
	ds.PollCandidate_PollCandidateListID(id).Lazy(&c.PollCandidateListID)
	ds.PollCandidate_UserID(id).Lazy(&c.UserID)
	ds.PollCandidate_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *pollCandidateBuilder) Preload(rel builderWrapperI) *pollCandidateBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *pollCandidateBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *pollCandidateBuilder) PollCandidateList() *pollCandidateListBuilder {
	return &pollCandidateListBuilder{
		builder: builder[pollCandidateListBuilder, *pollCandidateListBuilder, PollCandidateList]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollCandidateListID",
			relField: "PollCandidateList",
		},
	}
}

func (b *pollCandidateBuilder) User() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserID",
			relField: "User",
		},
	}
}

func (r *Fetch) PollCandidate(ids ...int) *pollCandidateBuilder {
	return &pollCandidateBuilder{
		builder: builder[pollCandidateBuilder, *pollCandidateBuilder, PollCandidate]{
			ids:   ids,
			fetch: r,
		},
	}
}

// PollCandidateList has all fields from poll_candidate_list.
type PollCandidateList struct {
	ID                int
	MeetingID         int
	OptionID          int
	PollCandidateIDs  maybe.Maybe[[]int]
	Meeting           *Meeting
	Option            *Option
	PollCandidateList []PollCandidate
}

type pollCandidateListBuilder struct {
	builder[pollCandidateListBuilder, *pollCandidateListBuilder, PollCandidateList]
}

func (b *pollCandidateListBuilder) lazy(ds *Fetch, id int) *PollCandidateList {
	c := PollCandidateList{}
	ds.PollCandidateList_ID(id).Lazy(&c.ID)
	ds.PollCandidateList_MeetingID(id).Lazy(&c.MeetingID)
	ds.PollCandidateList_OptionID(id).Lazy(&c.OptionID)
	ds.PollCandidateList_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
	return &c
}

func (b *pollCandidateListBuilder) Preload(rel builderWrapperI) *pollCandidateListBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *pollCandidateListBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *pollCandidateListBuilder) Option() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OptionID",
			relField: "Option",
		},
	}
}

func (b *pollCandidateListBuilder) PollCandidateList() *pollCandidateBuilder {
	return &pollCandidateBuilder{
		builder: builder[pollCandidateBuilder, *pollCandidateBuilder, PollCandidate]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollCandidateIDs",
			relField: "PollCandidateList",
			many:     true,
		},
	}
}

func (r *Fetch) PollCandidateList(ids ...int) *pollCandidateListBuilder {
	return &pollCandidateListBuilder{
		builder: builder[pollCandidateListBuilder, *pollCandidateListBuilder, PollCandidateList]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Projection has all fields from projection.
type Projection struct {
	Content            maybe.Maybe[json.RawMessage]
	ContentObjectID    string
	CurrentProjectorID maybe.Maybe[int]
	HistoryProjectorID maybe.Maybe[int]
	ID                 int
	MeetingID          int
	Options            maybe.Maybe[json.RawMessage]
	PreviewProjectorID maybe.Maybe[int]
	Stable             maybe.Maybe[bool]
	Type               maybe.Maybe[string]
	Weight             maybe.Maybe[int]
	CurrentProjector   *maybe.Maybe[Projector]
	HistoryProjector   *maybe.Maybe[Projector]
	Meeting            *Meeting
	PreviewProjector   *maybe.Maybe[Projector]
}

type projectionBuilder struct {
	builder[projectionBuilder, *projectionBuilder, Projection]
}

func (b *projectionBuilder) lazy(ds *Fetch, id int) *Projection {
	c := Projection{}
	ds.Projection_Content(id).Lazy(&c.Content)
	ds.Projection_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Projection_CurrentProjectorID(id).Lazy(&c.CurrentProjectorID)
	ds.Projection_HistoryProjectorID(id).Lazy(&c.HistoryProjectorID)
	ds.Projection_ID(id).Lazy(&c.ID)
	ds.Projection_MeetingID(id).Lazy(&c.MeetingID)
	ds.Projection_Options(id).Lazy(&c.Options)
	ds.Projection_PreviewProjectorID(id).Lazy(&c.PreviewProjectorID)
	ds.Projection_Stable(id).Lazy(&c.Stable)
	ds.Projection_Type(id).Lazy(&c.Type)
	ds.Projection_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *projectionBuilder) Preload(rel builderWrapperI) *projectionBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *projectionBuilder) CurrentProjector() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CurrentProjectorID",
			relField: "CurrentProjector",
		},
	}
}

func (b *projectionBuilder) HistoryProjector() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HistoryProjectorID",
			relField: "HistoryProjector",
		},
	}
}

func (b *projectionBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *projectionBuilder) PreviewProjector() *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PreviewProjectorID",
			relField: "PreviewProjector",
		},
	}
}

func (r *Fetch) Projection(ids ...int) *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Projector has all fields from projector.
type Projector struct {
	AspectRatioDenominator                             maybe.Maybe[int]
	AspectRatioNumerator                               maybe.Maybe[int]
	BackgroundColor                                    maybe.Maybe[string]
	ChyronBackgroundColor                              maybe.Maybe[string]
	ChyronBackgroundColor2                             maybe.Maybe[string]
	ChyronFontColor                                    maybe.Maybe[string]
	ChyronFontColor2                                   maybe.Maybe[string]
	Color                                              maybe.Maybe[string]
	CurrentProjectionIDs                               maybe.Maybe[[]int]
	HeaderBackgroundColor                              maybe.Maybe[string]
	HeaderFontColor                                    maybe.Maybe[string]
	HeaderH1Color                                      maybe.Maybe[string]
	HistoryProjectionIDs                               maybe.Maybe[[]int]
	ID                                                 int
	IsInternal                                         maybe.Maybe[bool]
	MeetingID                                          int
	Name                                               string
	PreviewProjectionIDs                               maybe.Maybe[[]int]
	Scale                                              maybe.Maybe[int]
	Scroll                                             maybe.Maybe[int]
	SequentialNumber                                   int
	ShowClock                                          maybe.Maybe[bool]
	ShowHeaderFooter                                   maybe.Maybe[bool]
	ShowLogo                                           maybe.Maybe[bool]
	ShowTitle                                          maybe.Maybe[bool]
	UsedAsDefaultProjectorForAgendaItemListInMeetingID maybe.Maybe[int]
	UsedAsDefaultProjectorForAmendmentInMeetingID      maybe.Maybe[int]
	UsedAsDefaultProjectorForAssignmentInMeetingID     maybe.Maybe[int]
	UsedAsDefaultProjectorForAssignmentPollInMeetingID maybe.Maybe[int]
	UsedAsDefaultProjectorForCountdownInMeetingID      maybe.Maybe[int]
	UsedAsDefaultProjectorForCurrentLosInMeetingID     maybe.Maybe[int]
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID maybe.Maybe[int]
	UsedAsDefaultProjectorForMediafileInMeetingID      maybe.Maybe[int]
	UsedAsDefaultProjectorForMessageInMeetingID        maybe.Maybe[int]
	UsedAsDefaultProjectorForMotionBlockInMeetingID    maybe.Maybe[int]
	UsedAsDefaultProjectorForMotionInMeetingID         maybe.Maybe[int]
	UsedAsDefaultProjectorForMotionPollInMeetingID     maybe.Maybe[int]
	UsedAsDefaultProjectorForPollInMeetingID           maybe.Maybe[int]
	UsedAsDefaultProjectorForTopicInMeetingID          maybe.Maybe[int]
	UsedAsReferenceProjectorMeetingID                  maybe.Maybe[int]
	Width                                              maybe.Maybe[int]
	CurrentProjectionList                              []Projection
	HistoryProjectionList                              []Projection
	Meeting                                            *Meeting
	PreviewProjectionList                              []Projection
	UsedAsDefaultProjectorForAgendaItemListInMeeting   *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForAmendmentInMeeting        *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForAssignmentInMeeting       *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForAssignmentPollInMeeting   *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForCountdownInMeeting        *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForCurrentLosInMeeting       *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForListOfSpeakersInMeeting   *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForMediafileInMeeting        *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForMessageInMeeting          *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForMotionBlockInMeeting      *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForMotionInMeeting           *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForMotionPollInMeeting       *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForPollInMeeting             *maybe.Maybe[Meeting]
	UsedAsDefaultProjectorForTopicInMeeting            *maybe.Maybe[Meeting]
	UsedAsReferenceProjectorMeeting                    *maybe.Maybe[Meeting]
}

type projectorBuilder struct {
	builder[projectorBuilder, *projectorBuilder, Projector]
}

func (b *projectorBuilder) lazy(ds *Fetch, id int) *Projector {
	c := Projector{}
	ds.Projector_AspectRatioDenominator(id).Lazy(&c.AspectRatioDenominator)
	ds.Projector_AspectRatioNumerator(id).Lazy(&c.AspectRatioNumerator)
	ds.Projector_BackgroundColor(id).Lazy(&c.BackgroundColor)
	ds.Projector_ChyronBackgroundColor(id).Lazy(&c.ChyronBackgroundColor)
	ds.Projector_ChyronBackgroundColor2(id).Lazy(&c.ChyronBackgroundColor2)
	ds.Projector_ChyronFontColor(id).Lazy(&c.ChyronFontColor)
	ds.Projector_ChyronFontColor2(id).Lazy(&c.ChyronFontColor2)
	ds.Projector_Color(id).Lazy(&c.Color)
	ds.Projector_CurrentProjectionIDs(id).Lazy(&c.CurrentProjectionIDs)
	ds.Projector_HeaderBackgroundColor(id).Lazy(&c.HeaderBackgroundColor)
	ds.Projector_HeaderFontColor(id).Lazy(&c.HeaderFontColor)
	ds.Projector_HeaderH1Color(id).Lazy(&c.HeaderH1Color)
	ds.Projector_HistoryProjectionIDs(id).Lazy(&c.HistoryProjectionIDs)
	ds.Projector_ID(id).Lazy(&c.ID)
	ds.Projector_IsInternal(id).Lazy(&c.IsInternal)
	ds.Projector_MeetingID(id).Lazy(&c.MeetingID)
	ds.Projector_Name(id).Lazy(&c.Name)
	ds.Projector_PreviewProjectionIDs(id).Lazy(&c.PreviewProjectionIDs)
	ds.Projector_Scale(id).Lazy(&c.Scale)
	ds.Projector_Scroll(id).Lazy(&c.Scroll)
	ds.Projector_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Projector_ShowClock(id).Lazy(&c.ShowClock)
	ds.Projector_ShowHeaderFooter(id).Lazy(&c.ShowHeaderFooter)
	ds.Projector_ShowLogo(id).Lazy(&c.ShowLogo)
	ds.Projector_ShowTitle(id).Lazy(&c.ShowTitle)
	ds.Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAgendaItemListInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAmendmentInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAssignmentInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAssignmentPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForCountdownInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForCountdownInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForCurrentLosInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForCurrentLosInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForListOfSpeakersInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMediafileInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMediafileInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMessageInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMessageInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMotionBlockInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMotionInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMotionPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForPollInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForTopicInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForTopicInMeetingID)
	ds.Projector_UsedAsReferenceProjectorMeetingID(id).Lazy(&c.UsedAsReferenceProjectorMeetingID)
	ds.Projector_Width(id).Lazy(&c.Width)
	return &c
}

func (b *projectorBuilder) Preload(rel builderWrapperI) *projectorBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *projectorBuilder) CurrentProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CurrentProjectionIDs",
			relField: "CurrentProjectionList",
			many:     true,
		},
	}
}

func (b *projectorBuilder) HistoryProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HistoryProjectionIDs",
			relField: "HistoryProjectionList",
			many:     true,
		},
	}
}

func (b *projectorBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *projectorBuilder) PreviewProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PreviewProjectionIDs",
			relField: "PreviewProjectionList",
			many:     true,
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForAgendaItemListInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForAgendaItemListInMeetingID",
			relField: "UsedAsDefaultProjectorForAgendaItemListInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForAmendmentInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForAmendmentInMeetingID",
			relField: "UsedAsDefaultProjectorForAmendmentInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForAssignmentInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForAssignmentInMeetingID",
			relField: "UsedAsDefaultProjectorForAssignmentInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForAssignmentPollInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForAssignmentPollInMeetingID",
			relField: "UsedAsDefaultProjectorForAssignmentPollInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForCountdownInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForCountdownInMeetingID",
			relField: "UsedAsDefaultProjectorForCountdownInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForCurrentLosInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForCurrentLosInMeetingID",
			relField: "UsedAsDefaultProjectorForCurrentLosInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForListOfSpeakersInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForListOfSpeakersInMeetingID",
			relField: "UsedAsDefaultProjectorForListOfSpeakersInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForMediafileInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForMediafileInMeetingID",
			relField: "UsedAsDefaultProjectorForMediafileInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForMessageInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForMessageInMeetingID",
			relField: "UsedAsDefaultProjectorForMessageInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForMotionBlockInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForMotionBlockInMeetingID",
			relField: "UsedAsDefaultProjectorForMotionBlockInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForMotionInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForMotionInMeetingID",
			relField: "UsedAsDefaultProjectorForMotionInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForMotionPollInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForMotionPollInMeetingID",
			relField: "UsedAsDefaultProjectorForMotionPollInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForPollInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForPollInMeetingID",
			relField: "UsedAsDefaultProjectorForPollInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsDefaultProjectorForTopicInMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsDefaultProjectorForTopicInMeetingID",
			relField: "UsedAsDefaultProjectorForTopicInMeeting",
		},
	}
}

func (b *projectorBuilder) UsedAsReferenceProjectorMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsReferenceProjectorMeetingID",
			relField: "UsedAsReferenceProjectorMeeting",
		},
	}
}

func (r *Fetch) Projector(ids ...int) *projectorBuilder {
	return &projectorBuilder{
		builder: builder[projectorBuilder, *projectorBuilder, Projector]{
			ids:   ids,
			fetch: r,
		},
	}
}

// ProjectorCountdown has all fields from projector_countdown.
type ProjectorCountdown struct {
	CountdownTime                          maybe.Maybe[float64]
	DefaultTime                            maybe.Maybe[int]
	Description                            maybe.Maybe[string]
	ID                                     int
	MeetingID                              int
	ProjectionIDs                          maybe.Maybe[[]int]
	Running                                maybe.Maybe[bool]
	Title                                  string
	UsedAsListOfSpeakersCountdownMeetingID maybe.Maybe[int]
	UsedAsPollCountdownMeetingID           maybe.Maybe[int]
	Meeting                                *Meeting
	ProjectionList                         []Projection
	UsedAsListOfSpeakersCountdownMeeting   *maybe.Maybe[Meeting]
	UsedAsPollCountdownMeeting             *maybe.Maybe[Meeting]
}

type projectorCountdownBuilder struct {
	builder[projectorCountdownBuilder, *projectorCountdownBuilder, ProjectorCountdown]
}

func (b *projectorCountdownBuilder) lazy(ds *Fetch, id int) *ProjectorCountdown {
	c := ProjectorCountdown{}
	ds.ProjectorCountdown_CountdownTime(id).Lazy(&c.CountdownTime)
	ds.ProjectorCountdown_DefaultTime(id).Lazy(&c.DefaultTime)
	ds.ProjectorCountdown_Description(id).Lazy(&c.Description)
	ds.ProjectorCountdown_ID(id).Lazy(&c.ID)
	ds.ProjectorCountdown_MeetingID(id).Lazy(&c.MeetingID)
	ds.ProjectorCountdown_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.ProjectorCountdown_Running(id).Lazy(&c.Running)
	ds.ProjectorCountdown_Title(id).Lazy(&c.Title)
	ds.ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(id).Lazy(&c.UsedAsListOfSpeakersCountdownMeetingID)
	ds.ProjectorCountdown_UsedAsPollCountdownMeetingID(id).Lazy(&c.UsedAsPollCountdownMeetingID)
	return &c
}

func (b *projectorCountdownBuilder) Preload(rel builderWrapperI) *projectorCountdownBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *projectorCountdownBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *projectorCountdownBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (b *projectorCountdownBuilder) UsedAsListOfSpeakersCountdownMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsListOfSpeakersCountdownMeetingID",
			relField: "UsedAsListOfSpeakersCountdownMeeting",
		},
	}
}

func (b *projectorCountdownBuilder) UsedAsPollCountdownMeeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UsedAsPollCountdownMeetingID",
			relField: "UsedAsPollCountdownMeeting",
		},
	}
}

func (r *Fetch) ProjectorCountdown(ids ...int) *projectorCountdownBuilder {
	return &projectorCountdownBuilder{
		builder: builder[projectorCountdownBuilder, *projectorCountdownBuilder, ProjectorCountdown]{
			ids:   ids,
			fetch: r,
		},
	}
}

// ProjectorMessage has all fields from projector_message.
type ProjectorMessage struct {
	ID             int
	MeetingID      int
	Message        string
	ProjectionIDs  maybe.Maybe[[]int]
	Meeting        *Meeting
	ProjectionList []Projection
}

type projectorMessageBuilder struct {
	builder[projectorMessageBuilder, *projectorMessageBuilder, ProjectorMessage]
}

func (b *projectorMessageBuilder) lazy(ds *Fetch, id int) *ProjectorMessage {
	c := ProjectorMessage{}
	ds.ProjectorMessage_ID(id).Lazy(&c.ID)
	ds.ProjectorMessage_MeetingID(id).Lazy(&c.MeetingID)
	ds.ProjectorMessage_Message(id).Lazy(&c.Message)
	ds.ProjectorMessage_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	return &c
}

func (b *projectorMessageBuilder) Preload(rel builderWrapperI) *projectorMessageBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *projectorMessageBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *projectorMessageBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (r *Fetch) ProjectorMessage(ids ...int) *projectorMessageBuilder {
	return &projectorMessageBuilder{
		builder: builder[projectorMessageBuilder, *projectorMessageBuilder, ProjectorMessage]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Speaker has all fields from speaker.
type Speaker struct {
	Answer                         maybe.Maybe[bool]
	BeginTime                      maybe.Maybe[int]
	EndTime                        maybe.Maybe[int]
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	MeetingUserID                  maybe.Maybe[int]
	Note                           maybe.Maybe[string]
	PauseTime                      maybe.Maybe[int]
	PointOfOrder                   maybe.Maybe[bool]
	PointOfOrderCategoryID         maybe.Maybe[int]
	SpeechState                    maybe.Maybe[dstypes.Speaker_SpeechState]
	StructureLevelListOfSpeakersID maybe.Maybe[int]
	TotalPause                     maybe.Maybe[int]
	UnpauseTime                    maybe.Maybe[int]
	Weight                         maybe.Maybe[int]
	ListOfSpeakers                 *ListOfSpeakers
	Meeting                        *Meeting
	MeetingUser                    *maybe.Maybe[MeetingUser]
	PointOfOrderCategory           *maybe.Maybe[PointOfOrderCategory]
	StructureLevelListOfSpeakers   *maybe.Maybe[StructureLevelListOfSpeakers]
}

type speakerBuilder struct {
	builder[speakerBuilder, *speakerBuilder, Speaker]
}

func (b *speakerBuilder) lazy(ds *Fetch, id int) *Speaker {
	c := Speaker{}
	ds.Speaker_Answer(id).Lazy(&c.Answer)
	ds.Speaker_BeginTime(id).Lazy(&c.BeginTime)
	ds.Speaker_EndTime(id).Lazy(&c.EndTime)
	ds.Speaker_ID(id).Lazy(&c.ID)
	ds.Speaker_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Speaker_MeetingID(id).Lazy(&c.MeetingID)
	ds.Speaker_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.Speaker_Note(id).Lazy(&c.Note)
	ds.Speaker_PauseTime(id).Lazy(&c.PauseTime)
	ds.Speaker_PointOfOrder(id).Lazy(&c.PointOfOrder)
	ds.Speaker_PointOfOrderCategoryID(id).Lazy(&c.PointOfOrderCategoryID)
	ds.Speaker_SpeechState(id).Lazy(&c.SpeechState)
	ds.Speaker_StructureLevelListOfSpeakersID(id).Lazy(&c.StructureLevelListOfSpeakersID)
	ds.Speaker_TotalPause(id).Lazy(&c.TotalPause)
	ds.Speaker_UnpauseTime(id).Lazy(&c.UnpauseTime)
	ds.Speaker_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *speakerBuilder) Preload(rel builderWrapperI) *speakerBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *speakerBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *speakerBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *speakerBuilder) MeetingUser() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserID",
			relField: "MeetingUser",
		},
	}
}

func (b *speakerBuilder) PointOfOrderCategory() *pointOfOrderCategoryBuilder {
	return &pointOfOrderCategoryBuilder{
		builder: builder[pointOfOrderCategoryBuilder, *pointOfOrderCategoryBuilder, PointOfOrderCategory]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PointOfOrderCategoryID",
			relField: "PointOfOrderCategory",
		},
	}
}

func (b *speakerBuilder) StructureLevelListOfSpeakers() *structureLevelListOfSpeakersBuilder {
	return &structureLevelListOfSpeakersBuilder{
		builder: builder[structureLevelListOfSpeakersBuilder, *structureLevelListOfSpeakersBuilder, StructureLevelListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelListOfSpeakersID",
			relField: "StructureLevelListOfSpeakers",
		},
	}
}

func (r *Fetch) Speaker(ids ...int) *speakerBuilder {
	return &speakerBuilder{
		builder: builder[speakerBuilder, *speakerBuilder, Speaker]{
			ids:   ids,
			fetch: r,
		},
	}
}

// StructureLevel has all fields from structure_level.
type StructureLevel struct {
	Color                            maybe.Maybe[string]
	DefaultTime                      maybe.Maybe[int]
	ID                               int
	MeetingID                        int
	MeetingUserIDs                   maybe.Maybe[[]int]
	Name                             string
	StructureLevelListOfSpeakersIDs  maybe.Maybe[[]int]
	Meeting                          *Meeting
	MeetingUserList                  []MeetingUser
	StructureLevelListOfSpeakersList []StructureLevelListOfSpeakers
}

type structureLevelBuilder struct {
	builder[structureLevelBuilder, *structureLevelBuilder, StructureLevel]
}

func (b *structureLevelBuilder) lazy(ds *Fetch, id int) *StructureLevel {
	c := StructureLevel{}
	ds.StructureLevel_Color(id).Lazy(&c.Color)
	ds.StructureLevel_DefaultTime(id).Lazy(&c.DefaultTime)
	ds.StructureLevel_ID(id).Lazy(&c.ID)
	ds.StructureLevel_MeetingID(id).Lazy(&c.MeetingID)
	ds.StructureLevel_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.StructureLevel_Name(id).Lazy(&c.Name)
	ds.StructureLevel_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
	return &c
}

func (b *structureLevelBuilder) Preload(rel builderWrapperI) *structureLevelBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *structureLevelBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *structureLevelBuilder) MeetingUserList() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserIDs",
			relField: "MeetingUserList",
			many:     true,
		},
	}
}

func (b *structureLevelBuilder) StructureLevelListOfSpeakersList() *structureLevelListOfSpeakersBuilder {
	return &structureLevelListOfSpeakersBuilder{
		builder: builder[structureLevelListOfSpeakersBuilder, *structureLevelListOfSpeakersBuilder, StructureLevelListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelListOfSpeakersIDs",
			relField: "StructureLevelListOfSpeakersList",
			many:     true,
		},
	}
}

func (r *Fetch) StructureLevel(ids ...int) *structureLevelBuilder {
	return &structureLevelBuilder{
		builder: builder[structureLevelBuilder, *structureLevelBuilder, StructureLevel]{
			ids:   ids,
			fetch: r,
		},
	}
}

// StructureLevelListOfSpeakers has all fields from structure_level_list_of_speakers.
type StructureLevelListOfSpeakers struct {
	AdditionalTime   maybe.Maybe[float64]
	CurrentStartTime maybe.Maybe[int]
	ID               int
	InitialTime      int
	ListOfSpeakersID int
	MeetingID        int
	RemainingTime    float64
	SpeakerIDs       maybe.Maybe[[]int]
	StructureLevelID int
	ListOfSpeakers   *ListOfSpeakers
	Meeting          *Meeting
	SpeakerList      []Speaker
	StructureLevel   *StructureLevel
}

type structureLevelListOfSpeakersBuilder struct {
	builder[structureLevelListOfSpeakersBuilder, *structureLevelListOfSpeakersBuilder, StructureLevelListOfSpeakers]
}

func (b *structureLevelListOfSpeakersBuilder) lazy(ds *Fetch, id int) *StructureLevelListOfSpeakers {
	c := StructureLevelListOfSpeakers{}
	ds.StructureLevelListOfSpeakers_AdditionalTime(id).Lazy(&c.AdditionalTime)
	ds.StructureLevelListOfSpeakers_CurrentStartTime(id).Lazy(&c.CurrentStartTime)
	ds.StructureLevelListOfSpeakers_ID(id).Lazy(&c.ID)
	ds.StructureLevelListOfSpeakers_InitialTime(id).Lazy(&c.InitialTime)
	ds.StructureLevelListOfSpeakers_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.StructureLevelListOfSpeakers_MeetingID(id).Lazy(&c.MeetingID)
	ds.StructureLevelListOfSpeakers_RemainingTime(id).Lazy(&c.RemainingTime)
	ds.StructureLevelListOfSpeakers_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.StructureLevelListOfSpeakers_StructureLevelID(id).Lazy(&c.StructureLevelID)
	return &c
}

func (b *structureLevelListOfSpeakersBuilder) Preload(rel builderWrapperI) *structureLevelListOfSpeakersBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *structureLevelListOfSpeakersBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *structureLevelListOfSpeakersBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *structureLevelListOfSpeakersBuilder) SpeakerList() *speakerBuilder {
	return &speakerBuilder{
		builder: builder[speakerBuilder, *speakerBuilder, Speaker]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "SpeakerIDs",
			relField: "SpeakerList",
			many:     true,
		},
	}
}

func (b *structureLevelListOfSpeakersBuilder) StructureLevel() *structureLevelBuilder {
	return &structureLevelBuilder{
		builder: builder[structureLevelBuilder, *structureLevelBuilder, StructureLevel]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "StructureLevelID",
			relField: "StructureLevel",
		},
	}
}

func (r *Fetch) StructureLevelListOfSpeakers(ids ...int) *structureLevelListOfSpeakersBuilder {
	return &structureLevelListOfSpeakersBuilder{
		builder: builder[structureLevelListOfSpeakersBuilder, *structureLevelListOfSpeakersBuilder, StructureLevelListOfSpeakers]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Tag has all fields from tag.
type Tag struct {
	ID        int
	MeetingID int
	Name      string
	TaggedIDs maybe.Maybe[[]string]
	Meeting   *Meeting
}

type tagBuilder struct {
	builder[tagBuilder, *tagBuilder, Tag]
}

func (b *tagBuilder) lazy(ds *Fetch, id int) *Tag {
	c := Tag{}
	ds.Tag_ID(id).Lazy(&c.ID)
	ds.Tag_MeetingID(id).Lazy(&c.MeetingID)
	ds.Tag_Name(id).Lazy(&c.Name)
	ds.Tag_TaggedIDs(id).Lazy(&c.TaggedIDs)
	return &c
}

func (b *tagBuilder) Preload(rel builderWrapperI) *tagBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *tagBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (r *Fetch) Tag(ids ...int) *tagBuilder {
	return &tagBuilder{
		builder: builder[tagBuilder, *tagBuilder, Tag]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Theme has all fields from theme.
type Theme struct {
	Abstain                maybe.Maybe[string]
	Accent100              maybe.Maybe[string]
	Accent200              maybe.Maybe[string]
	Accent300              maybe.Maybe[string]
	Accent400              maybe.Maybe[string]
	Accent50               maybe.Maybe[string]
	Accent500              string
	Accent600              maybe.Maybe[string]
	Accent700              maybe.Maybe[string]
	Accent800              maybe.Maybe[string]
	Accent900              maybe.Maybe[string]
	AccentA100             maybe.Maybe[string]
	AccentA200             maybe.Maybe[string]
	AccentA400             maybe.Maybe[string]
	AccentA700             maybe.Maybe[string]
	Headbar                maybe.Maybe[string]
	ID                     int
	Name                   string
	No                     maybe.Maybe[string]
	OrganizationID         int
	Primary100             maybe.Maybe[string]
	Primary200             maybe.Maybe[string]
	Primary300             maybe.Maybe[string]
	Primary400             maybe.Maybe[string]
	Primary50              maybe.Maybe[string]
	Primary500             string
	Primary600             maybe.Maybe[string]
	Primary700             maybe.Maybe[string]
	Primary800             maybe.Maybe[string]
	Primary900             maybe.Maybe[string]
	PrimaryA100            maybe.Maybe[string]
	PrimaryA200            maybe.Maybe[string]
	PrimaryA400            maybe.Maybe[string]
	PrimaryA700            maybe.Maybe[string]
	ThemeForOrganizationID maybe.Maybe[int]
	Warn100                maybe.Maybe[string]
	Warn200                maybe.Maybe[string]
	Warn300                maybe.Maybe[string]
	Warn400                maybe.Maybe[string]
	Warn50                 maybe.Maybe[string]
	Warn500                string
	Warn600                maybe.Maybe[string]
	Warn700                maybe.Maybe[string]
	Warn800                maybe.Maybe[string]
	Warn900                maybe.Maybe[string]
	WarnA100               maybe.Maybe[string]
	WarnA200               maybe.Maybe[string]
	WarnA400               maybe.Maybe[string]
	WarnA700               maybe.Maybe[string]
	Yes                    maybe.Maybe[string]
	Organization           *Organization
	ThemeForOrganization   *maybe.Maybe[Organization]
}

type themeBuilder struct {
	builder[themeBuilder, *themeBuilder, Theme]
}

func (b *themeBuilder) lazy(ds *Fetch, id int) *Theme {
	c := Theme{}
	ds.Theme_Abstain(id).Lazy(&c.Abstain)
	ds.Theme_Accent100(id).Lazy(&c.Accent100)
	ds.Theme_Accent200(id).Lazy(&c.Accent200)
	ds.Theme_Accent300(id).Lazy(&c.Accent300)
	ds.Theme_Accent400(id).Lazy(&c.Accent400)
	ds.Theme_Accent50(id).Lazy(&c.Accent50)
	ds.Theme_Accent500(id).Lazy(&c.Accent500)
	ds.Theme_Accent600(id).Lazy(&c.Accent600)
	ds.Theme_Accent700(id).Lazy(&c.Accent700)
	ds.Theme_Accent800(id).Lazy(&c.Accent800)
	ds.Theme_Accent900(id).Lazy(&c.Accent900)
	ds.Theme_AccentA100(id).Lazy(&c.AccentA100)
	ds.Theme_AccentA200(id).Lazy(&c.AccentA200)
	ds.Theme_AccentA400(id).Lazy(&c.AccentA400)
	ds.Theme_AccentA700(id).Lazy(&c.AccentA700)
	ds.Theme_Headbar(id).Lazy(&c.Headbar)
	ds.Theme_ID(id).Lazy(&c.ID)
	ds.Theme_Name(id).Lazy(&c.Name)
	ds.Theme_No(id).Lazy(&c.No)
	ds.Theme_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Theme_Primary100(id).Lazy(&c.Primary100)
	ds.Theme_Primary200(id).Lazy(&c.Primary200)
	ds.Theme_Primary300(id).Lazy(&c.Primary300)
	ds.Theme_Primary400(id).Lazy(&c.Primary400)
	ds.Theme_Primary50(id).Lazy(&c.Primary50)
	ds.Theme_Primary500(id).Lazy(&c.Primary500)
	ds.Theme_Primary600(id).Lazy(&c.Primary600)
	ds.Theme_Primary700(id).Lazy(&c.Primary700)
	ds.Theme_Primary800(id).Lazy(&c.Primary800)
	ds.Theme_Primary900(id).Lazy(&c.Primary900)
	ds.Theme_PrimaryA100(id).Lazy(&c.PrimaryA100)
	ds.Theme_PrimaryA200(id).Lazy(&c.PrimaryA200)
	ds.Theme_PrimaryA400(id).Lazy(&c.PrimaryA400)
	ds.Theme_PrimaryA700(id).Lazy(&c.PrimaryA700)
	ds.Theme_ThemeForOrganizationID(id).Lazy(&c.ThemeForOrganizationID)
	ds.Theme_Warn100(id).Lazy(&c.Warn100)
	ds.Theme_Warn200(id).Lazy(&c.Warn200)
	ds.Theme_Warn300(id).Lazy(&c.Warn300)
	ds.Theme_Warn400(id).Lazy(&c.Warn400)
	ds.Theme_Warn50(id).Lazy(&c.Warn50)
	ds.Theme_Warn500(id).Lazy(&c.Warn500)
	ds.Theme_Warn600(id).Lazy(&c.Warn600)
	ds.Theme_Warn700(id).Lazy(&c.Warn700)
	ds.Theme_Warn800(id).Lazy(&c.Warn800)
	ds.Theme_Warn900(id).Lazy(&c.Warn900)
	ds.Theme_WarnA100(id).Lazy(&c.WarnA100)
	ds.Theme_WarnA200(id).Lazy(&c.WarnA200)
	ds.Theme_WarnA400(id).Lazy(&c.WarnA400)
	ds.Theme_WarnA700(id).Lazy(&c.WarnA700)
	ds.Theme_Yes(id).Lazy(&c.Yes)
	return &c
}

func (b *themeBuilder) Preload(rel builderWrapperI) *themeBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *themeBuilder) Organization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationID",
			relField: "Organization",
		},
	}
}

func (b *themeBuilder) ThemeForOrganization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ThemeForOrganizationID",
			relField: "ThemeForOrganization",
		},
	}
}

func (r *Fetch) Theme(ids ...int) *themeBuilder {
	return &themeBuilder{
		builder: builder[themeBuilder, *themeBuilder, Theme]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Topic has all fields from topic.
type Topic struct {
	AgendaItemID                   int
	AttachmentMeetingMediafileIDs  maybe.Maybe[[]int]
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	PollIDs                        maybe.Maybe[[]int]
	ProjectionIDs                  maybe.Maybe[[]int]
	SequentialNumber               int
	Text                           maybe.Maybe[string]
	Title                          string
	AgendaItem                     *AgendaItem
	AttachmentMeetingMediafileList []MeetingMediafile
	ListOfSpeakers                 *ListOfSpeakers
	Meeting                        *Meeting
	PollList                       []Poll
	ProjectionList                 []Projection
}

type topicBuilder struct {
	builder[topicBuilder, *topicBuilder, Topic]
}

func (b *topicBuilder) lazy(ds *Fetch, id int) *Topic {
	c := Topic{}
	ds.Topic_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Topic_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Topic_ID(id).Lazy(&c.ID)
	ds.Topic_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Topic_MeetingID(id).Lazy(&c.MeetingID)
	ds.Topic_PollIDs(id).Lazy(&c.PollIDs)
	ds.Topic_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Topic_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Topic_Text(id).Lazy(&c.Text)
	ds.Topic_Title(id).Lazy(&c.Title)
	return &c
}

func (b *topicBuilder) Preload(rel builderWrapperI) *topicBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *topicBuilder) AgendaItem() *agendaItemBuilder {
	return &agendaItemBuilder{
		builder: builder[agendaItemBuilder, *agendaItemBuilder, AgendaItem]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AgendaItemID",
			relField: "AgendaItem",
		},
	}
}

func (b *topicBuilder) AttachmentMeetingMediafileList() *meetingMediafileBuilder {
	return &meetingMediafileBuilder{
		builder: builder[meetingMediafileBuilder, *meetingMediafileBuilder, MeetingMediafile]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "AttachmentMeetingMediafileIDs",
			relField: "AttachmentMeetingMediafileList",
			many:     true,
		},
	}
}

func (b *topicBuilder) ListOfSpeakers() *listOfSpeakersBuilder {
	return &listOfSpeakersBuilder{
		builder: builder[listOfSpeakersBuilder, *listOfSpeakersBuilder, ListOfSpeakers]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ListOfSpeakersID",
			relField: "ListOfSpeakers",
		},
	}
}

func (b *topicBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *topicBuilder) PollList() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollIDs",
			relField: "PollList",
			many:     true,
		},
	}
}

func (b *topicBuilder) ProjectionList() *projectionBuilder {
	return &projectionBuilder{
		builder: builder[projectionBuilder, *projectionBuilder, Projection]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "ProjectionIDs",
			relField: "ProjectionList",
			many:     true,
		},
	}
}

func (r *Fetch) Topic(ids ...int) *topicBuilder {
	return &topicBuilder{
		builder: builder[topicBuilder, *topicBuilder, Topic]{
			ids:   ids,
			fetch: r,
		},
	}
}

// User has all fields from user.
type User struct {
	CanChangeOwnPassword        maybe.Maybe[bool]
	CommitteeIDs                maybe.Maybe[[]int]
	CommitteeManagementIDs      maybe.Maybe[[]int]
	DefaultPassword             maybe.Maybe[string]
	DefaultVoteWeight           maybe.Maybe[decimal.Decimal]
	DelegatedVoteIDs            maybe.Maybe[[]int]
	Email                       maybe.Maybe[string]
	External                    maybe.Maybe[bool]
	FirstName                   maybe.Maybe[string]
	GenderID                    maybe.Maybe[int]
	HistoryEntryIDs             maybe.Maybe[[]int]
	HistoryPositionIDs          maybe.Maybe[[]int]
	HomeCommitteeID             maybe.Maybe[int]
	ID                          int
	IsActive                    maybe.Maybe[bool]
	IsDemoUser                  maybe.Maybe[bool]
	IsPhysicalPerson            maybe.Maybe[bool]
	IsPresentInMeetingIDs       maybe.Maybe[[]int]
	LastEmailSent               maybe.Maybe[int]
	LastLogin                   maybe.Maybe[int]
	LastName                    maybe.Maybe[string]
	MeetingIDs                  maybe.Maybe[[]int]
	MeetingUserIDs              maybe.Maybe[[]int]
	MemberNumber                maybe.Maybe[string]
	OptionIDs                   maybe.Maybe[[]int]
	OrganizationID              int
	OrganizationManagementLevel maybe.Maybe[dstypes.User_OrganizationManagementLevel]
	Password                    maybe.Maybe[string]
	PollCandidateIDs            maybe.Maybe[[]int]
	PollVotedIDs                maybe.Maybe[[]int]
	Pronoun                     maybe.Maybe[string]
	SamlID                      maybe.Maybe[string]
	Title                       maybe.Maybe[string]
	Username                    string
	VoteIDs                     maybe.Maybe[[]int]
	CommitteeList               []Committee
	CommitteeManagementList     []Committee
	DelegatedVoteList           []Vote
	Gender                      *maybe.Maybe[Gender]
	HistoryEntryList            []HistoryEntry
	HistoryPositionList         []HistoryPosition
	HomeCommittee               *maybe.Maybe[Committee]
	IsPresentInMeetingList      []Meeting
	MeetingList                 []Meeting
	MeetingUserList             []MeetingUser
	OptionList                  []Option
	Organization                *Organization
	PollCandidateList           []PollCandidate
	PollVotedList               []Poll
	VoteList                    []Vote
}

type userBuilder struct {
	builder[userBuilder, *userBuilder, User]
}

func (b *userBuilder) lazy(ds *Fetch, id int) *User {
	c := User{}
	ds.User_CanChangeOwnPassword(id).Lazy(&c.CanChangeOwnPassword)
	ds.User_CommitteeIDs(id).Lazy(&c.CommitteeIDs)
	ds.User_CommitteeManagementIDs(id).Lazy(&c.CommitteeManagementIDs)
	ds.User_DefaultPassword(id).Lazy(&c.DefaultPassword)
	ds.User_DefaultVoteWeight(id).Lazy(&c.DefaultVoteWeight)
	ds.User_DelegatedVoteIDs(id).Lazy(&c.DelegatedVoteIDs)
	ds.User_Email(id).Lazy(&c.Email)
	ds.User_External(id).Lazy(&c.External)
	ds.User_FirstName(id).Lazy(&c.FirstName)
	ds.User_GenderID(id).Lazy(&c.GenderID)
	ds.User_HistoryEntryIDs(id).Lazy(&c.HistoryEntryIDs)
	ds.User_HistoryPositionIDs(id).Lazy(&c.HistoryPositionIDs)
	ds.User_HomeCommitteeID(id).Lazy(&c.HomeCommitteeID)
	ds.User_ID(id).Lazy(&c.ID)
	ds.User_IsActive(id).Lazy(&c.IsActive)
	ds.User_IsDemoUser(id).Lazy(&c.IsDemoUser)
	ds.User_IsPhysicalPerson(id).Lazy(&c.IsPhysicalPerson)
	ds.User_IsPresentInMeetingIDs(id).Lazy(&c.IsPresentInMeetingIDs)
	ds.User_LastEmailSent(id).Lazy(&c.LastEmailSent)
	ds.User_LastLogin(id).Lazy(&c.LastLogin)
	ds.User_LastName(id).Lazy(&c.LastName)
	ds.User_MeetingIDs(id).Lazy(&c.MeetingIDs)
	ds.User_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.User_MemberNumber(id).Lazy(&c.MemberNumber)
	ds.User_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.User_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.User_OrganizationManagementLevel(id).Lazy(&c.OrganizationManagementLevel)
	ds.User_Password(id).Lazy(&c.Password)
	ds.User_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
	ds.User_PollVotedIDs(id).Lazy(&c.PollVotedIDs)
	ds.User_Pronoun(id).Lazy(&c.Pronoun)
	ds.User_SamlID(id).Lazy(&c.SamlID)
	ds.User_Title(id).Lazy(&c.Title)
	ds.User_Username(id).Lazy(&c.Username)
	ds.User_VoteIDs(id).Lazy(&c.VoteIDs)
	return &c
}

func (b *userBuilder) Preload(rel builderWrapperI) *userBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *userBuilder) CommitteeList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CommitteeIDs",
			relField: "CommitteeList",
			many:     true,
		},
	}
}

func (b *userBuilder) CommitteeManagementList() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "CommitteeManagementIDs",
			relField: "CommitteeManagementList",
			many:     true,
		},
	}
}

func (b *userBuilder) DelegatedVoteList() *voteBuilder {
	return &voteBuilder{
		builder: builder[voteBuilder, *voteBuilder, Vote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DelegatedVoteIDs",
			relField: "DelegatedVoteList",
			many:     true,
		},
	}
}

func (b *userBuilder) Gender() *genderBuilder {
	return &genderBuilder{
		builder: builder[genderBuilder, *genderBuilder, Gender]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "GenderID",
			relField: "Gender",
		},
	}
}

func (b *userBuilder) HistoryEntryList() *historyEntryBuilder {
	return &historyEntryBuilder{
		builder: builder[historyEntryBuilder, *historyEntryBuilder, HistoryEntry]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HistoryEntryIDs",
			relField: "HistoryEntryList",
			many:     true,
		},
	}
}

func (b *userBuilder) HistoryPositionList() *historyPositionBuilder {
	return &historyPositionBuilder{
		builder: builder[historyPositionBuilder, *historyPositionBuilder, HistoryPosition]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HistoryPositionIDs",
			relField: "HistoryPositionList",
			many:     true,
		},
	}
}

func (b *userBuilder) HomeCommittee() *committeeBuilder {
	return &committeeBuilder{
		builder: builder[committeeBuilder, *committeeBuilder, Committee]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "HomeCommitteeID",
			relField: "HomeCommittee",
		},
	}
}

func (b *userBuilder) IsPresentInMeetingList() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "IsPresentInMeetingIDs",
			relField: "IsPresentInMeetingList",
			many:     true,
		},
	}
}

func (b *userBuilder) MeetingList() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingIDs",
			relField: "MeetingList",
			many:     true,
		},
	}
}

func (b *userBuilder) MeetingUserList() *meetingUserBuilder {
	return &meetingUserBuilder{
		builder: builder[meetingUserBuilder, *meetingUserBuilder, MeetingUser]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingUserIDs",
			relField: "MeetingUserList",
			many:     true,
		},
	}
}

func (b *userBuilder) OptionList() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OptionIDs",
			relField: "OptionList",
			many:     true,
		},
	}
}

func (b *userBuilder) Organization() *organizationBuilder {
	return &organizationBuilder{
		builder: builder[organizationBuilder, *organizationBuilder, Organization]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OrganizationID",
			relField: "Organization",
		},
	}
}

func (b *userBuilder) PollCandidateList() *pollCandidateBuilder {
	return &pollCandidateBuilder{
		builder: builder[pollCandidateBuilder, *pollCandidateBuilder, PollCandidate]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollCandidateIDs",
			relField: "PollCandidateList",
			many:     true,
		},
	}
}

func (b *userBuilder) PollVotedList() *pollBuilder {
	return &pollBuilder{
		builder: builder[pollBuilder, *pollBuilder, Poll]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "PollVotedIDs",
			relField: "PollVotedList",
			many:     true,
		},
	}
}

func (b *userBuilder) VoteList() *voteBuilder {
	return &voteBuilder{
		builder: builder[voteBuilder, *voteBuilder, Vote]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "VoteIDs",
			relField: "VoteList",
			many:     true,
		},
	}
}

func (r *Fetch) User(ids ...int) *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			ids:   ids,
			fetch: r,
		},
	}
}

// Vote has all fields from vote.
type Vote struct {
	DelegatedUserID maybe.Maybe[int]
	ID              int
	MeetingID       int
	OptionID        int
	UserID          maybe.Maybe[int]
	UserToken       string
	Value           string
	Weight          decimal.Decimal
	DelegatedUser   *maybe.Maybe[User]
	Meeting         *Meeting
	Option          *Option
	User            *maybe.Maybe[User]
}

type voteBuilder struct {
	builder[voteBuilder, *voteBuilder, Vote]
}

func (b *voteBuilder) lazy(ds *Fetch, id int) *Vote {
	c := Vote{}
	ds.Vote_DelegatedUserID(id).Lazy(&c.DelegatedUserID)
	ds.Vote_ID(id).Lazy(&c.ID)
	ds.Vote_MeetingID(id).Lazy(&c.MeetingID)
	ds.Vote_OptionID(id).Lazy(&c.OptionID)
	ds.Vote_UserID(id).Lazy(&c.UserID)
	ds.Vote_UserToken(id).Lazy(&c.UserToken)
	ds.Vote_Value(id).Lazy(&c.Value)
	ds.Vote_Weight(id).Lazy(&c.Weight)
	return &c
}

func (b *voteBuilder) Preload(rel builderWrapperI) *voteBuilder {
	b.builder.Preload(rel)
	return b
}

func (b *voteBuilder) DelegatedUser() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "DelegatedUserID",
			relField: "DelegatedUser",
		},
	}
}

func (b *voteBuilder) Meeting() *meetingBuilder {
	return &meetingBuilder{
		builder: builder[meetingBuilder, *meetingBuilder, Meeting]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "MeetingID",
			relField: "Meeting",
		},
	}
}

func (b *voteBuilder) Option() *optionBuilder {
	return &optionBuilder{
		builder: builder[optionBuilder, *optionBuilder, Option]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "OptionID",
			relField: "Option",
		},
	}
}

func (b *voteBuilder) User() *userBuilder {
	return &userBuilder{
		builder: builder[userBuilder, *userBuilder, User]{
			fetch:    b.fetch,
			parent:   b,
			idField:  "UserID",
			relField: "User",
		},
	}
}

func (r *Fetch) Vote(ids ...int) *voteBuilder {
	return &voteBuilder{
		builder: builder[voteBuilder, *voteBuilder, Vote]{
			ids:   ids,
			fetch: r,
		},
	}
}
