package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertingunreadstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertingunreadstatusDud struct { 
    

}

// Alertingunreadstatus
type Alertingunreadstatus struct { 
    // Unread - True is alert is unread, false if it has not been.
    Unread bool `json:"unread"`

}

// String returns a JSON representation of the model
func (o *Alertingunreadstatus) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertingunreadstatus) MarshalJSON() ([]byte, error) {
    type Alias Alertingunreadstatus

    if AlertingunreadstatusMarshalled {
        return []byte("{}"), nil
    }
    AlertingunreadstatusMarshalled = true

    return json.Marshal(&struct {
        
        Unread bool `json:"unread"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

