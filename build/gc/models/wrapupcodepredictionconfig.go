package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WrapupcodepredictionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WrapupcodepredictionconfigDud struct { 
    

}

// Wrapupcodepredictionconfig
type Wrapupcodepredictionconfig struct { 
    // Enabled - Copilot generated wrapup code prediction is enabled.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Wrapupcodepredictionconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wrapupcodepredictionconfig) MarshalJSON() ([]byte, error) {
    type Alias Wrapupcodepredictionconfig

    if WrapupcodepredictionconfigMarshalled {
        return []byte("{}"), nil
    }
    WrapupcodepredictionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

