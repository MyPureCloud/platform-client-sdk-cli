package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentescalationeventsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentescalationeventsettingsDud struct { }

// Agenticvirtualagentescalationeventsettings
type Agenticvirtualagentescalationeventsettings struct { }

// String returns a JSON representation of the model
func (o *Agenticvirtualagentescalationeventsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentescalationeventsettings) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentescalationeventsettings

    if AgenticvirtualagentescalationeventsettingsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentescalationeventsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

