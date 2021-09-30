package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingruleDud struct { 
    


    


    

}

// Routingrule
type Routingrule struct { 
    // Operator - matching operator.  MEETS_THRESHOLD matches any agent with a score at or above the rule's threshold.  ANY matches all specified agents, regardless of score.
    Operator string `json:"operator"`


    // Threshold - threshold required for routing attempt (generally an agent score).  may be null for operator ANY.
    Threshold int `json:"threshold"`


    // WaitSeconds - seconds to wait in this rule before moving to the next
    WaitSeconds float64 `json:"waitSeconds"`

}

// String returns a JSON representation of the model
func (o *Routingrule) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingrule) MarshalJSON() ([]byte, error) {
    type Alias Routingrule

    if RoutingruleMarshalled {
        return []byte("{}"), nil
    }
    RoutingruleMarshalled = true

    return json.Marshal(&struct { 
        Operator string `json:"operator"`
        
        Threshold int `json:"threshold"`
        
        WaitSeconds float64 `json:"waitSeconds"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

