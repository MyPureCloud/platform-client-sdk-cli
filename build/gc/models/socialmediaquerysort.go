package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaquerysortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaquerysortDud struct { }

// Socialmediaquerysort
type Socialmediaquerysort struct { }

// String returns a JSON representation of the model
func (o *Socialmediaquerysort) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaquerysort) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaquerysort

    if SocialmediaquerysortMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaquerysortMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

