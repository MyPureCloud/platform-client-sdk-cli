package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictiveroutingcustomkpiattributioneventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictiveroutingcustomkpiattributioneventDud struct { 
    


    


    


    


    


    


    

}

// Predictiveroutingcustomkpiattributionevent
type Predictiveroutingcustomkpiattributionevent struct { 
    // EventId - A unique (UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A timestamp as epoch representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ExternalContactId - The UUID of the external contact associated with this event
    ExternalContactId string `json:"externalContactId"`


    // ConversationId - The UUID of the conversation associated with this event
    ConversationId string `json:"conversationId"`


    // AgentId - The UUID of the agent associated with this event
    AgentId string `json:"agentId"`


    // KpiId - The UUID of the KPI associated with this event
    KpiId string `json:"kpiId"`


    // AssociatedValue - The value associated with this outcome attribution
    AssociatedValue float64 `json:"associatedValue"`

}

// String returns a JSON representation of the model
func (o *Predictiveroutingcustomkpiattributionevent) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictiveroutingcustomkpiattributionevent) MarshalJSON() ([]byte, error) {
    type Alias Predictiveroutingcustomkpiattributionevent

    if PredictiveroutingcustomkpiattributioneventMarshalled {
        return []byte("{}"), nil
    }
    PredictiveroutingcustomkpiattributioneventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ExternalContactId string `json:"externalContactId"`
        
        ConversationId string `json:"conversationId"`
        
        AgentId string `json:"agentId"`
        
        KpiId string `json:"kpiId"`
        
        AssociatedValue float64 `json:"associatedValue"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

