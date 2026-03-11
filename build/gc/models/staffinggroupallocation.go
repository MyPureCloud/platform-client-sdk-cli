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


    // BaseStartingFullTimeEquivalentCount - The weekly calculated starting full time equivalent count
    BaseStartingFullTimeEquivalentCount []float64 `json:"baseStartingFullTimeEquivalentCount"`


    // AttritionFullTimeEquivalentCount - The weekly projected attrition full time equivalent count
    AttritionFullTimeEquivalentCount []float64 `json:"attritionFullTimeEquivalentCount"`


    // StaffingGroupPlannedFullTimeEquivalentCount - The weekly calculated staffing group full time equivalent count
    StaffingGroupPlannedFullTimeEquivalentCount []float64 `json:"staffingGroupPlannedFullTimeEquivalentCount"`


    // EndOfMonthPlannedFullTimeEquivalentCount - The end of month planned full time equivalent count of this staffing group
    EndOfMonthPlannedFullTimeEquivalentCount []float64 `json:"endOfMonthPlannedFullTimeEquivalentCount"`


    // ShrinkageFullTimeEquivalentCount - The weekly projected shrinkage full time equivalent count of this staffing group
    ShrinkageFullTimeEquivalentCount []float64 `json:"shrinkageFullTimeEquivalentCount"`


    // NetFullTimeEquivalentCount - The weekly net full time equivalent count of this staffing group
    NetFullTimeEquivalentCount []float64 `json:"netFullTimeEquivalentCount"`


    // ExtraTimeUnderTimeFullTimeEquivalentCount - The weekly projected extra or under full time equivalent to the staffing group
    ExtraTimeUnderTimeFullTimeEquivalentCount []float64 `json:"extraTimeUnderTimeFullTimeEquivalentCount"`


    // TransfersFullTimeEquivalentCount - The weekly projected full time equivalent transfers of agents into or out of this staffing group
    TransfersFullTimeEquivalentCount []float64 `json:"transfersFullTimeEquivalentCount"`

}

// String returns a JSON representation of the model
func (o *Staffinggroupallocation) String() string {
    
     o.ShrinkagePercentages = []float64{0.0} 
     o.AttritionPercentages = []float64{0.0} 
     o.NewHiresFullTimeEquivalentCount = []float64{0.0} 
    
     o.PlanningGroupIds = []string{""} 
     o.BaseStartingFullTimeEquivalentCount = []float64{0.0} 
     o.AttritionFullTimeEquivalentCount = []float64{0.0} 
     o.StaffingGroupPlannedFullTimeEquivalentCount = []float64{0.0} 
     o.EndOfMonthPlannedFullTimeEquivalentCount = []float64{0.0} 
     o.ShrinkageFullTimeEquivalentCount = []float64{0.0} 
     o.NetFullTimeEquivalentCount = []float64{0.0} 
     o.ExtraTimeUnderTimeFullTimeEquivalentCount = []float64{0.0} 
     o.TransfersFullTimeEquivalentCount = []float64{0.0} 

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
        
        BaseStartingFullTimeEquivalentCount []float64 `json:"baseStartingFullTimeEquivalentCount"`
        
        AttritionFullTimeEquivalentCount []float64 `json:"attritionFullTimeEquivalentCount"`
        
        StaffingGroupPlannedFullTimeEquivalentCount []float64 `json:"staffingGroupPlannedFullTimeEquivalentCount"`
        
        EndOfMonthPlannedFullTimeEquivalentCount []float64 `json:"endOfMonthPlannedFullTimeEquivalentCount"`
        
        ShrinkageFullTimeEquivalentCount []float64 `json:"shrinkageFullTimeEquivalentCount"`
        
        NetFullTimeEquivalentCount []float64 `json:"netFullTimeEquivalentCount"`
        
        ExtraTimeUnderTimeFullTimeEquivalentCount []float64 `json:"extraTimeUnderTimeFullTimeEquivalentCount"`
        
        TransfersFullTimeEquivalentCount []float64 `json:"transfersFullTimeEquivalentCount"`
        *Alias
    }{

        


        
        ShrinkagePercentages: []float64{0.0},
        


        
        AttritionPercentages: []float64{0.0},
        


        
        NewHiresFullTimeEquivalentCount: []float64{0.0},
        


        


        
        PlanningGroupIds: []string{""},
        


        
        BaseStartingFullTimeEquivalentCount: []float64{0.0},
        


        
        AttritionFullTimeEquivalentCount: []float64{0.0},
        


        
        StaffingGroupPlannedFullTimeEquivalentCount: []float64{0.0},
        


        
        EndOfMonthPlannedFullTimeEquivalentCount: []float64{0.0},
        


        
        ShrinkageFullTimeEquivalentCount: []float64{0.0},
        


        
        NetFullTimeEquivalentCount: []float64{0.0},
        


        
        ExtraTimeUnderTimeFullTimeEquivalentCount: []float64{0.0},
        


        
        TransfersFullTimeEquivalentCount: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

