package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentguardrailinstructionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentguardrailinstructionDud struct { 
    


    

}

// Agenticvirtualagentguardrailinstruction - Custom guardrail rule for a virtual agent.
type Agenticvirtualagentguardrailinstruction struct { 
    // Instruction - Natural language rule describing user behavior to detect and block.
    Instruction string `json:"instruction"`


    // Enabled - Whether this custom guardrail rule is active.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentguardrailinstruction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentguardrailinstruction) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentguardrailinstruction

    if AgenticvirtualagentguardrailinstructionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentguardrailinstructionMarshalled = true

    return json.Marshal(&struct {
        
        Instruction string `json:"instruction"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

