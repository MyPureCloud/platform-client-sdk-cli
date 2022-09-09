package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MarkcontactaddressuncontactableactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MarkcontactaddressuncontactableactionsettingsDud struct { }

// Markcontactaddressuncontactableactionsettings
type Markcontactaddressuncontactableactionsettings struct { }

// String returns a JSON representation of the model
func (o *Markcontactaddressuncontactableactionsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Markcontactaddressuncontactableactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Markcontactaddressuncontactableactionsettings

    if MarkcontactaddressuncontactableactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    MarkcontactaddressuncontactableactionsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

