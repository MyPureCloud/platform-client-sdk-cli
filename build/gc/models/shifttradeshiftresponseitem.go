package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeshiftresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeshiftresponseitemDud struct { 
    


    


    


    

}

// Shifttradeshiftresponseitem
type Shifttradeshiftresponseitem struct { 
    // Id - The ID of the shift
    Id string `json:"id"`


    // StartDate - The start date/time for the shift in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date/time for the shift in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // WeekDate - The start week date of the user shift in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`

}

// String returns a JSON representation of the model
func (o *Shifttradeshiftresponseitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeshiftresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeshiftresponseitem

    if ShifttradeshiftresponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeshiftresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        WeekDate time.Time `json:"weekDate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

