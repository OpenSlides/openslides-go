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

// BallotPaperSelection represents the BallotPaperSelection enum type.
type BallotPaperSelection string

const (
	BallotPaperSelectionNUMBEROFDELEGATES       BallotPaperSelection = "NUMBER_OF_DELEGATES"
	BallotPaperSelectionNUMBEROFALLPARTICIPANTS BallotPaperSelection = "NUMBER_OF_ALL_PARTICIPANTS"
	BallotPaperSelectionCUSTOMNUMBER            BallotPaperSelection = "CUSTOM_NUMBER"
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

// Meeting_MotionPollProjectionNameOrderFirst represents the Meeting_MotionPollProjectionNameOrderFirst enum type.
type Meeting_MotionPollProjectionNameOrderFirst string

const (
	Meeting_MotionPollProjectionNameOrderFirstFirstName Meeting_MotionPollProjectionNameOrderFirst = "first_name"
	Meeting_MotionPollProjectionNameOrderFirstLastName  Meeting_MotionPollProjectionNameOrderFirst = "last_name"
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

// OnehundredPercentBases represents the OnehundredPercentBases enum type.
type OnehundredPercentBases string

const (
	OnehundredPercentBasesY               OnehundredPercentBases = "Y"
	OnehundredPercentBasesYN              OnehundredPercentBases = "YN"
	OnehundredPercentBasesYNA             OnehundredPercentBases = "YNA"
	OnehundredPercentBasesN               OnehundredPercentBases = "N"
	OnehundredPercentBasesValid           OnehundredPercentBases = "valid"
	OnehundredPercentBasesCast            OnehundredPercentBases = "cast"
	OnehundredPercentBasesEntitled        OnehundredPercentBases = "entitled"
	OnehundredPercentBasesEntitledPresent OnehundredPercentBases = "entitled_present"
	OnehundredPercentBasesDisabled        OnehundredPercentBases = "disabled"
)

// PollBackends represents the PollBackends enum type.
type PollBackends string

const (
	PollBackendsLong PollBackends = "long"
	PollBackendsFast PollBackends = "fast"
)

// Poll_Pollmethod represents the Poll_Pollmethod enum type.
type Poll_Pollmethod string

const (
	Poll_PollmethodY   Poll_Pollmethod = "Y"
	Poll_PollmethodYN  Poll_Pollmethod = "YN"
	Poll_PollmethodYNA Poll_Pollmethod = "YNA"
	Poll_PollmethodN   Poll_Pollmethod = "N"
)

// Poll_State represents the Poll_State enum type.
type Poll_State string

const (
	Poll_StateCreated   Poll_State = "created"
	Poll_StateStarted   Poll_State = "started"
	Poll_StateFinished  Poll_State = "finished"
	Poll_StatePublished Poll_State = "published"
)

// Poll_Type represents the Poll_Type enum type.
type Poll_Type string

const (
	Poll_TypeAnalog          Poll_Type = "analog"
	Poll_TypeNamed           Poll_Type = "named"
	Poll_TypePseudoanonymous Poll_Type = "pseudoanonymous"
	Poll_TypeCryptographic   Poll_Type = "cryptographic"
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
