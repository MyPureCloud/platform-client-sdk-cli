package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConstraintconflictmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConstraintconflictmessageDud struct { 
    


    

}

// Constraintconflictmessage
type Constraintconflictmessage struct { 
    // Message - Message for how to resolve a set of conflicted work plan constraints
    Message Workplanconstraintconflictmessage `json:"message"`


    // ConflictedConstraintMessages - Messages for the set of conflicted work plan constraints. Each element indicates the message of a work plan constraint that is conflicted in the set
    ConflictedConstraintMessages []Workplanconstraintmessage `json:"conflictedConstraintMessages"`

}

// String returns a JSON representation of the model
func (o *Constraintconflictmessage) String() string {
    
     o.ConflictedConstraintMessages = []Workplanconstraintmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Constraintconflictmessage) MarshalJSON() ([]byte, error) {
    type Alias Constraintconflictmessage

    if ConstraintconflictmessageMarshalled {
        return []byte("{}"), nil
    }
    ConstraintconflictmessageMarshalled = true

    return json.Marshal(&struct {
        
        Message Workplanconstraintconflictmessage `json:"message"`
        
        ConflictedConstraintMessages []Workplanconstraintmessage `json:"conflictedConstraintMessages"`
        *Alias
    }{

        


        
        ConflictedConstraintMessages: []Workplanconstraintmessage{{}},
        

        Alias: (*Alias)(u),
    })
}

