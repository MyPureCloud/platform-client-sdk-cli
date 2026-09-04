package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AllocationresultstemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AllocationresultstemplateDud struct { 
    


    


    

}

// Allocationresultstemplate
type Allocationresultstemplate struct { 
    // PlanningGroupId - The ID of the associated planning group
    PlanningGroupId string `json:"planningGroupId"`


    // AssignedAgentsPerInterval - Assigned agent allocation per interval used to generate the performance prediction
    AssignedAgentsPerInterval []float64 `json:"assignedAgentsPerInterval"`


    // HeadcountMultiplierPerInterval - Headcount multiplier per interval used to generate the performance prediction
    HeadcountMultiplierPerInterval []float64 `json:"headcountMultiplierPerInterval"`

}

// String returns a JSON representation of the model
func (o *Allocationresultstemplate) String() string {
    
     o.AssignedAgentsPerInterval = []float64{0.0} 
     o.HeadcountMultiplierPerInterval = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Allocationresultstemplate) MarshalJSON() ([]byte, error) {
    type Alias Allocationresultstemplate

    if AllocationresultstemplateMarshalled {
        return []byte("{}"), nil
    }
    AllocationresultstemplateMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        AssignedAgentsPerInterval []float64 `json:"assignedAgentsPerInterval"`
        
        HeadcountMultiplierPerInterval []float64 `json:"headcountMultiplierPerInterval"`
        *Alias
    }{

        


        
        AssignedAgentsPerInterval: []float64{0.0},
        


        
        HeadcountMultiplierPerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

