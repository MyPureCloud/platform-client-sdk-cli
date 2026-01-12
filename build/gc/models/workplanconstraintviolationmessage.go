package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanconstraintviolationmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanconstraintviolationmessageDud struct { 
    


    

}

// Workplanconstraintviolationmessage
type Workplanconstraintviolationmessage struct { 
    // VarType - Type of the work plan constraint in this message.
    VarType string `json:"type"`


    // Arguments - Arguments of the message that provide information about the constraint that is being conflicted with such as the value of the constraint.
    Arguments []Unavailabletimesviolationmessageargument `json:"arguments"`

}

// String returns a JSON representation of the model
func (o *Workplanconstraintviolationmessage) String() string {
    
     o.Arguments = []Unavailabletimesviolationmessageargument{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanconstraintviolationmessage) MarshalJSON() ([]byte, error) {
    type Alias Workplanconstraintviolationmessage

    if WorkplanconstraintviolationmessageMarshalled {
        return []byte("{}"), nil
    }
    WorkplanconstraintviolationmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Arguments []Unavailabletimesviolationmessageargument `json:"arguments"`
        *Alias
    }{

        


        
        Arguments: []Unavailabletimesviolationmessageargument{{}},
        

        Alias: (*Alias)(u),
    })
}

