package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredoutputinstructionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredoutputinstructionDud struct { 
    


    

}

// Agenticvirtualagentstructuredoutputinstruction
type Agenticvirtualagentstructuredoutputinstruction struct { 
    // VarType - Output instruction type discriminator.
    VarType string `json:"type"`


    // When
    When Agenticvirtualagentstructuredoutputconditiongroup `json:"when"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredoutputinstruction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredoutputinstruction) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredoutputinstruction

    if AgenticvirtualagentstructuredoutputinstructionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredoutputinstructionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        When Agenticvirtualagentstructuredoutputconditiongroup `json:"when"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

