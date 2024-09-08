package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BunotificationsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BunotificationsettingsrequestDud struct { 
    

}

// Bunotificationsettingsrequest
type Bunotificationsettingsrequest struct { 
    // Scheduling - Schedule notification settings
    Scheduling Buschedulenotificationssettingsrequest `json:"scheduling"`

}

// String returns a JSON representation of the model
func (o *Bunotificationsettingsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bunotificationsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bunotificationsettingsrequest

    if BunotificationsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    BunotificationsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Scheduling Buschedulenotificationssettingsrequest `json:"scheduling"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

