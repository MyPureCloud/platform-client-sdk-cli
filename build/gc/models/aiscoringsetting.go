package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AiscoringsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AiscoringsettingDud struct { 
    

}

// Aiscoringsetting
type Aiscoringsetting struct { 
    // Enabled - True if AI Scoring feature is configured.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Aiscoringsetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aiscoringsetting) MarshalJSON() ([]byte, error) {
    type Alias Aiscoringsetting

    if AiscoringsettingMarshalled {
        return []byte("{}"), nil
    }
    AiscoringsettingMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

