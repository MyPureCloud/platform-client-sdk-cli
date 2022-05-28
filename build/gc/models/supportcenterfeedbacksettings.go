package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterfeedbacksettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterfeedbacksettingsDud struct { 
    

}

// Supportcenterfeedbacksettings
type Supportcenterfeedbacksettings struct { 
    // Enabled - Whether or not requesting customer feedback on article content and article search results is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Supportcenterfeedbacksettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterfeedbacksettings) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterfeedbacksettings

    if SupportcenterfeedbacksettingsMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterfeedbacksettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

