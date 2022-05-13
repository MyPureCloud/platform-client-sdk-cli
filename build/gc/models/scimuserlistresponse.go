package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimuserlistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimuserlistresponseDud struct { 
    


    TotalResults int `json:"totalResults"`


    StartIndex int `json:"startIndex"`


    ItemsPerPage int `json:"itemsPerPage"`


    Resources []Scimv2user `json:"Resources"`

}

// Scimuserlistresponse - Defines a response for a list of SCIM users.
type Scimuserlistresponse struct { 
    // Schemas - The list of supported schemas.
    Schemas []string `json:"schemas"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimuserlistresponse) String() string {
     o.Schemas = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimuserlistresponse) MarshalJSON() ([]byte, error) {
    type Alias Scimuserlistresponse

    if ScimuserlistresponseMarshalled {
        return []byte("{}"), nil
    }
    ScimuserlistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Schemas []string `json:"schemas"`
        *Alias
    }{

        
        Schemas: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

