package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttoolinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttoolinputDud struct { 
    


    


    


    


    


    

}

// Agenticvirtualagenttoolinput - Input for a tool.
type Agenticvirtualagenttoolinput struct { 
    // TargetName - The unique name that identifies this input parameter within the tool
    TargetName string `json:"targetName"`


    // VarType - Input type name. The valid referenced type depends on the input source.
    VarType string `json:"type"`


    // Source - Source of the input value.
    Source string `json:"source"`


    // Required - Whether this input must be supplied.
    Required bool `json:"required"`


    // FallbackToUser - Whether the virtual agent should ask the user for this input value when it is not available from the configured source.
    FallbackToUser bool `json:"fallbackToUser"`


    // Mapping - Path used to extract this input from a previous tool output. Only valid when source is 'ToolOutput'. The path starts with a tool output type name, may contain only string property names or integer array indexes, and must resolve to a primitive value.
    Mapping []interface{} `json:"mapping"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttoolinput) String() string {
    
    
    
    
    
     o.Mapping = []interface{}{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttoolinput) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttoolinput

    if AgenticvirtualagenttoolinputMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttoolinputMarshalled = true

    return json.Marshal(&struct {
        
        TargetName string `json:"targetName"`
        
        VarType string `json:"type"`
        
        Source string `json:"source"`
        
        Required bool `json:"required"`
        
        FallbackToUser bool `json:"fallbackToUser"`
        
        Mapping []interface{} `json:"mapping"`
        *Alias
    }{

        


        


        


        


        


        
        Mapping: []interface{}{},
        

        Alias: (*Alias)(u),
    })
}

