// Code generated from meta collections. DO NOT EDIT.
package dstypes

// ActionWorker_State represents the ActionWorker_State enum type.
type ActionWorker_State string

const (
	ActionWorker_StateRunning ActionWorker_State = "running"
	ActionWorker_StateEnd     ActionWorker_State = "end"
	ActionWorker_StateAborted ActionWorker_State = "aborted"
)

// AgendaItem_Type represents the AgendaItem_Type enum type.
type AgendaItem_Type string

const (
	AgendaItem_TypeCommon   AgendaItem_Type = "common"
	AgendaItem_TypeInternal AgendaItem_Type = "internal"
	AgendaItem_TypeHidden   AgendaItem_Type = "hidden"
)

// Assignment_Phase represents the Assignment_Phase enum type.
type Assignment_Phase string

const (
	Assignment_PhaseSearch   Assignment_Phase = "search"
	Assignment_PhaseVoting   Assignment_Phase = "voting"
	Assignment_PhaseFinished Assignment_Phase = "finished"
)

// ImportPreview_Name represents the ImportPreview_Name enum type.
type ImportPreview_Name string

const (
	ImportPreview_NameAccount     ImportPreview_Name = "account"
	ImportPreview_NameParticipant ImportPreview_Name = "participant"
	ImportPreview_NameTopic       ImportPreview_Name = "topic"
	ImportPreview_NameCommittee   ImportPreview_Name = "committee"
	ImportPreview_NameMotion      ImportPreview_Name = "motion"
)

// ImportPreview_State represents the ImportPreview_State enum type.
type ImportPreview_State string

const (
	ImportPreview_StateWarning ImportPreview_State = "warning"
	ImportPreview_StateError   ImportPreview_State = "error"
	ImportPreview_StateDone    ImportPreview_State = "done"
)

// Languages represents the Languages enum type.
type Languages string

const (
	LanguagesEn Languages = "en"
	LanguagesDe Languages = "de"
	LanguagesIt Languages = "it"
	LanguagesEs Languages = "es"
	LanguagesRu Languages = "ru"
	LanguagesCs Languages = "cs"
	LanguagesFr Languages = "fr"
)

// MeetingPollSetting_Method represents the MeetingPollSetting_Method enum type.
type MeetingPollSetting_Method string

const (
	MeetingPollSetting_MethodApprovalYesNo              MeetingPollSetting_Method = "approval.yes_no"
	MeetingPollSetting_MethodApprovalYesNoAbstain       MeetingPollSetting_Method = "approval.yes_no_abstain"
	MeetingPollSetting_MethodSelectionYes               MeetingPollSetting_Method = "selection.yes"
	MeetingPollSetting_MethodSelectionNo                MeetingPollSetting_Method = "selection.no"
	MeetingPollSetting_MethodRatingScore                MeetingPollSetting_Method = "rating_score"
	MeetingPollSetting_MethodRatingApprovalYesNo        MeetingPollSetting_Method = "rating_approval.yes_no"
	MeetingPollSetting_MethodRatingApprovalYesNoAbstain MeetingPollSetting_Method = "rating_approval.yes_no_abstain"
)

// Meeting_AgendaItemCreation represents the Meeting_AgendaItemCreation enum type.
type Meeting_AgendaItemCreation string

const (
	Meeting_AgendaItemCreationAlways     Meeting_AgendaItemCreation = "always"
	Meeting_AgendaItemCreationNever      Meeting_AgendaItemCreation = "never"
	Meeting_AgendaItemCreationDefaultYes Meeting_AgendaItemCreation = "default_yes"
	Meeting_AgendaItemCreationDefaultNo  Meeting_AgendaItemCreation = "default_no"
)

// Meeting_AgendaNewItemsDefaultVisibility represents the Meeting_AgendaNewItemsDefaultVisibility enum type.
type Meeting_AgendaNewItemsDefaultVisibility string

const (
	Meeting_AgendaNewItemsDefaultVisibilityCommon   Meeting_AgendaNewItemsDefaultVisibility = "common"
	Meeting_AgendaNewItemsDefaultVisibilityInternal Meeting_AgendaNewItemsDefaultVisibility = "internal"
	Meeting_AgendaNewItemsDefaultVisibilityHidden   Meeting_AgendaNewItemsDefaultVisibility = "hidden"
)

// Meeting_AgendaNumeralSystem represents the Meeting_AgendaNumeralSystem enum type.
type Meeting_AgendaNumeralSystem string

const (
	Meeting_AgendaNumeralSystemArabic Meeting_AgendaNumeralSystem = "arabic"
	Meeting_AgendaNumeralSystemRoman  Meeting_AgendaNumeralSystem = "roman"
)

// Meeting_ApplauseType represents the Meeting_ApplauseType enum type.
type Meeting_ApplauseType string

const (
	Meeting_ApplauseTypeApplauseTypeBar       Meeting_ApplauseType = "applause-type-bar"
	Meeting_ApplauseTypeApplauseTypeParticles Meeting_ApplauseType = "applause-type-particles"
)

// Meeting_ExportCsvEncoding represents the Meeting_ExportCsvEncoding enum type.
type Meeting_ExportCsvEncoding string

const (
	Meeting_ExportCsvEncodingUtf8      Meeting_ExportCsvEncoding = "utf-8"
	Meeting_ExportCsvEncodingIso885915 Meeting_ExportCsvEncoding = "iso-8859-15"
)

// Meeting_ExportPdfPagenumberAlignment represents the Meeting_ExportPdfPagenumberAlignment enum type.
type Meeting_ExportPdfPagenumberAlignment string

const (
	Meeting_ExportPdfPagenumberAlignmentLeft   Meeting_ExportPdfPagenumberAlignment = "left"
	Meeting_ExportPdfPagenumberAlignmentRight  Meeting_ExportPdfPagenumberAlignment = "right"
	Meeting_ExportPdfPagenumberAlignmentCenter Meeting_ExportPdfPagenumberAlignment = "center"
)

// Meeting_ExportPdfPagesize represents the Meeting_ExportPdfPagesize enum type.
type Meeting_ExportPdfPagesize string

const (
	Meeting_ExportPdfPagesizeA4 Meeting_ExportPdfPagesize = "A4"
	Meeting_ExportPdfPagesizeA5 Meeting_ExportPdfPagesize = "A5"
)

// Meeting_MotionsAmendmentsTextMode represents the Meeting_MotionsAmendmentsTextMode enum type.
type Meeting_MotionsAmendmentsTextMode string

const (
	Meeting_MotionsAmendmentsTextModeFreestyle Meeting_MotionsAmendmentsTextMode = "freestyle"
	Meeting_MotionsAmendmentsTextModeFulltext  Meeting_MotionsAmendmentsTextMode = "fulltext"
	Meeting_MotionsAmendmentsTextModeParagraph Meeting_MotionsAmendmentsTextMode = "paragraph"
)

// Meeting_MotionsDefaultLineNumbering represents the Meeting_MotionsDefaultLineNumbering enum type.
type Meeting_MotionsDefaultLineNumbering string

const (
	Meeting_MotionsDefaultLineNumberingOutside Meeting_MotionsDefaultLineNumbering = "outside"
	Meeting_MotionsDefaultLineNumberingInline  Meeting_MotionsDefaultLineNumbering = "inline"
	Meeting_MotionsDefaultLineNumberingNone    Meeting_MotionsDefaultLineNumbering = "none"
)

// Meeting_MotionsDefaultSorting represents the Meeting_MotionsDefaultSorting enum type.
type Meeting_MotionsDefaultSorting string

const (
	Meeting_MotionsDefaultSortingNumber Meeting_MotionsDefaultSorting = "number"
	Meeting_MotionsDefaultSortingWeight Meeting_MotionsDefaultSorting = "weight"
)

// Meeting_MotionsNumberType represents the Meeting_MotionsNumberType enum type.
type Meeting_MotionsNumberType string

const (
	Meeting_MotionsNumberTypePerCategory      Meeting_MotionsNumberType = "per_category"
	Meeting_MotionsNumberTypeSeriallyNumbered Meeting_MotionsNumberType = "serially_numbered"
	Meeting_MotionsNumberTypeManually         Meeting_MotionsNumberType = "manually"
)

// Meeting_MotionsRecommendationTextMode represents the Meeting_MotionsRecommendationTextMode enum type.
type Meeting_MotionsRecommendationTextMode string

const (
	Meeting_MotionsRecommendationTextModeOriginal Meeting_MotionsRecommendationTextMode = "original"
	Meeting_MotionsRecommendationTextModeChanged  Meeting_MotionsRecommendationTextMode = "changed"
	Meeting_MotionsRecommendationTextModeDiff     Meeting_MotionsRecommendationTextMode = "diff"
	Meeting_MotionsRecommendationTextModeAgreed   Meeting_MotionsRecommendationTextMode = "agreed"
)

// Meeting_PollProjectionNameOrderFirst represents the Meeting_PollProjectionNameOrderFirst enum type.
type Meeting_PollProjectionNameOrderFirst string

const (
	Meeting_PollProjectionNameOrderFirstFirstName Meeting_PollProjectionNameOrderFirst = "first_name"
	Meeting_PollProjectionNameOrderFirstLastName  Meeting_PollProjectionNameOrderFirst = "last_name"
)

// Meeting_UsersPdfWlanEncryption represents the Meeting_UsersPdfWlanEncryption enum type.
type Meeting_UsersPdfWlanEncryption string

const (
	Meeting_UsersPdfWlanEncryptionempty  Meeting_UsersPdfWlanEncryption = ""
	Meeting_UsersPdfWlanEncryptionWEP    Meeting_UsersPdfWlanEncryption = "WEP"
	Meeting_UsersPdfWlanEncryptionWPA    Meeting_UsersPdfWlanEncryption = "WPA"
	Meeting_UsersPdfWlanEncryptionNopass Meeting_UsersPdfWlanEncryption = "nopass"
)

// MotionChangeRecommendation_Type represents the MotionChangeRecommendation_Type enum type.
type MotionChangeRecommendation_Type string

const (
	MotionChangeRecommendation_TypeReplacement MotionChangeRecommendation_Type = "replacement"
	MotionChangeRecommendation_TypeInsertion   MotionChangeRecommendation_Type = "insertion"
	MotionChangeRecommendation_TypeDeletion    MotionChangeRecommendation_Type = "deletion"
	MotionChangeRecommendation_TypeOther       MotionChangeRecommendation_Type = "other"
)

// MotionState_CssClass represents the MotionState_CssClass enum type.
type MotionState_CssClass string

const (
	MotionState_CssClassGrey      MotionState_CssClass = "grey"
	MotionState_CssClassRed       MotionState_CssClass = "red"
	MotionState_CssClassGreen     MotionState_CssClass = "green"
	MotionState_CssClassLightblue MotionState_CssClass = "lightblue"
	MotionState_CssClassYellow    MotionState_CssClass = "yellow"
)

// MotionState_MergeAmendmentIntoFinal represents the MotionState_MergeAmendmentIntoFinal enum type.
type MotionState_MergeAmendmentIntoFinal string

const (
	MotionState_MergeAmendmentIntoFinalDoNotMerge MotionState_MergeAmendmentIntoFinal = "do_not_merge"
	MotionState_MergeAmendmentIntoFinalUndefined  MotionState_MergeAmendmentIntoFinal = "undefined"
	MotionState_MergeAmendmentIntoFinalDoMerge    MotionState_MergeAmendmentIntoFinal = "do_merge"
)

// PollConfigApproval_OnehundredPercentBase represents the PollConfigApproval_OnehundredPercentBase enum type.
type PollConfigApproval_OnehundredPercentBase string

const (
	PollConfigApproval_OnehundredPercentBaseYesNo           PollConfigApproval_OnehundredPercentBase = "yes_no"
	PollConfigApproval_OnehundredPercentBaseYesNoAbstain    PollConfigApproval_OnehundredPercentBase = "yes_no_abstain"
	PollConfigApproval_OnehundredPercentBaseValid           PollConfigApproval_OnehundredPercentBase = "valid"
	PollConfigApproval_OnehundredPercentBaseCast            PollConfigApproval_OnehundredPercentBase = "cast"
	PollConfigApproval_OnehundredPercentBaseEntitled        PollConfigApproval_OnehundredPercentBase = "entitled"
	PollConfigApproval_OnehundredPercentBaseEntitledPresent PollConfigApproval_OnehundredPercentBase = "entitled_present"
	PollConfigApproval_OnehundredPercentBaseDisabled        PollConfigApproval_OnehundredPercentBase = "disabled"
)

// PollConfigRatingApproval_OnehundredPercentBase represents the PollConfigRatingApproval_OnehundredPercentBase enum type.
type PollConfigRatingApproval_OnehundredPercentBase string

const (
	PollConfigRatingApproval_OnehundredPercentBaseYesNo           PollConfigRatingApproval_OnehundredPercentBase = "yes_no"
	PollConfigRatingApproval_OnehundredPercentBaseYesNoAbstain    PollConfigRatingApproval_OnehundredPercentBase = "yes_no_abstain"
	PollConfigRatingApproval_OnehundredPercentBaseValid           PollConfigRatingApproval_OnehundredPercentBase = "valid"
	PollConfigRatingApproval_OnehundredPercentBaseCast            PollConfigRatingApproval_OnehundredPercentBase = "cast"
	PollConfigRatingApproval_OnehundredPercentBaseEntitled        PollConfigRatingApproval_OnehundredPercentBase = "entitled"
	PollConfigRatingApproval_OnehundredPercentBaseEntitledPresent PollConfigRatingApproval_OnehundredPercentBase = "entitled_present"
	PollConfigRatingApproval_OnehundredPercentBaseDisabled        PollConfigRatingApproval_OnehundredPercentBase = "disabled"
)

// PollConfigRatingScore_OnehundredPercentBase represents the PollConfigRatingScore_OnehundredPercentBase enum type.
type PollConfigRatingScore_OnehundredPercentBase string

const (
	PollConfigRatingScore_OnehundredPercentBaseNoGeneral       PollConfigRatingScore_OnehundredPercentBase = "no_general"
	PollConfigRatingScore_OnehundredPercentBaseValid           PollConfigRatingScore_OnehundredPercentBase = "valid"
	PollConfigRatingScore_OnehundredPercentBaseCast            PollConfigRatingScore_OnehundredPercentBase = "cast"
	PollConfigRatingScore_OnehundredPercentBaseEntitled        PollConfigRatingScore_OnehundredPercentBase = "entitled"
	PollConfigRatingScore_OnehundredPercentBaseEntitledPresent PollConfigRatingScore_OnehundredPercentBase = "entitled_present"
	PollConfigRatingScore_OnehundredPercentBaseDisabled        PollConfigRatingScore_OnehundredPercentBase = "disabled"
)

// PollConfigSelection_OnehundredPercentBase represents the PollConfigSelection_OnehundredPercentBase enum type.
type PollConfigSelection_OnehundredPercentBase string

const (
	PollConfigSelection_OnehundredPercentBaseNoGeneral       PollConfigSelection_OnehundredPercentBase = "no_general"
	PollConfigSelection_OnehundredPercentBaseValid           PollConfigSelection_OnehundredPercentBase = "valid"
	PollConfigSelection_OnehundredPercentBaseCast            PollConfigSelection_OnehundredPercentBase = "cast"
	PollConfigSelection_OnehundredPercentBaseEntitled        PollConfigSelection_OnehundredPercentBase = "entitled"
	PollConfigSelection_OnehundredPercentBaseEntitledPresent PollConfigSelection_OnehundredPercentBase = "entitled_present"
	PollConfigSelection_OnehundredPercentBaseDisabled        PollConfigSelection_OnehundredPercentBase = "disabled"
)

// PollVisibility represents the PollVisibility enum type.
type PollVisibility string

const (
	PollVisibilityManually PollVisibility = "manually"
	PollVisibilityNamed    PollVisibility = "named"
	PollVisibilityOpen     PollVisibility = "open"
	PollVisibilitySecret   PollVisibility = "secret"
)

// Poll_State represents the Poll_State enum type.
type Poll_State string

const (
	Poll_StateCreated  Poll_State = "created"
	Poll_StateStarted  Poll_State = "started"
	Poll_StateFinished Poll_State = "finished"
)

// RequiredMajority represents the RequiredMajority enum type.
type RequiredMajority string

const (
	RequiredMajorityNoMajority       RequiredMajority = "no_majority"
	RequiredMajorityTwoThirdMajority RequiredMajority = "two_third_majority"
	RequiredMajorityAbsoluteMajority RequiredMajority = "absolute_majority"
)

// RequiredMajoritySelection represents the RequiredMajoritySelection enum type.
type RequiredMajoritySelection string

const (
	RequiredMajoritySelectionNoMajority       RequiredMajoritySelection = "no_majority"
	RequiredMajoritySelectionTwoThirdMajority RequiredMajoritySelection = "two_third_majority"
	RequiredMajoritySelectionAbsoluteMajority RequiredMajoritySelection = "absolute_majority"
	RequiredMajoritySelectionSimpleMajority   RequiredMajoritySelection = "simple_majority"
)

// Speaker_SpeechState represents the Speaker_SpeechState enum type.
type Speaker_SpeechState string

const (
	Speaker_SpeechStateContribution       Speaker_SpeechState = "contribution"
	Speaker_SpeechStatePro                Speaker_SpeechState = "pro"
	Speaker_SpeechStateContra             Speaker_SpeechState = "contra"
	Speaker_SpeechStateIntervention       Speaker_SpeechState = "intervention"
	Speaker_SpeechStateInterposedQuestion Speaker_SpeechState = "interposed_question"
)

// User_OrganizationManagementLevel represents the User_OrganizationManagementLevel enum type.
type User_OrganizationManagementLevel string

const (
	User_OrganizationManagementLevelSuperadmin            User_OrganizationManagementLevel = "superadmin"
	User_OrganizationManagementLevelCanManageOrganization User_OrganizationManagementLevel = "can_manage_organization"
	User_OrganizationManagementLevelCanManageUsers        User_OrganizationManagementLevel = "can_manage_users"
)
