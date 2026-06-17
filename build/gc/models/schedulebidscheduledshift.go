package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidscheduledshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidscheduledshiftDud struct { 
    


    


    


    


    

}

// Schedulebidscheduledshift
type Schedulebidscheduledshift struct { 
    // WorkPlanShiftId - The ID of the work plan shift that was used in schedule generation
    WorkPlanShiftId string `json:"workPlanShiftId"`


    // WorkPlanId - The ID of the work plan from which the shift comes
    WorkPlanId string `json:"workPlanId"`


    // StartDate - The start date of the scheduled shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the shift in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Activities - The activities associated with this shift
    Activities []Schedulebidscheduledactivity `json:"activities"`

}

// String returns a JSON representation of the model
func (o *Schedulebidscheduledshift) String() string {
    
    
    
    
     o.Activities = []Schedulebidscheduledactivity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidscheduledshift) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidscheduledshift

    if SchedulebidscheduledshiftMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidscheduledshiftMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanShiftId string `json:"workPlanShiftId"`
        
        WorkPlanId string `json:"workPlanId"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Activities []Schedulebidscheduledactivity `json:"activities"`
        *Alias
    }{

        


        


        


        


        
        Activities: []Schedulebidscheduledactivity{{}},
        

        Alias: (*Alias)(u),
    })
}

