package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeanddatesubconditionrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeanddatesubconditionrangeDud struct { }

// Timeanddatesubconditionrange
type Timeanddatesubconditionrange struct { }

// String returns a JSON representation of the model
func (o *Timeanddatesubconditionrange) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeanddatesubconditionrange) MarshalJSON() ([]byte, error) {
    type Alias Timeanddatesubconditionrange

    if TimeanddatesubconditionrangeMarshalled {
        return []byte("{}"), nil
    }
    TimeanddatesubconditionrangeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

