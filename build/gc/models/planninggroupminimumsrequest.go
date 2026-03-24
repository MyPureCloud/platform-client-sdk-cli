package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggroupminimumsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggroupminimumsrequestDud struct { 
    


    

}

// Planninggroupminimumsrequest
type Planninggroupminimumsrequest struct { 
    // PlanningGroupId - The ID of the planning group to which the day of week minimum staff values apply
    PlanningGroupId string `json:"planningGroupId"`


    // DayOfWeekMinimums - The list of day of week minimum staff values for this planning group
    DayOfWeekMinimums []Dayofweekminimums `json:"dayOfWeekMinimums"`

}

// String returns a JSON representation of the model
func (o *Planninggroupminimumsrequest) String() string {
    
     o.DayOfWeekMinimums = []Dayofweekminimums{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggroupminimumsrequest) MarshalJSON() ([]byte, error) {
    type Alias Planninggroupminimumsrequest

    if PlanninggroupminimumsrequestMarshalled {
        return []byte("{}"), nil
    }
    PlanninggroupminimumsrequestMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        DayOfWeekMinimums []Dayofweekminimums `json:"dayOfWeekMinimums"`
        *Alias
    }{

        


        
        DayOfWeekMinimums: []Dayofweekminimums{{}},
        

        Alias: (*Alias)(u),
    })
}

