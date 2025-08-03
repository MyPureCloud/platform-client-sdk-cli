package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatestaffinggroupallocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatestaffinggroupallocationDud struct { 
    


    


    


    

}

// Createstaffinggroupallocation
type Createstaffinggroupallocation struct { 
    // StaffingGroupId - The ID of the staffing group used in this capacity plan
    StaffingGroupId string `json:"staffingGroupId"`


    // InitialShrinkagePercentage - The shrinkage percentage of the staffing group that can be used for all weeks in the planning period, in the scale of 0 - 100
    InitialShrinkagePercentage float64 `json:"initialShrinkagePercentage"`


    // InitialAttritionPercentage - The attrition percentage of the staffing group that can be used for all weeks in the planning period, in the scale of 0 - 100
    InitialAttritionPercentage float64 `json:"initialAttritionPercentage"`


    // StartingWeeklyFullTimeEquivalentCount - The weekly count of full time equivalent agents in the staffing group that can be used for the first week of the planning period
    StartingWeeklyFullTimeEquivalentCount float64 `json:"startingWeeklyFullTimeEquivalentCount"`

}

// String returns a JSON representation of the model
func (o *Createstaffinggroupallocation) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createstaffinggroupallocation) MarshalJSON() ([]byte, error) {
    type Alias Createstaffinggroupallocation

    if CreatestaffinggroupallocationMarshalled {
        return []byte("{}"), nil
    }
    CreatestaffinggroupallocationMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupId string `json:"staffingGroupId"`
        
        InitialShrinkagePercentage float64 `json:"initialShrinkagePercentage"`
        
        InitialAttritionPercentage float64 `json:"initialAttritionPercentage"`
        
        StartingWeeklyFullTimeEquivalentCount float64 `json:"startingWeeklyFullTimeEquivalentCount"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

