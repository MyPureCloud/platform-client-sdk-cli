package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentdataactionschemasMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentdataactionschemasDud struct { 
    


    

}

// Agenticvirtualagentdataactionschemas - Data action input and output JSON schemas.
type Agenticvirtualagentdataactionschemas struct { 
    // Inputs - Input JSON schema for the selected data action.
    Inputs map[string]interface{} `json:"inputs"`


    // Outputs - Output JSON schema for the selected data action.
    Outputs map[string]interface{} `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentdataactionschemas) String() string {
     o.Inputs = map[string]interface{}{"": Interface{}} 
     o.Outputs = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentdataactionschemas) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentdataactionschemas

    if AgenticvirtualagentdataactionschemasMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentdataactionschemasMarshalled = true

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

