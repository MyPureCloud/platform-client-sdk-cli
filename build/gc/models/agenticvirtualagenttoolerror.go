package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttoolerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttoolerrorDud struct { 
    


    

}

// Agenticvirtualagenttoolerror - Error handling configuration for a tool.
type Agenticvirtualagenttoolerror struct { 
    // VarType - Error type name as defined in the types list.
    VarType string `json:"type"`


    // Instruction - Instruction for how the virtual agent should handle this error.
    Instruction string `json:"instruction"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttoolerror) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttoolerror) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttoolerror

    if AgenticvirtualagenttoolerrorMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttoolerrorMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Instruction string `json:"instruction"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

