package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleinvitationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleinvitationDud struct { }

// Appleinvitation - Apple Messages for Business invitation template configuration
type Appleinvitation struct { }

// String returns a JSON representation of the model
func (o *Appleinvitation) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleinvitation) MarshalJSON() ([]byte, error) {
    type Alias Appleinvitation

    if AppleinvitationMarshalled {
        return []byte("{}"), nil
    }
    AppleinvitationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

