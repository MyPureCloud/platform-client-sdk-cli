package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScheduleseterrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScheduleseterrorDud struct { 
    


    


    

}

// Scheduleseterror
type Scheduleseterror struct { 
    // ErrorCode - Error code that indicates why schedule set optimization failed. At least one of workPlans or workPlanRotations is set if there is an error during optimization
    ErrorCode string `json:"errorCode"`


    // WorkPlans - Work plans involved in the optimization failure
    WorkPlans []Workplanreference `json:"workPlans"`


    // WorkPlanRotations - Work plan rotations involved in the optimization failure
    WorkPlanRotations []Workplanrotationreference `json:"workPlanRotations"`

}

// String returns a JSON representation of the model
func (o *Scheduleseterror) String() string {
    
     o.WorkPlans = []Workplanreference{{}} 
     o.WorkPlanRotations = []Workplanrotationreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scheduleseterror) MarshalJSON() ([]byte, error) {
    type Alias Scheduleseterror

    if ScheduleseterrorMarshalled {
        return []byte("{}"), nil
    }
    ScheduleseterrorMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        
        WorkPlans []Workplanreference `json:"workPlans"`
        
        WorkPlanRotations []Workplanrotationreference `json:"workPlanRotations"`
        *Alias
    }{

        


        
        WorkPlans: []Workplanreference{{}},
        


        
        WorkPlanRotations: []Workplanrotationreference{{}},
        

        Alias: (*Alias)(u),
    })
}

