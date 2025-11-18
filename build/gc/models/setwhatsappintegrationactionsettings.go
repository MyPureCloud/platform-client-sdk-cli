package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetwhatsappintegrationactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetwhatsappintegrationactionsettingsDud struct { }

// Setwhatsappintegrationactionsettings
type Setwhatsappintegrationactionsettings struct { }

// String returns a JSON representation of the model
func (o *Setwhatsappintegrationactionsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setwhatsappintegrationactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Setwhatsappintegrationactionsettings

    if SetwhatsappintegrationactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    SetwhatsappintegrationactionsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

