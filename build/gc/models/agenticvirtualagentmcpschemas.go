package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentmcpschemasMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentmcpschemasDud struct { 
    


    

}

// Agenticvirtualagentmcpschemas - MCP tool input and output JSON schemas.
type Agenticvirtualagentmcpschemas struct { 
    // Inputs - Input JSON schema for the selected MCP tool.
    Inputs map[string]interface{} `json:"inputs"`


    // Outputs - Output JSON schema for the selected MCP tool. Absent when the MCP tool declares no output schema.
    Outputs map[string]interface{} `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentmcpschemas) String() string {
     o.Inputs = map[string]interface{}{"": Interface{}} 
     o.Outputs = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentmcpschemas) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentmcpschemas

    if AgenticvirtualagentmcpschemasMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentmcpschemasMarshalled = true

    return json.Marshal(&struct {
        
        Inputs map[string]interface{} `json:"inputs"`
        
        Outputs map[string]interface{} `json:"outputs"`
        *Alias
    }{

        
        Inputs: map[string]interface{}{"": Interface{}},
        


        
        Outputs: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

