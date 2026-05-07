package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListvaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListvaluesDud struct { }

// Listvalues
type Listvalues struct { }

// String returns a JSON representation of the model
func (o *Listvalues) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listvalues) MarshalJSON() ([]byte, error) {
    type Alias Listvalues

    if ListvaluesMarshalled {
        return []byte("{}"), nil
    }
    ListvaluesMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

