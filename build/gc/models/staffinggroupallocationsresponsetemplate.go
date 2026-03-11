package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggroupallocationsresponsetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggroupallocationsresponsetemplateDud struct { 
    


    


    


    

}

// Staffinggroupallocationsresponsetemplate
type Staffinggroupallocationsresponsetemplate struct { 
    // StaffingGroupAllocations - List of staffing group allocations
    StaffingGroupAllocations []Staffinggroupallocation `json:"staffingGroupAllocations"`


    // Months - The list of months covered by this capacity plan, formatted as yyyy-MM
    Months []string `json:"months"`


    // PlanningGroupAllocations - The planning group allocations
    PlanningGroupAllocations []Capacityplanningplanninggroupallocation `json:"planningGroupAllocations"`


    // CapacityPlanMetricsSummary - The total summary of staffing allocation metrics for this capacity plan, for the selected granularity
    CapacityPlanMetricsSummary Capacityplanmetricssummary `json:"capacityPlanMetricsSummary"`

}

// String returns a JSON representation of the model
func (o *Staffinggroupallocationsresponsetemplate) String() string {
     o.StaffingGroupAllocations = []Staffinggroupallocation{{}} 
     o.Months = []string{""} 
     o.PlanningGroupAllocations = []Capacityplanningplanninggroupallocation{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggroupallocationsresponsetemplate) MarshalJSON() ([]byte, error) {
    type Alias Staffinggroupallocationsresponsetemplate

    if StaffinggroupallocationsresponsetemplateMarshalled {
        return []byte("{}"), nil
    }
    StaffinggroupallocationsresponsetemplateMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupAllocations []Staffinggroupallocation `json:"staffingGroupAllocations"`
        
        Months []string `json:"months"`
        
        PlanningGroupAllocations []Capacityplanningplanninggroupallocation `json:"planningGroupAllocations"`
        
        CapacityPlanMetricsSummary Capacityplanmetricssummary `json:"capacityPlanMetricsSummary"`
        *Alias
    }{

        
        StaffingGroupAllocations: []Staffinggroupallocation{{}},
        


        
        Months: []string{""},
        


        
        PlanningGroupAllocations: []Capacityplanningplanninggroupallocation{{}},
        


        

        Alias: (*Alias)(u),
    })
}

