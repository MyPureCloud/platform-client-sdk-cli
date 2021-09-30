package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateworkplanmessagesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateworkplanmessagesDud struct { 
    


    

}

// Validateworkplanmessages
type Validateworkplanmessages struct { 
    // ViolationMessages - Messages for work plan violating some rules such as no shifts in a work plan
    ViolationMessages []Workplanconfigurationviolationmessage `json:"violationMessages"`


    // ConstraintConflictMessage - This field is not null when there is a set of work plan constraints that conflict thus agent schedules cannot be generated
    ConstraintConflictMessage Constraintconflictmessage `json:"constraintConflictMessage"`

}

// String returns a JSON representation of the model
func (o *Validateworkplanmessages) String() string {
    
    
     o.ViolationMessages = []Workplanconfigurationviolationmessage{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateworkplanmessages) MarshalJSON() ([]byte, error) {
    type Alias Validateworkplanmessages

    if ValidateworkplanmessagesMarshalled {
        return []byte("{}"), nil
    }
    ValidateworkplanmessagesMarshalled = true

    return json.Marshal(&struct { 
        ViolationMessages []Workplanconfigurationviolationmessage `json:"violationMessages"`
        
        ConstraintConflictMessage Constraintconflictmessage `json:"constraintConflictMessage"`
        
        *Alias
    }{
        

        
        ViolationMessages: []Workplanconfigurationviolationmessage{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

