package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulereferenceformurouteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulereferenceformurouteDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Buschedulereferenceformuroute
type Buschedulereferenceformuroute struct { 
    // Id - The ID of the schedule
    Id string `json:"id"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // BusinessUnit - The start week date for this schedule
    BusinessUnit Businessunitreference `json:"businessUnit"`


    

}

// String returns a JSON representation of the model
func (o *Buschedulereferenceformuroute) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulereferenceformuroute) MarshalJSON() ([]byte, error) {
    type Alias Buschedulereferenceformuroute

    if BuschedulereferenceformurouteMarshalled {
        return []byte("{}"), nil
    }
    BuschedulereferenceformurouteMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

