package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentscheduleshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentscheduleshiftDud struct { 
    


    StartDate time.Time `json:"startDate"`


    LengthMinutes int `json:"lengthMinutes"`


    


    


    Schedule Buschedulereference `json:"schedule"`


    


    

}

// Buagentscheduleshift
type Buagentscheduleshift struct { 
    // Id - The ID of the shift
    Id string `json:"id"`


    


    


    // Activities - The activities associated with this shift
    Activities []Buagentscheduleactivity `json:"activities"`


    // ManuallyEdited - Whether this shift was manually edited. This is only set by clients and is used for rescheduling
    ManuallyEdited bool `json:"manuallyEdited"`


    


    // WorkPlanId - The ID of the work plan for which the work plan shift emanates from
    WorkPlanId string `json:"workPlanId"`


    // WorkPlanShiftId - The ID of the work plan shift that was used in schedule generation
    WorkPlanShiftId string `json:"workPlanShiftId"`

}

// String returns a JSON representation of the model
func (o *Buagentscheduleshift) String() string {
    
     o.Activities = []Buagentscheduleactivity{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentscheduleshift) MarshalJSON() ([]byte, error) {
    type Alias Buagentscheduleshift

    if BuagentscheduleshiftMarshalled {
        return []byte("{}"), nil
    }
    BuagentscheduleshiftMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Activities []Buagentscheduleactivity `json:"activities"`
        
        ManuallyEdited bool `json:"manuallyEdited"`
        
        WorkPlanId string `json:"workPlanId"`
        
        WorkPlanShiftId string `json:"workPlanShiftId"`
        *Alias
    }{

        


        


        


        
        Activities: []Buagentscheduleactivity{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

