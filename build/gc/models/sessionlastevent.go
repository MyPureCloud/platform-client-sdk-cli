package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionlasteventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionlasteventDud struct { 
    


    


    

}

// Sessionlastevent
type Sessionlastevent struct { 
    // Id - The ID of the last event.
    Id string `json:"id"`


    // EventName - The name of the event.
    EventName string `json:"eventName"`


    // CreatedDate - Timestamp indicating when the event was published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Sessionlastevent) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionlastevent) MarshalJSON() ([]byte, error) {
    type Alias Sessionlastevent

    if SessionlasteventMarshalled {
        return []byte("{}"), nil
    }
    SessionlasteventMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        EventName string `json:"eventName"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

