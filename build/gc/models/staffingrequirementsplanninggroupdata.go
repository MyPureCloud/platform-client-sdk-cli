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

}

// String returns a JSON representation of the model
func (o *Staffingrequirementsplanninggroupdata) String() string {
    
     o.StaffingRequirementsPerInterval = []float64{0.0} 

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
        *Alias
    }{

        


        
        StaffingRequirementsPerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

