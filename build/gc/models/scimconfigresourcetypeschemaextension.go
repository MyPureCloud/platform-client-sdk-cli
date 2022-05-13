package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimconfigresourcetypeschemaextensionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimconfigresourcetypeschemaextensionDud struct { 
    Schema string `json:"schema"`


    Required bool `json:"required"`

}

// Scimconfigresourcetypeschemaextension - Defines a SCIM resource type's schema extension.
type Scimconfigresourcetypeschemaextension struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetypeschemaextension) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimconfigresourcetypeschemaextension) MarshalJSON() ([]byte, error) {
    type Alias Scimconfigresourcetypeschemaextension

    if ScimconfigresourcetypeschemaextensionMarshalled {
        return []byte("{}"), nil
    }
    ScimconfigresourcetypeschemaextensionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

