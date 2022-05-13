package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmptyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmptyDud struct { }

// Empty
type Empty struct { }

// String returns a JSON representation of the model
func (o *Empty) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Empty) MarshalJSON() ([]byte, error) {
    type Alias Empty

    if EmptyMarshalled {
        return []byte("{}"), nil
    }
    EmptyMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

