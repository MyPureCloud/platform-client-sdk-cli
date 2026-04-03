package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulereferencewithbusinessunitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulereferencewithbusinessunitDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulereferencewithbusinessunit
type Schedulereferencewithbusinessunit struct { 
    // Id - The unique identifier of the schedule
    Id string `json:"id"`


    // WeekDate - The start date for this schedule in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // BusinessUnit - The reference of the associated business unit
    BusinessUnit Businessunitreference `json:"businessUnit"`


    

}

// String returns a JSON representation of the model
func (o *Schedulereferencewithbusinessunit) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulereferencewithbusinessunit) MarshalJSON() ([]byte, error) {
    type Alias Schedulereferencewithbusinessunit

    if SchedulereferencewithbusinessunitMarshalled {
        return []byte("{}"), nil
    }
    SchedulereferencewithbusinessunitMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

