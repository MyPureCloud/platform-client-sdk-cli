package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttooloutputinstructionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttooloutputinstructionDud struct { 
    


    

}

// Agenticvirtualagenttooloutputinstruction
type Agenticvirtualagenttooloutputinstruction struct { 
    // VarType - Output instruction type discriminator.
    VarType string `json:"type"`


    // Then - Instruction to follow when the condition is met.
    Then string `json:"then"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttooloutputinstruction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttooloutputinstruction) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttooloutputinstruction

    if AgenticvirtualagenttooloutputinstructionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttooloutputinstructionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Then string `json:"then"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

