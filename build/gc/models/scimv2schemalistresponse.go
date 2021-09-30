package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2schemalistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2schemalistresponseDud struct { 
    


    TotalResults int `json:"totalResults"`


    StartIndex int `json:"startIndex"`


    ItemsPerPage int `json:"itemsPerPage"`


    Resources []Scimv2schemadefinition `json:"Resources"`

}

// Scimv2schemalistresponse - Defines the list response for SCIM resource types.
type Scimv2schemalistresponse struct { 
    // Schemas - The list of supported schemas.
    Schemas []string `json:"schemas"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimv2schemalistresponse) String() string {
    
    
     o.Schemas = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2schemalistresponse) MarshalJSON() ([]byte, error) {
    type Alias Scimv2schemalistresponse

    if Scimv2schemalistresponseMarshalled {
        return []byte("{}"), nil
    }
    Scimv2schemalistresponseMarshalled = true

    return json.Marshal(&struct { 
        Schemas []string `json:"schemas"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        
        Schemas: []string{""},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

