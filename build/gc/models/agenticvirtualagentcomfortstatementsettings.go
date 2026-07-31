package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentcomfortstatementsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentcomfortstatementsettingsDud struct { 
    

}

// Agenticvirtualagentcomfortstatementsettings - Comfort statement settings for tool calls.
type Agenticvirtualagentcomfortstatementsettings struct { 
    // Enabled - Whether comfort statements are enabled during eligible tool calls.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentcomfortstatementsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentcomfortstatementsettings) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentcomfortstatementsettings

    if AgenticvirtualagentcomfortstatementsettingsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentcomfortstatementsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

