package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotsentimentsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotsentimentsettingsDud struct { }

// Copilotsentimentsettings
type Copilotsentimentsettings struct { }

// String returns a JSON representation of the model
func (o *Copilotsentimentsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotsentimentsettings) MarshalJSON() ([]byte, error) {
    type Alias Copilotsentimentsettings

    if CopilotsentimentsettingsMarshalled {
        return []byte("{}"), nil
    }
    CopilotsentimentsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

