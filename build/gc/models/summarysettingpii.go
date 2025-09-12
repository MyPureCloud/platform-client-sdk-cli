package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingpiiMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingpiiDud struct { 
    

}

// Summarysettingpii
type Summarysettingpii struct { 
    // All - Toggle PII visibility in summary.
    All bool `json:"all"`

}

// String returns a JSON representation of the model
func (o *Summarysettingpii) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysettingpii) MarshalJSON() ([]byte, error) {
    type Alias Summarysettingpii

    if SummarysettingpiiMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingpiiMarshalled = true

    return json.Marshal(&struct {
        
        All bool `json:"all"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

