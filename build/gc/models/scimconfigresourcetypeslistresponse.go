package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimconfigresourcetypeslistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimconfigresourcetypeslistresponseDud struct { 
    


    TotalResults int `json:"totalResults"`


    StartIndex int `json:"startIndex"`


    ItemsPerPage int `json:"itemsPerPage"`


    Resources []Scimconfigresourcetype `json:"Resources"`

}

// Scimconfigresourcetypeslistresponse - Defines a response for a list of SCIM resource types.
type Scimconfigresourcetypeslistresponse struct { 
    // Schemas - The list of supported schemas.
    Schemas []string `json:"schemas"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetypeslistresponse) String() string {
    
    
     o.Schemas = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimconfigresourcetypeslistresponse) MarshalJSON() ([]byte, error) {
    type Alias Scimconfigresourcetypeslistresponse

    if ScimconfigresourcetypeslistresponseMarshalled {
        return []byte("{}"), nil
    }
    ScimconfigresourcetypeslistresponseMarshalled = true

    return json.Marshal(&struct { 
        Schemas []string `json:"schemas"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        
        Schemas: []string{""},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

