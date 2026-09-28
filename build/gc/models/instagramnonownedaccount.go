package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramnonownedaccountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramnonownedaccountDud struct { }

// Instagramnonownedaccount
type Instagramnonownedaccount struct { }

// String returns a JSON representation of the model
func (o *Instagramnonownedaccount) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramnonownedaccount) MarshalJSON() ([]byte, error) {
    type Alias Instagramnonownedaccount

    if InstagramnonownedaccountMarshalled {
        return []byte("{}"), nil
    }
    InstagramnonownedaccountMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

