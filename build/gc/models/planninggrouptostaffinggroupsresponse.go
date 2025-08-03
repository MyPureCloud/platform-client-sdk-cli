package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggrouptostaffinggroupsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggrouptostaffinggroupsresponseDud struct { 
    


    

}

// Planninggrouptostaffinggroupsresponse
type Planninggrouptostaffinggroupsresponse struct { 
    // PlanningGroup - The ID of the planning group
    PlanningGroup Planninggroupreference `json:"planningGroup"`


    // StaffingGroups - The list of staffing groups that are associated with the planning group
    StaffingGroups []Staffinggroupreference `json:"staffingGroups"`

}

// String returns a JSON representation of the model
func (o *Planninggrouptostaffinggroupsresponse) String() string {
    
     o.StaffingGroups = []Staffinggroupreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggrouptostaffinggroupsresponse) MarshalJSON() ([]byte, error) {
    type Alias Planninggrouptostaffinggroupsresponse

    if PlanninggrouptostaffinggroupsresponseMarshalled {
        return []byte("{}"), nil
    }
    PlanninggrouptostaffinggroupsresponseMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroup Planninggroupreference `json:"planningGroup"`
        
        StaffingGroups []Staffinggroupreference `json:"staffingGroups"`
        *Alias
    }{

        


        
        StaffingGroups: []Staffinggroupreference{{}},
        

        Alias: (*Alias)(u),
    })
}

