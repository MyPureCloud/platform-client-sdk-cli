package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentschedulingpreferencesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentschedulingpreferencesrequestDud struct { 
    


    

}

// Queryagentschedulingpreferencesrequest
type Queryagentschedulingpreferencesrequest struct { 
    // StartDate - The earliest date to retrieve agent scheduling preferences. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The latest date to retrieve agent scheduling preferences. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Queryagentschedulingpreferencesrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentschedulingpreferencesrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryagentschedulingpreferencesrequest

    if QueryagentschedulingpreferencesrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryagentschedulingpreferencesrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

