package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanoverrideDud struct { 
    


    


    

}

// Workplanoverride
type Workplanoverride struct { 
    // StartDate - The start date in yyyy-MM-dd format of the work plan override. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // WeekCount - Number of weeks for the work plan override
    WeekCount int `json:"weekCount"`


    // WorkPlan - The work plan reference associated with this override
    WorkPlan Workplanreference `json:"workPlan"`

}

// String returns a JSON representation of the model
func (o *Workplanoverride) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanoverride) MarshalJSON() ([]byte, error) {
    type Alias Workplanoverride

    if WorkplanoverrideMarshalled {
        return []byte("{}"), nil
    }
    WorkplanoverrideMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        WeekCount int `json:"weekCount"`
        
        WorkPlan Workplanreference `json:"workPlan"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

