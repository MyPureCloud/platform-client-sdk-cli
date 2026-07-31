package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttoolMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttoolDud struct { 
    


    


    


    


    


    

}

// Agenticvirtualagenttool
type Agenticvirtualagenttool struct { 
    // VarType - Tool type discriminator.
    VarType string `json:"type"`


    // Name - Name of the tool. Use a clear, specific action name.
    Name string `json:"name"`


    // Description - Additional information about how this tool works to help the virtual agent decide when and how to use it.
    Description string `json:"description"`


    // InputInstructions - Additional instructions specific to using this tool.
    InputInstructions []string `json:"inputInstructions"`


    // OutputInstructions - Instructions that apply after successful tool execution based on tool output conditions.
    OutputInstructions []Agenticvirtualagenttooloutputinstruction `json:"outputInstructions"`


    // Target - Resource selected for this tool.
    Target Domainentityref `json:"target"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttool) String() string {
    
    
    
     o.InputInstructions = []string{""} 
     o.OutputInstructions = []Agenticvirtualagenttooloutputinstruction{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttool) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttool

    if AgenticvirtualagenttoolMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttoolMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        InputInstructions []string `json:"inputInstructions"`
        
        OutputInstructions []Agenticvirtualagenttooloutputinstruction `json:"outputInstructions"`
        
        Target Domainentityref `json:"target"`
        *Alias
    }{

        


        


        


        
        InputInstructions: []string{""},
        


        
        OutputInstructions: []Agenticvirtualagenttooloutputinstruction{{}},
        


        

        Alias: (*Alias)(u),
    })
}

