package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyconditionDud struct { }

// Policycondition
type Policycondition struct { }

// String returns a JSON representation of the model
func (o *Policycondition) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policycondition) MarshalJSON() ([]byte, error) {
    type Alias Policycondition

    if PolicyconditionMarshalled {
        return []byte("{}"), nil
    }
    PolicyconditionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

