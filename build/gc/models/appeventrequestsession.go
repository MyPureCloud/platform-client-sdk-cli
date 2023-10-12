package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppeventrequestsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppeventrequestsessionDud struct { 
    


    


    

}

// Appeventrequestsession
type Appeventrequestsession struct { 
    // Id - ID of the app session.
    Id string `json:"id"`


    // EventCount - The count of all events recorded during this session. This should be incremented for each event in the session, regardless of event name.
    EventCount int `json:"eventCount"`


    // ScreenviewCount - The count of all screen view events recorded during this session. This should be incremented for each screen_viewed event in the session.
    ScreenviewCount int `json:"screenviewCount"`

}

// String returns a JSON representation of the model
func (o *Appeventrequestsession) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appeventrequestsession) MarshalJSON() ([]byte, error) {
    type Alias Appeventrequestsession

    if AppeventrequestsessionMarshalled {
        return []byte("{}"), nil
    }
    AppeventrequestsessionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        EventCount int `json:"eventCount"`
        
        ScreenviewCount int `json:"screenviewCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

