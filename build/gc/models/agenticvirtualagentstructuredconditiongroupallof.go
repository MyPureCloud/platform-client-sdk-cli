package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredconditiongroupallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredconditiongroupallofDud struct { 
    


    

}

// Agenticvirtualagentstructuredconditiongroupallof - A group of structured input validation rules or nested groups.
type Agenticvirtualagentstructuredconditiongroupallof struct { 
    // Group - Logical operator used to combine the rules in this group.
    Group string `json:"group"`


    // Rules - Structured rules or nested condition groups in this group.
    Rules []map[string]interface{} `json:"rules"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredconditiongroupallof) String() string {
    
     o.Rules = []map[string]interface{}{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredconditiongroupallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredconditiongroupallof

    if AgenticvirtualagentstructuredconditiongroupallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredconditiongroupallofMarshalled = true

    return json.Marshal(&struct {
        
        Group string `json:"group"`
        
        Rules []map[string]interface{} `json:"rules"`
        *Alias
    }{

        


        
        Rules: []map[string]interface{}{{}},
        

        Alias: (*Alias)(u),
    })
}

