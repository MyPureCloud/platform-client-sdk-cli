package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappcallingconfigurationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappcallingconfigurationrequestDud struct { }

// Whatsappcallingconfigurationrequest
type Whatsappcallingconfigurationrequest struct { }

// String returns a JSON representation of the model
func (o *Whatsappcallingconfigurationrequest) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappcallingconfigurationrequest) MarshalJSON() ([]byte, error) {
    type Alias Whatsappcallingconfigurationrequest

    if WhatsappcallingconfigurationrequestMarshalled {
        return []byte("{}"), nil
    }
    WhatsappcallingconfigurationrequestMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

