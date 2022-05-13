package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanconstraintmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanconstraintmessageDud struct { 
    


    

}

// Workplanconstraintmessage
type Workplanconstraintmessage struct { 
    // VarType - Type of the work plan constraint in this message
    VarType string `json:"type"`


    // Arguments - Arguments of the message that provide information about the constraint that is being conflicted with, such as the value of the constraint
    Arguments []Workplanvalidationmessageargument `json:"arguments"`

}

// String returns a JSON representation of the model
func (o *Workplanconstraintmessage) String() string {
    
     o.Arguments = []Workplanvalidationmessageargument{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanconstraintmessage) MarshalJSON() ([]byte, error) {
    type Alias Workplanconstraintmessage

    if WorkplanconstraintmessageMarshalled {
        return []byte("{}"), nil
    }
    WorkplanconstraintmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Arguments []Workplanvalidationmessageargument `json:"arguments"`
        *Alias
    }{

        


        
        Arguments: []Workplanvalidationmessageargument{{}},
        

        Alias: (*Alias)(u),
    })
}

