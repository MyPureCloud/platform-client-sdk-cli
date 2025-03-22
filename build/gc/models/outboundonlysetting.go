package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutboundonlysettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutboundonlysettingDud struct { 
    

}

// Outboundonlysetting
type Outboundonlysetting struct { 
    // Outbound - Status for the Outbound Direction
    Outbound string `json:"outbound"`

}

// String returns a JSON representation of the model
func (o *Outboundonlysetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outboundonlysetting) MarshalJSON() ([]byte, error) {
    type Alias Outboundonlysetting

    if OutboundonlysettingMarshalled {
        return []byte("{}"), nil
    }
    OutboundonlysettingMarshalled = true

    return json.Marshal(&struct {
        
        Outbound string `json:"outbound"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

