package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappcallingconfigurationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappcallingconfigurationresponseDud struct { }

// Whatsappcallingconfigurationresponse
type Whatsappcallingconfigurationresponse struct { }

// String returns a JSON representation of the model
func (o *Whatsappcallingconfigurationresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappcallingconfigurationresponse) MarshalJSON() ([]byte, error) {
    type Alias Whatsappcallingconfigurationresponse

    if WhatsappcallingconfigurationresponseMarshalled {
        return []byte("{}"), nil
    }
    WhatsappcallingconfigurationresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

