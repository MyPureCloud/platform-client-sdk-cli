package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserpresenceeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserpresenceeventDud struct { 
    


    


    


    


    


    

}

// Userpresenceevent
type Userpresenceevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // UserId - The User ID of the user associated with this UserPresence
    UserId string `json:"userId"`


    // SourceId - The id (V4 UUID) of the presence source being updated
    SourceId string `json:"sourceId"`


    // PresenceDefinitionId - The id (UUID) of the presence definition that the user presence is associated with
    PresenceDefinitionId string `json:"presenceDefinitionId"`


    // Message - The message associated with the presence
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Userpresenceevent) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userpresenceevent) MarshalJSON() ([]byte, error) {
    type Alias Userpresenceevent

    if UserpresenceeventMarshalled {
        return []byte("{}"), nil
    }
    UserpresenceeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        UserId string `json:"userId"`
        
        SourceId string `json:"sourceId"`
        
        PresenceDefinitionId string `json:"presenceDefinitionId"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

