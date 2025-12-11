package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanningplanninggroupallocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanningplanninggroupallocationDud struct { 
    


    


    


    


    

}

// Capacityplanningplanninggroupallocation
type Capacityplanningplanninggroupallocation struct { 
    // PlanningGroupId - The planning group ID to which the capacity planning allocations apply
    PlanningGroupId string `json:"planningGroupId"`


    // RequiredStaffFullTimeEquivalentCount - The weekly required staff to this planning group
    RequiredStaffFullTimeEquivalentCount []float64 `json:"requiredStaffFullTimeEquivalentCount"`


    // StaffingGroupFullTimeEquivalentContributions - The weekly planned full time equivalent contributions from associated staffing groups
    StaffingGroupFullTimeEquivalentContributions []Staffinggroupfulltimeequivalentcontribution `json:"staffingGroupFullTimeEquivalentContributions"`


    // TotalPlannedFullTimeEquivalentCount - The total weekly full time equivalent planned for this planning group, based on the associated staffing groups
    TotalPlannedFullTimeEquivalentCount []float64 `json:"totalPlannedFullTimeEquivalentCount"`


    // OverUnderFullTimeEquivalentCount - The weekly difference between the total planned full time equivalent and the required staff
    OverUnderFullTimeEquivalentCount []float64 `json:"overUnderFullTimeEquivalentCount"`

}

// String returns a JSON representation of the model
func (o *Capacityplanningplanninggroupallocation) String() string {
    
     o.RequiredStaffFullTimeEquivalentCount = []float64{0.0} 
     o.StaffingGroupFullTimeEquivalentContributions = []Staffinggroupfulltimeequivalentcontribution{{}} 
     o.TotalPlannedFullTimeEquivalentCount = []float64{0.0} 
     o.OverUnderFullTimeEquivalentCount = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanningplanninggroupallocation) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanningplanninggroupallocation

    if CapacityplanningplanninggroupallocationMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanningplanninggroupallocationMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        RequiredStaffFullTimeEquivalentCount []float64 `json:"requiredStaffFullTimeEquivalentCount"`
        
        StaffingGroupFullTimeEquivalentContributions []Staffinggroupfulltimeequivalentcontribution `json:"staffingGroupFullTimeEquivalentContributions"`
        
        TotalPlannedFullTimeEquivalentCount []float64 `json:"totalPlannedFullTimeEquivalentCount"`
        
        OverUnderFullTimeEquivalentCount []float64 `json:"overUnderFullTimeEquivalentCount"`
        *Alias
    }{

        


        
        RequiredStaffFullTimeEquivalentCount: []float64{0.0},
        


        
        StaffingGroupFullTimeEquivalentContributions: []Staffinggroupfulltimeequivalentcontribution{{}},
        


        
        TotalPlannedFullTimeEquivalentCount: []float64{0.0},
        


        
        OverUnderFullTimeEquivalentCount: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

