package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredoutputconditiongroupallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredoutputconditiongroupallofDud struct { 
    


    

}

// Agenticvirtualagentstructuredoutputconditiongroupallof - A group of structured tool-output rules or nested groups.
type Agenticvirtualagentstructuredoutputconditiongroupallof struct { 
    // Group - Logical operator used to combine the rules in this group.
    Group string `json:"group"`


    // Rules - Structured output rules or nested condition groups in this group.
    Rules []map[string]interface{} `json:"rules"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredoutputconditiongroupallof) String() string {
    
     o.Rules = []map[string]interface{}{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredoutputconditiongroupallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredoutputconditiongroupallof

    if AgenticvirtualagentstructuredoutputconditiongroupallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredoutputconditiongroupallofMarshalled = true

    return json.Marshal(&struct {
        
        Group string `json:"group"`
        
        Rules []map[string]interface{} `json:"rules"`
        *Alias
    }{

        


        
        Rules: []map[string]interface{}{{}},
        

        Alias: (*Alias)(u),
    })
}

