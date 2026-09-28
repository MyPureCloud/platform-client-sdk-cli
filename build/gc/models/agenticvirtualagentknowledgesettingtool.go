package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentknowledgesettingtoolMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentknowledgesettingtoolDud struct { }

// Agenticvirtualagentknowledgesettingtool
type Agenticvirtualagentknowledgesettingtool struct { }

// String returns a JSON representation of the model
func (o *Agenticvirtualagentknowledgesettingtool) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentknowledgesettingtool) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentknowledgesettingtool

    if AgenticvirtualagentknowledgesettingtoolMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentknowledgesettingtoolMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

