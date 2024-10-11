package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanoverriderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanoverriderequestDud struct { 
    


    


    


    

}

// Workplanoverriderequest
type Workplanoverriderequest struct { 
    // Action - The action to perform on work plan override, defaults to add
    Action string `json:"action"`


    // StartDate - The start date in yyyy-MM-dd format for the updated work plan. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // WeekCount - The week count of the updated work plan, required if action is Add or Update
    WeekCount int `json:"weekCount"`


    // WorkPlanId - The updated work plan id
    WorkPlanId string `json:"workPlanId"`

}

// String returns a JSON representation of the model
func (o *Workplanoverriderequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanoverriderequest) MarshalJSON() ([]byte, error) {
    type Alias Workplanoverriderequest

    if WorkplanoverriderequestMarshalled {
        return []byte("{}"), nil
    }
    WorkplanoverriderequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        StartDate time.Time `json:"startDate"`
        
        WeekCount int `json:"weekCount"`
        
        WorkPlanId string `json:"workPlanId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

