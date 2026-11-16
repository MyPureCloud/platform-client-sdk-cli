package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentmcptoolallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentmcptoolallofDud struct { 
    


    


    


    


    

}

// Agenticvirtualagentmcptoolallof - MCP tool.
type Agenticvirtualagentmcptoolallof struct { 
    // Inputs - Inputs passed to this MCP tool.
    Inputs []Agenticvirtualagenttoolinput `json:"inputs"`


    // Output - Name of the output type this tool returns when it completes successfully. When the MCP tool declares no output schema, this must reference a string aliased type.
    Output string `json:"output"`


    // InputValidation - Conditions that must be checked before invoking this tool.
    InputValidation []Agenticvirtualagentinputvalidation `json:"inputValidation"`


    // Schemas
    Schemas Agenticvirtualagentmcpschemas `json:"schemas"`


    // ContextRetentionEnabled - Whether this tool's call and result are retained in the virtual agent's conversation context when older turns are pruned to fit the model's context window. When true, the tool's messages are kept so the virtual agent does not lose or re-request this information.
    ContextRetentionEnabled bool `json:"contextRetentionEnabled"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentmcptoolallof) String() string {
     o.Inputs = []Agenticvirtualagenttoolinput{{}} 
    
     o.InputValidation = []Agenticvirtualagentinputvalidation{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentmcptoolallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentmcptoolallof

    if AgenticvirtualagentmcptoolallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentmcptoolallofMarshalled = true

    return json.Marshal(&struct {
        
        Inputs []Agenticvirtualagenttoolinput `json:"inputs"`
        
        Output string `json:"output"`
        
        InputValidation []Agenticvirtualagentinputvalidation `json:"inputValidation"`
        
        Schemas Agenticvirtualagentmcpschemas `json:"schemas"`
        
        ContextRetentionEnabled bool `json:"contextRetentionEnabled"`
        *Alias
    }{

        
        Inputs: []Agenticvirtualagenttoolinput{{}},
        


        


        
        InputValidation: []Agenticvirtualagentinputvalidation{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

