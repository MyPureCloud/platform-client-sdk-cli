package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramhashtagsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramhashtagsDud struct { }

// Instagramhashtags
type Instagramhashtags struct { }

// String returns a JSON representation of the model
func (o *Instagramhashtags) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramhashtags) MarshalJSON() ([]byte, error) {
    type Alias Instagramhashtags

    if InstagramhashtagsMarshalled {
        return []byte("{}"), nil
    }
    InstagramhashtagsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

