package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanconstraintsviolationmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanconstraintsviolationmessageDud struct { 
    


    

}

// Workplanconstraintsviolationmessage
type Workplanconstraintsviolationmessage struct { 
    // VarType - Message for how to resolve a set of conflicted work plan constraints
    VarType string `json:"type"`


    // FixableConstraintsMessages - If type == 'ConstraintConflictWithPotentialFixes', messages for the set of conflicted work plan constraints. Each element indicates the message of potential fix to unavailable times to resolve work plan constraints
    FixableConstraintsMessages []Workplanconstraintviolationmessage `json:"fixableConstraintsMessages"`

}

// String returns a JSON representation of the model
func (o *Workplanconstraintsviolationmessage) String() string {
    
     o.FixableConstraintsMessages = []Workplanconstraintviolationmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanconstraintsviolationmessage) MarshalJSON() ([]byte, error) {
    type Alias Workplanconstraintsviolationmessage

    if WorkplanconstraintsviolationmessageMarshalled {
        return []byte("{}"), nil
    }
    WorkplanconstraintsviolationmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        FixableConstraintsMessages []Workplanconstraintviolationmessage `json:"fixableConstraintsMessages"`
        *Alias
    }{

        


        
        FixableConstraintsMessages: []Workplanconstraintviolationmessage{{}},
        

        Alias: (*Alias)(u),
    })
}

