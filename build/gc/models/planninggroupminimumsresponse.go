package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggroupminimumsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggroupminimumsresponseDud struct { 
    


    

}

// Planninggroupminimumsresponse
type Planninggroupminimumsresponse struct { 
    // PlanningGroup - The planning group to which the day of week minimum staff values apply
    PlanningGroup Planninggroupreference `json:"planningGroup"`


    // DayOfWeekMinimums - The list of day of week minimum staff values for this planning group
    DayOfWeekMinimums []Dayofweekminimums `json:"dayOfWeekMinimums"`

}

// String returns a JSON representation of the model
func (o *Planninggroupminimumsresponse) String() string {
    
     o.DayOfWeekMinimums = []Dayofweekminimums{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggroupminimumsresponse) MarshalJSON() ([]byte, error) {
    type Alias Planninggroupminimumsresponse

    if PlanninggroupminimumsresponseMarshalled {
        return []byte("{}"), nil
    }
    PlanninggroupminimumsresponseMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroup Planninggroupreference `json:"planningGroup"`
        
        DayOfWeekMinimums []Dayofweekminimums `json:"dayOfWeekMinimums"`
        *Alias
    }{

        


        
        DayOfWeekMinimums: []Dayofweekminimums{{}},
        

        Alias: (*Alias)(u),
    })
}

