package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffingrequirementsplanninggroupdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffingrequirementsplanninggroupdataDud struct { 
    


    


    


    

}

// Staffingrequirementsplanninggroupdata
type Staffingrequirementsplanninggroupdata struct { 
    // PlanningGroupId - The ID of the planning group to which this data applies
    PlanningGroupId string `json:"planningGroupId"`


    // StaffingRequirementsPerInterval - Staffing requirements per interval for this week forecast
    StaffingRequirementsPerInterval []float64 `json:"staffingRequirementsPerInterval"`


    // MinimumStaffPerInterval - Minimum Staff per interval for this week forecast
    MinimumStaffPerInterval []float64 `json:"minimumStaffPerInterval"`


    // EffectiveStaffPerInterval - Effective Staff per interval for this week forecast
    EffectiveStaffPerInterval []float64 `json:"effectiveStaffPerInterval"`

}

// String returns a JSON representation of the model
func (o *Staffingrequirementsplanninggroupdata) String() string {
    
     o.StaffingRequirementsPerInterval = []float64{0.0} 
     o.MinimumStaffPerInterval = []float64{0.0} 
     o.EffectiveStaffPerInterval = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffingrequirementsplanninggroupdata) MarshalJSON() ([]byte, error) {
    type Alias Staffingrequirementsplanninggroupdata

    if StaffingrequirementsplanninggroupdataMarshalled {
        return []byte("{}"), nil
    }
    StaffingrequirementsplanninggroupdataMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        StaffingRequirementsPerInterval []float64 `json:"staffingRequirementsPerInterval"`
        
        MinimumStaffPerInterval []float64 `json:"minimumStaffPerInterval"`
        
        EffectiveStaffPerInterval []float64 `json:"effectiveStaffPerInterval"`
        *Alias
    }{

        


        
        StaffingRequirementsPerInterval: []float64{0.0},
        


        
        MinimumStaffPerInterval: []float64{0.0},
        


        
        EffectiveStaffPerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

