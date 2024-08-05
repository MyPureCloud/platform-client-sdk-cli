package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftoffersrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftoffersrequestDud struct { 
    


    

}

// Alternativeshiftoffersrequest
type Alternativeshiftoffersrequest struct { 
    // Schedule - The existing schedule being used to find alternative shift offers
    Schedule Alternativeshiftschedulelookup `json:"schedule"`


    // QueryWeekDate - The start date for the week in this schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    QueryWeekDate time.Time `json:"queryWeekDate"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftoffersrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftoffersrequest) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftoffersrequest

    if AlternativeshiftoffersrequestMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftoffersrequestMarshalled = true

    return json.Marshal(&struct {
        
        Schedule Alternativeshiftschedulelookup `json:"schedule"`
        
        QueryWeekDate time.Time `json:"queryWeekDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

