package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArraynodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArraynodeDud struct { }

// Arraynode
type Arraynode struct { }

// String returns a JSON representation of the model
func (o *Arraynode) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Arraynode) MarshalJSON() ([]byte, error) {
    type Alias Arraynode

    if ArraynodeMarshalled {
        return []byte("{}"), nil
    }
    ArraynodeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

