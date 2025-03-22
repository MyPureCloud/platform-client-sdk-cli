package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DiagnosticssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DiagnosticssettingsDud struct { }

// Diagnosticssettings
type Diagnosticssettings struct { }

// String returns a JSON representation of the model
func (o *Diagnosticssettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Diagnosticssettings) MarshalJSON() ([]byte, error) {
    type Alias Diagnosticssettings

    if DiagnosticssettingsMarshalled {
        return []byte("{}"), nil
    }
    DiagnosticssettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

