package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrackingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrackingsettingsDud struct { }

// Trackingsettings - Settings that control how tracking data is collected and filtered.
type Trackingsettings struct { }

// String returns a JSON representation of the model
func (o *Trackingsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trackingsettings) MarshalJSON() ([]byte, error) {
    type Alias Trackingsettings

    if TrackingsettingsMarshalled {
        return []byte("{}"), nil
    }
    TrackingsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

