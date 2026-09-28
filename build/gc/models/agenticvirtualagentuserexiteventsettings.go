package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentuserexiteventsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentuserexiteventsettingsDud struct { }

// Agenticvirtualagentuserexiteventsettings
type Agenticvirtualagentuserexiteventsettings struct { }

// String returns a JSON representation of the model
func (o *Agenticvirtualagentuserexiteventsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentuserexiteventsettings) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentuserexiteventsettings

    if AgenticvirtualagentuserexiteventsettingsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentuserexiteventsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

