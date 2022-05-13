package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmschedulereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmschedulereferenceDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Wfmschedulereference
type Wfmschedulereference struct { 
    // Id - The ID of the WFM schedule
    Id string `json:"id"`


    // BusinessUnit - A reference to a Workforce Management Business Unit
    BusinessUnit Wfmbusinessunitreference `json:"businessUnit"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    

}

// String returns a JSON representation of the model
func (o *Wfmschedulereference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmschedulereference) MarshalJSON() ([]byte, error) {
    type Alias Wfmschedulereference

    if WfmschedulereferenceMarshalled {
        return []byte("{}"), nil
    }
    WfmschedulereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        BusinessUnit Wfmbusinessunitreference `json:"businessUnit"`
        
        WeekDate time.Time `json:"weekDate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

