package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastplanninggroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastplanninggroupDud struct { 
    


    


    

}

// Capacityplanforecastplanninggroup
type Capacityplanforecastplanninggroup struct { 
    // PlanningGroupId - The ID of the planning group to which the forecast applies
    PlanningGroupId string `json:"planningGroupId"`


    // Actual - Absolute forecast values for offered counts and average handle time for this planning group. Exactly one of actual or distribution must be provided
    Actual Capacityplanforecastactualvalues `json:"actual"`


    // Distribution - Hourly forecast distribution values for offered volume and average handle time for this planning group. Exactly one of actual or distribution must be provided
    Distribution Capacityplanforecastdistributionvalues `json:"distribution"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastplanninggroup) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastplanninggroup) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastplanninggroup

    if CapacityplanforecastplanninggroupMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastplanninggroupMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        Actual Capacityplanforecastactualvalues `json:"actual"`
        
        Distribution Capacityplanforecastdistributionvalues `json:"distribution"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

