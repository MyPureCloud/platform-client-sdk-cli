package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeactivityruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeactivityruleDud struct { 
    


    


    

}

// Shifttradeactivityrule
type Shifttradeactivityrule struct { 
    // ActivityCategory - The activity category to which to apply this rule
    ActivityCategory string `json:"activityCategory"`


    // Action - The action this rule invokes
    Action string `json:"action"`


    // ActivityCodeIdReplacement - The ID of the activity code with which to replace activities belonging to the original category if applicable (required if action == Replace, must be a default activity code ID)
    ActivityCodeIdReplacement string `json:"activityCodeIdReplacement"`

}

// String returns a JSON representation of the model
func (o *Shifttradeactivityrule) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeactivityrule) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeactivityrule

    if ShifttradeactivityruleMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeactivityruleMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCategory string `json:"activityCategory"`
        
        Action string `json:"action"`
        
        ActivityCodeIdReplacement string `json:"activityCodeIdReplacement"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

