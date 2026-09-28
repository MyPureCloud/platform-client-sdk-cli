package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentknowledgebasetoolMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentknowledgebasetoolDud struct { }

// Agenticvirtualagentknowledgebasetool
type Agenticvirtualagentknowledgebasetool struct { }

// String returns a JSON representation of the model
func (o *Agenticvirtualagentknowledgebasetool) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentknowledgebasetool) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentknowledgebasetool

    if AgenticvirtualagentknowledgebasetoolMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentknowledgebasetoolMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

