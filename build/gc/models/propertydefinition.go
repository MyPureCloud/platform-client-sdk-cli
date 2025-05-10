package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PropertydefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PropertydefinitionDud struct { 
    


    


    


    


    


    


    

}

// Propertydefinition
type Propertydefinition struct { 
    // Title
    Title string `json:"title"`


    // Description
    Description string `json:"description"`


    // VarType
    VarType []string `json:"type"`


    // Pattern
    Pattern string `json:"pattern"`


    // Format
    Format string `json:"format"`


    // Items
    Items Items `json:"items"`


    // Properties
    Properties map[string]Propertydefinition `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Propertydefinition) String() string {
    
    
     o.VarType = []string{""} 
    
    
    
     o.Properties = map[string]Propertydefinition{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Propertydefinition) MarshalJSON() ([]byte, error) {
    type Alias Propertydefinition

    if PropertydefinitionMarshalled {
        return []byte("{}"), nil
    }
    PropertydefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType []string `json:"type"`
        
        Pattern string `json:"pattern"`
        
        Format string `json:"format"`
        
        Items Items `json:"items"`
        
        Properties map[string]Propertydefinition `json:"properties"`
        *Alias
    }{

        


        


        
        VarType: []string{""},
        


        


        


        


        
        Properties: map[string]Propertydefinition{"": {}},
        

        Alias: (*Alias)(u),
    })
}

