package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DefinitionDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Definition
type Definition struct { 
    // Name
    Name string `json:"name"`


    // Title
    Title string `json:"title"`


    // Description
    Description string `json:"description"`


    // VarType
    VarType []string `json:"type"`


    // Pattern
    Pattern string `json:"pattern"`


    // Items
    Items Items `json:"items"`


    

}

// String returns a JSON representation of the model
func (o *Definition) String() string {
    
    
    
     o.VarType = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Definition) MarshalJSON() ([]byte, error) {
    type Alias Definition

    if DefinitionMarshalled {
        return []byte("{}"), nil
    }
    DefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType []string `json:"type"`
        
        Pattern string `json:"pattern"`
        
        Items Items `json:"items"`
        *Alias
    }{

        


        


        


        
        VarType: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

