package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InboundonlysettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InboundonlysettingDud struct { 
    

}

// Inboundonlysetting
type Inboundonlysetting struct { 
    // Inbound
    Inbound string `json:"inbound"`

}

// String returns a JSON representation of the model
func (o *Inboundonlysetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inboundonlysetting) MarshalJSON() ([]byte, error) {
    type Alias Inboundonlysetting

    if InboundonlysettingMarshalled {
        return []byte("{}"), nil
    }
    InboundonlysettingMarshalled = true

    return json.Marshal(&struct {
        
        Inbound string `json:"inbound"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

