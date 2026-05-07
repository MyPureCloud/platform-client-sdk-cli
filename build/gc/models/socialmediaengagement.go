package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaengagementMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaengagementDud struct { }

// Socialmediaengagement
type Socialmediaengagement struct { }

// String returns a JSON representation of the model
func (o *Socialmediaengagement) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaengagement) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaengagement

    if SocialmediaengagementMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaengagementMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

