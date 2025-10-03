package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContractpropertydefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContractpropertydefinitionDud struct { 
    


    


    


    


    


    


    

}

// Contractpropertydefinition
type Contractpropertydefinition struct { 
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
    Items Contractitems `json:"items"`


    // Properties
    Properties map[string]Contractpropertydefinition `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Contractpropertydefinition) String() string {
    
    
     o.VarType = []string{""} 
    
    
    
     o.Properties = map[string]Contractpropertydefinition{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contractpropertydefinition) MarshalJSON() ([]byte, error) {
    type Alias Contractpropertydefinition

    if ContractpropertydefinitionMarshalled {
        return []byte("{}"), nil
    }
    ContractpropertydefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType []string `json:"type"`
        
        Pattern string `json:"pattern"`
        
        Format string `json:"format"`
        
        Items Contractitems `json:"items"`
        
        Properties map[string]Contractpropertydefinition `json:"properties"`
        *Alias
    }{

        


        


        
        VarType: []string{""},
        


        


        


        


        
        Properties: map[string]Contractpropertydefinition{"": {}},
        

        Alias: (*Alias)(u),
    })
}

