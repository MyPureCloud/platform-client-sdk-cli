package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ViewfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ViewfilterDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Viewfilter
type Viewfilter struct { 
    // MediaTypes - The media types are used to filter the view
    MediaTypes []string `json:"mediaTypes"`


    // QueueIds - The queue ids are used to filter the view
    QueueIds []string `json:"queueIds"`


    // SkillIds - The skill ids are used to filter the view
    SkillIds []string `json:"skillIds"`


    // SkillGroups - The skill groups used to filter the view
    SkillGroups []string `json:"skillGroups"`


    // LanguageIds - The language ids are used to filter the view
    LanguageIds []string `json:"languageIds"`


    // LanguageGroups - The language groups used to filter the view
    LanguageGroups []string `json:"languageGroups"`


    // Directions - The directions are used to filter the view
    Directions []string `json:"directions"`


    // OriginatingDirections - The list of orginating directions used to filter the view
    OriginatingDirections []string `json:"originatingDirections"`


    // WrapUpCodes - The wrap up codes are used to filter the view
    WrapUpCodes []string `json:"wrapUpCodes"`


    // DnisList - The dnis list is used to filter the view
    DnisList []string `json:"dnisList"`


    // SessionDnisList - The list of session dnis used to filter the view
    SessionDnisList []string `json:"sessionDnisList"`


    // FilterQueuesByUserIds - The user ids are used to fetch associated queues for the view
    FilterQueuesByUserIds []string `json:"filterQueuesByUserIds"`


    // FilterUsersByQueueIds - The queue ids are used to fetch associated users for the view
    FilterUsersByQueueIds []string `json:"filterUsersByQueueIds"`


    // UserIds - The user ids are used to filter the view
    UserIds []string `json:"userIds"`


    // ManagementUnitIds - The management unit ids are used to filter the view
    ManagementUnitIds []string `json:"managementUnitIds"`


    // AddressTos - The address To values are used to filter the view
    AddressTos []string `json:"addressTos"`


    // AddressFroms - The address from values are used to filter the view
    AddressFroms []string `json:"addressFroms"`


    // OutboundCampaignIds - The outbound campaign ids are used to filter the view
    OutboundCampaignIds []string `json:"outboundCampaignIds"`


    // OutboundContactListIds - The outbound contact list ids are used to filter the view
    OutboundContactListIds []string `json:"outboundContactListIds"`


    // ContactIds - The contact ids are used to filter the view
    ContactIds []string `json:"contactIds"`


    // ExternalContactIds - The external contact ids are used to filter the view
    ExternalContactIds []string `json:"externalContactIds"`


    // ExternalOrgIds - The external org ids are used to filter the view
    ExternalOrgIds []string `json:"externalOrgIds"`


    // AniList - The ani list ids are used to filter the view
    AniList []string `json:"aniList"`


    // DurationsMilliseconds - The durations in milliseconds used to filter the view
    DurationsMilliseconds []Numericrange `json:"durationsMilliseconds"`


    // AcdDurationsMilliseconds - The acd durations in milliseconds used to filter the view
    AcdDurationsMilliseconds []Numericrange `json:"acdDurationsMilliseconds"`


    // TalkDurationsMilliseconds - The talk durations in milliseconds used to filter the view
    TalkDurationsMilliseconds []Numericrange `json:"talkDurationsMilliseconds"`


    // AcwDurationsMilliseconds - The acw durations in milliseconds used to filter the view
    AcwDurationsMilliseconds []Numericrange `json:"acwDurationsMilliseconds"`


    // HandleDurationsMilliseconds - The handle durations in milliseconds used to filter the view
    HandleDurationsMilliseconds []Numericrange `json:"handleDurationsMilliseconds"`


    // HoldDurationsMilliseconds - The hold durations in milliseconds used to filter the view
    HoldDurationsMilliseconds []Numericrange `json:"holdDurationsMilliseconds"`


    // AbandonDurationsMilliseconds - The abandon durations in milliseconds used to filter the view
    AbandonDurationsMilliseconds []Numericrange `json:"abandonDurationsMilliseconds"`


    // EvaluationScore - The evaluationScore is used to filter the view
    EvaluationScore Numericrange `json:"evaluationScore"`


    // EvaluationCriticalScore - The evaluationCriticalScore is used to filter the view
    EvaluationCriticalScore Numericrange `json:"evaluationCriticalScore"`


    // EvaluationFormIds - The evaluation form ids are used to filter the view
    EvaluationFormIds []string `json:"evaluationFormIds"`


    // EvaluatedAgentIds - The evaluated agent ids are used to filter the view
    EvaluatedAgentIds []string `json:"evaluatedAgentIds"`


    // EvaluatorIds - The evaluator ids are used to filter the view
    EvaluatorIds []string `json:"evaluatorIds"`


    // Transferred - Indicates filtering for transfers
    Transferred bool `json:"transferred"`


    // Abandoned - Indicates filtering for abandons
    Abandoned bool `json:"abandoned"`


    // Answered - Indicates filtering for answered interactions
    Answered bool `json:"answered"`


    // MessageTypes - The message media types used to filter the view
    MessageTypes []string `json:"messageTypes"`


    // DivisionIds - The divison Ids used to filter the view
    DivisionIds []string `json:"divisionIds"`


    // SurveyFormIds - The survey form ids used to filter the view
    SurveyFormIds []string `json:"surveyFormIds"`


    // SurveyTotalScore - The survey total score used to filter the view
    SurveyTotalScore Numericrange `json:"surveyTotalScore"`


    // SurveyNpsScore - The survey NPS score used to filter the view
    SurveyNpsScore Numericrange `json:"surveyNpsScore"`


    // Mos - The desired range for mos values
    Mos Numericrange `json:"mos"`


    // SurveyQuestionGroupScore - The survey question group score used to filter the view
    SurveyQuestionGroupScore Numericrange `json:"surveyQuestionGroupScore"`


    // SurveyPromoterScore - The survey promoter score used to filter the view
    SurveyPromoterScore Numericrange `json:"surveyPromoterScore"`


    // SurveyFormContextIds - The list of survey form context ids used to filter the view
    SurveyFormContextIds []string `json:"surveyFormContextIds"`


    // ConversationIds - The list of conversation ids used to filter the view
    ConversationIds []string `json:"conversationIds"`


    // SipCallIds - The list of SIP call ids used to filter the view
    SipCallIds []string `json:"sipCallIds"`


    // IsEnded - Indicates filtering for ended
    IsEnded bool `json:"isEnded"`


    // IsSurveyed - Indicates filtering for survey
    IsSurveyed bool `json:"isSurveyed"`


    // SurveyScores - The list of survey score ranges used to filter the view
    SurveyScores []Numericrange `json:"surveyScores"`


    // PromoterScores - The list of promoter score ranges used to filter the view
    PromoterScores []Numericrange `json:"promoterScores"`


    // IsCampaign - Indicates filtering for campaign
    IsCampaign bool `json:"isCampaign"`


    // SurveyStatuses - The list of survey statuses used to filter the view
    SurveyStatuses []string `json:"surveyStatuses"`


    // ConversationProperties - A grouping of conversation level filters
    ConversationProperties Conversationproperties `json:"conversationProperties"`


    // IsBlindTransferred - Indicates filtering for blind transferred
    IsBlindTransferred bool `json:"isBlindTransferred"`


    // IsConsulted - Indicates filtering for consulted
    IsConsulted bool `json:"isConsulted"`


    // IsConsultTransferred - Indicates filtering for consult transferred
    IsConsultTransferred bool `json:"isConsultTransferred"`


    // RemoteParticipants - The list of remote participants used to filter the view
    RemoteParticipants []string `json:"remoteParticipants"`


    // FlowIds - The list of flow Ids
    FlowIds []string `json:"flowIds"`


    // FlowOutcomeIds - A list of outcome ids of the flow
    FlowOutcomeIds []string `json:"flowOutcomeIds"`


    // FlowOutcomeValues - A list of outcome values of the flow
    FlowOutcomeValues []string `json:"flowOutcomeValues"`


    // FlowDestinationTypes - The list of destination types of the flow
    FlowDestinationTypes []string `json:"flowDestinationTypes"`


    // FlowDisconnectReasons - The list of reasons for the flow to disconnect
    FlowDisconnectReasons []string `json:"flowDisconnectReasons"`


    // FlowTypes - A list of types of the flow
    FlowTypes []string `json:"flowTypes"`


    // FlowEntryTypes - A list of types of the flow entry
    FlowEntryTypes []string `json:"flowEntryTypes"`


    // FlowEntryReasons - A list of reasons of flow entry
    FlowEntryReasons []string `json:"flowEntryReasons"`


    // FlowVersions - A list of versions of a flow
    FlowVersions []string `json:"flowVersions"`


    // GroupIds - A list of directory group ids
    GroupIds []string `json:"groupIds"`


    // HasJourneyCustomerId - Indicates filtering for journey customer id
    HasJourneyCustomerId bool `json:"hasJourneyCustomerId"`


    // HasJourneyActionMapId - Indicates filtering for Journey action map id
    HasJourneyActionMapId bool `json:"hasJourneyActionMapId"`


    // HasJourneyVisitId - Indicates filtering for Journey visit id
    HasJourneyVisitId bool `json:"hasJourneyVisitId"`


    // HasMedia - Indicates filtering for presence of MMS media
    HasMedia bool `json:"hasMedia"`


    // RoleIds - The role Ids used to filter the view
    RoleIds []string `json:"roleIds"`


    // ReportsTos - The report to user IDs used to filter the view
    ReportsTos []string `json:"reportsTos"`


    // LocationIds - The location Ids used to filter the view
    LocationIds []string `json:"locationIds"`


    // FlowOutTypes - A list of flow out types
    FlowOutTypes []string `json:"flowOutTypes"`


    // ProviderList - A list of providers
    ProviderList []string `json:"providerList"`


    // CallbackNumberList - A list of callback numbers or substrings of numbers (ex: [\"317\", \"13172222222\"])
    CallbackNumberList []string `json:"callbackNumberList"`


    // CallbackInterval - An interval of time to filter for scheduled callbacks. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    CallbackInterval string `json:"callbackInterval"`


    // UsedRoutingTypes - A list of routing types used
    UsedRoutingTypes []string `json:"usedRoutingTypes"`


    // RequestedRoutingTypes - A list of routing types requested
    RequestedRoutingTypes []string `json:"requestedRoutingTypes"`


    // HasAgentAssistId - Indicates filtering for agent assist id
    HasAgentAssistId bool `json:"hasAgentAssistId"`


    // Transcripts - A list of transcript contents requested
    Transcripts []Transcripts `json:"transcripts"`


    // TranscriptLanguages - A list of transcript languages requested
    TranscriptLanguages []string `json:"transcriptLanguages"`


    // ParticipantPurposes - A list of participant purpose requested
    ParticipantPurposes []string `json:"participantPurposes"`


    // ShowFirstQueue - Indicates filtering for first queue data
    ShowFirstQueue bool `json:"showFirstQueue"`


    // TeamIds - The team ids used to filter the view data
    TeamIds []string `json:"teamIds"`


    // FilterUsersByTeamIds - The team ids are used to fetch associated users for the view
    FilterUsersByTeamIds []string `json:"filterUsersByTeamIds"`


    // JourneyActionMapIds - The journey action map ids are used to fetch action maps for the associated view
    JourneyActionMapIds []string `json:"journeyActionMapIds"`


    // JourneyOutcomeIds - The journey outcome ids are used to fetch outcomes for the associated view
    JourneyOutcomeIds []string `json:"journeyOutcomeIds"`


    // JourneySegmentIds - The journey segment ids are used to fetch segments for the associated view
    JourneySegmentIds []string `json:"journeySegmentIds"`


    // JourneyActionMapTypes - The journey action map types are used to filter action map data for the associated view
    JourneyActionMapTypes []string `json:"journeyActionMapTypes"`


    // DevelopmentRoleList - The list of development roles used to filter agent development view
    DevelopmentRoleList []string `json:"developmentRoleList"`


    // DevelopmentTypeList - The list of development types used to filter agent development view
    DevelopmentTypeList []string `json:"developmentTypeList"`


    // DevelopmentStatusList - The list of development status used to filter agent development view
    DevelopmentStatusList []string `json:"developmentStatusList"`


    // DevelopmentModuleIds - The list of development moduleIds used to filter agent development view
    DevelopmentModuleIds []string `json:"developmentModuleIds"`


    // DevelopmentActivityOverdue - Indicates filtering for development activities
    DevelopmentActivityOverdue bool `json:"developmentActivityOverdue"`


    // CustomerSentimentScore - The customer sentiment score used to filter the view
    CustomerSentimentScore Numericrange `json:"customerSentimentScore"`


    // CustomerSentimentTrend - The customer sentiment trend used to filter the view
    CustomerSentimentTrend Numericrange `json:"customerSentimentTrend"`


    // FlowTransferTargets - The list of transfer targets used to filter flow data
    FlowTransferTargets []string `json:"flowTransferTargets"`


    // DevelopmentName - Filter for development name
    DevelopmentName string `json:"developmentName"`


    // TopicIds - Represents the topics detected in the transcript
    TopicIds []string `json:"topicIds"`


    // ExternalTags - The list of external Tags used to filter conversation data
    ExternalTags []string `json:"externalTags"`


    // IsNotResponding - Indicates filtering for not responding users
    IsNotResponding bool `json:"isNotResponding"`


    // IsAuthenticated - Indicates filtering for the authenticated chat
    IsAuthenticated bool `json:"isAuthenticated"`


    // BotIds - The list of bot IDs used to filter bot views
    BotIds []string `json:"botIds"`


    // BotVersions - The list of bot versions used to filter bot views
    BotVersions []string `json:"botVersions"`


    // BotMessageTypes - The list of bot message types used to filter bot views
    BotMessageTypes []string `json:"botMessageTypes"`


    // BotProviderList - The list of bot providers used to filter bot views
    BotProviderList []string `json:"botProviderList"`


    // BotProductList - The list of bot products used to filter bot views
    BotProductList []string `json:"botProductList"`


    // BotRecognitionFailureReasonList - The list of bot recognition failure reasons used to filter bot views
    BotRecognitionFailureReasonList []string `json:"botRecognitionFailureReasonList"`


    // BotIntentList - The list of bot intents used to filter bot views
    BotIntentList []string `json:"botIntentList"`


    // BotFinalIntentList - The list of bot final intents used to filter bot views
    BotFinalIntentList []string `json:"botFinalIntentList"`


    // BotSlotList - The list of bot slots used to filter bot views
    BotSlotList []string `json:"botSlotList"`


    // BotResultList - The list of bot results used to filter bot views
    BotResultList []string `json:"botResultList"`


    // BlockedReasons - The list of blocked reason used to filter action map constraints views
    BlockedReasons []string `json:"blockedReasons"`


    // IsRecorded - Indicates filtering for recorded
    IsRecorded bool `json:"isRecorded"`


    // HasEvaluation - Indicates filtering for evaluation
    HasEvaluation bool `json:"hasEvaluation"`


    // HasScoredEvaluation - Indicates filtering for scored evaluation
    HasScoredEvaluation bool `json:"hasScoredEvaluation"`


    // EmailDeliveryStatusList - The list of email delivery statuses used to filter views
    EmailDeliveryStatusList []string `json:"emailDeliveryStatusList"`


    // IsAgentOwnedCallback - Indicates filtering for agent owned callback interactions
    IsAgentOwnedCallback bool `json:"isAgentOwnedCallback"`


    // AgentCallbackOwnerIds - The list of callback owners used to filter interactions
    AgentCallbackOwnerIds []string `json:"agentCallbackOwnerIds"`


    // TranscriptTopics - The list of transcript topics requested in filter
    TranscriptTopics []Transcripttopics `json:"transcriptTopics"`


    // JourneyFrequencyCapReasons - The list of frequency cap reasons to filter offer constraints
    JourneyFrequencyCapReasons []string `json:"journeyFrequencyCapReasons"`


    // JourneyBlockingActionMapIds - The list of blocking action maps to filter offer constraints
    JourneyBlockingActionMapIds []string `json:"journeyBlockingActionMapIds"`


    // JourneyActionTargetIds - The list of action targets to filter offer constraints
    JourneyActionTargetIds []string `json:"journeyActionTargetIds"`


    // JourneyBlockingScheduleGroupIds - The list of blocking schedule groups to filter offer constraints
    JourneyBlockingScheduleGroupIds []string `json:"journeyBlockingScheduleGroupIds"`


    // JourneyBlockingEmergencyScheduleGroupIds - The list of emergency schedule groups to filter offer constraints
    JourneyBlockingEmergencyScheduleGroupIds []string `json:"journeyBlockingEmergencyScheduleGroupIds"`


    // JourneyUrlEqualConditions - The list of url equal conditions to filter offer constraints
    JourneyUrlEqualConditions []string `json:"journeyUrlEqualConditions"`


    // JourneyUrlNotEqualConditions - The list of url not equal conditions to filter offer constraints
    JourneyUrlNotEqualConditions []string `json:"journeyUrlNotEqualConditions"`


    // JourneyUrlStartsWithConditions - The list of url starts with conditions to filter offer constraints
    JourneyUrlStartsWithConditions []string `json:"journeyUrlStartsWithConditions"`


    // JourneyUrlEndsWithConditions - The list of url ends with conditions to filter offer constraints
    JourneyUrlEndsWithConditions []string `json:"journeyUrlEndsWithConditions"`


    // JourneyUrlContainsAnyConditions - The list of url contains any conditions to filter offer constraints
    JourneyUrlContainsAnyConditions []string `json:"journeyUrlContainsAnyConditions"`


    // JourneyUrlNotContainsAnyConditions - The list of url not contains any conditions to filter offer constraints
    JourneyUrlNotContainsAnyConditions []string `json:"journeyUrlNotContainsAnyConditions"`


    // JourneyUrlContainsAllConditions - The list of url contains all conditions to filter offer constraints
    JourneyUrlContainsAllConditions []string `json:"journeyUrlContainsAllConditions"`


    // JourneyUrlNotContainsAllConditions - The list of url not contains all conditions to filter offer constraints
    JourneyUrlNotContainsAllConditions []string `json:"journeyUrlNotContainsAllConditions"`


    // FlowMilestoneIds - The list of flow milestones to filter exports
    FlowMilestoneIds []string `json:"flowMilestoneIds"`


    // IsAssessmentPassed - Filter to indicate if Agent passed assessment or not
    IsAssessmentPassed bool `json:"isAssessmentPassed"`


    // ConversationInitiators - The list to filter based on Brands (Bot/User/Agent) or End User who initiated the first message in the conversation
    ConversationInitiators []string `json:"conversationInitiators"`


    // HasCustomerParticipated - Indicates if the customer has participated in an initiated conversation
    HasCustomerParticipated bool `json:"hasCustomerParticipated"`


    // IsAcdInteraction - Filter to indicate if interaction was ACD or non-ACD
    IsAcdInteraction bool `json:"isAcdInteraction"`

}

// String returns a JSON representation of the model
func (o *Viewfilter) String() string {
     o.MediaTypes = []string{""} 
     o.QueueIds = []string{""} 
     o.SkillIds = []string{""} 
     o.SkillGroups = []string{""} 
     o.LanguageIds = []string{""} 
     o.LanguageGroups = []string{""} 
     o.Directions = []string{""} 
     o.OriginatingDirections = []string{""} 
     o.WrapUpCodes = []string{""} 
     o.DnisList = []string{""} 
     o.SessionDnisList = []string{""} 
     o.FilterQueuesByUserIds = []string{""} 
     o.FilterUsersByQueueIds = []string{""} 
     o.UserIds = []string{""} 
     o.ManagementUnitIds = []string{""} 
     o.AddressTos = []string{""} 
     o.AddressFroms = []string{""} 
     o.OutboundCampaignIds = []string{""} 
     o.OutboundContactListIds = []string{""} 
     o.ContactIds = []string{""} 
     o.ExternalContactIds = []string{""} 
     o.ExternalOrgIds = []string{""} 
     o.AniList = []string{""} 
     o.DurationsMilliseconds = []Numericrange{{}} 
     o.AcdDurationsMilliseconds = []Numericrange{{}} 
     o.TalkDurationsMilliseconds = []Numericrange{{}} 
     o.AcwDurationsMilliseconds = []Numericrange{{}} 
     o.HandleDurationsMilliseconds = []Numericrange{{}} 
     o.HoldDurationsMilliseconds = []Numericrange{{}} 
     o.AbandonDurationsMilliseconds = []Numericrange{{}} 
    
    
     o.EvaluationFormIds = []string{""} 
     o.EvaluatedAgentIds = []string{""} 
     o.EvaluatorIds = []string{""} 
    
    
    
     o.MessageTypes = []string{""} 
     o.DivisionIds = []string{""} 
     o.SurveyFormIds = []string{""} 
    
    
    
    
    
     o.SurveyFormContextIds = []string{""} 
     o.ConversationIds = []string{""} 
     o.SipCallIds = []string{""} 
    
    
     o.SurveyScores = []Numericrange{{}} 
     o.PromoterScores = []Numericrange{{}} 
    
     o.SurveyStatuses = []string{""} 
    
    
    
    
     o.RemoteParticipants = []string{""} 
     o.FlowIds = []string{""} 
     o.FlowOutcomeIds = []string{""} 
     o.FlowOutcomeValues = []string{""} 
     o.FlowDestinationTypes = []string{""} 
     o.FlowDisconnectReasons = []string{""} 
     o.FlowTypes = []string{""} 
     o.FlowEntryTypes = []string{""} 
     o.FlowEntryReasons = []string{""} 
     o.FlowVersions = []string{""} 
     o.GroupIds = []string{""} 
    
    
    
    
     o.RoleIds = []string{""} 
     o.ReportsTos = []string{""} 
     o.LocationIds = []string{""} 
     o.FlowOutTypes = []string{""} 
     o.ProviderList = []string{""} 
     o.CallbackNumberList = []string{""} 
    
     o.UsedRoutingTypes = []string{""} 
     o.RequestedRoutingTypes = []string{""} 
    
     o.Transcripts = []Transcripts{{}} 
     o.TranscriptLanguages = []string{""} 
     o.ParticipantPurposes = []string{""} 
    
     o.TeamIds = []string{""} 
     o.FilterUsersByTeamIds = []string{""} 
     o.JourneyActionMapIds = []string{""} 
     o.JourneyOutcomeIds = []string{""} 
     o.JourneySegmentIds = []string{""} 
     o.JourneyActionMapTypes = []string{""} 
     o.DevelopmentRoleList = []string{""} 
     o.DevelopmentTypeList = []string{""} 
     o.DevelopmentStatusList = []string{""} 
     o.DevelopmentModuleIds = []string{""} 
    
    
    
     o.FlowTransferTargets = []string{""} 
    
     o.TopicIds = []string{""} 
     o.ExternalTags = []string{""} 
    
    
     o.BotIds = []string{""} 
     o.BotVersions = []string{""} 
     o.BotMessageTypes = []string{""} 
     o.BotProviderList = []string{""} 
     o.BotProductList = []string{""} 
     o.BotRecognitionFailureReasonList = []string{""} 
     o.BotIntentList = []string{""} 
     o.BotFinalIntentList = []string{""} 
     o.BotSlotList = []string{""} 
     o.BotResultList = []string{""} 
     o.BlockedReasons = []string{""} 
    
    
    
     o.EmailDeliveryStatusList = []string{""} 
    
     o.AgentCallbackOwnerIds = []string{""} 
     o.TranscriptTopics = []Transcripttopics{{}} 
     o.JourneyFrequencyCapReasons = []string{""} 
     o.JourneyBlockingActionMapIds = []string{""} 
     o.JourneyActionTargetIds = []string{""} 
     o.JourneyBlockingScheduleGroupIds = []string{""} 
     o.JourneyBlockingEmergencyScheduleGroupIds = []string{""} 
     o.JourneyUrlEqualConditions = []string{""} 
     o.JourneyUrlNotEqualConditions = []string{""} 
     o.JourneyUrlStartsWithConditions = []string{""} 
     o.JourneyUrlEndsWithConditions = []string{""} 
     o.JourneyUrlContainsAnyConditions = []string{""} 
     o.JourneyUrlNotContainsAnyConditions = []string{""} 
     o.JourneyUrlContainsAllConditions = []string{""} 
     o.JourneyUrlNotContainsAllConditions = []string{""} 
     o.FlowMilestoneIds = []string{""} 
    
     o.ConversationInitiators = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Viewfilter) MarshalJSON() ([]byte, error) {
    type Alias Viewfilter

    if ViewfilterMarshalled {
        return []byte("{}"), nil
    }
    ViewfilterMarshalled = true

    return json.Marshal(&struct {
        
        MediaTypes []string `json:"mediaTypes"`
        
        QueueIds []string `json:"queueIds"`
        
        SkillIds []string `json:"skillIds"`
        
        SkillGroups []string `json:"skillGroups"`
        
        LanguageIds []string `json:"languageIds"`
        
        LanguageGroups []string `json:"languageGroups"`
        
        Directions []string `json:"directions"`
        
        OriginatingDirections []string `json:"originatingDirections"`
        
        WrapUpCodes []string `json:"wrapUpCodes"`
        
        DnisList []string `json:"dnisList"`
        
        SessionDnisList []string `json:"sessionDnisList"`
        
        FilterQueuesByUserIds []string `json:"filterQueuesByUserIds"`
        
        FilterUsersByQueueIds []string `json:"filterUsersByQueueIds"`
        
        UserIds []string `json:"userIds"`
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        AddressTos []string `json:"addressTos"`
        
        AddressFroms []string `json:"addressFroms"`
        
        OutboundCampaignIds []string `json:"outboundCampaignIds"`
        
        OutboundContactListIds []string `json:"outboundContactListIds"`
        
        ContactIds []string `json:"contactIds"`
        
        ExternalContactIds []string `json:"externalContactIds"`
        
        ExternalOrgIds []string `json:"externalOrgIds"`
        
        AniList []string `json:"aniList"`
        
        DurationsMilliseconds []Numericrange `json:"durationsMilliseconds"`
        
        AcdDurationsMilliseconds []Numericrange `json:"acdDurationsMilliseconds"`
        
        TalkDurationsMilliseconds []Numericrange `json:"talkDurationsMilliseconds"`
        
        AcwDurationsMilliseconds []Numericrange `json:"acwDurationsMilliseconds"`
        
        HandleDurationsMilliseconds []Numericrange `json:"handleDurationsMilliseconds"`
        
        HoldDurationsMilliseconds []Numericrange `json:"holdDurationsMilliseconds"`
        
        AbandonDurationsMilliseconds []Numericrange `json:"abandonDurationsMilliseconds"`
        
        EvaluationScore Numericrange `json:"evaluationScore"`
        
        EvaluationCriticalScore Numericrange `json:"evaluationCriticalScore"`
        
        EvaluationFormIds []string `json:"evaluationFormIds"`
        
        EvaluatedAgentIds []string `json:"evaluatedAgentIds"`
        
        EvaluatorIds []string `json:"evaluatorIds"`
        
        Transferred bool `json:"transferred"`
        
        Abandoned bool `json:"abandoned"`
        
        Answered bool `json:"answered"`
        
        MessageTypes []string `json:"messageTypes"`
        
        DivisionIds []string `json:"divisionIds"`
        
        SurveyFormIds []string `json:"surveyFormIds"`
        
        SurveyTotalScore Numericrange `json:"surveyTotalScore"`
        
        SurveyNpsScore Numericrange `json:"surveyNpsScore"`
        
        Mos Numericrange `json:"mos"`
        
        SurveyQuestionGroupScore Numericrange `json:"surveyQuestionGroupScore"`
        
        SurveyPromoterScore Numericrange `json:"surveyPromoterScore"`
        
        SurveyFormContextIds []string `json:"surveyFormContextIds"`
        
        ConversationIds []string `json:"conversationIds"`
        
        SipCallIds []string `json:"sipCallIds"`
        
        IsEnded bool `json:"isEnded"`
        
        IsSurveyed bool `json:"isSurveyed"`
        
        SurveyScores []Numericrange `json:"surveyScores"`
        
        PromoterScores []Numericrange `json:"promoterScores"`
        
        IsCampaign bool `json:"isCampaign"`
        
        SurveyStatuses []string `json:"surveyStatuses"`
        
        ConversationProperties Conversationproperties `json:"conversationProperties"`
        
        IsBlindTransferred bool `json:"isBlindTransferred"`
        
        IsConsulted bool `json:"isConsulted"`
        
        IsConsultTransferred bool `json:"isConsultTransferred"`
        
        RemoteParticipants []string `json:"remoteParticipants"`
        
        FlowIds []string `json:"flowIds"`
        
        FlowOutcomeIds []string `json:"flowOutcomeIds"`
        
        FlowOutcomeValues []string `json:"flowOutcomeValues"`
        
        FlowDestinationTypes []string `json:"flowDestinationTypes"`
        
        FlowDisconnectReasons []string `json:"flowDisconnectReasons"`
        
        FlowTypes []string `json:"flowTypes"`
        
        FlowEntryTypes []string `json:"flowEntryTypes"`
        
        FlowEntryReasons []string `json:"flowEntryReasons"`
        
        FlowVersions []string `json:"flowVersions"`
        
        GroupIds []string `json:"groupIds"`
        
        HasJourneyCustomerId bool `json:"hasJourneyCustomerId"`
        
        HasJourneyActionMapId bool `json:"hasJourneyActionMapId"`
        
        HasJourneyVisitId bool `json:"hasJourneyVisitId"`
        
        HasMedia bool `json:"hasMedia"`
        
        RoleIds []string `json:"roleIds"`
        
        ReportsTos []string `json:"reportsTos"`
        
        LocationIds []string `json:"locationIds"`
        
        FlowOutTypes []string `json:"flowOutTypes"`
        
        ProviderList []string `json:"providerList"`
        
        CallbackNumberList []string `json:"callbackNumberList"`
        
        CallbackInterval string `json:"callbackInterval"`
        
        UsedRoutingTypes []string `json:"usedRoutingTypes"`
        
        RequestedRoutingTypes []string `json:"requestedRoutingTypes"`
        
        HasAgentAssistId bool `json:"hasAgentAssistId"`
        
        Transcripts []Transcripts `json:"transcripts"`
        
        TranscriptLanguages []string `json:"transcriptLanguages"`
        
        ParticipantPurposes []string `json:"participantPurposes"`
        
        ShowFirstQueue bool `json:"showFirstQueue"`
        
        TeamIds []string `json:"teamIds"`
        
        FilterUsersByTeamIds []string `json:"filterUsersByTeamIds"`
        
        JourneyActionMapIds []string `json:"journeyActionMapIds"`
        
        JourneyOutcomeIds []string `json:"journeyOutcomeIds"`
        
        JourneySegmentIds []string `json:"journeySegmentIds"`
        
        JourneyActionMapTypes []string `json:"journeyActionMapTypes"`
        
        DevelopmentRoleList []string `json:"developmentRoleList"`
        
        DevelopmentTypeList []string `json:"developmentTypeList"`
        
        DevelopmentStatusList []string `json:"developmentStatusList"`
        
        DevelopmentModuleIds []string `json:"developmentModuleIds"`
        
        DevelopmentActivityOverdue bool `json:"developmentActivityOverdue"`
        
        CustomerSentimentScore Numericrange `json:"customerSentimentScore"`
        
        CustomerSentimentTrend Numericrange `json:"customerSentimentTrend"`
        
        FlowTransferTargets []string `json:"flowTransferTargets"`
        
        DevelopmentName string `json:"developmentName"`
        
        TopicIds []string `json:"topicIds"`
        
        ExternalTags []string `json:"externalTags"`
        
        IsNotResponding bool `json:"isNotResponding"`
        
        IsAuthenticated bool `json:"isAuthenticated"`
        
        BotIds []string `json:"botIds"`
        
        BotVersions []string `json:"botVersions"`
        
        BotMessageTypes []string `json:"botMessageTypes"`
        
        BotProviderList []string `json:"botProviderList"`
        
        BotProductList []string `json:"botProductList"`
        
        BotRecognitionFailureReasonList []string `json:"botRecognitionFailureReasonList"`
        
        BotIntentList []string `json:"botIntentList"`
        
        BotFinalIntentList []string `json:"botFinalIntentList"`
        
        BotSlotList []string `json:"botSlotList"`
        
        BotResultList []string `json:"botResultList"`
        
        BlockedReasons []string `json:"blockedReasons"`
        
        IsRecorded bool `json:"isRecorded"`
        
        HasEvaluation bool `json:"hasEvaluation"`
        
        HasScoredEvaluation bool `json:"hasScoredEvaluation"`
        
        EmailDeliveryStatusList []string `json:"emailDeliveryStatusList"`
        
        IsAgentOwnedCallback bool `json:"isAgentOwnedCallback"`
        
        AgentCallbackOwnerIds []string `json:"agentCallbackOwnerIds"`
        
        TranscriptTopics []Transcripttopics `json:"transcriptTopics"`
        
        JourneyFrequencyCapReasons []string `json:"journeyFrequencyCapReasons"`
        
        JourneyBlockingActionMapIds []string `json:"journeyBlockingActionMapIds"`
        
        JourneyActionTargetIds []string `json:"journeyActionTargetIds"`
        
        JourneyBlockingScheduleGroupIds []string `json:"journeyBlockingScheduleGroupIds"`
        
        JourneyBlockingEmergencyScheduleGroupIds []string `json:"journeyBlockingEmergencyScheduleGroupIds"`
        
        JourneyUrlEqualConditions []string `json:"journeyUrlEqualConditions"`
        
        JourneyUrlNotEqualConditions []string `json:"journeyUrlNotEqualConditions"`
        
        JourneyUrlStartsWithConditions []string `json:"journeyUrlStartsWithConditions"`
        
        JourneyUrlEndsWithConditions []string `json:"journeyUrlEndsWithConditions"`
        
        JourneyUrlContainsAnyConditions []string `json:"journeyUrlContainsAnyConditions"`
        
        JourneyUrlNotContainsAnyConditions []string `json:"journeyUrlNotContainsAnyConditions"`
        
        JourneyUrlContainsAllConditions []string `json:"journeyUrlContainsAllConditions"`
        
        JourneyUrlNotContainsAllConditions []string `json:"journeyUrlNotContainsAllConditions"`
        
        FlowMilestoneIds []string `json:"flowMilestoneIds"`
        
        IsAssessmentPassed bool `json:"isAssessmentPassed"`
        
        ConversationInitiators []string `json:"conversationInitiators"`
        
        HasCustomerParticipated bool `json:"hasCustomerParticipated"`
        
        IsAcdInteraction bool `json:"isAcdInteraction"`
        *Alias
    }{

        
        MediaTypes: []string{""},
        


        
        QueueIds: []string{""},
        


        
        SkillIds: []string{""},
        


        
        SkillGroups: []string{""},
        


        
        LanguageIds: []string{""},
        


        
        LanguageGroups: []string{""},
        


        
        Directions: []string{""},
        


        
        OriginatingDirections: []string{""},
        


        
        WrapUpCodes: []string{""},
        


        
        DnisList: []string{""},
        


        
        SessionDnisList: []string{""},
        


        
        FilterQueuesByUserIds: []string{""},
        


        
        FilterUsersByQueueIds: []string{""},
        


        
        UserIds: []string{""},
        


        
        ManagementUnitIds: []string{""},
        


        
        AddressTos: []string{""},
        


        
        AddressFroms: []string{""},
        


        
        OutboundCampaignIds: []string{""},
        


        
        OutboundContactListIds: []string{""},
        


        
        ContactIds: []string{""},
        


        
        ExternalContactIds: []string{""},
        


        
        ExternalOrgIds: []string{""},
        


        
        AniList: []string{""},
        


        
        DurationsMilliseconds: []Numericrange{{}},
        


        
        AcdDurationsMilliseconds: []Numericrange{{}},
        


        
        TalkDurationsMilliseconds: []Numericrange{{}},
        


        
        AcwDurationsMilliseconds: []Numericrange{{}},
        


        
        HandleDurationsMilliseconds: []Numericrange{{}},
        


        
        HoldDurationsMilliseconds: []Numericrange{{}},
        


        
        AbandonDurationsMilliseconds: []Numericrange{{}},
        


        


        


        
        EvaluationFormIds: []string{""},
        


        
        EvaluatedAgentIds: []string{""},
        


        
        EvaluatorIds: []string{""},
        


        


        


        


        
        MessageTypes: []string{""},
        


        
        DivisionIds: []string{""},
        


        
        SurveyFormIds: []string{""},
        


        


        


        


        


        


        
        SurveyFormContextIds: []string{""},
        


        
        ConversationIds: []string{""},
        


        
        SipCallIds: []string{""},
        


        


        


        
        SurveyScores: []Numericrange{{}},
        


        
        PromoterScores: []Numericrange{{}},
        


        


        
        SurveyStatuses: []string{""},
        


        


        


        


        


        
        RemoteParticipants: []string{""},
        


        
        FlowIds: []string{""},
        


        
        FlowOutcomeIds: []string{""},
        


        
        FlowOutcomeValues: []string{""},
        


        
        FlowDestinationTypes: []string{""},
        


        
        FlowDisconnectReasons: []string{""},
        


        
        FlowTypes: []string{""},
        


        
        FlowEntryTypes: []string{""},
        


        
        FlowEntryReasons: []string{""},
        


        
        FlowVersions: []string{""},
        


        
        GroupIds: []string{""},
        


        


        


        


        


        
        RoleIds: []string{""},
        


        
        ReportsTos: []string{""},
        


        
        LocationIds: []string{""},
        


        
        FlowOutTypes: []string{""},
        


        
        ProviderList: []string{""},
        


        
        CallbackNumberList: []string{""},
        


        


        
        UsedRoutingTypes: []string{""},
        


        
        RequestedRoutingTypes: []string{""},
        


        


        
        Transcripts: []Transcripts{{}},
        


        
        TranscriptLanguages: []string{""},
        


        
        ParticipantPurposes: []string{""},
        


        


        
        TeamIds: []string{""},
        


        
        FilterUsersByTeamIds: []string{""},
        


        
        JourneyActionMapIds: []string{""},
        


        
        JourneyOutcomeIds: []string{""},
        


        
        JourneySegmentIds: []string{""},
        


        
        JourneyActionMapTypes: []string{""},
        


        
        DevelopmentRoleList: []string{""},
        


        
        DevelopmentTypeList: []string{""},
        


        
        DevelopmentStatusList: []string{""},
        


        
        DevelopmentModuleIds: []string{""},
        


        


        


        


        
        FlowTransferTargets: []string{""},
        


        


        
        TopicIds: []string{""},
        


        
        ExternalTags: []string{""},
        


        


        


        
        BotIds: []string{""},
        


        
        BotVersions: []string{""},
        


        
        BotMessageTypes: []string{""},
        


        
        BotProviderList: []string{""},
        


        
        BotProductList: []string{""},
        


        
        BotRecognitionFailureReasonList: []string{""},
        


        
        BotIntentList: []string{""},
        


        
        BotFinalIntentList: []string{""},
        


        
        BotSlotList: []string{""},
        


        
        BotResultList: []string{""},
        


        
        BlockedReasons: []string{""},
        


        


        


        


        
        EmailDeliveryStatusList: []string{""},
        


        


        
        AgentCallbackOwnerIds: []string{""},
        


        
        TranscriptTopics: []Transcripttopics{{}},
        


        
        JourneyFrequencyCapReasons: []string{""},
        


        
        JourneyBlockingActionMapIds: []string{""},
        


        
        JourneyActionTargetIds: []string{""},
        


        
        JourneyBlockingScheduleGroupIds: []string{""},
        


        
        JourneyBlockingEmergencyScheduleGroupIds: []string{""},
        


        
        JourneyUrlEqualConditions: []string{""},
        


        
        JourneyUrlNotEqualConditions: []string{""},
        


        
        JourneyUrlStartsWithConditions: []string{""},
        


        
        JourneyUrlEndsWithConditions: []string{""},
        


        
        JourneyUrlContainsAnyConditions: []string{""},
        


        
        JourneyUrlNotContainsAnyConditions: []string{""},
        


        
        JourneyUrlContainsAllConditions: []string{""},
        


        
        JourneyUrlNotContainsAllConditions: []string{""},
        


        
        FlowMilestoneIds: []string{""},
        


        


        
        ConversationInitiators: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

