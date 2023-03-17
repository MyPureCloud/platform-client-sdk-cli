package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsettingsassigneeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsettingsassigneeDud struct { }

// Evaluationsettingsassignee
type Evaluationsettingsassignee struct { }

// String returns a JSON representation of the model
func (o *Evaluationsettingsassignee) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsettingsassignee) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsettingsassignee

    if EvaluationsettingsassigneeMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsettingsassigneeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

