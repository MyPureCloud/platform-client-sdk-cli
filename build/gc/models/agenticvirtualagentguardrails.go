package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentguardrailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentguardrailsDud struct { 
    

}

// Agenticvirtualagentguardrails - Guardrail rules for a virtual agent.
type Agenticvirtualagentguardrails struct { 
    // Custom - Custom guardrail rules used to detect and block matching user behavior.
    Custom []Agenticvirtualagentguardrailinstruction `json:"custom"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentguardrails) String() string {
     o.Custom = []Agenticvirtualagentguardrailinstruction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentguardrails) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentguardrails

    if AgenticvirtualagentguardrailsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentguardrailsMarshalled = true

    return json.Marshal(&struct {
        
        Custom []Agenticvirtualagentguardrailinstruction `json:"custom"`
        *Alias
    }{

        
        Custom: []Agenticvirtualagentguardrailinstruction{{}},
        

        Alias: (*Alias)(u),
    })
}

