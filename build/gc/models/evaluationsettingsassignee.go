package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsettingsassigneeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsettingsassigneeDud struct { 
    


    

}

// Evaluationsettingsassignee
type Evaluationsettingsassignee struct { 
    // User - The user the dispute should be assigned to
    User Userreferencewithname `json:"user"`


    // VarType - The assignee type. Valid values: Original, Individual, None
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Evaluationsettingsassignee) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsettingsassignee) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsettingsassignee

    if EvaluationsettingsassigneeMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsettingsassigneeMarshalled = true

    return json.Marshal(&struct {
        
        User Userreferencewithname `json:"user"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

