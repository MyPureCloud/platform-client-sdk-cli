package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InitiatingshiftrequestitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InitiatingshiftrequestitemDud struct { 
    


    


    

}

// Initiatingshiftrequestitem
type Initiatingshiftrequestitem struct { 
    // Id - The ID of the shift that the initiating user wants to give up in this trade
    Id string `json:"id"`


    // Schedule - A reference to the associated schedule to which this initiating shift belongs
    Schedule Buschedulereference `json:"schedule"`


    // WeekDate - The start week date of this schedule in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`

}

// String returns a JSON representation of the model
func (o *Initiatingshiftrequestitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Initiatingshiftrequestitem) MarshalJSON() ([]byte, error) {
    type Alias Initiatingshiftrequestitem

    if InitiatingshiftrequestitemMarshalled {
        return []byte("{}"), nil
    }
    InitiatingshiftrequestitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Schedule Buschedulereference `json:"schedule"`
        
        WeekDate time.Time `json:"weekDate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

