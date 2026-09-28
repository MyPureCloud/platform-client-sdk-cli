package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentpythoninputvalidationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentpythoninputvalidationDud struct { 
    


    


    

}

// Agenticvirtualagentpythoninputvalidation
type Agenticvirtualagentpythoninputvalidation struct { 
    // VarType - Validation type discriminator.
    VarType string `json:"type"`


    // VarIf - Python condition that must evaluate to true before invoking the tool.
    VarIf string `json:"if"`


    // VarElse - Instruction for the virtual agent when the validation condition is not met.
    VarElse string `json:"else"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentpythoninputvalidation) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentpythoninputvalidation) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentpythoninputvalidation

    if AgenticvirtualagentpythoninputvalidationMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentpythoninputvalidationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        VarIf string `json:"if"`
        
        VarElse string `json:"else"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

