package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButtonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButtonDud struct { }

// Button
type Button struct { }

// String returns a JSON representation of the model
func (o *Button) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Button) MarshalJSON() ([]byte, error) {
    type Alias Button

    if ButtonMarshalled {
        return []byte("{}"), nil
    }
    ButtonMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

