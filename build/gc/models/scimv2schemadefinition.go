package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2schemadefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2schemadefinitionDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Description string `json:"description"`


    Attributes []Scimv2schemaattribute `json:"attributes"`


    Meta Scimmetadata `json:"meta"`

}

// Scimv2schemadefinition - Defines a SCIM schema.
type Scimv2schemadefinition struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimv2schemadefinition) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2schemadefinition) MarshalJSON() ([]byte, error) {
    type Alias Scimv2schemadefinition

    if Scimv2schemadefinitionMarshalled {
        return []byte("{}"), nil
    }
    Scimv2schemadefinitionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

