package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutoanswersettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutoanswersettingDud struct { 
    

}

// Autoanswersetting
type Autoanswersetting struct { 
    // Enabled - The media type status.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Autoanswersetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Autoanswersetting) MarshalJSON() ([]byte, error) {
    type Alias Autoanswersetting

    if AutoanswersettingMarshalled {
        return []byte("{}"), nil
    }
    AutoanswersettingMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

