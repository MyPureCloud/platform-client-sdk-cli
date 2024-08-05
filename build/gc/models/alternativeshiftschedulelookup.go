package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftschedulelookupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftschedulelookupDud struct { 
    


    

}

// Alternativeshiftschedulelookup
type Alternativeshiftschedulelookup struct { 
    // Id - The unique identifier of the schedule
    Id string `json:"id"`


    // WeekDate - The start date for this schedule in yyyy-MM-dd. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftschedulelookup) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftschedulelookup) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftschedulelookup

    if AlternativeshiftschedulelookupMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftschedulelookupMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

