package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictivescoringsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictivescoringsettingDud struct { 
    

}

// Predictivescoringsetting
type Predictivescoringsetting struct { 
    // Enabled - True if Predictive Scoring feature is configured.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Predictivescoringsetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictivescoringsetting) MarshalJSON() ([]byte, error) {
    type Alias Predictivescoringsetting

    if PredictivescoringsettingMarshalled {
        return []byte("{}"), nil
    }
    PredictivescoringsettingMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

