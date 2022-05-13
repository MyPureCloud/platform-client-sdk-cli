package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimserviceproviderconfigauthenticationschemeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimserviceproviderconfigauthenticationschemeDud struct { 
    Name string `json:"name"`


    Description string `json:"description"`


    SpecUri string `json:"specUri"`


    DocumentationUri string `json:"documentationUri"`


    VarType string `json:"type"`


    Primary bool `json:"primary"`

}

// Scimserviceproviderconfigauthenticationscheme - Defines an authentication scheme in the SCIM service provider's configuration.
type Scimserviceproviderconfigauthenticationscheme struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigauthenticationscheme) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimserviceproviderconfigauthenticationscheme) MarshalJSON() ([]byte, error) {
    type Alias Scimserviceproviderconfigauthenticationscheme

    if ScimserviceproviderconfigauthenticationschemeMarshalled {
        return []byte("{}"), nil
    }
    ScimserviceproviderconfigauthenticationschemeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

