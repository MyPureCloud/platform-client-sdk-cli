package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConfigurationoverridesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConfigurationoverridesDud struct { 
    

}

// Configurationoverrides
type Configurationoverrides struct { 
    // Priority - Indicates whether or not the contact will be placed in front of the queue or at the end of the queue.
    Priority bool `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Configurationoverrides) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Configurationoverrides) MarshalJSON() ([]byte, error) {
    type Alias Configurationoverrides

    if ConfigurationoverridesMarshalled {
        return []byte("{}"), nil
    }
    ConfigurationoverridesMarshalled = true

    return json.Marshal(&struct {
        
        Priority bool `json:"priority"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

