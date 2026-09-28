package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentdataactiontoolallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentdataactiontoolallofDud struct { 
    


    


    


    


    

}

// Agenticvirtualagentdataactiontoolallof - Data action tool.
type Agenticvirtualagentdataactiontoolallof struct { 
    // Errors - Error types this tool can raise.
    Errors []Agenticvirtualagenttoolerror `json:"errors"`


    // Inputs - Inputs passed to this data action tool.
    Inputs []Agenticvirtualagenttoolinput `json:"inputs"`


    // Output - Name of the output type this tool returns when it completes successfully.
    Output string `json:"output"`


    // InputValidation - Conditions that must be checked before invoking this tool.
    InputValidation []Agenticvirtualagentinputvalidation `json:"inputValidation"`


    // Schemas
    Schemas Agenticvirtualagentdataactionschemas `json:"schemas"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentdataactiontoolallof) String() string {
     o.Errors = []Agenticvirtualagenttoolerror{{}} 
     o.Inputs = []Agenticvirtualagenttoolinput{{}} 
    
     o.InputValidation = []Agenticvirtualagentinputvalidation{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentdataactiontoolallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentdataactiontoolallof

    if AgenticvirtualagentdataactiontoolallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentdataactiontoolallofMarshalled = true

    return json.Marshal(&struct {
        
        Errors []Agenticvirtualagenttoolerror `json:"errors"`
        
        Inputs []Agenticvirtualagenttoolinput `json:"inputs"`
        
        Output string `json:"output"`
        
        InputValidation []Agenticvirtualagentinputvalidation `json:"inputValidation"`
        
        Schemas Agenticvirtualagentdataactionschemas `json:"schemas"`
        *Alias
    }{

        
        Errors: []Agenticvirtualagenttoolerror{{}},
        


        
        Inputs: []Agenticvirtualagenttoolinput{{}},
        


        


        
        InputValidation: []Agenticvirtualagentinputvalidation{{}},
        


        

        Alias: (*Alias)(u),
    })
}

