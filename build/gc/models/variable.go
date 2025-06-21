package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VariableMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VariableDud struct { 
    


    


    


    

}

// Variable
type Variable struct { 
    // Name - The name of the variable.
    Name string `json:"name"`


    // VarType - The data type of the variable.
    VarType string `json:"type"`


    // Scope - The scope that determines the variable's usage context within Guides runtime.
    Scope string `json:"scope"`


    // Description - The description of the variable used by Guides runtime for input/output handling.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Variable) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Variable) MarshalJSON() ([]byte, error) {
    type Alias Variable

    if VariableMarshalled {
        return []byte("{}"), nil
    }
    VariableMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Scope string `json:"scope"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

