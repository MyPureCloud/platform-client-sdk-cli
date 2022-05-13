package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanconstraintconflictmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanconstraintconflictmessageDud struct { 
    


    

}

// Workplanconstraintconflictmessage
type Workplanconstraintconflictmessage struct { 
    // VarType - Type of constraint conflict that can be resolved by clients in order to generate agent schedules
    VarType string `json:"type"`


    // Arguments - The arguments to the type of the message that can help clients resolve validation issues
    Arguments []Workplanvalidationmessageargument `json:"arguments"`

}

// String returns a JSON representation of the model
func (o *Workplanconstraintconflictmessage) String() string {
    
     o.Arguments = []Workplanvalidationmessageargument{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanconstraintconflictmessage) MarshalJSON() ([]byte, error) {
    type Alias Workplanconstraintconflictmessage

    if WorkplanconstraintconflictmessageMarshalled {
        return []byte("{}"), nil
    }
    WorkplanconstraintconflictmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Arguments []Workplanvalidationmessageargument `json:"arguments"`
        *Alias
    }{

        


        
        Arguments: []Workplanvalidationmessageargument{{}},
        

        Alias: (*Alias)(u),
    })
}

