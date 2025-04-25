// Code generated from models.yml DO NOT EDIT.
package dsmodels

import (
	"encoding/json"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// ActionWorker has all fields from action_worker.
type ActionWorker struct {
	Created   int
	ID        int
	Name      string
	Result    json.RawMessage
	State     string
	Timestamp int
	UserID    int
	fetch     *Fetch
}

func (c *ActionWorker) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ActionWorker_Created(id).Lazy(&c.Created)
	ds.ActionWorker_ID(id).Lazy(&c.ID)
	ds.ActionWorker_Name(id).Lazy(&c.Name)
	ds.ActionWorker_Result(id).Lazy(&c.Result)
	ds.ActionWorker_State(id).Lazy(&c.State)
	ds.ActionWorker_Timestamp(id).Lazy(&c.Timestamp)
	ds.ActionWorker_UserID(id).Lazy(&c.UserID)
}

func (r *Fetch) ActionWorker(id int) *ValueCollection[ActionWorker, *ActionWorker] {
	return &ValueCollection[ActionWorker, *ActionWorker]{
		id:    id,
		fetch: r,
	}
}

// AgendaItem has all fields from agenda_item.
type AgendaItem struct {
	ChildIDs        []int
	Closed          bool
	Comment         string
	ContentObjectID string
	Duration        int
	ID              int
	IsHidden        bool
	IsInternal      bool
	ItemNumber      string
	Level           int
	MeetingID       int
	ParentID        dsfetch.Maybe[int]
	ProjectionIDs   []int
	TagIDs          []int
	Type            string
	Weight          int
	childList       *RelationList[AgendaItem, *AgendaItem]
	meeting         *ValueCollection[Meeting, *Meeting]
	parent          *MaybeRelation[AgendaItem, *AgendaItem]
	projectionList  *RelationList[Projection, *Projection]
	tagList         *RelationList[Tag, *Tag]
	fetch           *Fetch
}

func (c *AgendaItem) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *AgendaItem) ChildList() *RelationList[AgendaItem, *AgendaItem] {
	if c.childList == nil {
		refs := make([]*ValueCollection[AgendaItem, *AgendaItem], len(c.ChildIDs))
		for i, id := range c.ChildIDs {
			refs[i] = &ValueCollection[AgendaItem, *AgendaItem]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.childList = &RelationList[AgendaItem, *AgendaItem]{refs}
	}
	return c.childList
}

func (c *AgendaItem) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *AgendaItem) Parent() *MaybeRelation[AgendaItem, *AgendaItem] {
	if c.parent == nil {
		var ref dsfetch.Maybe[*ValueCollection[AgendaItem, *AgendaItem]]
		id, hasValue := c.ParentID.Value()
		if hasValue {
			value := &ValueCollection[AgendaItem, *AgendaItem]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.parent = &MaybeRelation[AgendaItem, *AgendaItem]{ref}
	}
	return c.parent
}

func (c *AgendaItem) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *AgendaItem) TagList() *RelationList[Tag, *Tag] {
	if c.tagList == nil {
		refs := make([]*ValueCollection[Tag, *Tag], len(c.TagIDs))
		for i, id := range c.TagIDs {
			refs[i] = &ValueCollection[Tag, *Tag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.tagList = &RelationList[Tag, *Tag]{refs}
	}
	return c.tagList
}

func (r *Fetch) AgendaItem(id int) *ValueCollection[AgendaItem, *AgendaItem] {
	return &ValueCollection[AgendaItem, *AgendaItem]{
		id:    id,
		fetch: r,
	}
}

// Assignment has all fields from assignment.
type Assignment struct {
	AgendaItemID                   dsfetch.Maybe[int]
	AttachmentMeetingMediafileIDs  []int
	CandidateIDs                   []int
	DefaultPollDescription         string
	Description                    string
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	NumberPollCandidates           bool
	OpenPosts                      int
	Phase                          string
	PollIDs                        []int
	ProjectionIDs                  []int
	SequentialNumber               int
	TagIDs                         []int
	Title                          string
	agendaItem                     *MaybeRelation[AgendaItem, *AgendaItem]
	attachmentMeetingMediafileList *RelationList[MeetingMediafile, *MeetingMediafile]
	candidateList                  *RelationList[AssignmentCandidate, *AssignmentCandidate]
	listOfSpeakers                 *ValueCollection[ListOfSpeakers, *ListOfSpeakers]
	meeting                        *ValueCollection[Meeting, *Meeting]
	pollList                       *RelationList[Poll, *Poll]
	projectionList                 *RelationList[Projection, *Projection]
	tagList                        *RelationList[Tag, *Tag]
	fetch                          *Fetch
}

func (c *Assignment) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Assignment_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Assignment_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Assignment_CandidateIDs(id).Lazy(&c.CandidateIDs)
	ds.Assignment_DefaultPollDescription(id).Lazy(&c.DefaultPollDescription)
	ds.Assignment_Description(id).Lazy(&c.Description)
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
}

func (c *Assignment) AgendaItem() *MaybeRelation[AgendaItem, *AgendaItem] {
	if c.agendaItem == nil {
		var ref dsfetch.Maybe[*ValueCollection[AgendaItem, *AgendaItem]]
		id, hasValue := c.AgendaItemID.Value()
		if hasValue {
			value := &ValueCollection[AgendaItem, *AgendaItem]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.agendaItem = &MaybeRelation[AgendaItem, *AgendaItem]{ref}
	}
	return c.agendaItem
}

func (c *Assignment) AttachmentMeetingMediafileList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.attachmentMeetingMediafileList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.AttachmentMeetingMediafileIDs))
		for i, id := range c.AttachmentMeetingMediafileIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.attachmentMeetingMediafileList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.attachmentMeetingMediafileList
}

func (c *Assignment) CandidateList() *RelationList[AssignmentCandidate, *AssignmentCandidate] {
	if c.candidateList == nil {
		refs := make([]*ValueCollection[AssignmentCandidate, *AssignmentCandidate], len(c.CandidateIDs))
		for i, id := range c.CandidateIDs {
			refs[i] = &ValueCollection[AssignmentCandidate, *AssignmentCandidate]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.candidateList = &RelationList[AssignmentCandidate, *AssignmentCandidate]{refs}
	}
	return c.candidateList
}

func (c *Assignment) ListOfSpeakers() *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		c.listOfSpeakers = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
			id:    c.ListOfSpeakersID,
			fetch: c.fetch,
		}
	}
	return c.listOfSpeakers
}

func (c *Assignment) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Assignment) PollList() *RelationList[Poll, *Poll] {
	if c.pollList == nil {
		refs := make([]*ValueCollection[Poll, *Poll], len(c.PollIDs))
		for i, id := range c.PollIDs {
			refs[i] = &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollList = &RelationList[Poll, *Poll]{refs}
	}
	return c.pollList
}

func (c *Assignment) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *Assignment) TagList() *RelationList[Tag, *Tag] {
	if c.tagList == nil {
		refs := make([]*ValueCollection[Tag, *Tag], len(c.TagIDs))
		for i, id := range c.TagIDs {
			refs[i] = &ValueCollection[Tag, *Tag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.tagList = &RelationList[Tag, *Tag]{refs}
	}
	return c.tagList
}

func (r *Fetch) Assignment(id int) *ValueCollection[Assignment, *Assignment] {
	return &ValueCollection[Assignment, *Assignment]{
		id:    id,
		fetch: r,
	}
}

// AssignmentCandidate has all fields from assignment_candidate.
type AssignmentCandidate struct {
	AssignmentID  int
	ID            int
	MeetingID     int
	MeetingUserID dsfetch.Maybe[int]
	Weight        int
	assignment    *ValueCollection[Assignment, *Assignment]
	meeting       *ValueCollection[Meeting, *Meeting]
	meetingUser   *MaybeRelation[MeetingUser, *MeetingUser]
	fetch         *Fetch
}

func (c *AssignmentCandidate) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.AssignmentCandidate_AssignmentID(id).Lazy(&c.AssignmentID)
	ds.AssignmentCandidate_ID(id).Lazy(&c.ID)
	ds.AssignmentCandidate_MeetingID(id).Lazy(&c.MeetingID)
	ds.AssignmentCandidate_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.AssignmentCandidate_Weight(id).Lazy(&c.Weight)
}

func (c *AssignmentCandidate) Assignment() *ValueCollection[Assignment, *Assignment] {
	if c.assignment == nil {
		c.assignment = &ValueCollection[Assignment, *Assignment]{
			id:    c.AssignmentID,
			fetch: c.fetch,
		}
	}
	return c.assignment
}

func (c *AssignmentCandidate) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *AssignmentCandidate) MeetingUser() *MaybeRelation[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingUser, *MeetingUser]]
		id, hasValue := c.MeetingUserID.Value()
		if hasValue {
			value := &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.meetingUser = &MaybeRelation[MeetingUser, *MeetingUser]{ref}
	}
	return c.meetingUser
}

func (r *Fetch) AssignmentCandidate(id int) *ValueCollection[AssignmentCandidate, *AssignmentCandidate] {
	return &ValueCollection[AssignmentCandidate, *AssignmentCandidate]{
		id:    id,
		fetch: r,
	}
}

// ChatGroup has all fields from chat_group.
type ChatGroup struct {
	ChatMessageIDs  []int
	ID              int
	MeetingID       int
	Name            string
	ReadGroupIDs    []int
	Weight          int
	WriteGroupIDs   []int
	chatMessageList *RelationList[ChatMessage, *ChatMessage]
	meeting         *ValueCollection[Meeting, *Meeting]
	readGroupList   *RelationList[Group, *Group]
	writeGroupList  *RelationList[Group, *Group]
	fetch           *Fetch
}

func (c *ChatGroup) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ChatGroup_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.ChatGroup_ID(id).Lazy(&c.ID)
	ds.ChatGroup_MeetingID(id).Lazy(&c.MeetingID)
	ds.ChatGroup_Name(id).Lazy(&c.Name)
	ds.ChatGroup_ReadGroupIDs(id).Lazy(&c.ReadGroupIDs)
	ds.ChatGroup_Weight(id).Lazy(&c.Weight)
	ds.ChatGroup_WriteGroupIDs(id).Lazy(&c.WriteGroupIDs)
}

func (c *ChatGroup) ChatMessageList() *RelationList[ChatMessage, *ChatMessage] {
	if c.chatMessageList == nil {
		refs := make([]*ValueCollection[ChatMessage, *ChatMessage], len(c.ChatMessageIDs))
		for i, id := range c.ChatMessageIDs {
			refs[i] = &ValueCollection[ChatMessage, *ChatMessage]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.chatMessageList = &RelationList[ChatMessage, *ChatMessage]{refs}
	}
	return c.chatMessageList
}

func (c *ChatGroup) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *ChatGroup) ReadGroupList() *RelationList[Group, *Group] {
	if c.readGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.ReadGroupIDs))
		for i, id := range c.ReadGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.readGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.readGroupList
}

func (c *ChatGroup) WriteGroupList() *RelationList[Group, *Group] {
	if c.writeGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.WriteGroupIDs))
		for i, id := range c.WriteGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.writeGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.writeGroupList
}

func (r *Fetch) ChatGroup(id int) *ValueCollection[ChatGroup, *ChatGroup] {
	return &ValueCollection[ChatGroup, *ChatGroup]{
		id:    id,
		fetch: r,
	}
}

// ChatMessage has all fields from chat_message.
type ChatMessage struct {
	ChatGroupID   int
	Content       string
	Created       int
	ID            int
	MeetingID     int
	MeetingUserID dsfetch.Maybe[int]
	chatGroup     *ValueCollection[ChatGroup, *ChatGroup]
	meeting       *ValueCollection[Meeting, *Meeting]
	meetingUser   *MaybeRelation[MeetingUser, *MeetingUser]
	fetch         *Fetch
}

func (c *ChatMessage) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ChatMessage_ChatGroupID(id).Lazy(&c.ChatGroupID)
	ds.ChatMessage_Content(id).Lazy(&c.Content)
	ds.ChatMessage_Created(id).Lazy(&c.Created)
	ds.ChatMessage_ID(id).Lazy(&c.ID)
	ds.ChatMessage_MeetingID(id).Lazy(&c.MeetingID)
	ds.ChatMessage_MeetingUserID(id).Lazy(&c.MeetingUserID)
}

func (c *ChatMessage) ChatGroup() *ValueCollection[ChatGroup, *ChatGroup] {
	if c.chatGroup == nil {
		c.chatGroup = &ValueCollection[ChatGroup, *ChatGroup]{
			id:    c.ChatGroupID,
			fetch: c.fetch,
		}
	}
	return c.chatGroup
}

func (c *ChatMessage) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *ChatMessage) MeetingUser() *MaybeRelation[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingUser, *MeetingUser]]
		id, hasValue := c.MeetingUserID.Value()
		if hasValue {
			value := &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.meetingUser = &MaybeRelation[MeetingUser, *MeetingUser]{ref}
	}
	return c.meetingUser
}

func (r *Fetch) ChatMessage(id int) *ValueCollection[ChatMessage, *ChatMessage] {
	return &ValueCollection[ChatMessage, *ChatMessage]{
		id:    id,
		fetch: r,
	}
}

// Committee has all fields from committee.
type Committee struct {
	DefaultMeetingID                    dsfetch.Maybe[int]
	Description                         string
	ExternalID                          string
	ForwardToCommitteeIDs               []int
	ID                                  int
	ManagerIDs                          []int
	MeetingIDs                          []int
	Name                                string
	OrganizationID                      int
	OrganizationTagIDs                  []int
	ReceiveForwardingsFromCommitteeIDs  []int
	UserIDs                             []int
	defaultMeeting                      *MaybeRelation[Meeting, *Meeting]
	forwardToCommitteeList              *RelationList[Committee, *Committee]
	managerList                         *RelationList[User, *User]
	meetingList                         *RelationList[Meeting, *Meeting]
	organization                        *ValueCollection[Organization, *Organization]
	organizationTagList                 *RelationList[OrganizationTag, *OrganizationTag]
	receiveForwardingsFromCommitteeList *RelationList[Committee, *Committee]
	userList                            *RelationList[User, *User]
	fetch                               *Fetch
}

func (c *Committee) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Committee_DefaultMeetingID(id).Lazy(&c.DefaultMeetingID)
	ds.Committee_Description(id).Lazy(&c.Description)
	ds.Committee_ExternalID(id).Lazy(&c.ExternalID)
	ds.Committee_ForwardToCommitteeIDs(id).Lazy(&c.ForwardToCommitteeIDs)
	ds.Committee_ID(id).Lazy(&c.ID)
	ds.Committee_ManagerIDs(id).Lazy(&c.ManagerIDs)
	ds.Committee_MeetingIDs(id).Lazy(&c.MeetingIDs)
	ds.Committee_Name(id).Lazy(&c.Name)
	ds.Committee_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Committee_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Committee_ReceiveForwardingsFromCommitteeIDs(id).Lazy(&c.ReceiveForwardingsFromCommitteeIDs)
	ds.Committee_UserIDs(id).Lazy(&c.UserIDs)
}

func (c *Committee) DefaultMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.defaultMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.DefaultMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.defaultMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.defaultMeeting
}

func (c *Committee) ForwardToCommitteeList() *RelationList[Committee, *Committee] {
	if c.forwardToCommitteeList == nil {
		refs := make([]*ValueCollection[Committee, *Committee], len(c.ForwardToCommitteeIDs))
		for i, id := range c.ForwardToCommitteeIDs {
			refs[i] = &ValueCollection[Committee, *Committee]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.forwardToCommitteeList = &RelationList[Committee, *Committee]{refs}
	}
	return c.forwardToCommitteeList
}

func (c *Committee) ManagerList() *RelationList[User, *User] {
	if c.managerList == nil {
		refs := make([]*ValueCollection[User, *User], len(c.ManagerIDs))
		for i, id := range c.ManagerIDs {
			refs[i] = &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.managerList = &RelationList[User, *User]{refs}
	}
	return c.managerList
}

func (c *Committee) MeetingList() *RelationList[Meeting, *Meeting] {
	if c.meetingList == nil {
		refs := make([]*ValueCollection[Meeting, *Meeting], len(c.MeetingIDs))
		for i, id := range c.MeetingIDs {
			refs[i] = &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingList = &RelationList[Meeting, *Meeting]{refs}
	}
	return c.meetingList
}

func (c *Committee) Organization() *ValueCollection[Organization, *Organization] {
	if c.organization == nil {
		c.organization = &ValueCollection[Organization, *Organization]{
			id:    c.OrganizationID,
			fetch: c.fetch,
		}
	}
	return c.organization
}

func (c *Committee) OrganizationTagList() *RelationList[OrganizationTag, *OrganizationTag] {
	if c.organizationTagList == nil {
		refs := make([]*ValueCollection[OrganizationTag, *OrganizationTag], len(c.OrganizationTagIDs))
		for i, id := range c.OrganizationTagIDs {
			refs[i] = &ValueCollection[OrganizationTag, *OrganizationTag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.organizationTagList = &RelationList[OrganizationTag, *OrganizationTag]{refs}
	}
	return c.organizationTagList
}

func (c *Committee) ReceiveForwardingsFromCommitteeList() *RelationList[Committee, *Committee] {
	if c.receiveForwardingsFromCommitteeList == nil {
		refs := make([]*ValueCollection[Committee, *Committee], len(c.ReceiveForwardingsFromCommitteeIDs))
		for i, id := range c.ReceiveForwardingsFromCommitteeIDs {
			refs[i] = &ValueCollection[Committee, *Committee]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.receiveForwardingsFromCommitteeList = &RelationList[Committee, *Committee]{refs}
	}
	return c.receiveForwardingsFromCommitteeList
}

func (c *Committee) UserList() *RelationList[User, *User] {
	if c.userList == nil {
		refs := make([]*ValueCollection[User, *User], len(c.UserIDs))
		for i, id := range c.UserIDs {
			refs[i] = &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.userList = &RelationList[User, *User]{refs}
	}
	return c.userList
}

func (r *Fetch) Committee(id int) *ValueCollection[Committee, *Committee] {
	return &ValueCollection[Committee, *Committee]{
		id:    id,
		fetch: r,
	}
}

// Gender has all fields from gender.
type Gender struct {
	ID             int
	Name           string
	OrganizationID int
	UserIDs        []int
	organization   *ValueCollection[Organization, *Organization]
	userList       *RelationList[User, *User]
	fetch          *Fetch
}

func (c *Gender) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Gender_ID(id).Lazy(&c.ID)
	ds.Gender_Name(id).Lazy(&c.Name)
	ds.Gender_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Gender_UserIDs(id).Lazy(&c.UserIDs)
}

func (c *Gender) Organization() *ValueCollection[Organization, *Organization] {
	if c.organization == nil {
		c.organization = &ValueCollection[Organization, *Organization]{
			id:    c.OrganizationID,
			fetch: c.fetch,
		}
	}
	return c.organization
}

func (c *Gender) UserList() *RelationList[User, *User] {
	if c.userList == nil {
		refs := make([]*ValueCollection[User, *User], len(c.UserIDs))
		for i, id := range c.UserIDs {
			refs[i] = &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.userList = &RelationList[User, *User]{refs}
	}
	return c.userList
}

func (r *Fetch) Gender(id int) *ValueCollection[Gender, *Gender] {
	return &ValueCollection[Gender, *Gender]{
		id:    id,
		fetch: r,
	}
}

// Group has all fields from group.
type Group struct {
	AdminGroupForMeetingID                   dsfetch.Maybe[int]
	AnonymousGroupForMeetingID               dsfetch.Maybe[int]
	DefaultGroupForMeetingID                 dsfetch.Maybe[int]
	ExternalID                               string
	ID                                       int
	MeetingID                                int
	MeetingMediafileAccessGroupIDs           []int
	MeetingMediafileInheritedAccessGroupIDs  []int
	MeetingUserIDs                           []int
	Name                                     string
	Permissions                              []string
	PollIDs                                  []int
	ReadChatGroupIDs                         []int
	ReadCommentSectionIDs                    []int
	UsedAsAssignmentPollDefaultID            dsfetch.Maybe[int]
	UsedAsMotionPollDefaultID                dsfetch.Maybe[int]
	UsedAsPollDefaultID                      dsfetch.Maybe[int]
	UsedAsTopicPollDefaultID                 dsfetch.Maybe[int]
	Weight                                   int
	WriteChatGroupIDs                        []int
	WriteCommentSectionIDs                   []int
	adminGroupForMeeting                     *MaybeRelation[Meeting, *Meeting]
	anonymousGroupForMeeting                 *MaybeRelation[Meeting, *Meeting]
	defaultGroupForMeeting                   *MaybeRelation[Meeting, *Meeting]
	meeting                                  *ValueCollection[Meeting, *Meeting]
	meetingMediafileAccessGroupList          *RelationList[MeetingMediafile, *MeetingMediafile]
	meetingMediafileInheritedAccessGroupList *RelationList[MeetingMediafile, *MeetingMediafile]
	meetingUserList                          *RelationList[MeetingUser, *MeetingUser]
	pollList                                 *RelationList[Poll, *Poll]
	readChatGroupList                        *RelationList[ChatGroup, *ChatGroup]
	readCommentSectionList                   *RelationList[MotionCommentSection, *MotionCommentSection]
	usedAsAssignmentPollDefault              *MaybeRelation[Meeting, *Meeting]
	usedAsMotionPollDefault                  *MaybeRelation[Meeting, *Meeting]
	usedAsPollDefault                        *MaybeRelation[Meeting, *Meeting]
	usedAsTopicPollDefault                   *MaybeRelation[Meeting, *Meeting]
	writeChatGroupList                       *RelationList[ChatGroup, *ChatGroup]
	writeCommentSectionList                  *RelationList[MotionCommentSection, *MotionCommentSection]
	fetch                                    *Fetch
}

func (c *Group) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Group) AdminGroupForMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.adminGroupForMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.AdminGroupForMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.adminGroupForMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.adminGroupForMeeting
}

func (c *Group) AnonymousGroupForMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.anonymousGroupForMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.AnonymousGroupForMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.anonymousGroupForMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.anonymousGroupForMeeting
}

func (c *Group) DefaultGroupForMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.defaultGroupForMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.DefaultGroupForMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.defaultGroupForMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.defaultGroupForMeeting
}

func (c *Group) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Group) MeetingMediafileAccessGroupList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.meetingMediafileAccessGroupList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.MeetingMediafileAccessGroupIDs))
		for i, id := range c.MeetingMediafileAccessGroupIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingMediafileAccessGroupList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.meetingMediafileAccessGroupList
}

func (c *Group) MeetingMediafileInheritedAccessGroupList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.meetingMediafileInheritedAccessGroupList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.MeetingMediafileInheritedAccessGroupIDs))
		for i, id := range c.MeetingMediafileInheritedAccessGroupIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingMediafileInheritedAccessGroupList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.meetingMediafileInheritedAccessGroupList
}

func (c *Group) MeetingUserList() *RelationList[MeetingUser, *MeetingUser] {
	if c.meetingUserList == nil {
		refs := make([]*ValueCollection[MeetingUser, *MeetingUser], len(c.MeetingUserIDs))
		for i, id := range c.MeetingUserIDs {
			refs[i] = &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingUserList = &RelationList[MeetingUser, *MeetingUser]{refs}
	}
	return c.meetingUserList
}

func (c *Group) PollList() *RelationList[Poll, *Poll] {
	if c.pollList == nil {
		refs := make([]*ValueCollection[Poll, *Poll], len(c.PollIDs))
		for i, id := range c.PollIDs {
			refs[i] = &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollList = &RelationList[Poll, *Poll]{refs}
	}
	return c.pollList
}

func (c *Group) ReadChatGroupList() *RelationList[ChatGroup, *ChatGroup] {
	if c.readChatGroupList == nil {
		refs := make([]*ValueCollection[ChatGroup, *ChatGroup], len(c.ReadChatGroupIDs))
		for i, id := range c.ReadChatGroupIDs {
			refs[i] = &ValueCollection[ChatGroup, *ChatGroup]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.readChatGroupList = &RelationList[ChatGroup, *ChatGroup]{refs}
	}
	return c.readChatGroupList
}

func (c *Group) ReadCommentSectionList() *RelationList[MotionCommentSection, *MotionCommentSection] {
	if c.readCommentSectionList == nil {
		refs := make([]*ValueCollection[MotionCommentSection, *MotionCommentSection], len(c.ReadCommentSectionIDs))
		for i, id := range c.ReadCommentSectionIDs {
			refs[i] = &ValueCollection[MotionCommentSection, *MotionCommentSection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.readCommentSectionList = &RelationList[MotionCommentSection, *MotionCommentSection]{refs}
	}
	return c.readCommentSectionList
}

func (c *Group) UsedAsAssignmentPollDefault() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsAssignmentPollDefault == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsAssignmentPollDefaultID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsAssignmentPollDefault = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsAssignmentPollDefault
}

func (c *Group) UsedAsMotionPollDefault() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsMotionPollDefault == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsMotionPollDefaultID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsMotionPollDefault = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsMotionPollDefault
}

func (c *Group) UsedAsPollDefault() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsPollDefault == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsPollDefaultID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsPollDefault = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsPollDefault
}

func (c *Group) UsedAsTopicPollDefault() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsTopicPollDefault == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsTopicPollDefaultID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsTopicPollDefault = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsTopicPollDefault
}

func (c *Group) WriteChatGroupList() *RelationList[ChatGroup, *ChatGroup] {
	if c.writeChatGroupList == nil {
		refs := make([]*ValueCollection[ChatGroup, *ChatGroup], len(c.WriteChatGroupIDs))
		for i, id := range c.WriteChatGroupIDs {
			refs[i] = &ValueCollection[ChatGroup, *ChatGroup]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.writeChatGroupList = &RelationList[ChatGroup, *ChatGroup]{refs}
	}
	return c.writeChatGroupList
}

func (c *Group) WriteCommentSectionList() *RelationList[MotionCommentSection, *MotionCommentSection] {
	if c.writeCommentSectionList == nil {
		refs := make([]*ValueCollection[MotionCommentSection, *MotionCommentSection], len(c.WriteCommentSectionIDs))
		for i, id := range c.WriteCommentSectionIDs {
			refs[i] = &ValueCollection[MotionCommentSection, *MotionCommentSection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.writeCommentSectionList = &RelationList[MotionCommentSection, *MotionCommentSection]{refs}
	}
	return c.writeCommentSectionList
}

func (r *Fetch) Group(id int) *ValueCollection[Group, *Group] {
	return &ValueCollection[Group, *Group]{
		id:    id,
		fetch: r,
	}
}

// ImportPreview has all fields from import_preview.
type ImportPreview struct {
	Created int
	ID      int
	Name    string
	Result  json.RawMessage
	State   string
	fetch   *Fetch
}

func (c *ImportPreview) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ImportPreview_Created(id).Lazy(&c.Created)
	ds.ImportPreview_ID(id).Lazy(&c.ID)
	ds.ImportPreview_Name(id).Lazy(&c.Name)
	ds.ImportPreview_Result(id).Lazy(&c.Result)
	ds.ImportPreview_State(id).Lazy(&c.State)
}

func (r *Fetch) ImportPreview(id int) *ValueCollection[ImportPreview, *ImportPreview] {
	return &ValueCollection[ImportPreview, *ImportPreview]{
		id:    id,
		fetch: r,
	}
}

// ListOfSpeakers has all fields from list_of_speakers.
type ListOfSpeakers struct {
	Closed                           bool
	ContentObjectID                  string
	ID                               int
	MeetingID                        int
	ModeratorNotes                   string
	ProjectionIDs                    []int
	SequentialNumber                 int
	SpeakerIDs                       []int
	StructureLevelListOfSpeakersIDs  []int
	meeting                          *ValueCollection[Meeting, *Meeting]
	projectionList                   *RelationList[Projection, *Projection]
	speakerList                      *RelationList[Speaker, *Speaker]
	structureLevelListOfSpeakersList *RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]
	fetch                            *Fetch
}

func (c *ListOfSpeakers) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ListOfSpeakers_Closed(id).Lazy(&c.Closed)
	ds.ListOfSpeakers_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.ListOfSpeakers_ID(id).Lazy(&c.ID)
	ds.ListOfSpeakers_MeetingID(id).Lazy(&c.MeetingID)
	ds.ListOfSpeakers_ModeratorNotes(id).Lazy(&c.ModeratorNotes)
	ds.ListOfSpeakers_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.ListOfSpeakers_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.ListOfSpeakers_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.ListOfSpeakers_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
}

func (c *ListOfSpeakers) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *ListOfSpeakers) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *ListOfSpeakers) SpeakerList() *RelationList[Speaker, *Speaker] {
	if c.speakerList == nil {
		refs := make([]*ValueCollection[Speaker, *Speaker], len(c.SpeakerIDs))
		for i, id := range c.SpeakerIDs {
			refs[i] = &ValueCollection[Speaker, *Speaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.speakerList = &RelationList[Speaker, *Speaker]{refs}
	}
	return c.speakerList
}

func (c *ListOfSpeakers) StructureLevelListOfSpeakersList() *RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	if c.structureLevelListOfSpeakersList == nil {
		refs := make([]*ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers], len(c.StructureLevelListOfSpeakersIDs))
		for i, id := range c.StructureLevelListOfSpeakersIDs {
			refs[i] = &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.structureLevelListOfSpeakersList = &RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{refs}
	}
	return c.structureLevelListOfSpeakersList
}

func (r *Fetch) ListOfSpeakers(id int) *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	return &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
		id:    id,
		fetch: r,
	}
}

// Mediafile has all fields from mediafile.
type Mediafile struct {
	ChildIDs                            []int
	CreateTimestamp                     int
	Filename                            string
	Filesize                            int
	ID                                  int
	IsDirectory                         bool
	MeetingMediafileIDs                 []int
	Mimetype                            string
	OwnerID                             string
	ParentID                            dsfetch.Maybe[int]
	PdfInformation                      json.RawMessage
	PublishedToMeetingsInOrganizationID dsfetch.Maybe[int]
	Title                               string
	Token                               string
	childList                           *RelationList[Mediafile, *Mediafile]
	meetingMediafileList                *RelationList[MeetingMediafile, *MeetingMediafile]
	parent                              *MaybeRelation[Mediafile, *Mediafile]
	publishedToMeetingsInOrganization   *MaybeRelation[Organization, *Organization]
	fetch                               *Fetch
}

func (c *Mediafile) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Mediafile) ChildList() *RelationList[Mediafile, *Mediafile] {
	if c.childList == nil {
		refs := make([]*ValueCollection[Mediafile, *Mediafile], len(c.ChildIDs))
		for i, id := range c.ChildIDs {
			refs[i] = &ValueCollection[Mediafile, *Mediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.childList = &RelationList[Mediafile, *Mediafile]{refs}
	}
	return c.childList
}

func (c *Mediafile) MeetingMediafileList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.meetingMediafileList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.MeetingMediafileIDs))
		for i, id := range c.MeetingMediafileIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingMediafileList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.meetingMediafileList
}

func (c *Mediafile) Parent() *MaybeRelation[Mediafile, *Mediafile] {
	if c.parent == nil {
		var ref dsfetch.Maybe[*ValueCollection[Mediafile, *Mediafile]]
		id, hasValue := c.ParentID.Value()
		if hasValue {
			value := &ValueCollection[Mediafile, *Mediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.parent = &MaybeRelation[Mediafile, *Mediafile]{ref}
	}
	return c.parent
}

func (c *Mediafile) PublishedToMeetingsInOrganization() *MaybeRelation[Organization, *Organization] {
	if c.publishedToMeetingsInOrganization == nil {
		var ref dsfetch.Maybe[*ValueCollection[Organization, *Organization]]
		id, hasValue := c.PublishedToMeetingsInOrganizationID.Value()
		if hasValue {
			value := &ValueCollection[Organization, *Organization]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.publishedToMeetingsInOrganization = &MaybeRelation[Organization, *Organization]{ref}
	}
	return c.publishedToMeetingsInOrganization
}

func (r *Fetch) Mediafile(id int) *ValueCollection[Mediafile, *Mediafile] {
	return &ValueCollection[Mediafile, *Mediafile]{
		id:    id,
		fetch: r,
	}
}

// Meeting has all fields from meeting.
type Meeting struct {
	AdminGroupID                                 dsfetch.Maybe[int]
	AgendaEnableNumbering                        bool
	AgendaItemCreation                           string
	AgendaItemIDs                                []int
	AgendaNewItemsDefaultVisibility              string
	AgendaNumberPrefix                           string
	AgendaNumeralSystem                          string
	AgendaShowInternalItemsOnProjector           bool
	AgendaShowSubtitles                          bool
	AgendaShowTopicNavigationOnDetailView        bool
	AllProjectionIDs                             []int
	AnonymousGroupID                             dsfetch.Maybe[int]
	ApplauseEnable                               bool
	ApplauseMaxAmount                            int
	ApplauseMinAmount                            int
	ApplauseParticleImageUrl                     string
	ApplauseShowLevel                            bool
	ApplauseTimeout                              int
	ApplauseType                                 string
	AssignmentCandidateIDs                       []int
	AssignmentIDs                                []int
	AssignmentPollAddCandidatesToListOfSpeakers  bool
	AssignmentPollBallotPaperNumber              int
	AssignmentPollBallotPaperSelection           string
	AssignmentPollDefaultBackend                 string
	AssignmentPollDefaultGroupIDs                []int
	AssignmentPollDefaultMethod                  string
	AssignmentPollDefaultOnehundredPercentBase   string
	AssignmentPollDefaultType                    string
	AssignmentPollEnableMaxVotesPerOption        bool
	AssignmentPollSortPollResultByVotes          bool
	AssignmentsExportPreamble                    string
	AssignmentsExportTitle                       string
	ChatGroupIDs                                 []int
	ChatMessageIDs                               []int
	CommitteeID                                  int
	ConferenceAutoConnect                        bool
	ConferenceAutoConnectNextSpeakers            int
	ConferenceEnableHelpdesk                     bool
	ConferenceLosRestriction                     bool
	ConferenceOpenMicrophone                     bool
	ConferenceOpenVideo                          bool
	ConferenceShow                               bool
	ConferenceStreamPosterUrl                    string
	ConferenceStreamUrl                          string
	CustomTranslations                           json.RawMessage
	DefaultGroupID                               int
	DefaultMeetingForCommitteeID                 dsfetch.Maybe[int]
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
	Description                                  string
	EnableAnonymous                              bool
	EndTime                                      int
	ExportCsvEncoding                            string
	ExportCsvSeparator                           string
	ExportPdfFontsize                            int
	ExportPdfLineHeight                          float32
	ExportPdfPageMarginBottom                    int
	ExportPdfPageMarginLeft                      int
	ExportPdfPageMarginRight                     int
	ExportPdfPageMarginTop                       int
	ExportPdfPagenumberAlignment                 string
	ExportPdfPagesize                            string
	ExternalID                                   string
	FontBoldID                                   dsfetch.Maybe[int]
	FontBoldItalicID                             dsfetch.Maybe[int]
	FontChyronSpeakerNameID                      dsfetch.Maybe[int]
	FontItalicID                                 dsfetch.Maybe[int]
	FontMonospaceID                              dsfetch.Maybe[int]
	FontProjectorH1ID                            dsfetch.Maybe[int]
	FontProjectorH2ID                            dsfetch.Maybe[int]
	FontRegularID                                dsfetch.Maybe[int]
	ForwardedMotionIDs                           []int
	GroupIDs                                     []int
	ID                                           int
	ImportedAt                                   int
	IsActiveInOrganizationID                     dsfetch.Maybe[int]
	IsArchivedInOrganizationID                   dsfetch.Maybe[int]
	JitsiDomain                                  string
	JitsiRoomName                                string
	JitsiRoomPassword                            string
	Language                                     string
	ListOfSpeakersAllowMultipleSpeakers          bool
	ListOfSpeakersAmountLastOnProjector          int
	ListOfSpeakersAmountNextOnProjector          int
	ListOfSpeakersCanCreatePointOfOrderForOthers bool
	ListOfSpeakersCanSetContributionSelf         bool
	ListOfSpeakersClosingDisablesPointOfOrder    bool
	ListOfSpeakersCountdownID                    dsfetch.Maybe[int]
	ListOfSpeakersCoupleCountdown                bool
	ListOfSpeakersDefaultStructureLevelTime      int
	ListOfSpeakersEnableInterposedQuestion       bool
	ListOfSpeakersEnablePointOfOrderCategories   bool
	ListOfSpeakersEnablePointOfOrderSpeakers     bool
	ListOfSpeakersEnableProContraSpeech          bool
	ListOfSpeakersHideContributionCount          bool
	ListOfSpeakersIDs                            []int
	ListOfSpeakersInitiallyClosed                bool
	ListOfSpeakersInterventionTime               int
	ListOfSpeakersPresentUsersOnly               bool
	ListOfSpeakersShowAmountOfSpeakersOnSlide    bool
	ListOfSpeakersShowFirstContribution          bool
	ListOfSpeakersSpeakerNoteForEveryone         bool
	Location                                     string
	LockedFromInside                             bool
	LogoPdfBallotPaperID                         dsfetch.Maybe[int]
	LogoPdfFooterLID                             dsfetch.Maybe[int]
	LogoPdfFooterRID                             dsfetch.Maybe[int]
	LogoPdfHeaderLID                             dsfetch.Maybe[int]
	LogoPdfHeaderRID                             dsfetch.Maybe[int]
	LogoProjectorHeaderID                        dsfetch.Maybe[int]
	LogoProjectorMainID                          dsfetch.Maybe[int]
	LogoWebHeaderID                              dsfetch.Maybe[int]
	MediafileIDs                                 []int
	MeetingMediafileIDs                          []int
	MeetingUserIDs                               []int
	MotionBlockIDs                               []int
	MotionCategoryIDs                            []int
	MotionChangeRecommendationIDs                []int
	MotionCommentIDs                             []int
	MotionCommentSectionIDs                      []int
	MotionEditorIDs                              []int
	MotionIDs                                    []int
	MotionPollBallotPaperNumber                  int
	MotionPollBallotPaperSelection               string
	MotionPollDefaultBackend                     string
	MotionPollDefaultGroupIDs                    []int
	MotionPollDefaultMethod                      string
	MotionPollDefaultOnehundredPercentBase       string
	MotionPollDefaultType                        string
	MotionPollProjectionMaxColumns               int
	MotionPollProjectionNameOrderFirst           string
	MotionStateIDs                               []int
	MotionSubmitterIDs                           []int
	MotionWorkflowIDs                            []int
	MotionWorkingGroupSpeakerIDs                 []int
	MotionsAmendmentsEnabled                     bool
	MotionsAmendmentsInMainList                  bool
	MotionsAmendmentsMultipleParagraphs          bool
	MotionsAmendmentsOfAmendments                bool
	MotionsAmendmentsPrefix                      string
	MotionsAmendmentsTextMode                    string
	MotionsBlockSlideColumns                     int
	MotionsCreateEnableAdditionalSubmitterText   bool
	MotionsDefaultAmendmentWorkflowID            int
	MotionsDefaultLineNumbering                  string
	MotionsDefaultSorting                        string
	MotionsDefaultWorkflowID                     int
	MotionsEnableEditor                          bool
	MotionsEnableOriginMotionDisplay             bool
	MotionsEnableReasonOnProjector               bool
	MotionsEnableRecommendationOnProjector       bool
	MotionsEnableSideboxOnProjector              bool
	MotionsEnableTextOnProjector                 bool
	MotionsEnableWorkingGroupSpeaker             bool
	MotionsExportFollowRecommendation            bool
	MotionsExportPreamble                        string
	MotionsExportSubmitterRecommendation         bool
	MotionsExportTitle                           string
	MotionsHideMetadataBackground                bool
	MotionsLineLength                            int
	MotionsNumberMinDigits                       int
	MotionsNumberType                            string
	MotionsNumberWithBlank                       bool
	MotionsOriginMotionToggleDefault             bool
	MotionsPreamble                              string
	MotionsReasonRequired                        bool
	MotionsRecommendationTextMode                string
	MotionsRecommendationsBy                     string
	MotionsShowReferringMotions                  bool
	MotionsShowSequentialNumber                  bool
	MotionsSupportersMinAmount                   int
	Name                                         string
	OptionIDs                                    []int
	OrganizationTagIDs                           []int
	PersonalNoteIDs                              []int
	PointOfOrderCategoryIDs                      []int
	PollBallotPaperNumber                        int
	PollBallotPaperSelection                     string
	PollCandidateIDs                             []int
	PollCandidateListIDs                         []int
	PollCountdownID                              dsfetch.Maybe[int]
	PollCoupleCountdown                          bool
	PollDefaultBackend                           string
	PollDefaultGroupIDs                          []int
	PollDefaultMethod                            string
	PollDefaultOnehundredPercentBase             string
	PollDefaultType                              string
	PollIDs                                      []int
	PollSortPollResultByVotes                    bool
	PresentUserIDs                               []int
	ProjectionIDs                                []int
	ProjectorCountdownDefaultTime                int
	ProjectorCountdownIDs                        []int
	ProjectorCountdownWarningTime                int
	ProjectorIDs                                 []int
	ProjectorMessageIDs                          []int
	ReferenceProjectorID                         int
	SpeakerIDs                                   []int
	StartTime                                    int
	StructureLevelIDs                            []int
	StructureLevelListOfSpeakersIDs              []int
	TagIDs                                       []int
	TemplateForOrganizationID                    dsfetch.Maybe[int]
	TopicIDs                                     []int
	TopicPollDefaultGroupIDs                     []int
	UserIDs                                      []int
	UsersAllowSelfSetPresent                     bool
	UsersEmailBody                               string
	UsersEmailReplyto                            string
	UsersEmailSender                             string
	UsersEmailSubject                            string
	UsersEnablePresenceView                      bool
	UsersEnableVoteDelegations                   bool
	UsersEnableVoteWeight                        bool
	UsersForbidDelegatorAsSubmitter              bool
	UsersForbidDelegatorAsSupporter              bool
	UsersForbidDelegatorInListOfSpeakers         bool
	UsersForbidDelegatorToVote                   bool
	UsersPdfWelcometext                          string
	UsersPdfWelcometitle                         string
	UsersPdfWlanEncryption                       string
	UsersPdfWlanPassword                         string
	UsersPdfWlanSsid                             string
	VoteIDs                                      []int
	WelcomeText                                  string
	WelcomeTitle                                 string
	adminGroup                                   *MaybeRelation[Group, *Group]
	agendaItemList                               *RelationList[AgendaItem, *AgendaItem]
	allProjectionList                            *RelationList[Projection, *Projection]
	anonymousGroup                               *MaybeRelation[Group, *Group]
	assignmentCandidateList                      *RelationList[AssignmentCandidate, *AssignmentCandidate]
	assignmentList                               *RelationList[Assignment, *Assignment]
	assignmentPollDefaultGroupList               *RelationList[Group, *Group]
	chatGroupList                                *RelationList[ChatGroup, *ChatGroup]
	chatMessageList                              *RelationList[ChatMessage, *ChatMessage]
	committee                                    *ValueCollection[Committee, *Committee]
	defaultGroup                                 *ValueCollection[Group, *Group]
	defaultMeetingForCommittee                   *MaybeRelation[Committee, *Committee]
	defaultProjectorAgendaItemListList           *RelationList[Projector, *Projector]
	defaultProjectorAmendmentList                *RelationList[Projector, *Projector]
	defaultProjectorAssignmentList               *RelationList[Projector, *Projector]
	defaultProjectorAssignmentPollList           *RelationList[Projector, *Projector]
	defaultProjectorCountdownList                *RelationList[Projector, *Projector]
	defaultProjectorCurrentLosList               *RelationList[Projector, *Projector]
	defaultProjectorListOfSpeakersList           *RelationList[Projector, *Projector]
	defaultProjectorMediafileList                *RelationList[Projector, *Projector]
	defaultProjectorMessageList                  *RelationList[Projector, *Projector]
	defaultProjectorMotionBlockList              *RelationList[Projector, *Projector]
	defaultProjectorMotionList                   *RelationList[Projector, *Projector]
	defaultProjectorMotionPollList               *RelationList[Projector, *Projector]
	defaultProjectorPollList                     *RelationList[Projector, *Projector]
	defaultProjectorTopicList                    *RelationList[Projector, *Projector]
	fontBold                                     *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontBoldItalic                               *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontChyronSpeakerName                        *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontItalic                                   *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontMonospace                                *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontProjectorH1                              *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontProjectorH2                              *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	fontRegular                                  *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	forwardedMotionList                          *RelationList[Motion, *Motion]
	groupList                                    *RelationList[Group, *Group]
	isActiveInOrganization                       *MaybeRelation[Organization, *Organization]
	isArchivedInOrganization                     *MaybeRelation[Organization, *Organization]
	listOfSpeakersCountdown                      *MaybeRelation[ProjectorCountdown, *ProjectorCountdown]
	listOfSpeakersList                           *RelationList[ListOfSpeakers, *ListOfSpeakers]
	logoPdfBallotPaper                           *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoPdfFooterL                               *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoPdfFooterR                               *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoPdfHeaderL                               *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoPdfHeaderR                               *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoProjectorHeader                          *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoProjectorMain                            *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	logoWebHeader                                *MaybeRelation[MeetingMediafile, *MeetingMediafile]
	mediafileList                                *RelationList[Mediafile, *Mediafile]
	meetingMediafileList                         *RelationList[MeetingMediafile, *MeetingMediafile]
	meetingUserList                              *RelationList[MeetingUser, *MeetingUser]
	motionBlockList                              *RelationList[MotionBlock, *MotionBlock]
	motionCategoryList                           *RelationList[MotionCategory, *MotionCategory]
	motionChangeRecommendationList               *RelationList[MotionChangeRecommendation, *MotionChangeRecommendation]
	motionCommentList                            *RelationList[MotionComment, *MotionComment]
	motionCommentSectionList                     *RelationList[MotionCommentSection, *MotionCommentSection]
	motionEditorList                             *RelationList[MotionEditor, *MotionEditor]
	motionList                                   *RelationList[Motion, *Motion]
	motionPollDefaultGroupList                   *RelationList[Group, *Group]
	motionStateList                              *RelationList[MotionState, *MotionState]
	motionSubmitterList                          *RelationList[MotionSubmitter, *MotionSubmitter]
	motionWorkflowList                           *RelationList[MotionWorkflow, *MotionWorkflow]
	motionWorkingGroupSpeakerList                *RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]
	motionsDefaultAmendmentWorkflow              *ValueCollection[MotionWorkflow, *MotionWorkflow]
	motionsDefaultWorkflow                       *ValueCollection[MotionWorkflow, *MotionWorkflow]
	optionList                                   *RelationList[Option, *Option]
	organizationTagList                          *RelationList[OrganizationTag, *OrganizationTag]
	personalNoteList                             *RelationList[PersonalNote, *PersonalNote]
	pointOfOrderCategoryList                     *RelationList[PointOfOrderCategory, *PointOfOrderCategory]
	pollCandidateList                            *RelationList[PollCandidate, *PollCandidate]
	pollCandidateListList                        *RelationList[PollCandidateList, *PollCandidateList]
	pollCountdown                                *MaybeRelation[ProjectorCountdown, *ProjectorCountdown]
	pollDefaultGroupList                         *RelationList[Group, *Group]
	pollList                                     *RelationList[Poll, *Poll]
	presentUserList                              *RelationList[User, *User]
	projectionList                               *RelationList[Projection, *Projection]
	projectorCountdownList                       *RelationList[ProjectorCountdown, *ProjectorCountdown]
	projectorList                                *RelationList[Projector, *Projector]
	projectorMessageList                         *RelationList[ProjectorMessage, *ProjectorMessage]
	referenceProjector                           *ValueCollection[Projector, *Projector]
	speakerList                                  *RelationList[Speaker, *Speaker]
	structureLevelList                           *RelationList[StructureLevel, *StructureLevel]
	structureLevelListOfSpeakersList             *RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]
	tagList                                      *RelationList[Tag, *Tag]
	templateForOrganization                      *MaybeRelation[Organization, *Organization]
	topicList                                    *RelationList[Topic, *Topic]
	topicPollDefaultGroupList                    *RelationList[Group, *Group]
	voteList                                     *RelationList[Vote, *Vote]
	fetch                                        *Fetch
}

func (c *Meeting) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
	ds.Meeting_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.Meeting_StartTime(id).Lazy(&c.StartTime)
	ds.Meeting_StructureLevelIDs(id).Lazy(&c.StructureLevelIDs)
	ds.Meeting_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
	ds.Meeting_TagIDs(id).Lazy(&c.TagIDs)
	ds.Meeting_TemplateForOrganizationID(id).Lazy(&c.TemplateForOrganizationID)
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
}

func (c *Meeting) AdminGroup() *MaybeRelation[Group, *Group] {
	if c.adminGroup == nil {
		var ref dsfetch.Maybe[*ValueCollection[Group, *Group]]
		id, hasValue := c.AdminGroupID.Value()
		if hasValue {
			value := &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.adminGroup = &MaybeRelation[Group, *Group]{ref}
	}
	return c.adminGroup
}

func (c *Meeting) AgendaItemList() *RelationList[AgendaItem, *AgendaItem] {
	if c.agendaItemList == nil {
		refs := make([]*ValueCollection[AgendaItem, *AgendaItem], len(c.AgendaItemIDs))
		for i, id := range c.AgendaItemIDs {
			refs[i] = &ValueCollection[AgendaItem, *AgendaItem]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.agendaItemList = &RelationList[AgendaItem, *AgendaItem]{refs}
	}
	return c.agendaItemList
}

func (c *Meeting) AllProjectionList() *RelationList[Projection, *Projection] {
	if c.allProjectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.AllProjectionIDs))
		for i, id := range c.AllProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.allProjectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.allProjectionList
}

func (c *Meeting) AnonymousGroup() *MaybeRelation[Group, *Group] {
	if c.anonymousGroup == nil {
		var ref dsfetch.Maybe[*ValueCollection[Group, *Group]]
		id, hasValue := c.AnonymousGroupID.Value()
		if hasValue {
			value := &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.anonymousGroup = &MaybeRelation[Group, *Group]{ref}
	}
	return c.anonymousGroup
}

func (c *Meeting) AssignmentCandidateList() *RelationList[AssignmentCandidate, *AssignmentCandidate] {
	if c.assignmentCandidateList == nil {
		refs := make([]*ValueCollection[AssignmentCandidate, *AssignmentCandidate], len(c.AssignmentCandidateIDs))
		for i, id := range c.AssignmentCandidateIDs {
			refs[i] = &ValueCollection[AssignmentCandidate, *AssignmentCandidate]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.assignmentCandidateList = &RelationList[AssignmentCandidate, *AssignmentCandidate]{refs}
	}
	return c.assignmentCandidateList
}

func (c *Meeting) AssignmentList() *RelationList[Assignment, *Assignment] {
	if c.assignmentList == nil {
		refs := make([]*ValueCollection[Assignment, *Assignment], len(c.AssignmentIDs))
		for i, id := range c.AssignmentIDs {
			refs[i] = &ValueCollection[Assignment, *Assignment]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.assignmentList = &RelationList[Assignment, *Assignment]{refs}
	}
	return c.assignmentList
}

func (c *Meeting) AssignmentPollDefaultGroupList() *RelationList[Group, *Group] {
	if c.assignmentPollDefaultGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.AssignmentPollDefaultGroupIDs))
		for i, id := range c.AssignmentPollDefaultGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.assignmentPollDefaultGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.assignmentPollDefaultGroupList
}

func (c *Meeting) ChatGroupList() *RelationList[ChatGroup, *ChatGroup] {
	if c.chatGroupList == nil {
		refs := make([]*ValueCollection[ChatGroup, *ChatGroup], len(c.ChatGroupIDs))
		for i, id := range c.ChatGroupIDs {
			refs[i] = &ValueCollection[ChatGroup, *ChatGroup]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.chatGroupList = &RelationList[ChatGroup, *ChatGroup]{refs}
	}
	return c.chatGroupList
}

func (c *Meeting) ChatMessageList() *RelationList[ChatMessage, *ChatMessage] {
	if c.chatMessageList == nil {
		refs := make([]*ValueCollection[ChatMessage, *ChatMessage], len(c.ChatMessageIDs))
		for i, id := range c.ChatMessageIDs {
			refs[i] = &ValueCollection[ChatMessage, *ChatMessage]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.chatMessageList = &RelationList[ChatMessage, *ChatMessage]{refs}
	}
	return c.chatMessageList
}

func (c *Meeting) Committee() *ValueCollection[Committee, *Committee] {
	if c.committee == nil {
		c.committee = &ValueCollection[Committee, *Committee]{
			id:    c.CommitteeID,
			fetch: c.fetch,
		}
	}
	return c.committee
}

func (c *Meeting) DefaultGroup() *ValueCollection[Group, *Group] {
	if c.defaultGroup == nil {
		c.defaultGroup = &ValueCollection[Group, *Group]{
			id:    c.DefaultGroupID,
			fetch: c.fetch,
		}
	}
	return c.defaultGroup
}

func (c *Meeting) DefaultMeetingForCommittee() *MaybeRelation[Committee, *Committee] {
	if c.defaultMeetingForCommittee == nil {
		var ref dsfetch.Maybe[*ValueCollection[Committee, *Committee]]
		id, hasValue := c.DefaultMeetingForCommitteeID.Value()
		if hasValue {
			value := &ValueCollection[Committee, *Committee]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.defaultMeetingForCommittee = &MaybeRelation[Committee, *Committee]{ref}
	}
	return c.defaultMeetingForCommittee
}

func (c *Meeting) DefaultProjectorAgendaItemListList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorAgendaItemListList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorAgendaItemListIDs))
		for i, id := range c.DefaultProjectorAgendaItemListIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorAgendaItemListList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorAgendaItemListList
}

func (c *Meeting) DefaultProjectorAmendmentList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorAmendmentList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorAmendmentIDs))
		for i, id := range c.DefaultProjectorAmendmentIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorAmendmentList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorAmendmentList
}

func (c *Meeting) DefaultProjectorAssignmentList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorAssignmentList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorAssignmentIDs))
		for i, id := range c.DefaultProjectorAssignmentIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorAssignmentList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorAssignmentList
}

func (c *Meeting) DefaultProjectorAssignmentPollList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorAssignmentPollList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorAssignmentPollIDs))
		for i, id := range c.DefaultProjectorAssignmentPollIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorAssignmentPollList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorAssignmentPollList
}

func (c *Meeting) DefaultProjectorCountdownList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorCountdownList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorCountdownIDs))
		for i, id := range c.DefaultProjectorCountdownIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorCountdownList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorCountdownList
}

func (c *Meeting) DefaultProjectorCurrentLosList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorCurrentLosList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorCurrentLosIDs))
		for i, id := range c.DefaultProjectorCurrentLosIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorCurrentLosList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorCurrentLosList
}

func (c *Meeting) DefaultProjectorListOfSpeakersList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorListOfSpeakersList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorListOfSpeakersIDs))
		for i, id := range c.DefaultProjectorListOfSpeakersIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorListOfSpeakersList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorListOfSpeakersList
}

func (c *Meeting) DefaultProjectorMediafileList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorMediafileList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorMediafileIDs))
		for i, id := range c.DefaultProjectorMediafileIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorMediafileList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorMediafileList
}

func (c *Meeting) DefaultProjectorMessageList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorMessageList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorMessageIDs))
		for i, id := range c.DefaultProjectorMessageIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorMessageList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorMessageList
}

func (c *Meeting) DefaultProjectorMotionBlockList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorMotionBlockList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorMotionBlockIDs))
		for i, id := range c.DefaultProjectorMotionBlockIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorMotionBlockList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorMotionBlockList
}

func (c *Meeting) DefaultProjectorMotionList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorMotionList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorMotionIDs))
		for i, id := range c.DefaultProjectorMotionIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorMotionList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorMotionList
}

func (c *Meeting) DefaultProjectorMotionPollList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorMotionPollList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorMotionPollIDs))
		for i, id := range c.DefaultProjectorMotionPollIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorMotionPollList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorMotionPollList
}

func (c *Meeting) DefaultProjectorPollList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorPollList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorPollIDs))
		for i, id := range c.DefaultProjectorPollIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorPollList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorPollList
}

func (c *Meeting) DefaultProjectorTopicList() *RelationList[Projector, *Projector] {
	if c.defaultProjectorTopicList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.DefaultProjectorTopicIDs))
		for i, id := range c.DefaultProjectorTopicIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.defaultProjectorTopicList = &RelationList[Projector, *Projector]{refs}
	}
	return c.defaultProjectorTopicList
}

func (c *Meeting) FontBold() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontBold == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontBoldID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontBold = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontBold
}

func (c *Meeting) FontBoldItalic() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontBoldItalic == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontBoldItalicID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontBoldItalic = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontBoldItalic
}

func (c *Meeting) FontChyronSpeakerName() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontChyronSpeakerName == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontChyronSpeakerNameID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontChyronSpeakerName = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontChyronSpeakerName
}

func (c *Meeting) FontItalic() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontItalic == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontItalicID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontItalic = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontItalic
}

func (c *Meeting) FontMonospace() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontMonospace == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontMonospaceID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontMonospace = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontMonospace
}

func (c *Meeting) FontProjectorH1() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontProjectorH1 == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontProjectorH1ID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontProjectorH1 = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontProjectorH1
}

func (c *Meeting) FontProjectorH2() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontProjectorH2 == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontProjectorH2ID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontProjectorH2 = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontProjectorH2
}

func (c *Meeting) FontRegular() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.fontRegular == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.FontRegularID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.fontRegular = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.fontRegular
}

func (c *Meeting) ForwardedMotionList() *RelationList[Motion, *Motion] {
	if c.forwardedMotionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.ForwardedMotionIDs))
		for i, id := range c.ForwardedMotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.forwardedMotionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.forwardedMotionList
}

func (c *Meeting) GroupList() *RelationList[Group, *Group] {
	if c.groupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.GroupIDs))
		for i, id := range c.GroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.groupList = &RelationList[Group, *Group]{refs}
	}
	return c.groupList
}

func (c *Meeting) IsActiveInOrganization() *MaybeRelation[Organization, *Organization] {
	if c.isActiveInOrganization == nil {
		var ref dsfetch.Maybe[*ValueCollection[Organization, *Organization]]
		id, hasValue := c.IsActiveInOrganizationID.Value()
		if hasValue {
			value := &ValueCollection[Organization, *Organization]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.isActiveInOrganization = &MaybeRelation[Organization, *Organization]{ref}
	}
	return c.isActiveInOrganization
}

func (c *Meeting) IsArchivedInOrganization() *MaybeRelation[Organization, *Organization] {
	if c.isArchivedInOrganization == nil {
		var ref dsfetch.Maybe[*ValueCollection[Organization, *Organization]]
		id, hasValue := c.IsArchivedInOrganizationID.Value()
		if hasValue {
			value := &ValueCollection[Organization, *Organization]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.isArchivedInOrganization = &MaybeRelation[Organization, *Organization]{ref}
	}
	return c.isArchivedInOrganization
}

func (c *Meeting) ListOfSpeakersCountdown() *MaybeRelation[ProjectorCountdown, *ProjectorCountdown] {
	if c.listOfSpeakersCountdown == nil {
		var ref dsfetch.Maybe[*ValueCollection[ProjectorCountdown, *ProjectorCountdown]]
		id, hasValue := c.ListOfSpeakersCountdownID.Value()
		if hasValue {
			value := &ValueCollection[ProjectorCountdown, *ProjectorCountdown]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.listOfSpeakersCountdown = &MaybeRelation[ProjectorCountdown, *ProjectorCountdown]{ref}
	}
	return c.listOfSpeakersCountdown
}

func (c *Meeting) ListOfSpeakersList() *RelationList[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakersList == nil {
		refs := make([]*ValueCollection[ListOfSpeakers, *ListOfSpeakers], len(c.ListOfSpeakersIDs))
		for i, id := range c.ListOfSpeakersIDs {
			refs[i] = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.listOfSpeakersList = &RelationList[ListOfSpeakers, *ListOfSpeakers]{refs}
	}
	return c.listOfSpeakersList
}

func (c *Meeting) LogoPdfBallotPaper() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoPdfBallotPaper == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoPdfBallotPaperID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoPdfBallotPaper = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoPdfBallotPaper
}

func (c *Meeting) LogoPdfFooterL() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoPdfFooterL == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoPdfFooterLID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoPdfFooterL = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoPdfFooterL
}

func (c *Meeting) LogoPdfFooterR() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoPdfFooterR == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoPdfFooterRID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoPdfFooterR = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoPdfFooterR
}

func (c *Meeting) LogoPdfHeaderL() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoPdfHeaderL == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoPdfHeaderLID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoPdfHeaderL = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoPdfHeaderL
}

func (c *Meeting) LogoPdfHeaderR() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoPdfHeaderR == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoPdfHeaderRID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoPdfHeaderR = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoPdfHeaderR
}

func (c *Meeting) LogoProjectorHeader() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoProjectorHeader == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoProjectorHeaderID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoProjectorHeader = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoProjectorHeader
}

func (c *Meeting) LogoProjectorMain() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoProjectorMain == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoProjectorMainID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoProjectorMain = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoProjectorMain
}

func (c *Meeting) LogoWebHeader() *MaybeRelation[MeetingMediafile, *MeetingMediafile] {
	if c.logoWebHeader == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingMediafile, *MeetingMediafile]]
		id, hasValue := c.LogoWebHeaderID.Value()
		if hasValue {
			value := &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.logoWebHeader = &MaybeRelation[MeetingMediafile, *MeetingMediafile]{ref}
	}
	return c.logoWebHeader
}

func (c *Meeting) MediafileList() *RelationList[Mediafile, *Mediafile] {
	if c.mediafileList == nil {
		refs := make([]*ValueCollection[Mediafile, *Mediafile], len(c.MediafileIDs))
		for i, id := range c.MediafileIDs {
			refs[i] = &ValueCollection[Mediafile, *Mediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.mediafileList = &RelationList[Mediafile, *Mediafile]{refs}
	}
	return c.mediafileList
}

func (c *Meeting) MeetingMediafileList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.meetingMediafileList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.MeetingMediafileIDs))
		for i, id := range c.MeetingMediafileIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingMediafileList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.meetingMediafileList
}

func (c *Meeting) MeetingUserList() *RelationList[MeetingUser, *MeetingUser] {
	if c.meetingUserList == nil {
		refs := make([]*ValueCollection[MeetingUser, *MeetingUser], len(c.MeetingUserIDs))
		for i, id := range c.MeetingUserIDs {
			refs[i] = &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingUserList = &RelationList[MeetingUser, *MeetingUser]{refs}
	}
	return c.meetingUserList
}

func (c *Meeting) MotionBlockList() *RelationList[MotionBlock, *MotionBlock] {
	if c.motionBlockList == nil {
		refs := make([]*ValueCollection[MotionBlock, *MotionBlock], len(c.MotionBlockIDs))
		for i, id := range c.MotionBlockIDs {
			refs[i] = &ValueCollection[MotionBlock, *MotionBlock]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionBlockList = &RelationList[MotionBlock, *MotionBlock]{refs}
	}
	return c.motionBlockList
}

func (c *Meeting) MotionCategoryList() *RelationList[MotionCategory, *MotionCategory] {
	if c.motionCategoryList == nil {
		refs := make([]*ValueCollection[MotionCategory, *MotionCategory], len(c.MotionCategoryIDs))
		for i, id := range c.MotionCategoryIDs {
			refs[i] = &ValueCollection[MotionCategory, *MotionCategory]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionCategoryList = &RelationList[MotionCategory, *MotionCategory]{refs}
	}
	return c.motionCategoryList
}

func (c *Meeting) MotionChangeRecommendationList() *RelationList[MotionChangeRecommendation, *MotionChangeRecommendation] {
	if c.motionChangeRecommendationList == nil {
		refs := make([]*ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation], len(c.MotionChangeRecommendationIDs))
		for i, id := range c.MotionChangeRecommendationIDs {
			refs[i] = &ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionChangeRecommendationList = &RelationList[MotionChangeRecommendation, *MotionChangeRecommendation]{refs}
	}
	return c.motionChangeRecommendationList
}

func (c *Meeting) MotionCommentList() *RelationList[MotionComment, *MotionComment] {
	if c.motionCommentList == nil {
		refs := make([]*ValueCollection[MotionComment, *MotionComment], len(c.MotionCommentIDs))
		for i, id := range c.MotionCommentIDs {
			refs[i] = &ValueCollection[MotionComment, *MotionComment]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionCommentList = &RelationList[MotionComment, *MotionComment]{refs}
	}
	return c.motionCommentList
}

func (c *Meeting) MotionCommentSectionList() *RelationList[MotionCommentSection, *MotionCommentSection] {
	if c.motionCommentSectionList == nil {
		refs := make([]*ValueCollection[MotionCommentSection, *MotionCommentSection], len(c.MotionCommentSectionIDs))
		for i, id := range c.MotionCommentSectionIDs {
			refs[i] = &ValueCollection[MotionCommentSection, *MotionCommentSection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionCommentSectionList = &RelationList[MotionCommentSection, *MotionCommentSection]{refs}
	}
	return c.motionCommentSectionList
}

func (c *Meeting) MotionEditorList() *RelationList[MotionEditor, *MotionEditor] {
	if c.motionEditorList == nil {
		refs := make([]*ValueCollection[MotionEditor, *MotionEditor], len(c.MotionEditorIDs))
		for i, id := range c.MotionEditorIDs {
			refs[i] = &ValueCollection[MotionEditor, *MotionEditor]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionEditorList = &RelationList[MotionEditor, *MotionEditor]{refs}
	}
	return c.motionEditorList
}

func (c *Meeting) MotionList() *RelationList[Motion, *Motion] {
	if c.motionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.MotionIDs))
		for i, id := range c.MotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.motionList
}

func (c *Meeting) MotionPollDefaultGroupList() *RelationList[Group, *Group] {
	if c.motionPollDefaultGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.MotionPollDefaultGroupIDs))
		for i, id := range c.MotionPollDefaultGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionPollDefaultGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.motionPollDefaultGroupList
}

func (c *Meeting) MotionStateList() *RelationList[MotionState, *MotionState] {
	if c.motionStateList == nil {
		refs := make([]*ValueCollection[MotionState, *MotionState], len(c.MotionStateIDs))
		for i, id := range c.MotionStateIDs {
			refs[i] = &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionStateList = &RelationList[MotionState, *MotionState]{refs}
	}
	return c.motionStateList
}

func (c *Meeting) MotionSubmitterList() *RelationList[MotionSubmitter, *MotionSubmitter] {
	if c.motionSubmitterList == nil {
		refs := make([]*ValueCollection[MotionSubmitter, *MotionSubmitter], len(c.MotionSubmitterIDs))
		for i, id := range c.MotionSubmitterIDs {
			refs[i] = &ValueCollection[MotionSubmitter, *MotionSubmitter]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionSubmitterList = &RelationList[MotionSubmitter, *MotionSubmitter]{refs}
	}
	return c.motionSubmitterList
}

func (c *Meeting) MotionWorkflowList() *RelationList[MotionWorkflow, *MotionWorkflow] {
	if c.motionWorkflowList == nil {
		refs := make([]*ValueCollection[MotionWorkflow, *MotionWorkflow], len(c.MotionWorkflowIDs))
		for i, id := range c.MotionWorkflowIDs {
			refs[i] = &ValueCollection[MotionWorkflow, *MotionWorkflow]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionWorkflowList = &RelationList[MotionWorkflow, *MotionWorkflow]{refs}
	}
	return c.motionWorkflowList
}

func (c *Meeting) MotionWorkingGroupSpeakerList() *RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker] {
	if c.motionWorkingGroupSpeakerList == nil {
		refs := make([]*ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker], len(c.MotionWorkingGroupSpeakerIDs))
		for i, id := range c.MotionWorkingGroupSpeakerIDs {
			refs[i] = &ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionWorkingGroupSpeakerList = &RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{refs}
	}
	return c.motionWorkingGroupSpeakerList
}

func (c *Meeting) MotionsDefaultAmendmentWorkflow() *ValueCollection[MotionWorkflow, *MotionWorkflow] {
	if c.motionsDefaultAmendmentWorkflow == nil {
		c.motionsDefaultAmendmentWorkflow = &ValueCollection[MotionWorkflow, *MotionWorkflow]{
			id:    c.MotionsDefaultAmendmentWorkflowID,
			fetch: c.fetch,
		}
	}
	return c.motionsDefaultAmendmentWorkflow
}

func (c *Meeting) MotionsDefaultWorkflow() *ValueCollection[MotionWorkflow, *MotionWorkflow] {
	if c.motionsDefaultWorkflow == nil {
		c.motionsDefaultWorkflow = &ValueCollection[MotionWorkflow, *MotionWorkflow]{
			id:    c.MotionsDefaultWorkflowID,
			fetch: c.fetch,
		}
	}
	return c.motionsDefaultWorkflow
}

func (c *Meeting) OptionList() *RelationList[Option, *Option] {
	if c.optionList == nil {
		refs := make([]*ValueCollection[Option, *Option], len(c.OptionIDs))
		for i, id := range c.OptionIDs {
			refs[i] = &ValueCollection[Option, *Option]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.optionList = &RelationList[Option, *Option]{refs}
	}
	return c.optionList
}

func (c *Meeting) OrganizationTagList() *RelationList[OrganizationTag, *OrganizationTag] {
	if c.organizationTagList == nil {
		refs := make([]*ValueCollection[OrganizationTag, *OrganizationTag], len(c.OrganizationTagIDs))
		for i, id := range c.OrganizationTagIDs {
			refs[i] = &ValueCollection[OrganizationTag, *OrganizationTag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.organizationTagList = &RelationList[OrganizationTag, *OrganizationTag]{refs}
	}
	return c.organizationTagList
}

func (c *Meeting) PersonalNoteList() *RelationList[PersonalNote, *PersonalNote] {
	if c.personalNoteList == nil {
		refs := make([]*ValueCollection[PersonalNote, *PersonalNote], len(c.PersonalNoteIDs))
		for i, id := range c.PersonalNoteIDs {
			refs[i] = &ValueCollection[PersonalNote, *PersonalNote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.personalNoteList = &RelationList[PersonalNote, *PersonalNote]{refs}
	}
	return c.personalNoteList
}

func (c *Meeting) PointOfOrderCategoryList() *RelationList[PointOfOrderCategory, *PointOfOrderCategory] {
	if c.pointOfOrderCategoryList == nil {
		refs := make([]*ValueCollection[PointOfOrderCategory, *PointOfOrderCategory], len(c.PointOfOrderCategoryIDs))
		for i, id := range c.PointOfOrderCategoryIDs {
			refs[i] = &ValueCollection[PointOfOrderCategory, *PointOfOrderCategory]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pointOfOrderCategoryList = &RelationList[PointOfOrderCategory, *PointOfOrderCategory]{refs}
	}
	return c.pointOfOrderCategoryList
}

func (c *Meeting) PollCandidateList() *RelationList[PollCandidate, *PollCandidate] {
	if c.pollCandidateList == nil {
		refs := make([]*ValueCollection[PollCandidate, *PollCandidate], len(c.PollCandidateIDs))
		for i, id := range c.PollCandidateIDs {
			refs[i] = &ValueCollection[PollCandidate, *PollCandidate]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollCandidateList = &RelationList[PollCandidate, *PollCandidate]{refs}
	}
	return c.pollCandidateList
}

func (c *Meeting) PollCandidateListList() *RelationList[PollCandidateList, *PollCandidateList] {
	if c.pollCandidateListList == nil {
		refs := make([]*ValueCollection[PollCandidateList, *PollCandidateList], len(c.PollCandidateListIDs))
		for i, id := range c.PollCandidateListIDs {
			refs[i] = &ValueCollection[PollCandidateList, *PollCandidateList]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollCandidateListList = &RelationList[PollCandidateList, *PollCandidateList]{refs}
	}
	return c.pollCandidateListList
}

func (c *Meeting) PollCountdown() *MaybeRelation[ProjectorCountdown, *ProjectorCountdown] {
	if c.pollCountdown == nil {
		var ref dsfetch.Maybe[*ValueCollection[ProjectorCountdown, *ProjectorCountdown]]
		id, hasValue := c.PollCountdownID.Value()
		if hasValue {
			value := &ValueCollection[ProjectorCountdown, *ProjectorCountdown]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.pollCountdown = &MaybeRelation[ProjectorCountdown, *ProjectorCountdown]{ref}
	}
	return c.pollCountdown
}

func (c *Meeting) PollDefaultGroupList() *RelationList[Group, *Group] {
	if c.pollDefaultGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.PollDefaultGroupIDs))
		for i, id := range c.PollDefaultGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollDefaultGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.pollDefaultGroupList
}

func (c *Meeting) PollList() *RelationList[Poll, *Poll] {
	if c.pollList == nil {
		refs := make([]*ValueCollection[Poll, *Poll], len(c.PollIDs))
		for i, id := range c.PollIDs {
			refs[i] = &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollList = &RelationList[Poll, *Poll]{refs}
	}
	return c.pollList
}

func (c *Meeting) PresentUserList() *RelationList[User, *User] {
	if c.presentUserList == nil {
		refs := make([]*ValueCollection[User, *User], len(c.PresentUserIDs))
		for i, id := range c.PresentUserIDs {
			refs[i] = &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.presentUserList = &RelationList[User, *User]{refs}
	}
	return c.presentUserList
}

func (c *Meeting) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *Meeting) ProjectorCountdownList() *RelationList[ProjectorCountdown, *ProjectorCountdown] {
	if c.projectorCountdownList == nil {
		refs := make([]*ValueCollection[ProjectorCountdown, *ProjectorCountdown], len(c.ProjectorCountdownIDs))
		for i, id := range c.ProjectorCountdownIDs {
			refs[i] = &ValueCollection[ProjectorCountdown, *ProjectorCountdown]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectorCountdownList = &RelationList[ProjectorCountdown, *ProjectorCountdown]{refs}
	}
	return c.projectorCountdownList
}

func (c *Meeting) ProjectorList() *RelationList[Projector, *Projector] {
	if c.projectorList == nil {
		refs := make([]*ValueCollection[Projector, *Projector], len(c.ProjectorIDs))
		for i, id := range c.ProjectorIDs {
			refs[i] = &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectorList = &RelationList[Projector, *Projector]{refs}
	}
	return c.projectorList
}

func (c *Meeting) ProjectorMessageList() *RelationList[ProjectorMessage, *ProjectorMessage] {
	if c.projectorMessageList == nil {
		refs := make([]*ValueCollection[ProjectorMessage, *ProjectorMessage], len(c.ProjectorMessageIDs))
		for i, id := range c.ProjectorMessageIDs {
			refs[i] = &ValueCollection[ProjectorMessage, *ProjectorMessage]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectorMessageList = &RelationList[ProjectorMessage, *ProjectorMessage]{refs}
	}
	return c.projectorMessageList
}

func (c *Meeting) ReferenceProjector() *ValueCollection[Projector, *Projector] {
	if c.referenceProjector == nil {
		c.referenceProjector = &ValueCollection[Projector, *Projector]{
			id:    c.ReferenceProjectorID,
			fetch: c.fetch,
		}
	}
	return c.referenceProjector
}

func (c *Meeting) SpeakerList() *RelationList[Speaker, *Speaker] {
	if c.speakerList == nil {
		refs := make([]*ValueCollection[Speaker, *Speaker], len(c.SpeakerIDs))
		for i, id := range c.SpeakerIDs {
			refs[i] = &ValueCollection[Speaker, *Speaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.speakerList = &RelationList[Speaker, *Speaker]{refs}
	}
	return c.speakerList
}

func (c *Meeting) StructureLevelList() *RelationList[StructureLevel, *StructureLevel] {
	if c.structureLevelList == nil {
		refs := make([]*ValueCollection[StructureLevel, *StructureLevel], len(c.StructureLevelIDs))
		for i, id := range c.StructureLevelIDs {
			refs[i] = &ValueCollection[StructureLevel, *StructureLevel]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.structureLevelList = &RelationList[StructureLevel, *StructureLevel]{refs}
	}
	return c.structureLevelList
}

func (c *Meeting) StructureLevelListOfSpeakersList() *RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	if c.structureLevelListOfSpeakersList == nil {
		refs := make([]*ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers], len(c.StructureLevelListOfSpeakersIDs))
		for i, id := range c.StructureLevelListOfSpeakersIDs {
			refs[i] = &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.structureLevelListOfSpeakersList = &RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{refs}
	}
	return c.structureLevelListOfSpeakersList
}

func (c *Meeting) TagList() *RelationList[Tag, *Tag] {
	if c.tagList == nil {
		refs := make([]*ValueCollection[Tag, *Tag], len(c.TagIDs))
		for i, id := range c.TagIDs {
			refs[i] = &ValueCollection[Tag, *Tag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.tagList = &RelationList[Tag, *Tag]{refs}
	}
	return c.tagList
}

func (c *Meeting) TemplateForOrganization() *MaybeRelation[Organization, *Organization] {
	if c.templateForOrganization == nil {
		var ref dsfetch.Maybe[*ValueCollection[Organization, *Organization]]
		id, hasValue := c.TemplateForOrganizationID.Value()
		if hasValue {
			value := &ValueCollection[Organization, *Organization]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.templateForOrganization = &MaybeRelation[Organization, *Organization]{ref}
	}
	return c.templateForOrganization
}

func (c *Meeting) TopicList() *RelationList[Topic, *Topic] {
	if c.topicList == nil {
		refs := make([]*ValueCollection[Topic, *Topic], len(c.TopicIDs))
		for i, id := range c.TopicIDs {
			refs[i] = &ValueCollection[Topic, *Topic]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.topicList = &RelationList[Topic, *Topic]{refs}
	}
	return c.topicList
}

func (c *Meeting) TopicPollDefaultGroupList() *RelationList[Group, *Group] {
	if c.topicPollDefaultGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.TopicPollDefaultGroupIDs))
		for i, id := range c.TopicPollDefaultGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.topicPollDefaultGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.topicPollDefaultGroupList
}

func (c *Meeting) VoteList() *RelationList[Vote, *Vote] {
	if c.voteList == nil {
		refs := make([]*ValueCollection[Vote, *Vote], len(c.VoteIDs))
		for i, id := range c.VoteIDs {
			refs[i] = &ValueCollection[Vote, *Vote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.voteList = &RelationList[Vote, *Vote]{refs}
	}
	return c.voteList
}

func (r *Fetch) Meeting(id int) *ValueCollection[Meeting, *Meeting] {
	return &ValueCollection[Meeting, *Meeting]{
		id:    id,
		fetch: r,
	}
}

// MeetingMediafile has all fields from meeting_mediafile.
type MeetingMediafile struct {
	AccessGroupIDs                         []int
	AttachmentIDs                          []string
	ID                                     int
	InheritedAccessGroupIDs                []int
	IsPublic                               bool
	ListOfSpeakersID                       dsfetch.Maybe[int]
	MediafileID                            int
	MeetingID                              int
	ProjectionIDs                          []int
	UsedAsFontBoldInMeetingID              dsfetch.Maybe[int]
	UsedAsFontBoldItalicInMeetingID        dsfetch.Maybe[int]
	UsedAsFontChyronSpeakerNameInMeetingID dsfetch.Maybe[int]
	UsedAsFontItalicInMeetingID            dsfetch.Maybe[int]
	UsedAsFontMonospaceInMeetingID         dsfetch.Maybe[int]
	UsedAsFontProjectorH1InMeetingID       dsfetch.Maybe[int]
	UsedAsFontProjectorH2InMeetingID       dsfetch.Maybe[int]
	UsedAsFontRegularInMeetingID           dsfetch.Maybe[int]
	UsedAsLogoPdfBallotPaperInMeetingID    dsfetch.Maybe[int]
	UsedAsLogoPdfFooterLInMeetingID        dsfetch.Maybe[int]
	UsedAsLogoPdfFooterRInMeetingID        dsfetch.Maybe[int]
	UsedAsLogoPdfHeaderLInMeetingID        dsfetch.Maybe[int]
	UsedAsLogoPdfHeaderRInMeetingID        dsfetch.Maybe[int]
	UsedAsLogoProjectorHeaderInMeetingID   dsfetch.Maybe[int]
	UsedAsLogoProjectorMainInMeetingID     dsfetch.Maybe[int]
	UsedAsLogoWebHeaderInMeetingID         dsfetch.Maybe[int]
	accessGroupList                        *RelationList[Group, *Group]
	inheritedAccessGroupList               *RelationList[Group, *Group]
	listOfSpeakers                         *MaybeRelation[ListOfSpeakers, *ListOfSpeakers]
	mediafile                              *ValueCollection[Mediafile, *Mediafile]
	meeting                                *ValueCollection[Meeting, *Meeting]
	projectionList                         *RelationList[Projection, *Projection]
	usedAsFontBoldInMeeting                *MaybeRelation[Meeting, *Meeting]
	usedAsFontBoldItalicInMeeting          *MaybeRelation[Meeting, *Meeting]
	usedAsFontChyronSpeakerNameInMeeting   *MaybeRelation[Meeting, *Meeting]
	usedAsFontItalicInMeeting              *MaybeRelation[Meeting, *Meeting]
	usedAsFontMonospaceInMeeting           *MaybeRelation[Meeting, *Meeting]
	usedAsFontProjectorH1InMeeting         *MaybeRelation[Meeting, *Meeting]
	usedAsFontProjectorH2InMeeting         *MaybeRelation[Meeting, *Meeting]
	usedAsFontRegularInMeeting             *MaybeRelation[Meeting, *Meeting]
	usedAsLogoPdfBallotPaperInMeeting      *MaybeRelation[Meeting, *Meeting]
	usedAsLogoPdfFooterLInMeeting          *MaybeRelation[Meeting, *Meeting]
	usedAsLogoPdfFooterRInMeeting          *MaybeRelation[Meeting, *Meeting]
	usedAsLogoPdfHeaderLInMeeting          *MaybeRelation[Meeting, *Meeting]
	usedAsLogoPdfHeaderRInMeeting          *MaybeRelation[Meeting, *Meeting]
	usedAsLogoProjectorHeaderInMeeting     *MaybeRelation[Meeting, *Meeting]
	usedAsLogoProjectorMainInMeeting       *MaybeRelation[Meeting, *Meeting]
	usedAsLogoWebHeaderInMeeting           *MaybeRelation[Meeting, *Meeting]
	fetch                                  *Fetch
}

func (c *MeetingMediafile) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *MeetingMediafile) AccessGroupList() *RelationList[Group, *Group] {
	if c.accessGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.AccessGroupIDs))
		for i, id := range c.AccessGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.accessGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.accessGroupList
}

func (c *MeetingMediafile) InheritedAccessGroupList() *RelationList[Group, *Group] {
	if c.inheritedAccessGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.InheritedAccessGroupIDs))
		for i, id := range c.InheritedAccessGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.inheritedAccessGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.inheritedAccessGroupList
}

func (c *MeetingMediafile) ListOfSpeakers() *MaybeRelation[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		var ref dsfetch.Maybe[*ValueCollection[ListOfSpeakers, *ListOfSpeakers]]
		id, hasValue := c.ListOfSpeakersID.Value()
		if hasValue {
			value := &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.listOfSpeakers = &MaybeRelation[ListOfSpeakers, *ListOfSpeakers]{ref}
	}
	return c.listOfSpeakers
}

func (c *MeetingMediafile) Mediafile() *ValueCollection[Mediafile, *Mediafile] {
	if c.mediafile == nil {
		c.mediafile = &ValueCollection[Mediafile, *Mediafile]{
			id:    c.MediafileID,
			fetch: c.fetch,
		}
	}
	return c.mediafile
}

func (c *MeetingMediafile) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MeetingMediafile) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *MeetingMediafile) UsedAsFontBoldInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontBoldInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontBoldInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontBoldInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontBoldInMeeting
}

func (c *MeetingMediafile) UsedAsFontBoldItalicInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontBoldItalicInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontBoldItalicInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontBoldItalicInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontBoldItalicInMeeting
}

func (c *MeetingMediafile) UsedAsFontChyronSpeakerNameInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontChyronSpeakerNameInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontChyronSpeakerNameInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontChyronSpeakerNameInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontChyronSpeakerNameInMeeting
}

func (c *MeetingMediafile) UsedAsFontItalicInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontItalicInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontItalicInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontItalicInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontItalicInMeeting
}

func (c *MeetingMediafile) UsedAsFontMonospaceInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontMonospaceInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontMonospaceInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontMonospaceInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontMonospaceInMeeting
}

func (c *MeetingMediafile) UsedAsFontProjectorH1InMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontProjectorH1InMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontProjectorH1InMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontProjectorH1InMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontProjectorH1InMeeting
}

func (c *MeetingMediafile) UsedAsFontProjectorH2InMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontProjectorH2InMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontProjectorH2InMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontProjectorH2InMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontProjectorH2InMeeting
}

func (c *MeetingMediafile) UsedAsFontRegularInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsFontRegularInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsFontRegularInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsFontRegularInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsFontRegularInMeeting
}

func (c *MeetingMediafile) UsedAsLogoPdfBallotPaperInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoPdfBallotPaperInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoPdfBallotPaperInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoPdfBallotPaperInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoPdfBallotPaperInMeeting
}

func (c *MeetingMediafile) UsedAsLogoPdfFooterLInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoPdfFooterLInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoPdfFooterLInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoPdfFooterLInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoPdfFooterLInMeeting
}

func (c *MeetingMediafile) UsedAsLogoPdfFooterRInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoPdfFooterRInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoPdfFooterRInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoPdfFooterRInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoPdfFooterRInMeeting
}

func (c *MeetingMediafile) UsedAsLogoPdfHeaderLInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoPdfHeaderLInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoPdfHeaderLInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoPdfHeaderLInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoPdfHeaderLInMeeting
}

func (c *MeetingMediafile) UsedAsLogoPdfHeaderRInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoPdfHeaderRInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoPdfHeaderRInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoPdfHeaderRInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoPdfHeaderRInMeeting
}

func (c *MeetingMediafile) UsedAsLogoProjectorHeaderInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoProjectorHeaderInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoProjectorHeaderInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoProjectorHeaderInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoProjectorHeaderInMeeting
}

func (c *MeetingMediafile) UsedAsLogoProjectorMainInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoProjectorMainInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoProjectorMainInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoProjectorMainInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoProjectorMainInMeeting
}

func (c *MeetingMediafile) UsedAsLogoWebHeaderInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsLogoWebHeaderInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsLogoWebHeaderInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsLogoWebHeaderInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsLogoWebHeaderInMeeting
}

func (r *Fetch) MeetingMediafile(id int) *ValueCollection[MeetingMediafile, *MeetingMediafile] {
	return &ValueCollection[MeetingMediafile, *MeetingMediafile]{
		id:    id,
		fetch: r,
	}
}

// MeetingUser has all fields from meeting_user.
type MeetingUser struct {
	AboutMe                       string
	AssignmentCandidateIDs        []int
	ChatMessageIDs                []int
	Comment                       string
	GroupIDs                      []int
	ID                            int
	LockedOut                     bool
	MeetingID                     int
	MotionEditorIDs               []int
	MotionSubmitterIDs            []int
	MotionWorkingGroupSpeakerIDs  []int
	Number                        string
	PersonalNoteIDs               []int
	SpeakerIDs                    []int
	StructureLevelIDs             []int
	SupportedMotionIDs            []int
	UserID                        int
	VoteDelegatedToID             dsfetch.Maybe[int]
	VoteDelegationsFromIDs        []int
	VoteWeight                    string
	assignmentCandidateList       *RelationList[AssignmentCandidate, *AssignmentCandidate]
	chatMessageList               *RelationList[ChatMessage, *ChatMessage]
	groupList                     *RelationList[Group, *Group]
	meeting                       *ValueCollection[Meeting, *Meeting]
	motionEditorList              *RelationList[MotionEditor, *MotionEditor]
	motionSubmitterList           *RelationList[MotionSubmitter, *MotionSubmitter]
	motionWorkingGroupSpeakerList *RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]
	personalNoteList              *RelationList[PersonalNote, *PersonalNote]
	speakerList                   *RelationList[Speaker, *Speaker]
	structureLevelList            *RelationList[StructureLevel, *StructureLevel]
	supportedMotionList           *RelationList[Motion, *Motion]
	user                          *ValueCollection[User, *User]
	voteDelegatedTo               *MaybeRelation[MeetingUser, *MeetingUser]
	voteDelegationsFromList       *RelationList[MeetingUser, *MeetingUser]
	fetch                         *Fetch
}

func (c *MeetingUser) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
	ds.MeetingUser_MotionWorkingGroupSpeakerIDs(id).Lazy(&c.MotionWorkingGroupSpeakerIDs)
	ds.MeetingUser_Number(id).Lazy(&c.Number)
	ds.MeetingUser_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.MeetingUser_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.MeetingUser_StructureLevelIDs(id).Lazy(&c.StructureLevelIDs)
	ds.MeetingUser_SupportedMotionIDs(id).Lazy(&c.SupportedMotionIDs)
	ds.MeetingUser_UserID(id).Lazy(&c.UserID)
	ds.MeetingUser_VoteDelegatedToID(id).Lazy(&c.VoteDelegatedToID)
	ds.MeetingUser_VoteDelegationsFromIDs(id).Lazy(&c.VoteDelegationsFromIDs)
	ds.MeetingUser_VoteWeight(id).Lazy(&c.VoteWeight)
}

func (c *MeetingUser) AssignmentCandidateList() *RelationList[AssignmentCandidate, *AssignmentCandidate] {
	if c.assignmentCandidateList == nil {
		refs := make([]*ValueCollection[AssignmentCandidate, *AssignmentCandidate], len(c.AssignmentCandidateIDs))
		for i, id := range c.AssignmentCandidateIDs {
			refs[i] = &ValueCollection[AssignmentCandidate, *AssignmentCandidate]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.assignmentCandidateList = &RelationList[AssignmentCandidate, *AssignmentCandidate]{refs}
	}
	return c.assignmentCandidateList
}

func (c *MeetingUser) ChatMessageList() *RelationList[ChatMessage, *ChatMessage] {
	if c.chatMessageList == nil {
		refs := make([]*ValueCollection[ChatMessage, *ChatMessage], len(c.ChatMessageIDs))
		for i, id := range c.ChatMessageIDs {
			refs[i] = &ValueCollection[ChatMessage, *ChatMessage]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.chatMessageList = &RelationList[ChatMessage, *ChatMessage]{refs}
	}
	return c.chatMessageList
}

func (c *MeetingUser) GroupList() *RelationList[Group, *Group] {
	if c.groupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.GroupIDs))
		for i, id := range c.GroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.groupList = &RelationList[Group, *Group]{refs}
	}
	return c.groupList
}

func (c *MeetingUser) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MeetingUser) MotionEditorList() *RelationList[MotionEditor, *MotionEditor] {
	if c.motionEditorList == nil {
		refs := make([]*ValueCollection[MotionEditor, *MotionEditor], len(c.MotionEditorIDs))
		for i, id := range c.MotionEditorIDs {
			refs[i] = &ValueCollection[MotionEditor, *MotionEditor]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionEditorList = &RelationList[MotionEditor, *MotionEditor]{refs}
	}
	return c.motionEditorList
}

func (c *MeetingUser) MotionSubmitterList() *RelationList[MotionSubmitter, *MotionSubmitter] {
	if c.motionSubmitterList == nil {
		refs := make([]*ValueCollection[MotionSubmitter, *MotionSubmitter], len(c.MotionSubmitterIDs))
		for i, id := range c.MotionSubmitterIDs {
			refs[i] = &ValueCollection[MotionSubmitter, *MotionSubmitter]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionSubmitterList = &RelationList[MotionSubmitter, *MotionSubmitter]{refs}
	}
	return c.motionSubmitterList
}

func (c *MeetingUser) MotionWorkingGroupSpeakerList() *RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker] {
	if c.motionWorkingGroupSpeakerList == nil {
		refs := make([]*ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker], len(c.MotionWorkingGroupSpeakerIDs))
		for i, id := range c.MotionWorkingGroupSpeakerIDs {
			refs[i] = &ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionWorkingGroupSpeakerList = &RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{refs}
	}
	return c.motionWorkingGroupSpeakerList
}

func (c *MeetingUser) PersonalNoteList() *RelationList[PersonalNote, *PersonalNote] {
	if c.personalNoteList == nil {
		refs := make([]*ValueCollection[PersonalNote, *PersonalNote], len(c.PersonalNoteIDs))
		for i, id := range c.PersonalNoteIDs {
			refs[i] = &ValueCollection[PersonalNote, *PersonalNote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.personalNoteList = &RelationList[PersonalNote, *PersonalNote]{refs}
	}
	return c.personalNoteList
}

func (c *MeetingUser) SpeakerList() *RelationList[Speaker, *Speaker] {
	if c.speakerList == nil {
		refs := make([]*ValueCollection[Speaker, *Speaker], len(c.SpeakerIDs))
		for i, id := range c.SpeakerIDs {
			refs[i] = &ValueCollection[Speaker, *Speaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.speakerList = &RelationList[Speaker, *Speaker]{refs}
	}
	return c.speakerList
}

func (c *MeetingUser) StructureLevelList() *RelationList[StructureLevel, *StructureLevel] {
	if c.structureLevelList == nil {
		refs := make([]*ValueCollection[StructureLevel, *StructureLevel], len(c.StructureLevelIDs))
		for i, id := range c.StructureLevelIDs {
			refs[i] = &ValueCollection[StructureLevel, *StructureLevel]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.structureLevelList = &RelationList[StructureLevel, *StructureLevel]{refs}
	}
	return c.structureLevelList
}

func (c *MeetingUser) SupportedMotionList() *RelationList[Motion, *Motion] {
	if c.supportedMotionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.SupportedMotionIDs))
		for i, id := range c.SupportedMotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.supportedMotionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.supportedMotionList
}

func (c *MeetingUser) User() *ValueCollection[User, *User] {
	if c.user == nil {
		c.user = &ValueCollection[User, *User]{
			id:    c.UserID,
			fetch: c.fetch,
		}
	}
	return c.user
}

func (c *MeetingUser) VoteDelegatedTo() *MaybeRelation[MeetingUser, *MeetingUser] {
	if c.voteDelegatedTo == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingUser, *MeetingUser]]
		id, hasValue := c.VoteDelegatedToID.Value()
		if hasValue {
			value := &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.voteDelegatedTo = &MaybeRelation[MeetingUser, *MeetingUser]{ref}
	}
	return c.voteDelegatedTo
}

func (c *MeetingUser) VoteDelegationsFromList() *RelationList[MeetingUser, *MeetingUser] {
	if c.voteDelegationsFromList == nil {
		refs := make([]*ValueCollection[MeetingUser, *MeetingUser], len(c.VoteDelegationsFromIDs))
		for i, id := range c.VoteDelegationsFromIDs {
			refs[i] = &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.voteDelegationsFromList = &RelationList[MeetingUser, *MeetingUser]{refs}
	}
	return c.voteDelegationsFromList
}

func (r *Fetch) MeetingUser(id int) *ValueCollection[MeetingUser, *MeetingUser] {
	return &ValueCollection[MeetingUser, *MeetingUser]{
		id:    id,
		fetch: r,
	}
}

// Motion has all fields from motion.
type Motion struct {
	AdditionalSubmitter                           string
	AgendaItemID                                  dsfetch.Maybe[int]
	AllDerivedMotionIDs                           []int
	AllOriginIDs                                  []int
	AmendmentIDs                                  []int
	AmendmentParagraphs                           json.RawMessage
	AttachmentMeetingMediafileIDs                 []int
	BlockID                                       dsfetch.Maybe[int]
	CategoryID                                    dsfetch.Maybe[int]
	CategoryWeight                                int
	ChangeRecommendationIDs                       []int
	CommentIDs                                    []int
	Created                                       int
	DerivedMotionIDs                              []int
	EditorIDs                                     []int
	Forwarded                                     int
	ID                                            int
	IDenticalMotionIDs                            []int
	LastModified                                  int
	LeadMotionID                                  dsfetch.Maybe[int]
	ListOfSpeakersID                              int
	MeetingID                                     int
	ModifiedFinalVersion                          string
	Number                                        string
	NumberValue                                   int
	OptionIDs                                     []int
	OriginID                                      dsfetch.Maybe[int]
	OriginMeetingID                               dsfetch.Maybe[int]
	PersonalNoteIDs                               []int
	PollIDs                                       []int
	ProjectionIDs                                 []int
	Reason                                        string
	RecommendationExtension                       string
	RecommendationExtensionReferenceIDs           []string
	RecommendationID                              dsfetch.Maybe[int]
	ReferencedInMotionRecommendationExtensionIDs  []int
	ReferencedInMotionStateExtensionIDs           []int
	SequentialNumber                              int
	SortChildIDs                                  []int
	SortParentID                                  dsfetch.Maybe[int]
	SortWeight                                    int
	StartLineNumber                               int
	StateExtension                                string
	StateExtensionReferenceIDs                    []string
	StateID                                       int
	SubmitterIDs                                  []int
	SupporterMeetingUserIDs                       []int
	TagIDs                                        []int
	Text                                          string
	TextHash                                      string
	Title                                         string
	WorkflowTimestamp                             int
	WorkingGroupSpeakerIDs                        []int
	agendaItem                                    *MaybeRelation[AgendaItem, *AgendaItem]
	allDerivedMotionList                          *RelationList[Motion, *Motion]
	allOriginList                                 *RelationList[Motion, *Motion]
	amendmentList                                 *RelationList[Motion, *Motion]
	attachmentMeetingMediafileList                *RelationList[MeetingMediafile, *MeetingMediafile]
	block                                         *MaybeRelation[MotionBlock, *MotionBlock]
	category                                      *MaybeRelation[MotionCategory, *MotionCategory]
	changeRecommendationList                      *RelationList[MotionChangeRecommendation, *MotionChangeRecommendation]
	commentList                                   *RelationList[MotionComment, *MotionComment]
	derivedMotionList                             *RelationList[Motion, *Motion]
	editorList                                    *RelationList[MotionEditor, *MotionEditor]
	iDenticalMotionList                           *RelationList[Motion, *Motion]
	leadMotion                                    *MaybeRelation[Motion, *Motion]
	listOfSpeakers                                *ValueCollection[ListOfSpeakers, *ListOfSpeakers]
	meeting                                       *ValueCollection[Meeting, *Meeting]
	optionList                                    *RelationList[Option, *Option]
	origin                                        *MaybeRelation[Motion, *Motion]
	originMeeting                                 *MaybeRelation[Meeting, *Meeting]
	personalNoteList                              *RelationList[PersonalNote, *PersonalNote]
	pollList                                      *RelationList[Poll, *Poll]
	projectionList                                *RelationList[Projection, *Projection]
	recommendation                                *MaybeRelation[MotionState, *MotionState]
	referencedInMotionRecommendationExtensionList *RelationList[Motion, *Motion]
	referencedInMotionStateExtensionList          *RelationList[Motion, *Motion]
	sortChildList                                 *RelationList[Motion, *Motion]
	sortParent                                    *MaybeRelation[Motion, *Motion]
	state                                         *ValueCollection[MotionState, *MotionState]
	submitterList                                 *RelationList[MotionSubmitter, *MotionSubmitter]
	supporterMeetingUserList                      *RelationList[MeetingUser, *MeetingUser]
	tagList                                       *RelationList[Tag, *Tag]
	workingGroupSpeakerList                       *RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]
	fetch                                         *Fetch
}

func (c *Motion) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
	ds.Motion_EditorIDs(id).Lazy(&c.EditorIDs)
	ds.Motion_Forwarded(id).Lazy(&c.Forwarded)
	ds.Motion_ID(id).Lazy(&c.ID)
	ds.Motion_IDenticalMotionIDs(id).Lazy(&c.IDenticalMotionIDs)
	ds.Motion_LastModified(id).Lazy(&c.LastModified)
	ds.Motion_LeadMotionID(id).Lazy(&c.LeadMotionID)
	ds.Motion_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
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
	ds.Motion_SupporterMeetingUserIDs(id).Lazy(&c.SupporterMeetingUserIDs)
	ds.Motion_TagIDs(id).Lazy(&c.TagIDs)
	ds.Motion_Text(id).Lazy(&c.Text)
	ds.Motion_TextHash(id).Lazy(&c.TextHash)
	ds.Motion_Title(id).Lazy(&c.Title)
	ds.Motion_WorkflowTimestamp(id).Lazy(&c.WorkflowTimestamp)
	ds.Motion_WorkingGroupSpeakerIDs(id).Lazy(&c.WorkingGroupSpeakerIDs)
}

func (c *Motion) AgendaItem() *MaybeRelation[AgendaItem, *AgendaItem] {
	if c.agendaItem == nil {
		var ref dsfetch.Maybe[*ValueCollection[AgendaItem, *AgendaItem]]
		id, hasValue := c.AgendaItemID.Value()
		if hasValue {
			value := &ValueCollection[AgendaItem, *AgendaItem]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.agendaItem = &MaybeRelation[AgendaItem, *AgendaItem]{ref}
	}
	return c.agendaItem
}

func (c *Motion) AllDerivedMotionList() *RelationList[Motion, *Motion] {
	if c.allDerivedMotionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.AllDerivedMotionIDs))
		for i, id := range c.AllDerivedMotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.allDerivedMotionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.allDerivedMotionList
}

func (c *Motion) AllOriginList() *RelationList[Motion, *Motion] {
	if c.allOriginList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.AllOriginIDs))
		for i, id := range c.AllOriginIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.allOriginList = &RelationList[Motion, *Motion]{refs}
	}
	return c.allOriginList
}

func (c *Motion) AmendmentList() *RelationList[Motion, *Motion] {
	if c.amendmentList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.AmendmentIDs))
		for i, id := range c.AmendmentIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.amendmentList = &RelationList[Motion, *Motion]{refs}
	}
	return c.amendmentList
}

func (c *Motion) AttachmentMeetingMediafileList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.attachmentMeetingMediafileList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.AttachmentMeetingMediafileIDs))
		for i, id := range c.AttachmentMeetingMediafileIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.attachmentMeetingMediafileList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.attachmentMeetingMediafileList
}

func (c *Motion) Block() *MaybeRelation[MotionBlock, *MotionBlock] {
	if c.block == nil {
		var ref dsfetch.Maybe[*ValueCollection[MotionBlock, *MotionBlock]]
		id, hasValue := c.BlockID.Value()
		if hasValue {
			value := &ValueCollection[MotionBlock, *MotionBlock]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.block = &MaybeRelation[MotionBlock, *MotionBlock]{ref}
	}
	return c.block
}

func (c *Motion) Category() *MaybeRelation[MotionCategory, *MotionCategory] {
	if c.category == nil {
		var ref dsfetch.Maybe[*ValueCollection[MotionCategory, *MotionCategory]]
		id, hasValue := c.CategoryID.Value()
		if hasValue {
			value := &ValueCollection[MotionCategory, *MotionCategory]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.category = &MaybeRelation[MotionCategory, *MotionCategory]{ref}
	}
	return c.category
}

func (c *Motion) ChangeRecommendationList() *RelationList[MotionChangeRecommendation, *MotionChangeRecommendation] {
	if c.changeRecommendationList == nil {
		refs := make([]*ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation], len(c.ChangeRecommendationIDs))
		for i, id := range c.ChangeRecommendationIDs {
			refs[i] = &ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.changeRecommendationList = &RelationList[MotionChangeRecommendation, *MotionChangeRecommendation]{refs}
	}
	return c.changeRecommendationList
}

func (c *Motion) CommentList() *RelationList[MotionComment, *MotionComment] {
	if c.commentList == nil {
		refs := make([]*ValueCollection[MotionComment, *MotionComment], len(c.CommentIDs))
		for i, id := range c.CommentIDs {
			refs[i] = &ValueCollection[MotionComment, *MotionComment]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.commentList = &RelationList[MotionComment, *MotionComment]{refs}
	}
	return c.commentList
}

func (c *Motion) DerivedMotionList() *RelationList[Motion, *Motion] {
	if c.derivedMotionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.DerivedMotionIDs))
		for i, id := range c.DerivedMotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.derivedMotionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.derivedMotionList
}

func (c *Motion) EditorList() *RelationList[MotionEditor, *MotionEditor] {
	if c.editorList == nil {
		refs := make([]*ValueCollection[MotionEditor, *MotionEditor], len(c.EditorIDs))
		for i, id := range c.EditorIDs {
			refs[i] = &ValueCollection[MotionEditor, *MotionEditor]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.editorList = &RelationList[MotionEditor, *MotionEditor]{refs}
	}
	return c.editorList
}

func (c *Motion) IDenticalMotionList() *RelationList[Motion, *Motion] {
	if c.iDenticalMotionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.IDenticalMotionIDs))
		for i, id := range c.IDenticalMotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.iDenticalMotionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.iDenticalMotionList
}

func (c *Motion) LeadMotion() *MaybeRelation[Motion, *Motion] {
	if c.leadMotion == nil {
		var ref dsfetch.Maybe[*ValueCollection[Motion, *Motion]]
		id, hasValue := c.LeadMotionID.Value()
		if hasValue {
			value := &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.leadMotion = &MaybeRelation[Motion, *Motion]{ref}
	}
	return c.leadMotion
}

func (c *Motion) ListOfSpeakers() *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		c.listOfSpeakers = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
			id:    c.ListOfSpeakersID,
			fetch: c.fetch,
		}
	}
	return c.listOfSpeakers
}

func (c *Motion) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Motion) OptionList() *RelationList[Option, *Option] {
	if c.optionList == nil {
		refs := make([]*ValueCollection[Option, *Option], len(c.OptionIDs))
		for i, id := range c.OptionIDs {
			refs[i] = &ValueCollection[Option, *Option]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.optionList = &RelationList[Option, *Option]{refs}
	}
	return c.optionList
}

func (c *Motion) Origin() *MaybeRelation[Motion, *Motion] {
	if c.origin == nil {
		var ref dsfetch.Maybe[*ValueCollection[Motion, *Motion]]
		id, hasValue := c.OriginID.Value()
		if hasValue {
			value := &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.origin = &MaybeRelation[Motion, *Motion]{ref}
	}
	return c.origin
}

func (c *Motion) OriginMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.originMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.OriginMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.originMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.originMeeting
}

func (c *Motion) PersonalNoteList() *RelationList[PersonalNote, *PersonalNote] {
	if c.personalNoteList == nil {
		refs := make([]*ValueCollection[PersonalNote, *PersonalNote], len(c.PersonalNoteIDs))
		for i, id := range c.PersonalNoteIDs {
			refs[i] = &ValueCollection[PersonalNote, *PersonalNote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.personalNoteList = &RelationList[PersonalNote, *PersonalNote]{refs}
	}
	return c.personalNoteList
}

func (c *Motion) PollList() *RelationList[Poll, *Poll] {
	if c.pollList == nil {
		refs := make([]*ValueCollection[Poll, *Poll], len(c.PollIDs))
		for i, id := range c.PollIDs {
			refs[i] = &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollList = &RelationList[Poll, *Poll]{refs}
	}
	return c.pollList
}

func (c *Motion) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *Motion) Recommendation() *MaybeRelation[MotionState, *MotionState] {
	if c.recommendation == nil {
		var ref dsfetch.Maybe[*ValueCollection[MotionState, *MotionState]]
		id, hasValue := c.RecommendationID.Value()
		if hasValue {
			value := &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.recommendation = &MaybeRelation[MotionState, *MotionState]{ref}
	}
	return c.recommendation
}

func (c *Motion) ReferencedInMotionRecommendationExtensionList() *RelationList[Motion, *Motion] {
	if c.referencedInMotionRecommendationExtensionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.ReferencedInMotionRecommendationExtensionIDs))
		for i, id := range c.ReferencedInMotionRecommendationExtensionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.referencedInMotionRecommendationExtensionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.referencedInMotionRecommendationExtensionList
}

func (c *Motion) ReferencedInMotionStateExtensionList() *RelationList[Motion, *Motion] {
	if c.referencedInMotionStateExtensionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.ReferencedInMotionStateExtensionIDs))
		for i, id := range c.ReferencedInMotionStateExtensionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.referencedInMotionStateExtensionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.referencedInMotionStateExtensionList
}

func (c *Motion) SortChildList() *RelationList[Motion, *Motion] {
	if c.sortChildList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.SortChildIDs))
		for i, id := range c.SortChildIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.sortChildList = &RelationList[Motion, *Motion]{refs}
	}
	return c.sortChildList
}

func (c *Motion) SortParent() *MaybeRelation[Motion, *Motion] {
	if c.sortParent == nil {
		var ref dsfetch.Maybe[*ValueCollection[Motion, *Motion]]
		id, hasValue := c.SortParentID.Value()
		if hasValue {
			value := &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.sortParent = &MaybeRelation[Motion, *Motion]{ref}
	}
	return c.sortParent
}

func (c *Motion) State() *ValueCollection[MotionState, *MotionState] {
	if c.state == nil {
		c.state = &ValueCollection[MotionState, *MotionState]{
			id:    c.StateID,
			fetch: c.fetch,
		}
	}
	return c.state
}

func (c *Motion) SubmitterList() *RelationList[MotionSubmitter, *MotionSubmitter] {
	if c.submitterList == nil {
		refs := make([]*ValueCollection[MotionSubmitter, *MotionSubmitter], len(c.SubmitterIDs))
		for i, id := range c.SubmitterIDs {
			refs[i] = &ValueCollection[MotionSubmitter, *MotionSubmitter]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.submitterList = &RelationList[MotionSubmitter, *MotionSubmitter]{refs}
	}
	return c.submitterList
}

func (c *Motion) SupporterMeetingUserList() *RelationList[MeetingUser, *MeetingUser] {
	if c.supporterMeetingUserList == nil {
		refs := make([]*ValueCollection[MeetingUser, *MeetingUser], len(c.SupporterMeetingUserIDs))
		for i, id := range c.SupporterMeetingUserIDs {
			refs[i] = &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.supporterMeetingUserList = &RelationList[MeetingUser, *MeetingUser]{refs}
	}
	return c.supporterMeetingUserList
}

func (c *Motion) TagList() *RelationList[Tag, *Tag] {
	if c.tagList == nil {
		refs := make([]*ValueCollection[Tag, *Tag], len(c.TagIDs))
		for i, id := range c.TagIDs {
			refs[i] = &ValueCollection[Tag, *Tag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.tagList = &RelationList[Tag, *Tag]{refs}
	}
	return c.tagList
}

func (c *Motion) WorkingGroupSpeakerList() *RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker] {
	if c.workingGroupSpeakerList == nil {
		refs := make([]*ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker], len(c.WorkingGroupSpeakerIDs))
		for i, id := range c.WorkingGroupSpeakerIDs {
			refs[i] = &ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.workingGroupSpeakerList = &RelationList[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{refs}
	}
	return c.workingGroupSpeakerList
}

func (r *Fetch) Motion(id int) *ValueCollection[Motion, *Motion] {
	return &ValueCollection[Motion, *Motion]{
		id:    id,
		fetch: r,
	}
}

// MotionBlock has all fields from motion_block.
type MotionBlock struct {
	AgendaItemID     dsfetch.Maybe[int]
	ID               int
	Internal         bool
	ListOfSpeakersID int
	MeetingID        int
	MotionIDs        []int
	ProjectionIDs    []int
	SequentialNumber int
	Title            string
	agendaItem       *MaybeRelation[AgendaItem, *AgendaItem]
	listOfSpeakers   *ValueCollection[ListOfSpeakers, *ListOfSpeakers]
	meeting          *ValueCollection[Meeting, *Meeting]
	motionList       *RelationList[Motion, *Motion]
	projectionList   *RelationList[Projection, *Projection]
	fetch            *Fetch
}

func (c *MotionBlock) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionBlock_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.MotionBlock_ID(id).Lazy(&c.ID)
	ds.MotionBlock_Internal(id).Lazy(&c.Internal)
	ds.MotionBlock_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.MotionBlock_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionBlock_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionBlock_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.MotionBlock_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionBlock_Title(id).Lazy(&c.Title)
}

func (c *MotionBlock) AgendaItem() *MaybeRelation[AgendaItem, *AgendaItem] {
	if c.agendaItem == nil {
		var ref dsfetch.Maybe[*ValueCollection[AgendaItem, *AgendaItem]]
		id, hasValue := c.AgendaItemID.Value()
		if hasValue {
			value := &ValueCollection[AgendaItem, *AgendaItem]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.agendaItem = &MaybeRelation[AgendaItem, *AgendaItem]{ref}
	}
	return c.agendaItem
}

func (c *MotionBlock) ListOfSpeakers() *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		c.listOfSpeakers = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
			id:    c.ListOfSpeakersID,
			fetch: c.fetch,
		}
	}
	return c.listOfSpeakers
}

func (c *MotionBlock) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionBlock) MotionList() *RelationList[Motion, *Motion] {
	if c.motionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.MotionIDs))
		for i, id := range c.MotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.motionList
}

func (c *MotionBlock) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (r *Fetch) MotionBlock(id int) *ValueCollection[MotionBlock, *MotionBlock] {
	return &ValueCollection[MotionBlock, *MotionBlock]{
		id:    id,
		fetch: r,
	}
}

// MotionCategory has all fields from motion_category.
type MotionCategory struct {
	ChildIDs         []int
	ID               int
	Level            int
	MeetingID        int
	MotionIDs        []int
	Name             string
	ParentID         dsfetch.Maybe[int]
	Prefix           string
	SequentialNumber int
	Weight           int
	childList        *RelationList[MotionCategory, *MotionCategory]
	meeting          *ValueCollection[Meeting, *Meeting]
	motionList       *RelationList[Motion, *Motion]
	parent           *MaybeRelation[MotionCategory, *MotionCategory]
	fetch            *Fetch
}

func (c *MotionCategory) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *MotionCategory) ChildList() *RelationList[MotionCategory, *MotionCategory] {
	if c.childList == nil {
		refs := make([]*ValueCollection[MotionCategory, *MotionCategory], len(c.ChildIDs))
		for i, id := range c.ChildIDs {
			refs[i] = &ValueCollection[MotionCategory, *MotionCategory]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.childList = &RelationList[MotionCategory, *MotionCategory]{refs}
	}
	return c.childList
}

func (c *MotionCategory) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionCategory) MotionList() *RelationList[Motion, *Motion] {
	if c.motionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.MotionIDs))
		for i, id := range c.MotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.motionList
}

func (c *MotionCategory) Parent() *MaybeRelation[MotionCategory, *MotionCategory] {
	if c.parent == nil {
		var ref dsfetch.Maybe[*ValueCollection[MotionCategory, *MotionCategory]]
		id, hasValue := c.ParentID.Value()
		if hasValue {
			value := &ValueCollection[MotionCategory, *MotionCategory]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.parent = &MaybeRelation[MotionCategory, *MotionCategory]{ref}
	}
	return c.parent
}

func (r *Fetch) MotionCategory(id int) *ValueCollection[MotionCategory, *MotionCategory] {
	return &ValueCollection[MotionCategory, *MotionCategory]{
		id:    id,
		fetch: r,
	}
}

// MotionChangeRecommendation has all fields from motion_change_recommendation.
type MotionChangeRecommendation struct {
	CreationTime     int
	ID               int
	Internal         bool
	LineFrom         int
	LineTo           int
	MeetingID        int
	MotionID         int
	OtherDescription string
	Rejected         bool
	Text             string
	Type             string
	meeting          *ValueCollection[Meeting, *Meeting]
	motion           *ValueCollection[Motion, *Motion]
	fetch            *Fetch
}

func (c *MotionChangeRecommendation) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *MotionChangeRecommendation) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionChangeRecommendation) Motion() *ValueCollection[Motion, *Motion] {
	if c.motion == nil {
		c.motion = &ValueCollection[Motion, *Motion]{
			id:    c.MotionID,
			fetch: c.fetch,
		}
	}
	return c.motion
}

func (r *Fetch) MotionChangeRecommendation(id int) *ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation] {
	return &ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation]{
		id:    id,
		fetch: r,
	}
}

// MotionComment has all fields from motion_comment.
type MotionComment struct {
	Comment   string
	ID        int
	MeetingID int
	MotionID  int
	SectionID int
	meeting   *ValueCollection[Meeting, *Meeting]
	motion    *ValueCollection[Motion, *Motion]
	section   *ValueCollection[MotionCommentSection, *MotionCommentSection]
	fetch     *Fetch
}

func (c *MotionComment) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionComment_Comment(id).Lazy(&c.Comment)
	ds.MotionComment_ID(id).Lazy(&c.ID)
	ds.MotionComment_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionComment_MotionID(id).Lazy(&c.MotionID)
	ds.MotionComment_SectionID(id).Lazy(&c.SectionID)
}

func (c *MotionComment) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionComment) Motion() *ValueCollection[Motion, *Motion] {
	if c.motion == nil {
		c.motion = &ValueCollection[Motion, *Motion]{
			id:    c.MotionID,
			fetch: c.fetch,
		}
	}
	return c.motion
}

func (c *MotionComment) Section() *ValueCollection[MotionCommentSection, *MotionCommentSection] {
	if c.section == nil {
		c.section = &ValueCollection[MotionCommentSection, *MotionCommentSection]{
			id:    c.SectionID,
			fetch: c.fetch,
		}
	}
	return c.section
}

func (r *Fetch) MotionComment(id int) *ValueCollection[MotionComment, *MotionComment] {
	return &ValueCollection[MotionComment, *MotionComment]{
		id:    id,
		fetch: r,
	}
}

// MotionCommentSection has all fields from motion_comment_section.
type MotionCommentSection struct {
	CommentIDs        []int
	ID                int
	MeetingID         int
	Name              string
	ReadGroupIDs      []int
	SequentialNumber  int
	SubmitterCanWrite bool
	Weight            int
	WriteGroupIDs     []int
	commentList       *RelationList[MotionComment, *MotionComment]
	meeting           *ValueCollection[Meeting, *Meeting]
	readGroupList     *RelationList[Group, *Group]
	writeGroupList    *RelationList[Group, *Group]
	fetch             *Fetch
}

func (c *MotionCommentSection) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionCommentSection_CommentIDs(id).Lazy(&c.CommentIDs)
	ds.MotionCommentSection_ID(id).Lazy(&c.ID)
	ds.MotionCommentSection_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionCommentSection_Name(id).Lazy(&c.Name)
	ds.MotionCommentSection_ReadGroupIDs(id).Lazy(&c.ReadGroupIDs)
	ds.MotionCommentSection_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionCommentSection_SubmitterCanWrite(id).Lazy(&c.SubmitterCanWrite)
	ds.MotionCommentSection_Weight(id).Lazy(&c.Weight)
	ds.MotionCommentSection_WriteGroupIDs(id).Lazy(&c.WriteGroupIDs)
}

func (c *MotionCommentSection) CommentList() *RelationList[MotionComment, *MotionComment] {
	if c.commentList == nil {
		refs := make([]*ValueCollection[MotionComment, *MotionComment], len(c.CommentIDs))
		for i, id := range c.CommentIDs {
			refs[i] = &ValueCollection[MotionComment, *MotionComment]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.commentList = &RelationList[MotionComment, *MotionComment]{refs}
	}
	return c.commentList
}

func (c *MotionCommentSection) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionCommentSection) ReadGroupList() *RelationList[Group, *Group] {
	if c.readGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.ReadGroupIDs))
		for i, id := range c.ReadGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.readGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.readGroupList
}

func (c *MotionCommentSection) WriteGroupList() *RelationList[Group, *Group] {
	if c.writeGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.WriteGroupIDs))
		for i, id := range c.WriteGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.writeGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.writeGroupList
}

func (r *Fetch) MotionCommentSection(id int) *ValueCollection[MotionCommentSection, *MotionCommentSection] {
	return &ValueCollection[MotionCommentSection, *MotionCommentSection]{
		id:    id,
		fetch: r,
	}
}

// MotionEditor has all fields from motion_editor.
type MotionEditor struct {
	ID            int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        int
	meeting       *ValueCollection[Meeting, *Meeting]
	meetingUser   *ValueCollection[MeetingUser, *MeetingUser]
	motion        *ValueCollection[Motion, *Motion]
	fetch         *Fetch
}

func (c *MotionEditor) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionEditor_ID(id).Lazy(&c.ID)
	ds.MotionEditor_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionEditor_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionEditor_MotionID(id).Lazy(&c.MotionID)
	ds.MotionEditor_Weight(id).Lazy(&c.Weight)
}

func (c *MotionEditor) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionEditor) MeetingUser() *ValueCollection[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		c.meetingUser = &ValueCollection[MeetingUser, *MeetingUser]{
			id:    c.MeetingUserID,
			fetch: c.fetch,
		}
	}
	return c.meetingUser
}

func (c *MotionEditor) Motion() *ValueCollection[Motion, *Motion] {
	if c.motion == nil {
		c.motion = &ValueCollection[Motion, *Motion]{
			id:    c.MotionID,
			fetch: c.fetch,
		}
	}
	return c.motion
}

func (r *Fetch) MotionEditor(id int) *ValueCollection[MotionEditor, *MotionEditor] {
	return &ValueCollection[MotionEditor, *MotionEditor]{
		id:    id,
		fetch: r,
	}
}

// MotionState has all fields from motion_state.
type MotionState struct {
	AllowCreatePoll                  bool
	AllowMotionForwarding            bool
	AllowSubmitterEdit               bool
	AllowSupport                     bool
	CssClass                         string
	FirstStateOfWorkflowID           dsfetch.Maybe[int]
	ID                               int
	IsInternal                       bool
	MeetingID                        int
	MergeAmendmentIntoFinal          string
	MotionIDs                        []int
	MotionRecommendationIDs          []int
	Name                             string
	NextStateIDs                     []int
	PreviousStateIDs                 []int
	RecommendationLabel              string
	Restrictions                     []string
	SetNumber                        bool
	SetWorkflowTimestamp             bool
	ShowRecommendationExtensionField bool
	ShowStateExtensionField          bool
	SubmitterWithdrawBackIDs         []int
	SubmitterWithdrawStateID         dsfetch.Maybe[int]
	Weight                           int
	WorkflowID                       int
	firstStateOfWorkflow             *MaybeRelation[MotionWorkflow, *MotionWorkflow]
	meeting                          *ValueCollection[Meeting, *Meeting]
	motionList                       *RelationList[Motion, *Motion]
	motionRecommendationList         *RelationList[Motion, *Motion]
	nextStateList                    *RelationList[MotionState, *MotionState]
	previousStateList                *RelationList[MotionState, *MotionState]
	submitterWithdrawBackList        *RelationList[MotionState, *MotionState]
	submitterWithdrawState           *MaybeRelation[MotionState, *MotionState]
	workflow                         *ValueCollection[MotionWorkflow, *MotionWorkflow]
	fetch                            *Fetch
}

func (c *MotionState) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
	ds.MotionState_SubmitterWithdrawBackIDs(id).Lazy(&c.SubmitterWithdrawBackIDs)
	ds.MotionState_SubmitterWithdrawStateID(id).Lazy(&c.SubmitterWithdrawStateID)
	ds.MotionState_Weight(id).Lazy(&c.Weight)
	ds.MotionState_WorkflowID(id).Lazy(&c.WorkflowID)
}

func (c *MotionState) FirstStateOfWorkflow() *MaybeRelation[MotionWorkflow, *MotionWorkflow] {
	if c.firstStateOfWorkflow == nil {
		var ref dsfetch.Maybe[*ValueCollection[MotionWorkflow, *MotionWorkflow]]
		id, hasValue := c.FirstStateOfWorkflowID.Value()
		if hasValue {
			value := &ValueCollection[MotionWorkflow, *MotionWorkflow]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.firstStateOfWorkflow = &MaybeRelation[MotionWorkflow, *MotionWorkflow]{ref}
	}
	return c.firstStateOfWorkflow
}

func (c *MotionState) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionState) MotionList() *RelationList[Motion, *Motion] {
	if c.motionList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.MotionIDs))
		for i, id := range c.MotionIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionList = &RelationList[Motion, *Motion]{refs}
	}
	return c.motionList
}

func (c *MotionState) MotionRecommendationList() *RelationList[Motion, *Motion] {
	if c.motionRecommendationList == nil {
		refs := make([]*ValueCollection[Motion, *Motion], len(c.MotionRecommendationIDs))
		for i, id := range c.MotionRecommendationIDs {
			refs[i] = &ValueCollection[Motion, *Motion]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.motionRecommendationList = &RelationList[Motion, *Motion]{refs}
	}
	return c.motionRecommendationList
}

func (c *MotionState) NextStateList() *RelationList[MotionState, *MotionState] {
	if c.nextStateList == nil {
		refs := make([]*ValueCollection[MotionState, *MotionState], len(c.NextStateIDs))
		for i, id := range c.NextStateIDs {
			refs[i] = &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.nextStateList = &RelationList[MotionState, *MotionState]{refs}
	}
	return c.nextStateList
}

func (c *MotionState) PreviousStateList() *RelationList[MotionState, *MotionState] {
	if c.previousStateList == nil {
		refs := make([]*ValueCollection[MotionState, *MotionState], len(c.PreviousStateIDs))
		for i, id := range c.PreviousStateIDs {
			refs[i] = &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.previousStateList = &RelationList[MotionState, *MotionState]{refs}
	}
	return c.previousStateList
}

func (c *MotionState) SubmitterWithdrawBackList() *RelationList[MotionState, *MotionState] {
	if c.submitterWithdrawBackList == nil {
		refs := make([]*ValueCollection[MotionState, *MotionState], len(c.SubmitterWithdrawBackIDs))
		for i, id := range c.SubmitterWithdrawBackIDs {
			refs[i] = &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.submitterWithdrawBackList = &RelationList[MotionState, *MotionState]{refs}
	}
	return c.submitterWithdrawBackList
}

func (c *MotionState) SubmitterWithdrawState() *MaybeRelation[MotionState, *MotionState] {
	if c.submitterWithdrawState == nil {
		var ref dsfetch.Maybe[*ValueCollection[MotionState, *MotionState]]
		id, hasValue := c.SubmitterWithdrawStateID.Value()
		if hasValue {
			value := &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.submitterWithdrawState = &MaybeRelation[MotionState, *MotionState]{ref}
	}
	return c.submitterWithdrawState
}

func (c *MotionState) Workflow() *ValueCollection[MotionWorkflow, *MotionWorkflow] {
	if c.workflow == nil {
		c.workflow = &ValueCollection[MotionWorkflow, *MotionWorkflow]{
			id:    c.WorkflowID,
			fetch: c.fetch,
		}
	}
	return c.workflow
}

func (r *Fetch) MotionState(id int) *ValueCollection[MotionState, *MotionState] {
	return &ValueCollection[MotionState, *MotionState]{
		id:    id,
		fetch: r,
	}
}

// MotionSubmitter has all fields from motion_submitter.
type MotionSubmitter struct {
	ID            int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        int
	meeting       *ValueCollection[Meeting, *Meeting]
	meetingUser   *ValueCollection[MeetingUser, *MeetingUser]
	motion        *ValueCollection[Motion, *Motion]
	fetch         *Fetch
}

func (c *MotionSubmitter) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionSubmitter_ID(id).Lazy(&c.ID)
	ds.MotionSubmitter_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionSubmitter_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionSubmitter_MotionID(id).Lazy(&c.MotionID)
	ds.MotionSubmitter_Weight(id).Lazy(&c.Weight)
}

func (c *MotionSubmitter) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionSubmitter) MeetingUser() *ValueCollection[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		c.meetingUser = &ValueCollection[MeetingUser, *MeetingUser]{
			id:    c.MeetingUserID,
			fetch: c.fetch,
		}
	}
	return c.meetingUser
}

func (c *MotionSubmitter) Motion() *ValueCollection[Motion, *Motion] {
	if c.motion == nil {
		c.motion = &ValueCollection[Motion, *Motion]{
			id:    c.MotionID,
			fetch: c.fetch,
		}
	}
	return c.motion
}

func (r *Fetch) MotionSubmitter(id int) *ValueCollection[MotionSubmitter, *MotionSubmitter] {
	return &ValueCollection[MotionSubmitter, *MotionSubmitter]{
		id:    id,
		fetch: r,
	}
}

// MotionWorkflow has all fields from motion_workflow.
type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID dsfetch.Maybe[int]
	DefaultWorkflowMeetingID          dsfetch.Maybe[int]
	FirstStateID                      int
	ID                                int
	MeetingID                         int
	Name                              string
	SequentialNumber                  int
	StateIDs                          []int
	defaultAmendmentWorkflowMeeting   *MaybeRelation[Meeting, *Meeting]
	defaultWorkflowMeeting            *MaybeRelation[Meeting, *Meeting]
	firstState                        *ValueCollection[MotionState, *MotionState]
	meeting                           *ValueCollection[Meeting, *Meeting]
	stateList                         *RelationList[MotionState, *MotionState]
	fetch                             *Fetch
}

func (c *MotionWorkflow) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionWorkflow_DefaultAmendmentWorkflowMeetingID(id).Lazy(&c.DefaultAmendmentWorkflowMeetingID)
	ds.MotionWorkflow_DefaultWorkflowMeetingID(id).Lazy(&c.DefaultWorkflowMeetingID)
	ds.MotionWorkflow_FirstStateID(id).Lazy(&c.FirstStateID)
	ds.MotionWorkflow_ID(id).Lazy(&c.ID)
	ds.MotionWorkflow_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionWorkflow_Name(id).Lazy(&c.Name)
	ds.MotionWorkflow_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionWorkflow_StateIDs(id).Lazy(&c.StateIDs)
}

func (c *MotionWorkflow) DefaultAmendmentWorkflowMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.defaultAmendmentWorkflowMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.DefaultAmendmentWorkflowMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.defaultAmendmentWorkflowMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.defaultAmendmentWorkflowMeeting
}

func (c *MotionWorkflow) DefaultWorkflowMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.defaultWorkflowMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.DefaultWorkflowMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.defaultWorkflowMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.defaultWorkflowMeeting
}

func (c *MotionWorkflow) FirstState() *ValueCollection[MotionState, *MotionState] {
	if c.firstState == nil {
		c.firstState = &ValueCollection[MotionState, *MotionState]{
			id:    c.FirstStateID,
			fetch: c.fetch,
		}
	}
	return c.firstState
}

func (c *MotionWorkflow) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionWorkflow) StateList() *RelationList[MotionState, *MotionState] {
	if c.stateList == nil {
		refs := make([]*ValueCollection[MotionState, *MotionState], len(c.StateIDs))
		for i, id := range c.StateIDs {
			refs[i] = &ValueCollection[MotionState, *MotionState]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.stateList = &RelationList[MotionState, *MotionState]{refs}
	}
	return c.stateList
}

func (r *Fetch) MotionWorkflow(id int) *ValueCollection[MotionWorkflow, *MotionWorkflow] {
	return &ValueCollection[MotionWorkflow, *MotionWorkflow]{
		id:    id,
		fetch: r,
	}
}

// MotionWorkingGroupSpeaker has all fields from motion_working_group_speaker.
type MotionWorkingGroupSpeaker struct {
	ID            int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        int
	meeting       *ValueCollection[Meeting, *Meeting]
	meetingUser   *ValueCollection[MeetingUser, *MeetingUser]
	motion        *ValueCollection[Motion, *Motion]
	fetch         *Fetch
}

func (c *MotionWorkingGroupSpeaker) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionWorkingGroupSpeaker_ID(id).Lazy(&c.ID)
	ds.MotionWorkingGroupSpeaker_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionWorkingGroupSpeaker_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionWorkingGroupSpeaker_MotionID(id).Lazy(&c.MotionID)
	ds.MotionWorkingGroupSpeaker_Weight(id).Lazy(&c.Weight)
}

func (c *MotionWorkingGroupSpeaker) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *MotionWorkingGroupSpeaker) MeetingUser() *ValueCollection[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		c.meetingUser = &ValueCollection[MeetingUser, *MeetingUser]{
			id:    c.MeetingUserID,
			fetch: c.fetch,
		}
	}
	return c.meetingUser
}

func (c *MotionWorkingGroupSpeaker) Motion() *ValueCollection[Motion, *Motion] {
	if c.motion == nil {
		c.motion = &ValueCollection[Motion, *Motion]{
			id:    c.MotionID,
			fetch: c.fetch,
		}
	}
	return c.motion
}

func (r *Fetch) MotionWorkingGroupSpeaker(id int) *ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker] {
	return &ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{
		id:    id,
		fetch: r,
	}
}

// Option has all fields from option.
type Option struct {
	Abstain                    string
	ContentObjectID            dsfetch.Maybe[string]
	ID                         int
	MeetingID                  int
	No                         string
	PollID                     dsfetch.Maybe[int]
	Text                       string
	UsedAsGlobalOptionInPollID dsfetch.Maybe[int]
	VoteIDs                    []int
	Weight                     int
	Yes                        string
	meeting                    *ValueCollection[Meeting, *Meeting]
	poll                       *MaybeRelation[Poll, *Poll]
	usedAsGlobalOptionInPoll   *MaybeRelation[Poll, *Poll]
	voteList                   *RelationList[Vote, *Vote]
	fetch                      *Fetch
}

func (c *Option) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Option) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Option) Poll() *MaybeRelation[Poll, *Poll] {
	if c.poll == nil {
		var ref dsfetch.Maybe[*ValueCollection[Poll, *Poll]]
		id, hasValue := c.PollID.Value()
		if hasValue {
			value := &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.poll = &MaybeRelation[Poll, *Poll]{ref}
	}
	return c.poll
}

func (c *Option) UsedAsGlobalOptionInPoll() *MaybeRelation[Poll, *Poll] {
	if c.usedAsGlobalOptionInPoll == nil {
		var ref dsfetch.Maybe[*ValueCollection[Poll, *Poll]]
		id, hasValue := c.UsedAsGlobalOptionInPollID.Value()
		if hasValue {
			value := &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsGlobalOptionInPoll = &MaybeRelation[Poll, *Poll]{ref}
	}
	return c.usedAsGlobalOptionInPoll
}

func (c *Option) VoteList() *RelationList[Vote, *Vote] {
	if c.voteList == nil {
		refs := make([]*ValueCollection[Vote, *Vote], len(c.VoteIDs))
		for i, id := range c.VoteIDs {
			refs[i] = &ValueCollection[Vote, *Vote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.voteList = &RelationList[Vote, *Vote]{refs}
	}
	return c.voteList
}

func (r *Fetch) Option(id int) *ValueCollection[Option, *Option] {
	return &ValueCollection[Option, *Option]{
		id:    id,
		fetch: r,
	}
}

// Organization has all fields from organization.
type Organization struct {
	ActiveMeetingIDs           []int
	ArchivedMeetingIDs         []int
	CommitteeIDs               []int
	DefaultLanguage            string
	Description                string
	EnableAnonymous            bool
	EnableChat                 bool
	EnableElectronicVoting     bool
	GenderIDs                  []int
	ID                         int
	LegalNotice                string
	LimitOfMeetings            int
	LimitOfUsers               int
	LoginText                  string
	MediafileIDs               []int
	Name                       string
	OrganizationTagIDs         []int
	PrivacyPolicy              string
	PublishedMediafileIDs      []int
	RequireDuplicateFrom       bool
	ResetPasswordVerboseErrors bool
	SamlAttrMapping            json.RawMessage
	SamlEnabled                bool
	SamlLoginButtonText        string
	SamlMetadataIDp            string
	SamlMetadataSp             string
	SamlPrivateKey             string
	TemplateMeetingIDs         []int
	ThemeID                    int
	ThemeIDs                   []int
	Url                        string
	UserIDs                    []int
	UsersEmailBody             string
	UsersEmailReplyto          string
	UsersEmailSender           string
	UsersEmailSubject          string
	VoteDecryptPublicMainKey   string
	activeMeetingList          *RelationList[Meeting, *Meeting]
	archivedMeetingList        *RelationList[Meeting, *Meeting]
	committeeList              *RelationList[Committee, *Committee]
	genderList                 *RelationList[Gender, *Gender]
	mediafileList              *RelationList[Mediafile, *Mediafile]
	organizationTagList        *RelationList[OrganizationTag, *OrganizationTag]
	publishedMediafileList     *RelationList[Mediafile, *Mediafile]
	templateMeetingList        *RelationList[Meeting, *Meeting]
	theme                      *ValueCollection[Theme, *Theme]
	themeList                  *RelationList[Theme, *Theme]
	userList                   *RelationList[User, *User]
	fetch                      *Fetch
}

func (c *Organization) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Organization_ActiveMeetingIDs(id).Lazy(&c.ActiveMeetingIDs)
	ds.Organization_ArchivedMeetingIDs(id).Lazy(&c.ArchivedMeetingIDs)
	ds.Organization_CommitteeIDs(id).Lazy(&c.CommitteeIDs)
	ds.Organization_DefaultLanguage(id).Lazy(&c.DefaultLanguage)
	ds.Organization_Description(id).Lazy(&c.Description)
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
	ds.Organization_SamlAttrMapping(id).Lazy(&c.SamlAttrMapping)
	ds.Organization_SamlEnabled(id).Lazy(&c.SamlEnabled)
	ds.Organization_SamlLoginButtonText(id).Lazy(&c.SamlLoginButtonText)
	ds.Organization_SamlMetadataIDp(id).Lazy(&c.SamlMetadataIDp)
	ds.Organization_SamlMetadataSp(id).Lazy(&c.SamlMetadataSp)
	ds.Organization_SamlPrivateKey(id).Lazy(&c.SamlPrivateKey)
	ds.Organization_TemplateMeetingIDs(id).Lazy(&c.TemplateMeetingIDs)
	ds.Organization_ThemeID(id).Lazy(&c.ThemeID)
	ds.Organization_ThemeIDs(id).Lazy(&c.ThemeIDs)
	ds.Organization_Url(id).Lazy(&c.Url)
	ds.Organization_UserIDs(id).Lazy(&c.UserIDs)
	ds.Organization_UsersEmailBody(id).Lazy(&c.UsersEmailBody)
	ds.Organization_UsersEmailReplyto(id).Lazy(&c.UsersEmailReplyto)
	ds.Organization_UsersEmailSender(id).Lazy(&c.UsersEmailSender)
	ds.Organization_UsersEmailSubject(id).Lazy(&c.UsersEmailSubject)
	ds.Organization_VoteDecryptPublicMainKey(id).Lazy(&c.VoteDecryptPublicMainKey)
}

func (c *Organization) ActiveMeetingList() *RelationList[Meeting, *Meeting] {
	if c.activeMeetingList == nil {
		refs := make([]*ValueCollection[Meeting, *Meeting], len(c.ActiveMeetingIDs))
		for i, id := range c.ActiveMeetingIDs {
			refs[i] = &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.activeMeetingList = &RelationList[Meeting, *Meeting]{refs}
	}
	return c.activeMeetingList
}

func (c *Organization) ArchivedMeetingList() *RelationList[Meeting, *Meeting] {
	if c.archivedMeetingList == nil {
		refs := make([]*ValueCollection[Meeting, *Meeting], len(c.ArchivedMeetingIDs))
		for i, id := range c.ArchivedMeetingIDs {
			refs[i] = &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.archivedMeetingList = &RelationList[Meeting, *Meeting]{refs}
	}
	return c.archivedMeetingList
}

func (c *Organization) CommitteeList() *RelationList[Committee, *Committee] {
	if c.committeeList == nil {
		refs := make([]*ValueCollection[Committee, *Committee], len(c.CommitteeIDs))
		for i, id := range c.CommitteeIDs {
			refs[i] = &ValueCollection[Committee, *Committee]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.committeeList = &RelationList[Committee, *Committee]{refs}
	}
	return c.committeeList
}

func (c *Organization) GenderList() *RelationList[Gender, *Gender] {
	if c.genderList == nil {
		refs := make([]*ValueCollection[Gender, *Gender], len(c.GenderIDs))
		for i, id := range c.GenderIDs {
			refs[i] = &ValueCollection[Gender, *Gender]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.genderList = &RelationList[Gender, *Gender]{refs}
	}
	return c.genderList
}

func (c *Organization) MediafileList() *RelationList[Mediafile, *Mediafile] {
	if c.mediafileList == nil {
		refs := make([]*ValueCollection[Mediafile, *Mediafile], len(c.MediafileIDs))
		for i, id := range c.MediafileIDs {
			refs[i] = &ValueCollection[Mediafile, *Mediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.mediafileList = &RelationList[Mediafile, *Mediafile]{refs}
	}
	return c.mediafileList
}

func (c *Organization) OrganizationTagList() *RelationList[OrganizationTag, *OrganizationTag] {
	if c.organizationTagList == nil {
		refs := make([]*ValueCollection[OrganizationTag, *OrganizationTag], len(c.OrganizationTagIDs))
		for i, id := range c.OrganizationTagIDs {
			refs[i] = &ValueCollection[OrganizationTag, *OrganizationTag]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.organizationTagList = &RelationList[OrganizationTag, *OrganizationTag]{refs}
	}
	return c.organizationTagList
}

func (c *Organization) PublishedMediafileList() *RelationList[Mediafile, *Mediafile] {
	if c.publishedMediafileList == nil {
		refs := make([]*ValueCollection[Mediafile, *Mediafile], len(c.PublishedMediafileIDs))
		for i, id := range c.PublishedMediafileIDs {
			refs[i] = &ValueCollection[Mediafile, *Mediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.publishedMediafileList = &RelationList[Mediafile, *Mediafile]{refs}
	}
	return c.publishedMediafileList
}

func (c *Organization) TemplateMeetingList() *RelationList[Meeting, *Meeting] {
	if c.templateMeetingList == nil {
		refs := make([]*ValueCollection[Meeting, *Meeting], len(c.TemplateMeetingIDs))
		for i, id := range c.TemplateMeetingIDs {
			refs[i] = &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.templateMeetingList = &RelationList[Meeting, *Meeting]{refs}
	}
	return c.templateMeetingList
}

func (c *Organization) Theme() *ValueCollection[Theme, *Theme] {
	if c.theme == nil {
		c.theme = &ValueCollection[Theme, *Theme]{
			id:    c.ThemeID,
			fetch: c.fetch,
		}
	}
	return c.theme
}

func (c *Organization) ThemeList() *RelationList[Theme, *Theme] {
	if c.themeList == nil {
		refs := make([]*ValueCollection[Theme, *Theme], len(c.ThemeIDs))
		for i, id := range c.ThemeIDs {
			refs[i] = &ValueCollection[Theme, *Theme]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.themeList = &RelationList[Theme, *Theme]{refs}
	}
	return c.themeList
}

func (c *Organization) UserList() *RelationList[User, *User] {
	if c.userList == nil {
		refs := make([]*ValueCollection[User, *User], len(c.UserIDs))
		for i, id := range c.UserIDs {
			refs[i] = &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.userList = &RelationList[User, *User]{refs}
	}
	return c.userList
}

func (r *Fetch) Organization(id int) *ValueCollection[Organization, *Organization] {
	return &ValueCollection[Organization, *Organization]{
		id:    id,
		fetch: r,
	}
}

// OrganizationTag has all fields from organization_tag.
type OrganizationTag struct {
	Color          string
	ID             int
	Name           string
	OrganizationID int
	TaggedIDs      []string
	organization   *ValueCollection[Organization, *Organization]
	fetch          *Fetch
}

func (c *OrganizationTag) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.OrganizationTag_Color(id).Lazy(&c.Color)
	ds.OrganizationTag_ID(id).Lazy(&c.ID)
	ds.OrganizationTag_Name(id).Lazy(&c.Name)
	ds.OrganizationTag_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.OrganizationTag_TaggedIDs(id).Lazy(&c.TaggedIDs)
}

func (c *OrganizationTag) Organization() *ValueCollection[Organization, *Organization] {
	if c.organization == nil {
		c.organization = &ValueCollection[Organization, *Organization]{
			id:    c.OrganizationID,
			fetch: c.fetch,
		}
	}
	return c.organization
}

func (r *Fetch) OrganizationTag(id int) *ValueCollection[OrganizationTag, *OrganizationTag] {
	return &ValueCollection[OrganizationTag, *OrganizationTag]{
		id:    id,
		fetch: r,
	}
}

// PersonalNote has all fields from personal_note.
type PersonalNote struct {
	ContentObjectID dsfetch.Maybe[string]
	ID              int
	MeetingID       int
	MeetingUserID   int
	Note            string
	Star            bool
	meeting         *ValueCollection[Meeting, *Meeting]
	meetingUser     *ValueCollection[MeetingUser, *MeetingUser]
	fetch           *Fetch
}

func (c *PersonalNote) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PersonalNote_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.PersonalNote_ID(id).Lazy(&c.ID)
	ds.PersonalNote_MeetingID(id).Lazy(&c.MeetingID)
	ds.PersonalNote_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.PersonalNote_Note(id).Lazy(&c.Note)
	ds.PersonalNote_Star(id).Lazy(&c.Star)
}

func (c *PersonalNote) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *PersonalNote) MeetingUser() *ValueCollection[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		c.meetingUser = &ValueCollection[MeetingUser, *MeetingUser]{
			id:    c.MeetingUserID,
			fetch: c.fetch,
		}
	}
	return c.meetingUser
}

func (r *Fetch) PersonalNote(id int) *ValueCollection[PersonalNote, *PersonalNote] {
	return &ValueCollection[PersonalNote, *PersonalNote]{
		id:    id,
		fetch: r,
	}
}

// PointOfOrderCategory has all fields from point_of_order_category.
type PointOfOrderCategory struct {
	ID          int
	MeetingID   int
	Rank        int
	SpeakerIDs  []int
	Text        string
	meeting     *ValueCollection[Meeting, *Meeting]
	speakerList *RelationList[Speaker, *Speaker]
	fetch       *Fetch
}

func (c *PointOfOrderCategory) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PointOfOrderCategory_ID(id).Lazy(&c.ID)
	ds.PointOfOrderCategory_MeetingID(id).Lazy(&c.MeetingID)
	ds.PointOfOrderCategory_Rank(id).Lazy(&c.Rank)
	ds.PointOfOrderCategory_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.PointOfOrderCategory_Text(id).Lazy(&c.Text)
}

func (c *PointOfOrderCategory) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *PointOfOrderCategory) SpeakerList() *RelationList[Speaker, *Speaker] {
	if c.speakerList == nil {
		refs := make([]*ValueCollection[Speaker, *Speaker], len(c.SpeakerIDs))
		for i, id := range c.SpeakerIDs {
			refs[i] = &ValueCollection[Speaker, *Speaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.speakerList = &RelationList[Speaker, *Speaker]{refs}
	}
	return c.speakerList
}

func (r *Fetch) PointOfOrderCategory(id int) *ValueCollection[PointOfOrderCategory, *PointOfOrderCategory] {
	return &ValueCollection[PointOfOrderCategory, *PointOfOrderCategory]{
		id:    id,
		fetch: r,
	}
}

// Poll has all fields from poll.
type Poll struct {
	Backend               string
	ContentObjectID       string
	CryptKey              string
	CryptSignature        string
	Description           string
	EntitledGroupIDs      []int
	EntitledUsersAtStop   json.RawMessage
	GlobalAbstain         bool
	GlobalNo              bool
	GlobalOptionID        dsfetch.Maybe[int]
	GlobalYes             bool
	HasVotedUserIDs       []int
	ID                    int
	IsPseudoanonymized    bool
	MaxVotesAmount        int
	MaxVotesPerOption     int
	MeetingID             int
	MinVotesAmount        int
	OnehundredPercentBase string
	OptionIDs             []int
	Pollmethod            string
	ProjectionIDs         []int
	SequentialNumber      int
	State                 string
	Title                 string
	Type                  string
	VotedIDs              []int
	VotesRaw              string
	VotesSignature        string
	Votescast             string
	Votesinvalid          string
	Votesvalid            string
	entitledGroupList     *RelationList[Group, *Group]
	globalOption          *MaybeRelation[Option, *Option]
	meeting               *ValueCollection[Meeting, *Meeting]
	optionList            *RelationList[Option, *Option]
	projectionList        *RelationList[Projection, *Projection]
	votedList             *RelationList[User, *User]
	fetch                 *Fetch
}

func (c *Poll) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Poll_Backend(id).Lazy(&c.Backend)
	ds.Poll_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Poll_CryptKey(id).Lazy(&c.CryptKey)
	ds.Poll_CryptSignature(id).Lazy(&c.CryptSignature)
	ds.Poll_Description(id).Lazy(&c.Description)
	ds.Poll_EntitledGroupIDs(id).Lazy(&c.EntitledGroupIDs)
	ds.Poll_EntitledUsersAtStop(id).Lazy(&c.EntitledUsersAtStop)
	ds.Poll_GlobalAbstain(id).Lazy(&c.GlobalAbstain)
	ds.Poll_GlobalNo(id).Lazy(&c.GlobalNo)
	ds.Poll_GlobalOptionID(id).Lazy(&c.GlobalOptionID)
	ds.Poll_GlobalYes(id).Lazy(&c.GlobalYes)
	ds.Poll_HasVotedUserIDs(id).Lazy(&c.HasVotedUserIDs)
	ds.Poll_ID(id).Lazy(&c.ID)
	ds.Poll_IsPseudoanonymized(id).Lazy(&c.IsPseudoanonymized)
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
	ds.Poll_VotesRaw(id).Lazy(&c.VotesRaw)
	ds.Poll_VotesSignature(id).Lazy(&c.VotesSignature)
	ds.Poll_Votescast(id).Lazy(&c.Votescast)
	ds.Poll_Votesinvalid(id).Lazy(&c.Votesinvalid)
	ds.Poll_Votesvalid(id).Lazy(&c.Votesvalid)
}

func (c *Poll) EntitledGroupList() *RelationList[Group, *Group] {
	if c.entitledGroupList == nil {
		refs := make([]*ValueCollection[Group, *Group], len(c.EntitledGroupIDs))
		for i, id := range c.EntitledGroupIDs {
			refs[i] = &ValueCollection[Group, *Group]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.entitledGroupList = &RelationList[Group, *Group]{refs}
	}
	return c.entitledGroupList
}

func (c *Poll) GlobalOption() *MaybeRelation[Option, *Option] {
	if c.globalOption == nil {
		var ref dsfetch.Maybe[*ValueCollection[Option, *Option]]
		id, hasValue := c.GlobalOptionID.Value()
		if hasValue {
			value := &ValueCollection[Option, *Option]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.globalOption = &MaybeRelation[Option, *Option]{ref}
	}
	return c.globalOption
}

func (c *Poll) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Poll) OptionList() *RelationList[Option, *Option] {
	if c.optionList == nil {
		refs := make([]*ValueCollection[Option, *Option], len(c.OptionIDs))
		for i, id := range c.OptionIDs {
			refs[i] = &ValueCollection[Option, *Option]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.optionList = &RelationList[Option, *Option]{refs}
	}
	return c.optionList
}

func (c *Poll) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *Poll) VotedList() *RelationList[User, *User] {
	if c.votedList == nil {
		refs := make([]*ValueCollection[User, *User], len(c.VotedIDs))
		for i, id := range c.VotedIDs {
			refs[i] = &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.votedList = &RelationList[User, *User]{refs}
	}
	return c.votedList
}

func (r *Fetch) Poll(id int) *ValueCollection[Poll, *Poll] {
	return &ValueCollection[Poll, *Poll]{
		id:    id,
		fetch: r,
	}
}

// PollCandidate has all fields from poll_candidate.
type PollCandidate struct {
	ID                  int
	MeetingID           int
	PollCandidateListID int
	UserID              dsfetch.Maybe[int]
	Weight              int
	meeting             *ValueCollection[Meeting, *Meeting]
	pollCandidateList   *ValueCollection[PollCandidateList, *PollCandidateList]
	user                *MaybeRelation[User, *User]
	fetch               *Fetch
}

func (c *PollCandidate) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PollCandidate_ID(id).Lazy(&c.ID)
	ds.PollCandidate_MeetingID(id).Lazy(&c.MeetingID)
	ds.PollCandidate_PollCandidateListID(id).Lazy(&c.PollCandidateListID)
	ds.PollCandidate_UserID(id).Lazy(&c.UserID)
	ds.PollCandidate_Weight(id).Lazy(&c.Weight)
}

func (c *PollCandidate) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *PollCandidate) PollCandidateList() *ValueCollection[PollCandidateList, *PollCandidateList] {
	if c.pollCandidateList == nil {
		c.pollCandidateList = &ValueCollection[PollCandidateList, *PollCandidateList]{
			id:    c.PollCandidateListID,
			fetch: c.fetch,
		}
	}
	return c.pollCandidateList
}

func (c *PollCandidate) User() *MaybeRelation[User, *User] {
	if c.user == nil {
		var ref dsfetch.Maybe[*ValueCollection[User, *User]]
		id, hasValue := c.UserID.Value()
		if hasValue {
			value := &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.user = &MaybeRelation[User, *User]{ref}
	}
	return c.user
}

func (r *Fetch) PollCandidate(id int) *ValueCollection[PollCandidate, *PollCandidate] {
	return &ValueCollection[PollCandidate, *PollCandidate]{
		id:    id,
		fetch: r,
	}
}

// PollCandidateList has all fields from poll_candidate_list.
type PollCandidateList struct {
	ID                int
	MeetingID         int
	OptionID          int
	PollCandidateIDs  []int
	meeting           *ValueCollection[Meeting, *Meeting]
	option            *ValueCollection[Option, *Option]
	pollCandidateList *RelationList[PollCandidate, *PollCandidate]
	fetch             *Fetch
}

func (c *PollCandidateList) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PollCandidateList_ID(id).Lazy(&c.ID)
	ds.PollCandidateList_MeetingID(id).Lazy(&c.MeetingID)
	ds.PollCandidateList_OptionID(id).Lazy(&c.OptionID)
	ds.PollCandidateList_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
}

func (c *PollCandidateList) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *PollCandidateList) Option() *ValueCollection[Option, *Option] {
	if c.option == nil {
		c.option = &ValueCollection[Option, *Option]{
			id:    c.OptionID,
			fetch: c.fetch,
		}
	}
	return c.option
}

func (c *PollCandidateList) PollCandidateList() *RelationList[PollCandidate, *PollCandidate] {
	if c.pollCandidateList == nil {
		refs := make([]*ValueCollection[PollCandidate, *PollCandidate], len(c.PollCandidateIDs))
		for i, id := range c.PollCandidateIDs {
			refs[i] = &ValueCollection[PollCandidate, *PollCandidate]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollCandidateList = &RelationList[PollCandidate, *PollCandidate]{refs}
	}
	return c.pollCandidateList
}

func (r *Fetch) PollCandidateList(id int) *ValueCollection[PollCandidateList, *PollCandidateList] {
	return &ValueCollection[PollCandidateList, *PollCandidateList]{
		id:    id,
		fetch: r,
	}
}

// Projection has all fields from projection.
type Projection struct {
	Content            json.RawMessage
	ContentObjectID    string
	CurrentProjectorID dsfetch.Maybe[int]
	HistoryProjectorID dsfetch.Maybe[int]
	ID                 int
	MeetingID          int
	Options            json.RawMessage
	PreviewProjectorID dsfetch.Maybe[int]
	Stable             bool
	Type               string
	Weight             int
	currentProjector   *MaybeRelation[Projector, *Projector]
	historyProjector   *MaybeRelation[Projector, *Projector]
	meeting            *ValueCollection[Meeting, *Meeting]
	previewProjector   *MaybeRelation[Projector, *Projector]
	fetch              *Fetch
}

func (c *Projection) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Projection) CurrentProjector() *MaybeRelation[Projector, *Projector] {
	if c.currentProjector == nil {
		var ref dsfetch.Maybe[*ValueCollection[Projector, *Projector]]
		id, hasValue := c.CurrentProjectorID.Value()
		if hasValue {
			value := &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.currentProjector = &MaybeRelation[Projector, *Projector]{ref}
	}
	return c.currentProjector
}

func (c *Projection) HistoryProjector() *MaybeRelation[Projector, *Projector] {
	if c.historyProjector == nil {
		var ref dsfetch.Maybe[*ValueCollection[Projector, *Projector]]
		id, hasValue := c.HistoryProjectorID.Value()
		if hasValue {
			value := &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.historyProjector = &MaybeRelation[Projector, *Projector]{ref}
	}
	return c.historyProjector
}

func (c *Projection) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Projection) PreviewProjector() *MaybeRelation[Projector, *Projector] {
	if c.previewProjector == nil {
		var ref dsfetch.Maybe[*ValueCollection[Projector, *Projector]]
		id, hasValue := c.PreviewProjectorID.Value()
		if hasValue {
			value := &ValueCollection[Projector, *Projector]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.previewProjector = &MaybeRelation[Projector, *Projector]{ref}
	}
	return c.previewProjector
}

func (r *Fetch) Projection(id int) *ValueCollection[Projection, *Projection] {
	return &ValueCollection[Projection, *Projection]{
		id:    id,
		fetch: r,
	}
}

// Projector has all fields from projector.
type Projector struct {
	AspectRatioDenominator                             int
	AspectRatioNumerator                               int
	BackgroundColor                                    string
	ChyronBackgroundColor                              string
	ChyronBackgroundColor2                             string
	ChyronFontColor                                    string
	ChyronFontColor2                                   string
	Color                                              string
	CurrentProjectionIDs                               []int
	HeaderBackgroundColor                              string
	HeaderFontColor                                    string
	HeaderH1Color                                      string
	HistoryProjectionIDs                               []int
	ID                                                 int
	IsInternal                                         bool
	MeetingID                                          int
	Name                                               string
	PreviewProjectionIDs                               []int
	Scale                                              int
	Scroll                                             int
	SequentialNumber                                   int
	ShowClock                                          bool
	ShowHeaderFooter                                   bool
	ShowLogo                                           bool
	ShowTitle                                          bool
	UsedAsDefaultProjectorForAgendaItemListInMeetingID dsfetch.Maybe[int]
	UsedAsDefaultProjectorForAmendmentInMeetingID      dsfetch.Maybe[int]
	UsedAsDefaultProjectorForAssignmentInMeetingID     dsfetch.Maybe[int]
	UsedAsDefaultProjectorForAssignmentPollInMeetingID dsfetch.Maybe[int]
	UsedAsDefaultProjectorForCountdownInMeetingID      dsfetch.Maybe[int]
	UsedAsDefaultProjectorForCurrentLosInMeetingID     dsfetch.Maybe[int]
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID dsfetch.Maybe[int]
	UsedAsDefaultProjectorForMediafileInMeetingID      dsfetch.Maybe[int]
	UsedAsDefaultProjectorForMessageInMeetingID        dsfetch.Maybe[int]
	UsedAsDefaultProjectorForMotionBlockInMeetingID    dsfetch.Maybe[int]
	UsedAsDefaultProjectorForMotionInMeetingID         dsfetch.Maybe[int]
	UsedAsDefaultProjectorForMotionPollInMeetingID     dsfetch.Maybe[int]
	UsedAsDefaultProjectorForPollInMeetingID           dsfetch.Maybe[int]
	UsedAsDefaultProjectorForTopicInMeetingID          dsfetch.Maybe[int]
	UsedAsReferenceProjectorMeetingID                  dsfetch.Maybe[int]
	Width                                              int
	currentProjectionList                              *RelationList[Projection, *Projection]
	historyProjectionList                              *RelationList[Projection, *Projection]
	meeting                                            *ValueCollection[Meeting, *Meeting]
	previewProjectionList                              *RelationList[Projection, *Projection]
	usedAsDefaultProjectorForAgendaItemListInMeeting   *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForAmendmentInMeeting        *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForAssignmentInMeeting       *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForAssignmentPollInMeeting   *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForCountdownInMeeting        *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForCurrentLosInMeeting       *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForListOfSpeakersInMeeting   *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForMediafileInMeeting        *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForMessageInMeeting          *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForMotionBlockInMeeting      *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForMotionInMeeting           *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForMotionPollInMeeting       *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForPollInMeeting             *MaybeRelation[Meeting, *Meeting]
	usedAsDefaultProjectorForTopicInMeeting            *MaybeRelation[Meeting, *Meeting]
	usedAsReferenceProjectorMeeting                    *MaybeRelation[Meeting, *Meeting]
	fetch                                              *Fetch
}

func (c *Projector) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Projector) CurrentProjectionList() *RelationList[Projection, *Projection] {
	if c.currentProjectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.CurrentProjectionIDs))
		for i, id := range c.CurrentProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.currentProjectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.currentProjectionList
}

func (c *Projector) HistoryProjectionList() *RelationList[Projection, *Projection] {
	if c.historyProjectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.HistoryProjectionIDs))
		for i, id := range c.HistoryProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.historyProjectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.historyProjectionList
}

func (c *Projector) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Projector) PreviewProjectionList() *RelationList[Projection, *Projection] {
	if c.previewProjectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.PreviewProjectionIDs))
		for i, id := range c.PreviewProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.previewProjectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.previewProjectionList
}

func (c *Projector) UsedAsDefaultProjectorForAgendaItemListInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForAgendaItemListInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForAgendaItemListInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForAgendaItemListInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForAgendaItemListInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForAmendmentInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForAmendmentInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForAmendmentInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForAmendmentInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForAmendmentInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForAssignmentInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForAssignmentInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForAssignmentInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForAssignmentInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForAssignmentInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForAssignmentPollInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForAssignmentPollInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForAssignmentPollInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForAssignmentPollInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForAssignmentPollInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForCountdownInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForCountdownInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForCountdownInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForCountdownInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForCountdownInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForCurrentLosInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForCurrentLosInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForCurrentLosInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForCurrentLosInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForCurrentLosInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForListOfSpeakersInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForListOfSpeakersInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForListOfSpeakersInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForListOfSpeakersInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForListOfSpeakersInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForMediafileInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForMediafileInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForMediafileInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForMediafileInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForMediafileInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForMessageInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForMessageInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForMessageInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForMessageInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForMessageInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForMotionBlockInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForMotionBlockInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForMotionBlockInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForMotionBlockInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForMotionBlockInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForMotionInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForMotionInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForMotionInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForMotionInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForMotionInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForMotionPollInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForMotionPollInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForMotionPollInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForMotionPollInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForMotionPollInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForPollInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForPollInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForPollInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForPollInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForPollInMeeting
}

func (c *Projector) UsedAsDefaultProjectorForTopicInMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsDefaultProjectorForTopicInMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsDefaultProjectorForTopicInMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsDefaultProjectorForTopicInMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsDefaultProjectorForTopicInMeeting
}

func (c *Projector) UsedAsReferenceProjectorMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsReferenceProjectorMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsReferenceProjectorMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsReferenceProjectorMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsReferenceProjectorMeeting
}

func (r *Fetch) Projector(id int) *ValueCollection[Projector, *Projector] {
	return &ValueCollection[Projector, *Projector]{
		id:    id,
		fetch: r,
	}
}

// ProjectorCountdown has all fields from projector_countdown.
type ProjectorCountdown struct {
	CountdownTime                          float32
	DefaultTime                            int
	Description                            string
	ID                                     int
	MeetingID                              int
	ProjectionIDs                          []int
	Running                                bool
	Title                                  string
	UsedAsListOfSpeakersCountdownMeetingID dsfetch.Maybe[int]
	UsedAsPollCountdownMeetingID           dsfetch.Maybe[int]
	meeting                                *ValueCollection[Meeting, *Meeting]
	projectionList                         *RelationList[Projection, *Projection]
	usedAsListOfSpeakersCountdownMeeting   *MaybeRelation[Meeting, *Meeting]
	usedAsPollCountdownMeeting             *MaybeRelation[Meeting, *Meeting]
	fetch                                  *Fetch
}

func (c *ProjectorCountdown) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *ProjectorCountdown) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *ProjectorCountdown) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (c *ProjectorCountdown) UsedAsListOfSpeakersCountdownMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsListOfSpeakersCountdownMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsListOfSpeakersCountdownMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsListOfSpeakersCountdownMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsListOfSpeakersCountdownMeeting
}

func (c *ProjectorCountdown) UsedAsPollCountdownMeeting() *MaybeRelation[Meeting, *Meeting] {
	if c.usedAsPollCountdownMeeting == nil {
		var ref dsfetch.Maybe[*ValueCollection[Meeting, *Meeting]]
		id, hasValue := c.UsedAsPollCountdownMeetingID.Value()
		if hasValue {
			value := &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.usedAsPollCountdownMeeting = &MaybeRelation[Meeting, *Meeting]{ref}
	}
	return c.usedAsPollCountdownMeeting
}

func (r *Fetch) ProjectorCountdown(id int) *ValueCollection[ProjectorCountdown, *ProjectorCountdown] {
	return &ValueCollection[ProjectorCountdown, *ProjectorCountdown]{
		id:    id,
		fetch: r,
	}
}

// ProjectorMessage has all fields from projector_message.
type ProjectorMessage struct {
	ID             int
	MeetingID      int
	Message        string
	ProjectionIDs  []int
	meeting        *ValueCollection[Meeting, *Meeting]
	projectionList *RelationList[Projection, *Projection]
	fetch          *Fetch
}

func (c *ProjectorMessage) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ProjectorMessage_ID(id).Lazy(&c.ID)
	ds.ProjectorMessage_MeetingID(id).Lazy(&c.MeetingID)
	ds.ProjectorMessage_Message(id).Lazy(&c.Message)
	ds.ProjectorMessage_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
}

func (c *ProjectorMessage) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *ProjectorMessage) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (r *Fetch) ProjectorMessage(id int) *ValueCollection[ProjectorMessage, *ProjectorMessage] {
	return &ValueCollection[ProjectorMessage, *ProjectorMessage]{
		id:    id,
		fetch: r,
	}
}

// Speaker has all fields from speaker.
type Speaker struct {
	BeginTime                      int
	EndTime                        int
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	MeetingUserID                  dsfetch.Maybe[int]
	Note                           string
	PauseTime                      int
	PointOfOrder                   bool
	PointOfOrderCategoryID         dsfetch.Maybe[int]
	SpeechState                    string
	StructureLevelListOfSpeakersID dsfetch.Maybe[int]
	TotalPause                     int
	UnpauseTime                    int
	Weight                         int
	listOfSpeakers                 *ValueCollection[ListOfSpeakers, *ListOfSpeakers]
	meeting                        *ValueCollection[Meeting, *Meeting]
	meetingUser                    *MaybeRelation[MeetingUser, *MeetingUser]
	pointOfOrderCategory           *MaybeRelation[PointOfOrderCategory, *PointOfOrderCategory]
	structureLevelListOfSpeakers   *MaybeRelation[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]
	fetch                          *Fetch
}

func (c *Speaker) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Speaker) ListOfSpeakers() *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		c.listOfSpeakers = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
			id:    c.ListOfSpeakersID,
			fetch: c.fetch,
		}
	}
	return c.listOfSpeakers
}

func (c *Speaker) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Speaker) MeetingUser() *MaybeRelation[MeetingUser, *MeetingUser] {
	if c.meetingUser == nil {
		var ref dsfetch.Maybe[*ValueCollection[MeetingUser, *MeetingUser]]
		id, hasValue := c.MeetingUserID.Value()
		if hasValue {
			value := &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.meetingUser = &MaybeRelation[MeetingUser, *MeetingUser]{ref}
	}
	return c.meetingUser
}

func (c *Speaker) PointOfOrderCategory() *MaybeRelation[PointOfOrderCategory, *PointOfOrderCategory] {
	if c.pointOfOrderCategory == nil {
		var ref dsfetch.Maybe[*ValueCollection[PointOfOrderCategory, *PointOfOrderCategory]]
		id, hasValue := c.PointOfOrderCategoryID.Value()
		if hasValue {
			value := &ValueCollection[PointOfOrderCategory, *PointOfOrderCategory]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.pointOfOrderCategory = &MaybeRelation[PointOfOrderCategory, *PointOfOrderCategory]{ref}
	}
	return c.pointOfOrderCategory
}

func (c *Speaker) StructureLevelListOfSpeakers() *MaybeRelation[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	if c.structureLevelListOfSpeakers == nil {
		var ref dsfetch.Maybe[*ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]]
		id, hasValue := c.StructureLevelListOfSpeakersID.Value()
		if hasValue {
			value := &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.structureLevelListOfSpeakers = &MaybeRelation[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{ref}
	}
	return c.structureLevelListOfSpeakers
}

func (r *Fetch) Speaker(id int) *ValueCollection[Speaker, *Speaker] {
	return &ValueCollection[Speaker, *Speaker]{
		id:    id,
		fetch: r,
	}
}

// StructureLevel has all fields from structure_level.
type StructureLevel struct {
	Color                            string
	DefaultTime                      int
	ID                               int
	MeetingID                        int
	MeetingUserIDs                   []int
	Name                             string
	StructureLevelListOfSpeakersIDs  []int
	meeting                          *ValueCollection[Meeting, *Meeting]
	meetingUserList                  *RelationList[MeetingUser, *MeetingUser]
	structureLevelListOfSpeakersList *RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]
	fetch                            *Fetch
}

func (c *StructureLevel) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.StructureLevel_Color(id).Lazy(&c.Color)
	ds.StructureLevel_DefaultTime(id).Lazy(&c.DefaultTime)
	ds.StructureLevel_ID(id).Lazy(&c.ID)
	ds.StructureLevel_MeetingID(id).Lazy(&c.MeetingID)
	ds.StructureLevel_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.StructureLevel_Name(id).Lazy(&c.Name)
	ds.StructureLevel_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
}

func (c *StructureLevel) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *StructureLevel) MeetingUserList() *RelationList[MeetingUser, *MeetingUser] {
	if c.meetingUserList == nil {
		refs := make([]*ValueCollection[MeetingUser, *MeetingUser], len(c.MeetingUserIDs))
		for i, id := range c.MeetingUserIDs {
			refs[i] = &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingUserList = &RelationList[MeetingUser, *MeetingUser]{refs}
	}
	return c.meetingUserList
}

func (c *StructureLevel) StructureLevelListOfSpeakersList() *RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	if c.structureLevelListOfSpeakersList == nil {
		refs := make([]*ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers], len(c.StructureLevelListOfSpeakersIDs))
		for i, id := range c.StructureLevelListOfSpeakersIDs {
			refs[i] = &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.structureLevelListOfSpeakersList = &RelationList[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{refs}
	}
	return c.structureLevelListOfSpeakersList
}

func (r *Fetch) StructureLevel(id int) *ValueCollection[StructureLevel, *StructureLevel] {
	return &ValueCollection[StructureLevel, *StructureLevel]{
		id:    id,
		fetch: r,
	}
}

// StructureLevelListOfSpeakers has all fields from structure_level_list_of_speakers.
type StructureLevelListOfSpeakers struct {
	AdditionalTime   float32
	CurrentStartTime int
	ID               int
	InitialTime      int
	ListOfSpeakersID int
	MeetingID        int
	RemainingTime    float32
	SpeakerIDs       []int
	StructureLevelID int
	listOfSpeakers   *ValueCollection[ListOfSpeakers, *ListOfSpeakers]
	meeting          *ValueCollection[Meeting, *Meeting]
	speakerList      *RelationList[Speaker, *Speaker]
	structureLevel   *ValueCollection[StructureLevel, *StructureLevel]
	fetch            *Fetch
}

func (c *StructureLevelListOfSpeakers) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.StructureLevelListOfSpeakers_AdditionalTime(id).Lazy(&c.AdditionalTime)
	ds.StructureLevelListOfSpeakers_CurrentStartTime(id).Lazy(&c.CurrentStartTime)
	ds.StructureLevelListOfSpeakers_ID(id).Lazy(&c.ID)
	ds.StructureLevelListOfSpeakers_InitialTime(id).Lazy(&c.InitialTime)
	ds.StructureLevelListOfSpeakers_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.StructureLevelListOfSpeakers_MeetingID(id).Lazy(&c.MeetingID)
	ds.StructureLevelListOfSpeakers_RemainingTime(id).Lazy(&c.RemainingTime)
	ds.StructureLevelListOfSpeakers_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.StructureLevelListOfSpeakers_StructureLevelID(id).Lazy(&c.StructureLevelID)
}

func (c *StructureLevelListOfSpeakers) ListOfSpeakers() *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		c.listOfSpeakers = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
			id:    c.ListOfSpeakersID,
			fetch: c.fetch,
		}
	}
	return c.listOfSpeakers
}

func (c *StructureLevelListOfSpeakers) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *StructureLevelListOfSpeakers) SpeakerList() *RelationList[Speaker, *Speaker] {
	if c.speakerList == nil {
		refs := make([]*ValueCollection[Speaker, *Speaker], len(c.SpeakerIDs))
		for i, id := range c.SpeakerIDs {
			refs[i] = &ValueCollection[Speaker, *Speaker]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.speakerList = &RelationList[Speaker, *Speaker]{refs}
	}
	return c.speakerList
}

func (c *StructureLevelListOfSpeakers) StructureLevel() *ValueCollection[StructureLevel, *StructureLevel] {
	if c.structureLevel == nil {
		c.structureLevel = &ValueCollection[StructureLevel, *StructureLevel]{
			id:    c.StructureLevelID,
			fetch: c.fetch,
		}
	}
	return c.structureLevel
}

func (r *Fetch) StructureLevelListOfSpeakers(id int) *ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	return &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
		id:    id,
		fetch: r,
	}
}

// Tag has all fields from tag.
type Tag struct {
	ID        int
	MeetingID int
	Name      string
	TaggedIDs []string
	meeting   *ValueCollection[Meeting, *Meeting]
	fetch     *Fetch
}

func (c *Tag) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Tag_ID(id).Lazy(&c.ID)
	ds.Tag_MeetingID(id).Lazy(&c.MeetingID)
	ds.Tag_Name(id).Lazy(&c.Name)
	ds.Tag_TaggedIDs(id).Lazy(&c.TaggedIDs)
}

func (c *Tag) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (r *Fetch) Tag(id int) *ValueCollection[Tag, *Tag] {
	return &ValueCollection[Tag, *Tag]{
		id:    id,
		fetch: r,
	}
}

// Theme has all fields from theme.
type Theme struct {
	Abstain                string
	Accent100              string
	Accent200              string
	Accent300              string
	Accent400              string
	Accent50               string
	Accent500              string
	Accent600              string
	Accent700              string
	Accent800              string
	Accent900              string
	AccentA100             string
	AccentA200             string
	AccentA400             string
	AccentA700             string
	Headbar                string
	ID                     int
	Name                   string
	No                     string
	OrganizationID         int
	Primary100             string
	Primary200             string
	Primary300             string
	Primary400             string
	Primary50              string
	Primary500             string
	Primary600             string
	Primary700             string
	Primary800             string
	Primary900             string
	PrimaryA100            string
	PrimaryA200            string
	PrimaryA400            string
	PrimaryA700            string
	ThemeForOrganizationID dsfetch.Maybe[int]
	Warn100                string
	Warn200                string
	Warn300                string
	Warn400                string
	Warn50                 string
	Warn500                string
	Warn600                string
	Warn700                string
	Warn800                string
	Warn900                string
	WarnA100               string
	WarnA200               string
	WarnA400               string
	WarnA700               string
	Yes                    string
	organization           *ValueCollection[Organization, *Organization]
	themeForOrganization   *MaybeRelation[Organization, *Organization]
	fetch                  *Fetch
}

func (c *Theme) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Theme) Organization() *ValueCollection[Organization, *Organization] {
	if c.organization == nil {
		c.organization = &ValueCollection[Organization, *Organization]{
			id:    c.OrganizationID,
			fetch: c.fetch,
		}
	}
	return c.organization
}

func (c *Theme) ThemeForOrganization() *MaybeRelation[Organization, *Organization] {
	if c.themeForOrganization == nil {
		var ref dsfetch.Maybe[*ValueCollection[Organization, *Organization]]
		id, hasValue := c.ThemeForOrganizationID.Value()
		if hasValue {
			value := &ValueCollection[Organization, *Organization]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.themeForOrganization = &MaybeRelation[Organization, *Organization]{ref}
	}
	return c.themeForOrganization
}

func (r *Fetch) Theme(id int) *ValueCollection[Theme, *Theme] {
	return &ValueCollection[Theme, *Theme]{
		id:    id,
		fetch: r,
	}
}

// Topic has all fields from topic.
type Topic struct {
	AgendaItemID                   int
	AttachmentMeetingMediafileIDs  []int
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	PollIDs                        []int
	ProjectionIDs                  []int
	SequentialNumber               int
	Text                           string
	Title                          string
	agendaItem                     *ValueCollection[AgendaItem, *AgendaItem]
	attachmentMeetingMediafileList *RelationList[MeetingMediafile, *MeetingMediafile]
	listOfSpeakers                 *ValueCollection[ListOfSpeakers, *ListOfSpeakers]
	meeting                        *ValueCollection[Meeting, *Meeting]
	pollList                       *RelationList[Poll, *Poll]
	projectionList                 *RelationList[Projection, *Projection]
	fetch                          *Fetch
}

func (c *Topic) lazy(ds *Fetch, id int) {
	c.fetch = ds
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
}

func (c *Topic) AgendaItem() *ValueCollection[AgendaItem, *AgendaItem] {
	if c.agendaItem == nil {
		c.agendaItem = &ValueCollection[AgendaItem, *AgendaItem]{
			id:    c.AgendaItemID,
			fetch: c.fetch,
		}
	}
	return c.agendaItem
}

func (c *Topic) AttachmentMeetingMediafileList() *RelationList[MeetingMediafile, *MeetingMediafile] {
	if c.attachmentMeetingMediafileList == nil {
		refs := make([]*ValueCollection[MeetingMediafile, *MeetingMediafile], len(c.AttachmentMeetingMediafileIDs))
		for i, id := range c.AttachmentMeetingMediafileIDs {
			refs[i] = &ValueCollection[MeetingMediafile, *MeetingMediafile]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.attachmentMeetingMediafileList = &RelationList[MeetingMediafile, *MeetingMediafile]{refs}
	}
	return c.attachmentMeetingMediafileList
}

func (c *Topic) ListOfSpeakers() *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	if c.listOfSpeakers == nil {
		c.listOfSpeakers = &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
			id:    c.ListOfSpeakersID,
			fetch: c.fetch,
		}
	}
	return c.listOfSpeakers
}

func (c *Topic) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Topic) PollList() *RelationList[Poll, *Poll] {
	if c.pollList == nil {
		refs := make([]*ValueCollection[Poll, *Poll], len(c.PollIDs))
		for i, id := range c.PollIDs {
			refs[i] = &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollList = &RelationList[Poll, *Poll]{refs}
	}
	return c.pollList
}

func (c *Topic) ProjectionList() *RelationList[Projection, *Projection] {
	if c.projectionList == nil {
		refs := make([]*ValueCollection[Projection, *Projection], len(c.ProjectionIDs))
		for i, id := range c.ProjectionIDs {
			refs[i] = &ValueCollection[Projection, *Projection]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.projectionList = &RelationList[Projection, *Projection]{refs}
	}
	return c.projectionList
}

func (r *Fetch) Topic(id int) *ValueCollection[Topic, *Topic] {
	return &ValueCollection[Topic, *Topic]{
		id:    id,
		fetch: r,
	}
}

// User has all fields from user.
type User struct {
	CanChangeOwnPassword        bool
	CommitteeIDs                []int
	CommitteeManagementIDs      []int
	DefaultPassword             string
	DefaultVoteWeight           string
	DelegatedVoteIDs            []int
	Email                       string
	FirstName                   string
	GenderID                    dsfetch.Maybe[int]
	ID                          int
	IsActive                    bool
	IsDemoUser                  bool
	IsPhysicalPerson            bool
	IsPresentInMeetingIDs       []int
	LastEmailSent               int
	LastLogin                   int
	LastName                    string
	MeetingIDs                  []int
	MeetingUserIDs              []int
	MemberNumber                string
	OptionIDs                   []int
	OrganizationID              int
	OrganizationManagementLevel string
	Password                    string
	PollCandidateIDs            []int
	PollVotedIDs                []int
	Pronoun                     string
	SamlID                      string
	Title                       string
	Username                    string
	VoteIDs                     []int
	committeeList               *RelationList[Committee, *Committee]
	committeeManagementList     *RelationList[Committee, *Committee]
	delegatedVoteList           *RelationList[Vote, *Vote]
	gender                      *MaybeRelation[Gender, *Gender]
	isPresentInMeetingList      *RelationList[Meeting, *Meeting]
	meetingUserList             *RelationList[MeetingUser, *MeetingUser]
	optionList                  *RelationList[Option, *Option]
	organization                *ValueCollection[Organization, *Organization]
	pollCandidateList           *RelationList[PollCandidate, *PollCandidate]
	pollVotedList               *RelationList[Poll, *Poll]
	voteList                    *RelationList[Vote, *Vote]
	fetch                       *Fetch
}

func (c *User) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.User_CanChangeOwnPassword(id).Lazy(&c.CanChangeOwnPassword)
	ds.User_CommitteeIDs(id).Lazy(&c.CommitteeIDs)
	ds.User_CommitteeManagementIDs(id).Lazy(&c.CommitteeManagementIDs)
	ds.User_DefaultPassword(id).Lazy(&c.DefaultPassword)
	ds.User_DefaultVoteWeight(id).Lazy(&c.DefaultVoteWeight)
	ds.User_DelegatedVoteIDs(id).Lazy(&c.DelegatedVoteIDs)
	ds.User_Email(id).Lazy(&c.Email)
	ds.User_FirstName(id).Lazy(&c.FirstName)
	ds.User_GenderID(id).Lazy(&c.GenderID)
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
}

func (c *User) CommitteeList() *RelationList[Committee, *Committee] {
	if c.committeeList == nil {
		refs := make([]*ValueCollection[Committee, *Committee], len(c.CommitteeIDs))
		for i, id := range c.CommitteeIDs {
			refs[i] = &ValueCollection[Committee, *Committee]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.committeeList = &RelationList[Committee, *Committee]{refs}
	}
	return c.committeeList
}

func (c *User) CommitteeManagementList() *RelationList[Committee, *Committee] {
	if c.committeeManagementList == nil {
		refs := make([]*ValueCollection[Committee, *Committee], len(c.CommitteeManagementIDs))
		for i, id := range c.CommitteeManagementIDs {
			refs[i] = &ValueCollection[Committee, *Committee]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.committeeManagementList = &RelationList[Committee, *Committee]{refs}
	}
	return c.committeeManagementList
}

func (c *User) DelegatedVoteList() *RelationList[Vote, *Vote] {
	if c.delegatedVoteList == nil {
		refs := make([]*ValueCollection[Vote, *Vote], len(c.DelegatedVoteIDs))
		for i, id := range c.DelegatedVoteIDs {
			refs[i] = &ValueCollection[Vote, *Vote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.delegatedVoteList = &RelationList[Vote, *Vote]{refs}
	}
	return c.delegatedVoteList
}

func (c *User) Gender() *MaybeRelation[Gender, *Gender] {
	if c.gender == nil {
		var ref dsfetch.Maybe[*ValueCollection[Gender, *Gender]]
		id, hasValue := c.GenderID.Value()
		if hasValue {
			value := &ValueCollection[Gender, *Gender]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.gender = &MaybeRelation[Gender, *Gender]{ref}
	}
	return c.gender
}

func (c *User) IsPresentInMeetingList() *RelationList[Meeting, *Meeting] {
	if c.isPresentInMeetingList == nil {
		refs := make([]*ValueCollection[Meeting, *Meeting], len(c.IsPresentInMeetingIDs))
		for i, id := range c.IsPresentInMeetingIDs {
			refs[i] = &ValueCollection[Meeting, *Meeting]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.isPresentInMeetingList = &RelationList[Meeting, *Meeting]{refs}
	}
	return c.isPresentInMeetingList
}

func (c *User) MeetingUserList() *RelationList[MeetingUser, *MeetingUser] {
	if c.meetingUserList == nil {
		refs := make([]*ValueCollection[MeetingUser, *MeetingUser], len(c.MeetingUserIDs))
		for i, id := range c.MeetingUserIDs {
			refs[i] = &ValueCollection[MeetingUser, *MeetingUser]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.meetingUserList = &RelationList[MeetingUser, *MeetingUser]{refs}
	}
	return c.meetingUserList
}

func (c *User) OptionList() *RelationList[Option, *Option] {
	if c.optionList == nil {
		refs := make([]*ValueCollection[Option, *Option], len(c.OptionIDs))
		for i, id := range c.OptionIDs {
			refs[i] = &ValueCollection[Option, *Option]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.optionList = &RelationList[Option, *Option]{refs}
	}
	return c.optionList
}

func (c *User) Organization() *ValueCollection[Organization, *Organization] {
	if c.organization == nil {
		c.organization = &ValueCollection[Organization, *Organization]{
			id:    c.OrganizationID,
			fetch: c.fetch,
		}
	}
	return c.organization
}

func (c *User) PollCandidateList() *RelationList[PollCandidate, *PollCandidate] {
	if c.pollCandidateList == nil {
		refs := make([]*ValueCollection[PollCandidate, *PollCandidate], len(c.PollCandidateIDs))
		for i, id := range c.PollCandidateIDs {
			refs[i] = &ValueCollection[PollCandidate, *PollCandidate]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollCandidateList = &RelationList[PollCandidate, *PollCandidate]{refs}
	}
	return c.pollCandidateList
}

func (c *User) PollVotedList() *RelationList[Poll, *Poll] {
	if c.pollVotedList == nil {
		refs := make([]*ValueCollection[Poll, *Poll], len(c.PollVotedIDs))
		for i, id := range c.PollVotedIDs {
			refs[i] = &ValueCollection[Poll, *Poll]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.pollVotedList = &RelationList[Poll, *Poll]{refs}
	}
	return c.pollVotedList
}

func (c *User) VoteList() *RelationList[Vote, *Vote] {
	if c.voteList == nil {
		refs := make([]*ValueCollection[Vote, *Vote], len(c.VoteIDs))
		for i, id := range c.VoteIDs {
			refs[i] = &ValueCollection[Vote, *Vote]{
				id:    id,
				fetch: c.fetch,
			}
		}
		c.voteList = &RelationList[Vote, *Vote]{refs}
	}
	return c.voteList
}

func (r *Fetch) User(id int) *ValueCollection[User, *User] {
	return &ValueCollection[User, *User]{
		id:    id,
		fetch: r,
	}
}

// Vote has all fields from vote.
type Vote struct {
	DelegatedUserID dsfetch.Maybe[int]
	ID              int
	MeetingID       int
	OptionID        int
	UserID          dsfetch.Maybe[int]
	UserToken       string
	Value           string
	Weight          string
	delegatedUser   *MaybeRelation[User, *User]
	meeting         *ValueCollection[Meeting, *Meeting]
	option          *ValueCollection[Option, *Option]
	user            *MaybeRelation[User, *User]
	fetch           *Fetch
}

func (c *Vote) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Vote_DelegatedUserID(id).Lazy(&c.DelegatedUserID)
	ds.Vote_ID(id).Lazy(&c.ID)
	ds.Vote_MeetingID(id).Lazy(&c.MeetingID)
	ds.Vote_OptionID(id).Lazy(&c.OptionID)
	ds.Vote_UserID(id).Lazy(&c.UserID)
	ds.Vote_UserToken(id).Lazy(&c.UserToken)
	ds.Vote_Value(id).Lazy(&c.Value)
	ds.Vote_Weight(id).Lazy(&c.Weight)
}

func (c *Vote) DelegatedUser() *MaybeRelation[User, *User] {
	if c.delegatedUser == nil {
		var ref dsfetch.Maybe[*ValueCollection[User, *User]]
		id, hasValue := c.DelegatedUserID.Value()
		if hasValue {
			value := &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.delegatedUser = &MaybeRelation[User, *User]{ref}
	}
	return c.delegatedUser
}

func (c *Vote) Meeting() *ValueCollection[Meeting, *Meeting] {
	if c.meeting == nil {
		c.meeting = &ValueCollection[Meeting, *Meeting]{
			id:    c.MeetingID,
			fetch: c.fetch,
		}
	}
	return c.meeting
}

func (c *Vote) Option() *ValueCollection[Option, *Option] {
	if c.option == nil {
		c.option = &ValueCollection[Option, *Option]{
			id:    c.OptionID,
			fetch: c.fetch,
		}
	}
	return c.option
}

func (c *Vote) User() *MaybeRelation[User, *User] {
	if c.user == nil {
		var ref dsfetch.Maybe[*ValueCollection[User, *User]]
		id, hasValue := c.UserID.Value()
		if hasValue {
			value := &ValueCollection[User, *User]{
				id:    id,
				fetch: c.fetch,
			}
			ref.Set(value)
		}
		c.user = &MaybeRelation[User, *User]{ref}
	}
	return c.user
}

func (r *Fetch) Vote(id int) *ValueCollection[Vote, *Vote] {
	return &ValueCollection[Vote, *Vote]{
		id:    id,
		fetch: r,
	}
}
