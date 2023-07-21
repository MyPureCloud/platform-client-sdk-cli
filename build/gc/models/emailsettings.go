package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailsettingsDud struct { 
    

}

// Emailsettings
type Emailsettings struct { 
    // MultipleRouteDestinationsOnInboundEmailEnabled - This setting allows a single inbound email that contains multiple routes configured in Genesys Cloud to create a conversation per route. When this setting is disabled only a single conversation will be created
    MultipleRouteDestinationsOnInboundEmailEnabled bool `json:"multipleRouteDestinationsOnInboundEmailEnabled"`

}

// String returns a JSON representation of the model
func (o *Emailsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailsettings) MarshalJSON() ([]byte, error) {
    type Alias Emailsettings

    if EmailsettingsMarshalled {
        return []byte("{}"), nil
    }
    EmailsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MultipleRouteDestinationsOnInboundEmailEnabled bool `json:"multipleRouteDestinationsOnInboundEmailEnabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

