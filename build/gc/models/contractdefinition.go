package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContractdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContractdefinitionDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contractdefinition
type Contractdefinition struct { 
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
    Items Contractitems `json:"items"`


    

}

// String returns a JSON representation of the model
func (o *Contractdefinition) String() string {
    
    
    
     o.VarType = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contractdefinition) MarshalJSON() ([]byte, error) {
    type Alias Contractdefinition

    if ContractdefinitionMarshalled {
        return []byte("{}"), nil
    }
    ContractdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType []string `json:"type"`
        
        Pattern string `json:"pattern"`
        
        Items Contractitems `json:"items"`
        *Alias
    }{

        


        


        


        
        VarType: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

