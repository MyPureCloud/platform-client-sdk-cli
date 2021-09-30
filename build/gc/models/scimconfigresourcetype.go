package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimconfigresourcetypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimconfigresourcetypeDud struct { 
    Id string `json:"id"`


    Schemas []string `json:"schemas"`


    Name string `json:"name"`


    Description string `json:"description"`


    Schema string `json:"schema"`


    SchemaExtensions []Scimconfigresourcetypeschemaextension `json:"schemaExtensions"`


    Endpoint string `json:"endpoint"`


    Meta Scimmetadata `json:"meta"`

}

// Scimconfigresourcetype - Defines a SCIM resource.
type Scimconfigresourcetype struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetype) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimconfigresourcetype) MarshalJSON() ([]byte, error) {
    type Alias Scimconfigresourcetype

    if ScimconfigresourcetypeMarshalled {
        return []byte("{}"), nil
    }
    ScimconfigresourcetypeMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

