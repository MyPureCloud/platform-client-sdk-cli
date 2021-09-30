package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchemaDud struct { 
    Title string `json:"title"`


    Description string `json:"description"`


    VarType []string `json:"type"`


    Items Items `json:"items"`


    Pattern string `json:"pattern"`

}

// Schema
type Schema struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Schema) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schema) MarshalJSON() ([]byte, error) {
    type Alias Schema

    if SchemaMarshalled {
        return []byte("{}"), nil
    }
    SchemaMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

