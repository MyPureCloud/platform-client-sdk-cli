package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarygenerationconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarygenerationconfigDud struct { 
    

}

// Summarygenerationconfig
type Summarygenerationconfig struct { 
    // Enabled - Copilot generated summary is enabled.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Summarygenerationconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarygenerationconfig) MarshalJSON() ([]byte, error) {
    type Alias Summarygenerationconfig

    if SummarygenerationconfigMarshalled {
        return []byte("{}"), nil
    }
    SummarygenerationconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

