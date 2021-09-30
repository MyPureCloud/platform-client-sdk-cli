package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoidDud struct { }

// Void
type Void struct { }

// String returns a JSON representation of the model
func (o *Void) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Void) MarshalJSON() ([]byte, error) {
    type Alias Void

    if VoidMarshalled {
        return []byte("{}"), nil
    }
    VoidMarshalled = true

    return json.Marshal(&struct { 
        *Alias
    }{
        
        Alias: (*Alias)(u),
    })
}

