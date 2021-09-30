package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JsonnodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JsonnodeDud struct { }

// Jsonnode
type Jsonnode struct { }

// String returns a JSON representation of the model
func (o *Jsonnode) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Jsonnode) MarshalJSON() ([]byte, error) {
    type Alias Jsonnode

    if JsonnodeMarshalled {
        return []byte("{}"), nil
    }
    JsonnodeMarshalled = true

    return json.Marshal(&struct { 
        *Alias
    }{
        
        Alias: (*Alias)(u),
    })
}

