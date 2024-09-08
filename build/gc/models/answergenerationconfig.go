package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnswergenerationconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnswergenerationconfigDud struct { 
    

}

// Answergenerationconfig
type Answergenerationconfig struct { 
    // Enabled - Answer generation is enabled.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Answergenerationconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Answergenerationconfig) MarshalJSON() ([]byte, error) {
    type Alias Answergenerationconfig

    if AnswergenerationconfigMarshalled {
        return []byte("{}"), nil
    }
    AnswergenerationconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

