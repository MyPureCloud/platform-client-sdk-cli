package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitoringsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitoringsessionDud struct { 
    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Screenmonitoringsession
type Screenmonitoringsession struct { 
    // SourceUser - The user who initiated the screen monitoring session
    SourceUser Addressableentityref `json:"sourceUser"`


    // TargetUser - The user being monitored (for agent-level monitoring)
    TargetUser Addressableentityref `json:"targetUser"`


    // Conversation - The conversation being monitored (for conversation-level monitoring)
    Conversation Addressableentityref `json:"conversation"`


    // ParticipantId - The ID of the participant being monitored (for conversation-level monitoring)
    ParticipantId string `json:"participantId"`


    // MonitoringType - The type of screen monitoring session
    MonitoringType string `json:"monitoringType"`


    // DateCreated - The date and time when the screen monitoring session was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // ScreenMonitoringId - The unique identifier for this screen monitoring session
    ScreenMonitoringId string `json:"screenMonitoringId"`


    

}

// String returns a JSON representation of the model
func (o *Screenmonitoringsession) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitoringsession) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitoringsession

    if ScreenmonitoringsessionMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitoringsessionMarshalled = true

    return json.Marshal(&struct {
        
        SourceUser Addressableentityref `json:"sourceUser"`
        
        TargetUser Addressableentityref `json:"targetUser"`
        
        Conversation Addressableentityref `json:"conversation"`
        
        ParticipantId string `json:"participantId"`
        
        MonitoringType string `json:"monitoringType"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        ScreenMonitoringId string `json:"screenMonitoringId"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

