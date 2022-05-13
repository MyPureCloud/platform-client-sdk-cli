package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimserviceproviderconfigfilterfeatureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimserviceproviderconfigfilterfeatureDud struct { 
    Supported bool `json:"supported"`


    MaxResults int `json:"maxResults"`

}

// Scimserviceproviderconfigfilterfeature - Defines a \"filter\" request in the SCIM service provider's configuration.
type Scimserviceproviderconfigfilterfeature struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigfilterfeature) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimserviceproviderconfigfilterfeature) MarshalJSON() ([]byte, error) {
    type Alias Scimserviceproviderconfigfilterfeature

    if ScimserviceproviderconfigfilterfeatureMarshalled {
        return []byte("{}"), nil
    }
    ScimserviceproviderconfigfilterfeatureMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

