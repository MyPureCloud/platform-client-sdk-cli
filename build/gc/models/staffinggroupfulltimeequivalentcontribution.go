package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggroupfulltimeequivalentcontributionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggroupfulltimeequivalentcontributionDud struct { 
    


    

}

// Staffinggroupfulltimeequivalentcontribution
type Staffinggroupfulltimeequivalentcontribution struct { 
    // StaffingGroupId - The staffing group ID receiving full time equivalent agents from the planning group
    StaffingGroupId string `json:"staffingGroupId"`


    // PlannedFullTimeEquivalentContribution - The weekly planned full time contribution from the planning group to this staffing group
    PlannedFullTimeEquivalentContribution []float64 `json:"plannedFullTimeEquivalentContribution"`

}

// String returns a JSON representation of the model
func (o *Staffinggroupfulltimeequivalentcontribution) String() string {
    
     o.PlannedFullTimeEquivalentContribution = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggroupfulltimeequivalentcontribution) MarshalJSON() ([]byte, error) {
    type Alias Staffinggroupfulltimeequivalentcontribution

    if StaffinggroupfulltimeequivalentcontributionMarshalled {
        return []byte("{}"), nil
    }
    StaffinggroupfulltimeequivalentcontributionMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupId string `json:"staffingGroupId"`
        
        PlannedFullTimeEquivalentContribution []float64 `json:"plannedFullTimeEquivalentContribution"`
        *Alias
    }{

        


        
        PlannedFullTimeEquivalentContribution: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

