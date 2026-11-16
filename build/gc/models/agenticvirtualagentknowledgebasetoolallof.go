package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentknowledgebasetoolallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentknowledgebasetoolallofDud struct { 
    


    

}

// Agenticvirtualagentknowledgebasetoolallof - Knowledge base tool.
type Agenticvirtualagentknowledgebasetoolallof struct { 
    // Inputs - Inputs for this knowledge tool. Knowledge tool inputs are used for gating/preconditions only and are not sent with the underlying tool at runtime.
    Inputs []Agenticvirtualagenttoolinput `json:"inputs"`


    // InputValidation - Conditions that must be checked before invoking this tool.
    InputValidation []Agenticvirtualagentinputvalidation `json:"inputValidation"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentknowledgebasetoolallof) String() string {
     o.Inputs = []Agenticvirtualagenttoolinput{{}} 
     o.InputValidation = []Agenticvirtualagentinputvalidation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentknowledgebasetoolallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentknowledgebasetoolallof

    if AgenticvirtualagentknowledgebasetoolallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentknowledgebasetoolallofMarshalled = true

    return json.Marshal(&struct {
        
        Inputs []Agenticvirtualagenttoolinput `json:"inputs"`
        
        InputValidation []Agenticvirtualagentinputvalidation `json:"inputValidation"`
        *Alias
    }{

        
        Inputs: []Agenticvirtualagenttoolinput{{}},
        


        
        InputValidation: []Agenticvirtualagentinputvalidation{{}},
        

        Alias: (*Alias)(u),
    })
}

