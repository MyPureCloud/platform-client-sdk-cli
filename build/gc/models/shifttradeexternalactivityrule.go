package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeexternalactivityruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeexternalactivityruleDud struct { 
    


    


    

}

// Shifttradeexternalactivityrule
type Shifttradeexternalactivityrule struct { 
    // ExternalActivityType - The external activity type to which to apply this rule
    ExternalActivityType string `json:"externalActivityType"`


    // Action - The action this rule invokes
    Action string `json:"action"`


    // ActivityCodeIdReplacement - The ID of the activity code with which to replace to replace external activities (required if action == KeepWithAgent, must be a default activity code ID)
    ActivityCodeIdReplacement string `json:"activityCodeIdReplacement"`

}

// String returns a JSON representation of the model
func (o *Shifttradeexternalactivityrule) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeexternalactivityrule) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeexternalactivityrule

    if ShifttradeexternalactivityruleMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeexternalactivityruleMarshalled = true

    return json.Marshal(&struct {
        
        ExternalActivityType string `json:"externalActivityType"`
        
        Action string `json:"action"`
        
        ActivityCodeIdReplacement string `json:"activityCodeIdReplacement"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

