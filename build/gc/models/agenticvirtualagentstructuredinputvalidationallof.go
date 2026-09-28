package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredinputvalidationallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredinputvalidationallofDud struct { 
    


    


    

}

// Agenticvirtualagentstructuredinputvalidationallof - Input validation using a structured condition group.
type Agenticvirtualagentstructuredinputvalidationallof struct { 
    // VarType - Validation type discriminator.
    VarType string `json:"type"`


    // VarIf
    VarIf Agenticvirtualagentstructuredconditiongroup `json:"if"`


    // VarElse - Instruction for the virtual agent when the validation condition is not met.
    VarElse string `json:"else"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredinputvalidationallof) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredinputvalidationallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredinputvalidationallof

    if AgenticvirtualagentstructuredinputvalidationallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredinputvalidationallofMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        VarIf Agenticvirtualagentstructuredconditiongroup `json:"if"`
        
        VarElse string `json:"else"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

