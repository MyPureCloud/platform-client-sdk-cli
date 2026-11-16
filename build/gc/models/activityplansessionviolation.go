package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplansessionviolationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplansessionviolationDud struct { 
    


    

}

// Activityplansessionviolation
type Activityplansessionviolation struct { 
    // VarType - The type of violation for this session
    VarType string `json:"type"`


    // Details - The detail of the violation
    Details Activityplanviolationdetails `json:"details"`

}

// String returns a JSON representation of the model
func (o *Activityplansessionviolation) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplansessionviolation) MarshalJSON() ([]byte, error) {
    type Alias Activityplansessionviolation

    if ActivityplansessionviolationMarshalled {
        return []byte("{}"), nil
    }
    ActivityplansessionviolationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Details Activityplanviolationdetails `json:"details"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

