package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BugetcurrentagentschedulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BugetcurrentagentschedulerequestDud struct { 
    


    

}

// Bugetcurrentagentschedulerequest
type Bugetcurrentagentschedulerequest struct { 
    // StartDate - Start date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - End date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Bugetcurrentagentschedulerequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bugetcurrentagentschedulerequest) MarshalJSON() ([]byte, error) {
    type Alias Bugetcurrentagentschedulerequest

    if BugetcurrentagentschedulerequestMarshalled {
        return []byte("{}"), nil
    }
    BugetcurrentagentschedulerequestMarshalled = true

    return json.Marshal(&struct { 
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

