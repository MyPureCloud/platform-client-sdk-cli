package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ObservationvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ObservationvalueDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Observationvalue
type Observationvalue struct { 
    // ObservationDate - The time at which the observation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ObservationDate time.Time `json:"observationDate"`


    // ConversationId - Unique identifier for the conversation
    ConversationId string `json:"conversationId"`


    // SessionId - The unique identifier of this session
    SessionId string `json:"sessionId"`


    // RequestedRoutingSkillIds - Unique identifier for a skill requested for an interaction
    RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`


    // RequestedLanguageId - Unique identifier for the language requested for an interaction
    RequestedLanguageId string `json:"requestedLanguageId"`


    // RoutingPriority - Routing priority for the current interaction
    RoutingPriority int `json:"routingPriority"`


    // ParticipantName - A human readable name identifying the participant
    ParticipantName string `json:"participantName"`


    // UserId - Unique identifier for the user
    UserId string `json:"userId"`


    // Direction - The direction of the communication
    Direction string `json:"direction"`


    // ConvertedFrom - Session media type that was converted from in case of a media type conversion
    ConvertedFrom string `json:"convertedFrom"`


    // ConvertedTo - Session media type that was converted to in case of a media type conversion
    ConvertedTo string `json:"convertedTo"`


    // AddressFrom - The address that initiated an action
    AddressFrom string `json:"addressFrom"`


    // AddressTo - The address receiving an action
    AddressTo string `json:"addressTo"`


    // Ani - Automatic Number Identification (caller's number)
    Ani string `json:"ani"`


    // Dnis - Dialed number identification service (number dialed by the calling party)
    Dnis string `json:"dnis"`


    // TeamId - The team id the user is a member of
    TeamId string `json:"teamId"`


    // RequestedRoutings - All routing types for requested/attempted routing methods
    RequestedRoutings []string `json:"requestedRoutings"`


    // UsedRouting - Complete routing method
    UsedRouting string `json:"usedRouting"`


    // ScoredAgents
    ScoredAgents []Analyticsscoredagent `json:"scoredAgents"`

}

// String returns a JSON representation of the model
func (o *Observationvalue) String() string {
    
    
    
     o.RequestedRoutingSkillIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
     o.RequestedRoutings = []string{""} 
    
     o.ScoredAgents = []Analyticsscoredagent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Observationvalue) MarshalJSON() ([]byte, error) {
    type Alias Observationvalue

    if ObservationvalueMarshalled {
        return []byte("{}"), nil
    }
    ObservationvalueMarshalled = true

    return json.Marshal(&struct {
        
        ObservationDate time.Time `json:"observationDate"`
        
        ConversationId string `json:"conversationId"`
        
        SessionId string `json:"sessionId"`
        
        RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`
        
        RequestedLanguageId string `json:"requestedLanguageId"`
        
        RoutingPriority int `json:"routingPriority"`
        
        ParticipantName string `json:"participantName"`
        
        UserId string `json:"userId"`
        
        Direction string `json:"direction"`
        
        ConvertedFrom string `json:"convertedFrom"`
        
        ConvertedTo string `json:"convertedTo"`
        
        AddressFrom string `json:"addressFrom"`
        
        AddressTo string `json:"addressTo"`
        
        Ani string `json:"ani"`
        
        Dnis string `json:"dnis"`
        
        TeamId string `json:"teamId"`
        
        RequestedRoutings []string `json:"requestedRoutings"`
        
        UsedRouting string `json:"usedRouting"`
        
        ScoredAgents []Analyticsscoredagent `json:"scoredAgents"`
        *Alias
    }{

        


        


        


        
        RequestedRoutingSkillIds: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        
        RequestedRoutings: []string{""},
        


        


        
        ScoredAgents: []Analyticsscoredagent{{}},
        

        Alias: (*Alias)(u),
    })
}

