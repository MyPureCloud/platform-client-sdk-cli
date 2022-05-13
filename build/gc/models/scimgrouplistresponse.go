package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimgrouplistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimgrouplistresponseDud struct { 
    


    TotalResults int `json:"totalResults"`


    StartIndex int `json:"startIndex"`


    ItemsPerPage int `json:"itemsPerPage"`


    Resources []Scimv2group `json:"Resources"`

}

// Scimgrouplistresponse - Defines a response for a list of SCIM groups.
type Scimgrouplistresponse struct { 
    // Schemas - The list of supported schemas.
    Schemas []string `json:"schemas"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimgrouplistresponse) String() string {
     o.Schemas = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimgrouplistresponse) MarshalJSON() ([]byte, error) {
    type Alias Scimgrouplistresponse

    if ScimgrouplistresponseMarshalled {
        return []byte("{}"), nil
    }
    ScimgrouplistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Schemas []string `json:"schemas"`
        *Alias
    }{

        
        Schemas: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

