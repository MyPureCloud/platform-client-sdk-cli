package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggroupallocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggroupallocationDud struct { 
    


    


    


    


    


    

}

// Staffinggroupallocation
type Staffinggroupallocation struct { 
    // StaffingGroupId - The staffing group to which the result allocation belongs
    StaffingGroupId string `json:"staffingGroupId"`


    // ShrinkagePercentages - The weekly projected shrinkage percentage of staffing group, in the scale of 0 - 100
    ShrinkagePercentages []float64 `json:"shrinkagePercentages"`


    // AttritionPercentages - The weekly projected attrition percentage of the staffing group, in the scale of 0 - 100
    AttritionPercentages []float64 `json:"attritionPercentages"`


    // NewHiresFullTimeEquivalentCount - The weekly projected full time equivalent agents of new hire agents added to the staffing group
    NewHiresFullTimeEquivalentCount []float64 `json:"newHiresFullTimeEquivalentCount"`


    // StartingWeeklyFullTimeEquivalentCount - The weekly count of full time equivalent agents that can be used for the first week of the capacity plan
    StartingWeeklyFullTimeEquivalentCount float64 `json:"startingWeeklyFullTimeEquivalentCount"`


    // PlanningGroupIds - The IDs of the planning groups associated with this staffing group
    PlanningGroupIds []string `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Staffinggroupallocation) String() string {
    
     o.ShrinkagePercentages = []float64{0.0} 
     o.AttritionPercentages = []float64{0.0} 
     o.NewHiresFullTimeEquivalentCount = []float64{0.0} 
    
     o.PlanningGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggroupallocation) MarshalJSON() ([]byte, error) {
    type Alias Staffinggroupallocation

    if StaffinggroupallocationMarshalled {
        return []byte("{}"), nil
    }
    StaffinggroupallocationMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupId string `json:"staffingGroupId"`
        
        ShrinkagePercentages []float64 `json:"shrinkagePercentages"`
        
        AttritionPercentages []float64 `json:"attritionPercentages"`
        
        NewHiresFullTimeEquivalentCount []float64 `json:"newHiresFullTimeEquivalentCount"`
        
        StartingWeeklyFullTimeEquivalentCount float64 `json:"startingWeeklyFullTimeEquivalentCount"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        *Alias
    }{

        


        
        ShrinkagePercentages: []float64{0.0},
        


        
        AttritionPercentages: []float64{0.0},
        


        
        NewHiresFullTimeEquivalentCount: []float64{0.0},
        


        


        
        PlanningGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

