package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CardbodytextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CardbodytextDud struct { }

// Cardbodytext
type Cardbodytext struct { }

// String returns a JSON representation of the model
func (o *Cardbodytext) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cardbodytext) MarshalJSON() ([]byte, error) {
    type Alias Cardbodytext

    if CardbodytextMarshalled {
        return []byte("{}"), nil
    }
    CardbodytextMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

