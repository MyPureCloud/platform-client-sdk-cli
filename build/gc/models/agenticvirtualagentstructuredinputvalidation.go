package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredinputvalidationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredinputvalidationDud struct { 
    


    


    

}

// Agenticvirtualagentstructuredinputvalidation
type Agenticvirtualagentstructuredinputvalidation struct { 
    // VarType - Validation type discriminator.
    VarType string `json:"type"`


    // VarIf
    VarIf Agenticvirtualagentstructuredconditiongroup `json:"if"`


    // VarElse - Instruction for the virtual agent when the validation condition is not met.
    VarElse string `json:"else"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredinputvalidation) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredinputvalidation) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredinputvalidation

    if AgenticvirtualagentstructuredinputvalidationMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredinputvalidationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        VarIf Agenticvirtualagentstructuredconditiongroup `json:"if"`
        
        VarElse string `json:"else"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

