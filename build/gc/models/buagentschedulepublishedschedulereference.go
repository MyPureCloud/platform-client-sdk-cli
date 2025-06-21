package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulepublishedschedulereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulepublishedschedulereferenceDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Buagentschedulepublishedschedulereference
type Buagentschedulepublishedschedulereference struct { 
    // Id - The ID of the schedule
    Id string `json:"id"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks encompassed by the schedule
    WeekCount int `json:"weekCount"`


    

}

// String returns a JSON representation of the model
func (o *Buagentschedulepublishedschedulereference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulepublishedschedulereference) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulepublishedschedulereference

    if BuagentschedulepublishedschedulereferenceMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulepublishedschedulereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

