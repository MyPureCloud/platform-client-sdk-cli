package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DonotsendactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DonotsendactionsettingsDud struct { }

// Donotsendactionsettings
type Donotsendactionsettings struct { }

// String returns a JSON representation of the model
func (o *Donotsendactionsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Donotsendactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Donotsendactionsettings

    if DonotsendactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    DonotsendactionsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

