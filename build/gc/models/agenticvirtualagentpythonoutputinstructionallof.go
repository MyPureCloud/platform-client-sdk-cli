package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentpythonoutputinstructionallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentpythonoutputinstructionallofDud struct { 
    


    

}

// Agenticvirtualagentpythonoutputinstructionallof - Instruction for handling tool output based on a Python condition.
type Agenticvirtualagentpythonoutputinstructionallof struct { 
    // VarType - Output instruction type discriminator.
    VarType string `json:"type"`


    // When - Python condition evaluated against the successful tool result to determine when this instruction applies.
    When string `json:"when"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentpythonoutputinstructionallof) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentpythonoutputinstructionallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentpythonoutputinstructionallof

    if AgenticvirtualagentpythonoutputinstructionallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentpythonoutputinstructionallofMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        When string `json:"when"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

