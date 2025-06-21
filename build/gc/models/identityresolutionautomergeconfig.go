package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentityresolutionautomergeconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentityresolutionautomergeconfigDud struct { 
    

}

// Identityresolutionautomergeconfig
type Identityresolutionautomergeconfig struct { 
    // AuthenticatedWebMessaging - Whether automerging is enabled for Authenticated Webmessaging conversations in this channel.
    AuthenticatedWebMessaging bool `json:"authenticatedWebMessaging"`

}

// String returns a JSON representation of the model
func (o *Identityresolutionautomergeconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identityresolutionautomergeconfig) MarshalJSON() ([]byte, error) {
    type Alias Identityresolutionautomergeconfig

    if IdentityresolutionautomergeconfigMarshalled {
        return []byte("{}"), nil
    }
    IdentityresolutionautomergeconfigMarshalled = true

    return json.Marshal(&struct {
        
        AuthenticatedWebMessaging bool `json:"authenticatedWebMessaging"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

