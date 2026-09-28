package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredoutputinstructionallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredoutputinstructionallofDud struct { 
    


    

}

// Agenticvirtualagentstructuredoutputinstructionallof - Instruction for handling tool output based on a structured condition group.
type Agenticvirtualagentstructuredoutputinstructionallof struct { 
    // VarType - Output instruction type discriminator.
    VarType string `json:"type"`


    // When
    When Agenticvirtualagentstructuredoutputconditiongroup `json:"when"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredoutputinstructionallof) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredoutputinstructionallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredoutputinstructionallof

    if AgenticvirtualagentstructuredoutputinstructionallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredoutputinstructionallofMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        When Agenticvirtualagentstructuredoutputconditiongroup `json:"when"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

