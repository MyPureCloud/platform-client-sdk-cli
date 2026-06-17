package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappbusinessscopedidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappbusinessscopedidDud struct { }

// Whatsappbusinessscopedid - A WhatsAppBusinessScopedId record
type Whatsappbusinessscopedid struct { }

// String returns a JSON representation of the model
func (o *Whatsappbusinessscopedid) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappbusinessscopedid) MarshalJSON() ([]byte, error) {
    type Alias Whatsappbusinessscopedid

    if WhatsappbusinessscopedidMarshalled {
        return []byte("{}"), nil
    }
    WhatsappbusinessscopedidMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

