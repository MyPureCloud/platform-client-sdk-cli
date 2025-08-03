package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryplanninggrouptostaffinggroupsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryplanninggrouptostaffinggroupsrequestDud struct { 
    


    

}

// Queryplanninggrouptostaffinggroupsrequest
type Queryplanninggrouptostaffinggroupsrequest struct { 
    // PlanningGroupIds - The list of planning group IDs to request capacity group associations
    PlanningGroupIds []string `json:"planningGroupIds"`


    // StaffingGroupIds - The list of staffing group IDs to request capacity group associations
    StaffingGroupIds []string `json:"staffingGroupIds"`

}

// String returns a JSON representation of the model
func (o *Queryplanninggrouptostaffinggroupsrequest) String() string {
     o.PlanningGroupIds = []string{""} 
     o.StaffingGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryplanninggrouptostaffinggroupsrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryplanninggrouptostaffinggroupsrequest

    if QueryplanninggrouptostaffinggroupsrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryplanninggrouptostaffinggroupsrequestMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        
        StaffingGroupIds []string `json:"staffingGroupIds"`
        *Alias
    }{

        
        PlanningGroupIds: []string{""},
        


        
        StaffingGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

