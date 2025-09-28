package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleauthenticationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleauthenticationDud struct { }

// Appleauthentication
type Appleauthentication struct { }

// String returns a JSON representation of the model
func (o *Appleauthentication) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleauthentication) MarshalJSON() ([]byte, error) {
    type Alias Appleauthentication

    if AppleauthenticationMarshalled {
        return []byte("{}"), nil
    }
    AppleauthenticationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

