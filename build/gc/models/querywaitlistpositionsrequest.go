package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerywaitlistpositionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerywaitlistpositionsrequestDud struct { 
    

}

// Querywaitlistpositionsrequest
type Querywaitlistpositionsrequest struct { 
    // TimeOffRequests - The list of the time off request ids for which to fetch the daily waitlist positions
    TimeOffRequests []Usertimeoffrequestreference `json:"timeOffRequests"`

}

// String returns a JSON representation of the model
func (o *Querywaitlistpositionsrequest) String() string {
    
    
     o.TimeOffRequests = []Usertimeoffrequestreference{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querywaitlistpositionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Querywaitlistpositionsrequest

    if QuerywaitlistpositionsrequestMarshalled {
        return []byte("{}"), nil
    }
    QuerywaitlistpositionsrequestMarshalled = true

    return json.Marshal(&struct { 
        TimeOffRequests []Usertimeoffrequestreference `json:"timeOffRequests"`
        
        *Alias
    }{
        

        
        TimeOffRequests: []Usertimeoffrequestreference{{}},
        

        
        Alias: (*Alias)(u),
    })
}

