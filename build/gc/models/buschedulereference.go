package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulereferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Buschedulereference
type Buschedulereference struct { 
    // Id - The ID of the schedule
    Id string `json:"id"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    

}

// String returns a JSON representation of the model
func (o *Buschedulereference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulereference) MarshalJSON() ([]byte, error) {
    type Alias Buschedulereference

    if BuschedulereferenceMarshalled {
        return []byte("{}"), nil
    }
    BuschedulereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

