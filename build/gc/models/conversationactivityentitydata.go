package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityentitydataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityentitydataDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Conversationactivityentitydata
type Conversationactivityentitydata struct { 
    // ActivityDate - The time at which the activity was observed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ActivityDate time.Time `json:"activityDate"`


    // Metric - Activity metric
    Metric string `json:"metric"`


    // ActiveRouting - Active routing method
    ActiveRouting string `json:"activeRouting"`


    // AddressFrom - The address that initiated an action
    AddressFrom string `json:"addressFrom"`


    // AddressTo - The address receiving an action
    AddressTo string `json:"addressTo"`


    // Ani - Automatic Number Identification (caller's number)
    Ani string `json:"ani"`


    // ConversationId - Unique identifier for the conversation
    ConversationId string `json:"conversationId"`


    // ConvertedFrom - Session media type that was converted from in case of a media type conversion
    ConvertedFrom string `json:"convertedFrom"`


    // ConvertedTo - Session media type that was converted to in case of a media type conversion
    ConvertedTo string `json:"convertedTo"`


    // Direction - The direction of the communication
    Direction string `json:"direction"`


    // Dnis - Dialed number identification service (number dialed by the calling party)
    Dnis string `json:"dnis"`


    // MediaType - The session media type
    MediaType string `json:"mediaType"`


    // ParticipantName - A human readable name identifying the participant
    ParticipantName string `json:"participantName"`


    // QueueId - Queue identifier
    QueueId string `json:"queueId"`


    // RequestedLanguageId - Unique identifier for the language requested for an interaction
    RequestedLanguageId string `json:"requestedLanguageId"`


    // RequestedRoutingSkillIds - Unique identifier(s) for skill(s) requested for an interaction
    RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`


    // RequestedRoutings - Routing type(s) for requested/attempted routing methods.
    RequestedRoutings []string `json:"requestedRoutings"`


    // RoutingPriority - Routing priority for the current interaction
    RoutingPriority int `json:"routingPriority"`


    // SessionId - The unique identifier of this session
    SessionId string `json:"sessionId"`


    // TeamId - The team ID the user is a member of
    TeamId string `json:"teamId"`


    // UsedRouting - Complete routing method
    UsedRouting string `json:"usedRouting"`


    // UserId - Unique identifier for the user
    UserId string `json:"userId"`


    // ScoredAgents - Scored agents
    ScoredAgents []Conversationactivityscoredagent `json:"scoredAgents"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityentitydata) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RequestedRoutingSkillIds = []string{""} 
     o.RequestedRoutings = []string{""} 
    
    
    
    
    
     o.ScoredAgents = []Conversationactivityscoredagent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityentitydata) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityentitydata

    if ConversationactivityentitydataMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityentitydataMarshalled = true

    return json.Marshal(&struct {
        
        ActivityDate time.Time `json:"activityDate"`
        
        Metric string `json:"metric"`
        
        ActiveRouting string `json:"activeRouting"`
        
        AddressFrom string `json:"addressFrom"`
        
        AddressTo string `json:"addressTo"`
        
        Ani string `json:"ani"`
        
        ConversationId string `json:"conversationId"`
        
        ConvertedFrom string `json:"convertedFrom"`
        
        ConvertedTo string `json:"convertedTo"`
        
        Direction string `json:"direction"`
        
        Dnis string `json:"dnis"`
        
        MediaType string `json:"mediaType"`
        
        ParticipantName string `json:"participantName"`
        
        QueueId string `json:"queueId"`
        
        RequestedLanguageId string `json:"requestedLanguageId"`
        
        RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`
        
        RequestedRoutings []string `json:"requestedRoutings"`
        
        RoutingPriority int `json:"routingPriority"`
        
        SessionId string `json:"sessionId"`
        
        TeamId string `json:"teamId"`
        
        UsedRouting string `json:"usedRouting"`
        
        UserId string `json:"userId"`
        
        ScoredAgents []Conversationactivityscoredagent `json:"scoredAgents"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        RequestedRoutingSkillIds: []string{""},
        


        
        RequestedRoutings: []string{""},
        


        


        


        


        


        


        
        ScoredAgents: []Conversationactivityscoredagent{{}},
        

        Alias: (*Alias)(u),
    })
}

