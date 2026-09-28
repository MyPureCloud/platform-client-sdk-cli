package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredoutputruleallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredoutputruleallofDud struct { 
    


    


    

}

// Agenticvirtualagentstructuredoutputruleallof - A single structured tool-output rule. The mapping is a path into the tool output type.
type Agenticvirtualagentstructuredoutputruleallof struct { 
    // Mapping - Path into the tool output type this rule applies to. Each element is a field name (string) or an array index (integer).
    Mapping []interface{} `json:"mapping"`


    // Operator - Operator to apply to the value at the mapped path.
    Operator string `json:"operator"`


    // Value - Value to compare against. May be a string, integer, number, boolean, or null. Not required for 'IsNull', 'IsNotNull', 'IsEmpty', or 'IsNotEmpty' operators.
    Value interface{} `json:"value"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredoutputruleallof) String() string {
     o.Mapping = []interface{}{} 
    
     o.Value = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredoutputruleallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredoutputruleallof

    if AgenticvirtualagentstructuredoutputruleallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredoutputruleallofMarshalled = true

    return json.Marshal(&struct {
        
        Mapping []interface{} `json:"mapping"`
        
        Operator string `json:"operator"`
        
        Value interface{} `json:"value"`
        *Alias
    }{

        
        Mapping: []interface{}{},
        


        


        
        Value: Interface{},
        

        Alias: (*Alias)(u),
    })
}

