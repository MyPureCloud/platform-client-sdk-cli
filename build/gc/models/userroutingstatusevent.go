package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserroutingstatuseventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserroutingstatuseventDud struct { 
    


    


    


    


    

}

// Userroutingstatusevent
type Userroutingstatusevent struct { 
    // EventId - A unique (UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A timestamp as epoch representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // AgentId - Unique identifier of the agent.
    AgentId string `json:"agentId"`


    // Status - The agent's current routing status.
    Status string `json:"status"`


    // SourceId - The agent's source platform Id.
    SourceId string `json:"sourceId"`

}

// String returns a JSON representation of the model
func (o *Userroutingstatusevent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userroutingstatusevent) MarshalJSON() ([]byte, error) {
    type Alias Userroutingstatusevent

    if UserroutingstatuseventMarshalled {
        return []byte("{}"), nil
    }
    UserroutingstatuseventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        AgentId string `json:"agentId"`
        
        Status string `json:"status"`
        
        SourceId string `json:"sourceId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

