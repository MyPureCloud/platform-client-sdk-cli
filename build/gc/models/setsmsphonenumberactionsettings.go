package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetsmsphonenumberactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetsmsphonenumberactionsettingsDud struct { 
    

}

// Setsmsphonenumberactionsettings
type Setsmsphonenumberactionsettings struct { 
    // SenderSmsPhoneNumber - The string address for the sms phone number.
    SenderSmsPhoneNumber string `json:"senderSmsPhoneNumber"`

}

// String returns a JSON representation of the model
func (o *Setsmsphonenumberactionsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setsmsphonenumberactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Setsmsphonenumberactionsettings

    if SetsmsphonenumberactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    SetsmsphonenumberactionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        SenderSmsPhoneNumber string `json:"senderSmsPhoneNumber"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

