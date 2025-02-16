package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsconversationsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsconversationsegmentDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsconversationsegment
type Analyticsconversationsegment struct { 
    // AudioMuted - Flag indicating if audio is muted or not (true/false)
    AudioMuted bool `json:"audioMuted"`


    // Conference - Indicates whether the segment was a conference
    Conference bool `json:"conference"`


    // DestinationConversationId - The unique identifier of a new conversation when a conversation is ended for a conference
    DestinationConversationId string `json:"destinationConversationId"`


    // DestinationSessionId - The unique identifier of a new session when a session is ended for a conference
    DestinationSessionId string `json:"destinationSessionId"`


    // DisconnectType - The session disconnect type
    DisconnectType string `json:"disconnectType"`


    // ErrorCode - A code corresponding to the error that occurred
    ErrorCode string `json:"errorCode"`


    // GroupId - Unique identifier for a Genesys Cloud group
    GroupId string `json:"groupId"`


    // Q850ResponseCodes - Q.850 response code(s)
    Q850ResponseCodes []int `json:"q850ResponseCodes"`


    // QueueId - Queue identifier
    QueueId string `json:"queueId"`


    // RequestedLanguageId - Unique identifier for the language requested for an interaction
    RequestedLanguageId string `json:"requestedLanguageId"`


    // RequestedRoutingSkillIds - Unique identifier(s) for skill(s) requested for an interaction
    RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`


    // RequestedRoutingUserIds - Unique identifier(s) for agent(s) requested for an interaction
    RequestedRoutingUserIds []string `json:"requestedRoutingUserIds"`


    // SegmentEnd - The end time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SegmentEnd time.Time `json:"segmentEnd"`


    // SegmentStart - The start time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SegmentStart time.Time `json:"segmentStart"`


    // SegmentType - The activity that takes place in the segment, such as hold or interact
    SegmentType string `json:"segmentType"`


    // SipResponseCodes - SIP response code(s)
    SipResponseCodes []int `json:"sipResponseCodes"`


    // SourceConversationId - The unique identifier of the previous conversation when a new conversation is created for a conference
    SourceConversationId string `json:"sourceConversationId"`


    // SourceSessionId - The unique identifier of the previous session when a new session is created for a conference
    SourceSessionId string `json:"sourceSessionId"`


    // Subject - The subject for the initial email that started this conversation
    Subject string `json:"subject"`


    // VideoMuted - Flag indicating if video is muted/paused or not (true/false)
    VideoMuted bool `json:"videoMuted"`


    // WrapUpCode - Wrap up code
    WrapUpCode string `json:"wrapUpCode"`


    // WrapUpNote - Note entered by an agent during after-call work
    WrapUpNote string `json:"wrapUpNote"`


    // WrapUpTags - Tag(s) assigned during after-call work
    WrapUpTags []string `json:"wrapUpTags"`


    // ScoredAgents - Scored agents
    ScoredAgents []Analyticsscoredagent `json:"scoredAgents"`


    // Properties - Additional segment properties
    Properties []Analyticsproperty `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationsegment) String() string {
    
    
    
    
    
    
    
     o.Q850ResponseCodes = []int{0} 
    
    
     o.RequestedRoutingSkillIds = []string{""} 
     o.RequestedRoutingUserIds = []string{""} 
    
    
    
     o.SipResponseCodes = []int{0} 
    
    
    
    
    
    
     o.WrapUpTags = []string{""} 
     o.ScoredAgents = []Analyticsscoredagent{{}} 
     o.Properties = []Analyticsproperty{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsconversationsegment) MarshalJSON() ([]byte, error) {
    type Alias Analyticsconversationsegment

    if AnalyticsconversationsegmentMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsconversationsegmentMarshalled = true

    return json.Marshal(&struct {
        
        AudioMuted bool `json:"audioMuted"`
        
        Conference bool `json:"conference"`
        
        DestinationConversationId string `json:"destinationConversationId"`
        
        DestinationSessionId string `json:"destinationSessionId"`
        
        DisconnectType string `json:"disconnectType"`
        
        ErrorCode string `json:"errorCode"`
        
        GroupId string `json:"groupId"`
        
        Q850ResponseCodes []int `json:"q850ResponseCodes"`
        
        QueueId string `json:"queueId"`
        
        RequestedLanguageId string `json:"requestedLanguageId"`
        
        RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`
        
        RequestedRoutingUserIds []string `json:"requestedRoutingUserIds"`
        
        SegmentEnd time.Time `json:"segmentEnd"`
        
        SegmentStart time.Time `json:"segmentStart"`
        
        SegmentType string `json:"segmentType"`
        
        SipResponseCodes []int `json:"sipResponseCodes"`
        
        SourceConversationId string `json:"sourceConversationId"`
        
        SourceSessionId string `json:"sourceSessionId"`
        
        Subject string `json:"subject"`
        
        VideoMuted bool `json:"videoMuted"`
        
        WrapUpCode string `json:"wrapUpCode"`
        
        WrapUpNote string `json:"wrapUpNote"`
        
        WrapUpTags []string `json:"wrapUpTags"`
        
        ScoredAgents []Analyticsscoredagent `json:"scoredAgents"`
        
        Properties []Analyticsproperty `json:"properties"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Q850ResponseCodes: []int{0},
        


        


        


        
        RequestedRoutingSkillIds: []string{""},
        


        
        RequestedRoutingUserIds: []string{""},
        


        


        


        


        
        SipResponseCodes: []int{0},
        


        


        


        


        


        


        


        
        WrapUpTags: []string{""},
        


        
        ScoredAgents: []Analyticsscoredagent{{}},
        


        
        Properties: []Analyticsproperty{{}},
        

        Alias: (*Alias)(u),
    })
}

