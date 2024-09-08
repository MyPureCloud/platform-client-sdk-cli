package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BunotificationsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BunotificationsettingsresponseDud struct { 
    

}

// Bunotificationsettingsresponse
type Bunotificationsettingsresponse struct { 
    // Scheduling - Schedule notification settings
    Scheduling Buschedulenotificationssettingsresponse `json:"scheduling"`

}

// String returns a JSON representation of the model
func (o *Bunotificationsettingsresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bunotificationsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bunotificationsettingsresponse

    if BunotificationsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BunotificationsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Scheduling Buschedulenotificationssettingsresponse `json:"scheduling"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

