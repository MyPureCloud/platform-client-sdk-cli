package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationclearsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationclearsettingsDud struct { 
    

}

// Conversationclearsettings
type Conversationclearsettings struct { 
    // Enabled - whether or not conversation clear setting is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Conversationclearsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationclearsettings) MarshalJSON() ([]byte, error) {
    type Alias Conversationclearsettings

    if ConversationclearsettingsMarshalled {
        return []byte("{}"), nil
    }
    ConversationclearsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

