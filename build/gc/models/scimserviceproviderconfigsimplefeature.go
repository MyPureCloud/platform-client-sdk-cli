package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimserviceproviderconfigsimplefeatureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimserviceproviderconfigsimplefeatureDud struct { 
    Supported bool `json:"supported"`

}

// Scimserviceproviderconfigsimplefeature - Defines a request in the SCIM service provider's configuration.
type Scimserviceproviderconfigsimplefeature struct { 
    

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigsimplefeature) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimserviceproviderconfigsimplefeature) MarshalJSON() ([]byte, error) {
    type Alias Scimserviceproviderconfigsimplefeature

    if ScimserviceproviderconfigsimplefeatureMarshalled {
        return []byte("{}"), nil
    }
    ScimserviceproviderconfigsimplefeatureMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

