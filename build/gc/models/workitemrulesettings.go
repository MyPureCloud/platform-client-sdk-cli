package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemrulesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemrulesettingsDud struct { }

// Workitemrulesettings
type Workitemrulesettings struct { }

// String returns a JSON representation of the model
func (o *Workitemrulesettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemrulesettings) MarshalJSON() ([]byte, error) {
    type Alias Workitemrulesettings

    if WorkitemrulesettingsMarshalled {
        return []byte("{}"), nil
    }
    WorkitemrulesettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

