package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsettingsDud struct { }

// Evaluationsettings
type Evaluationsettings struct { }

// String returns a JSON representation of the model
func (o *Evaluationsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsettings) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsettings

    if EvaluationsettingsMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

