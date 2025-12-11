package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanmetricssummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanmetricssummaryDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Capacityplanmetricssummary
type Capacityplanmetricssummary struct { 
    // RequiredStaffFullTimeEquivalentCount - The total staff requirements for all planning groups in the capacity plan, aggregated by the selected time granularity
    RequiredStaffFullTimeEquivalentCount []float64 `json:"requiredStaffFullTimeEquivalentCount"`


    // PlannedFullTimeEquivalentCount - The planned full time equivalent for all staffing groups in the capacity plan, aggregated by the selected time granularity
    PlannedFullTimeEquivalentCount []float64 `json:"plannedFullTimeEquivalentCount"`


    // StaffingOverUnderFullTimeEquivalentCount - The difference between the summary of planning group required full time equivalent and planned full time equivalent, aggregated by the selected time granularity
    StaffingOverUnderFullTimeEquivalentCount []float64 `json:"staffingOverUnderFullTimeEquivalentCount"`


    // StartingFullTimeEquivalentCount - The total starting full time equivalent for all staffing groups in the capacity plan, aggregated by the selected time granularity
    StartingFullTimeEquivalentCount []float64 `json:"startingFullTimeEquivalentCount"`


    // AttritionFullTimeEquivalentCount - The sum of all staffing group attrition full time equivalent in the capacity plan, aggregated by the selected time granularity
    AttritionFullTimeEquivalentCount []float64 `json:"attritionFullTimeEquivalentCount"`


    // AttritionPercentage - The total attrition percentage of staffing groups in the capacity plan in the scale of 0.0-100.0, aggregated by the selected time granularity
    AttritionPercentage []float64 `json:"attritionPercentage"`


    // NewHireFullTimeEquivalentCount - The sum of all staffing group new hire full time equivalent in the capacity plan, aggregated by the selected time granularity
    NewHireFullTimeEquivalentCount []float64 `json:"newHireFullTimeEquivalentCount"`


    // TransfersFullTimeEquivalentCount - The sum of all staffing group projected transfers of agents into or out of this capacity plan, aggregated by the selected time granularity
    TransfersFullTimeEquivalentCount []float64 `json:"transfersFullTimeEquivalentCount"`


    // ExtraTimeUnderTimeFullTimeEquivalentCount - The sum of all staffing group extra or under time full time equivalent count in the capacity plan, aggregated by the selected time granularity
    ExtraTimeUnderTimeFullTimeEquivalentCount []float64 `json:"extraTimeUnderTimeFullTimeEquivalentCount"`


    // ShrinkageFullTimeEquivalentCount - The sum of all staffing groups shrinkage full time equivalent, aggregated by the selected time granularity
    ShrinkageFullTimeEquivalentCount []float64 `json:"shrinkageFullTimeEquivalentCount"`


    // ShrinkagePercentage - The total shrinkage percentage of all staffing groups in the capacity plan in the scale of 0.0-100.0, aggregated by the selected time granularity
    ShrinkagePercentage []float64 `json:"shrinkagePercentage"`


    // EndOfMonthPlannedFullTimeEquivalentCount - The total sum of all staffing group end of month planned full time equivalent for this capacity plan, aggregated by the selected time granularity
    EndOfMonthPlannedFullTimeEquivalentCount []float64 `json:"endOfMonthPlannedFullTimeEquivalentCount"`


    // NetFullTimeEquivalentCount - The sum of all staffing groups net full time equivalent in the capacity plan, aggregated by the selected time granularity
    NetFullTimeEquivalentCount []float64 `json:"netFullTimeEquivalentCount"`

}

// String returns a JSON representation of the model
func (o *Capacityplanmetricssummary) String() string {
     o.RequiredStaffFullTimeEquivalentCount = []float64{0.0} 
     o.PlannedFullTimeEquivalentCount = []float64{0.0} 
     o.StaffingOverUnderFullTimeEquivalentCount = []float64{0.0} 
     o.StartingFullTimeEquivalentCount = []float64{0.0} 
     o.AttritionFullTimeEquivalentCount = []float64{0.0} 
     o.AttritionPercentage = []float64{0.0} 
     o.NewHireFullTimeEquivalentCount = []float64{0.0} 
     o.TransfersFullTimeEquivalentCount = []float64{0.0} 
     o.ExtraTimeUnderTimeFullTimeEquivalentCount = []float64{0.0} 
     o.ShrinkageFullTimeEquivalentCount = []float64{0.0} 
     o.ShrinkagePercentage = []float64{0.0} 
     o.EndOfMonthPlannedFullTimeEquivalentCount = []float64{0.0} 
     o.NetFullTimeEquivalentCount = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanmetricssummary) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanmetricssummary

    if CapacityplanmetricssummaryMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanmetricssummaryMarshalled = true

    return json.Marshal(&struct {
        
        RequiredStaffFullTimeEquivalentCount []float64 `json:"requiredStaffFullTimeEquivalentCount"`
        
        PlannedFullTimeEquivalentCount []float64 `json:"plannedFullTimeEquivalentCount"`
        
        StaffingOverUnderFullTimeEquivalentCount []float64 `json:"staffingOverUnderFullTimeEquivalentCount"`
        
        StartingFullTimeEquivalentCount []float64 `json:"startingFullTimeEquivalentCount"`
        
        AttritionFullTimeEquivalentCount []float64 `json:"attritionFullTimeEquivalentCount"`
        
        AttritionPercentage []float64 `json:"attritionPercentage"`
        
        NewHireFullTimeEquivalentCount []float64 `json:"newHireFullTimeEquivalentCount"`
        
        TransfersFullTimeEquivalentCount []float64 `json:"transfersFullTimeEquivalentCount"`
        
        ExtraTimeUnderTimeFullTimeEquivalentCount []float64 `json:"extraTimeUnderTimeFullTimeEquivalentCount"`
        
        ShrinkageFullTimeEquivalentCount []float64 `json:"shrinkageFullTimeEquivalentCount"`
        
        ShrinkagePercentage []float64 `json:"shrinkagePercentage"`
        
        EndOfMonthPlannedFullTimeEquivalentCount []float64 `json:"endOfMonthPlannedFullTimeEquivalentCount"`
        
        NetFullTimeEquivalentCount []float64 `json:"netFullTimeEquivalentCount"`
        *Alias
    }{

        
        RequiredStaffFullTimeEquivalentCount: []float64{0.0},
        


        
        PlannedFullTimeEquivalentCount: []float64{0.0},
        


        
        StaffingOverUnderFullTimeEquivalentCount: []float64{0.0},
        


        
        StartingFullTimeEquivalentCount: []float64{0.0},
        


        
        AttritionFullTimeEquivalentCount: []float64{0.0},
        


        
        AttritionPercentage: []float64{0.0},
        


        
        NewHireFullTimeEquivalentCount: []float64{0.0},
        


        
        TransfersFullTimeEquivalentCount: []float64{0.0},
        


        
        ExtraTimeUnderTimeFullTimeEquivalentCount: []float64{0.0},
        


        
        ShrinkageFullTimeEquivalentCount: []float64{0.0},
        


        
        ShrinkagePercentage: []float64{0.0},
        


        
        EndOfMonthPlannedFullTimeEquivalentCount: []float64{0.0},
        


        
        NetFullTimeEquivalentCount: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

