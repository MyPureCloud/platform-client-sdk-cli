package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitorsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitorsettingsDud struct { 
    


    

}

// Screenmonitorsettings
type Screenmonitorsettings struct { 
    // EnableAgentNotifications
    EnableAgentNotifications bool `json:"enableAgentNotifications"`


    // MaxSimultaneousScreenMonitoringSessions
    MaxSimultaneousScreenMonitoringSessions int `json:"maxSimultaneousScreenMonitoringSessions"`

}

// String returns a JSON representation of the model
func (o *Screenmonitorsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitorsettings) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitorsettings

    if ScreenmonitorsettingsMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitorsettingsMarshalled = true

    return json.Marshal(&struct {
        
        EnableAgentNotifications bool `json:"enableAgentNotifications"`
        
        MaxSimultaneousScreenMonitoringSessions int `json:"maxSimultaneousScreenMonitoringSessions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

