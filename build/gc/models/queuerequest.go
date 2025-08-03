package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueuerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueuerequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    MemberCount int `json:"memberCount"`


    UserMemberCount int `json:"userMemberCount"`


    JoinedMemberCount int `json:"joinedMemberCount"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Queuerequest
type Queuerequest struct { 
    


    // Name - The queue name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - The queue description.
    Description string `json:"description"`


    // DateCreated - The date the queue was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the queue. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the queue.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the queue.
    CreatedBy string `json:"createdBy"`


    


    


    


    // MediaSettings - The media settings for the queue.
    MediaSettings Queuemediasettings `json:"mediaSettings"`


    // RoutingRules - The routing rules for the queue, used for Preferred Agent Routing.
    RoutingRules []Routingrule `json:"routingRules"`


    // ConditionalGroupRouting - The Conditional Group Routing settings for the queue.
    ConditionalGroupRouting Conditionalgrouprouting `json:"conditionalGroupRouting"`


    // ConditionalGroupActivation - The Conditional Group Activation settings for the queue.
    ConditionalGroupActivation Conditionalgroupactivation `json:"conditionalGroupActivation"`


    // Bullseye - The bullseye settings for the queue.
    Bullseye Bullseye `json:"bullseye"`


    // ScoringMethod - The Scoring Method for the queue.
    ScoringMethod string `json:"scoringMethod"`


    // LastAgentRoutingMode - The Last Agent Routing Mode for the queue.
    LastAgentRoutingMode string `json:"lastAgentRoutingMode"`


    // AcwSettings - The ACW settings for the queue.
    AcwSettings Acwsettings `json:"acwSettings"`


    // SkillEvaluationMethod - The skill evaluation method to use when routing conversations.
    SkillEvaluationMethod string `json:"skillEvaluationMethod"`


    // MemberGroups - The groups of agents associated with the queue, if any.  Queue membership will update to match group membership changes.
    MemberGroups []Membergroup `json:"memberGroups"`


    // QueueFlow - The in-queue flow to use for call conversations waiting in queue.
    QueueFlow Domainentityref `json:"queueFlow"`


    // EmailInQueueFlow - The in-queue flow to use for email conversations waiting in queue.
    EmailInQueueFlow Domainentityref `json:"emailInQueueFlow"`


    // MessageInQueueFlow - The in-queue flow to use for message conversations waiting in queue.
    MessageInQueueFlow Domainentityref `json:"messageInQueueFlow"`


    // WhisperPrompt - The prompt used for whisper on the queue, if configured.
    WhisperPrompt Domainentityref `json:"whisperPrompt"`


    // OnHoldPrompt - The audio to be played when calls on this queue are on hold. If not configured, the default on-hold music will play.
    OnHoldPrompt Domainentityref `json:"onHoldPrompt"`


    // AutoAnswerOnly - Specifies whether the configured whisper should play for all ACD calls, or only for those which are auto-answered.
    AutoAnswerOnly bool `json:"autoAnswerOnly"`


    // CannedResponseLibraries - Canned response library IDs and mode with which they are associated with the queue
    CannedResponseLibraries Cannedresponselibraries `json:"cannedResponseLibraries"`


    // EnableTranscription - Indicates whether voice transcription is enabled for this queue.
    EnableTranscription bool `json:"enableTranscription"`


    // EnableAudioMonitoring - Indicates whether audio monitoring is enabled for this queue.
    EnableAudioMonitoring bool `json:"enableAudioMonitoring"`


    // EnableManualAssignment - Indicates whether manual assignment is enabled for this queue.
    EnableManualAssignment bool `json:"enableManualAssignment"`


    // AgentOwnedRouting - The Agent Owned Routing settings for the queue
    AgentOwnedRouting Agentownedrouting `json:"agentOwnedRouting"`


    // DirectRouting - The Direct Routing settings for the queue
    DirectRouting Directrouting `json:"directRouting"`


    // CallingPartyName - The name to use for caller identification for outbound calls from this queue.
    CallingPartyName string `json:"callingPartyName"`


    // CallingPartyNumber - The phone number to use for caller identification for outbound calls from this queue.
    CallingPartyNumber string `json:"callingPartyNumber"`


    // DefaultScripts - The default script Ids for the communication types.
    DefaultScripts map[string]Script `json:"defaultScripts"`


    // OutboundMessagingAddresses - The messaging addresses for the queue.
    OutboundMessagingAddresses Queuemessagingaddresses `json:"outboundMessagingAddresses"`


    // OutboundEmailAddress - The default email address to use for outbound email from this queue.
    OutboundEmailAddress Queueemailaddress `json:"outboundEmailAddress"`


    // PeerId - The ID of an associated external queue.
    PeerId string `json:"peerId"`


    // SuppressInQueueCallRecording - Indicates whether recording in-queue calls is suppressed for this queue.
    SuppressInQueueCallRecording bool `json:"suppressInQueueCallRecording"`


    

}

// String returns a JSON representation of the model
func (o *Queuerequest) String() string {
    
    
    
    
    
    
    
    
     o.RoutingRules = []Routingrule{{}} 
    
    
    
    
    
    
    
     o.MemberGroups = []Membergroup{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.DefaultScripts = map[string]Script{"": {}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queuerequest) MarshalJSON() ([]byte, error) {
    type Alias Queuerequest

    if QueuerequestMarshalled {
        return []byte("{}"), nil
    }
    QueuerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        MediaSettings Queuemediasettings `json:"mediaSettings"`
        
        RoutingRules []Routingrule `json:"routingRules"`
        
        ConditionalGroupRouting Conditionalgrouprouting `json:"conditionalGroupRouting"`
        
        ConditionalGroupActivation Conditionalgroupactivation `json:"conditionalGroupActivation"`
        
        Bullseye Bullseye `json:"bullseye"`
        
        ScoringMethod string `json:"scoringMethod"`
        
        LastAgentRoutingMode string `json:"lastAgentRoutingMode"`
        
        AcwSettings Acwsettings `json:"acwSettings"`
        
        SkillEvaluationMethod string `json:"skillEvaluationMethod"`
        
        MemberGroups []Membergroup `json:"memberGroups"`
        
        QueueFlow Domainentityref `json:"queueFlow"`
        
        EmailInQueueFlow Domainentityref `json:"emailInQueueFlow"`
        
        MessageInQueueFlow Domainentityref `json:"messageInQueueFlow"`
        
        WhisperPrompt Domainentityref `json:"whisperPrompt"`
        
        OnHoldPrompt Domainentityref `json:"onHoldPrompt"`
        
        AutoAnswerOnly bool `json:"autoAnswerOnly"`
        
        CannedResponseLibraries Cannedresponselibraries `json:"cannedResponseLibraries"`
        
        EnableTranscription bool `json:"enableTranscription"`
        
        EnableAudioMonitoring bool `json:"enableAudioMonitoring"`
        
        EnableManualAssignment bool `json:"enableManualAssignment"`
        
        AgentOwnedRouting Agentownedrouting `json:"agentOwnedRouting"`
        
        DirectRouting Directrouting `json:"directRouting"`
        
        CallingPartyName string `json:"callingPartyName"`
        
        CallingPartyNumber string `json:"callingPartyNumber"`
        
        DefaultScripts map[string]Script `json:"defaultScripts"`
        
        OutboundMessagingAddresses Queuemessagingaddresses `json:"outboundMessagingAddresses"`
        
        OutboundEmailAddress Queueemailaddress `json:"outboundEmailAddress"`
        
        PeerId string `json:"peerId"`
        
        SuppressInQueueCallRecording bool `json:"suppressInQueueCallRecording"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        RoutingRules: []Routingrule{{}},
        


        


        


        


        


        


        


        


        
        MemberGroups: []Membergroup{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        DefaultScripts: map[string]Script{"": {}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

