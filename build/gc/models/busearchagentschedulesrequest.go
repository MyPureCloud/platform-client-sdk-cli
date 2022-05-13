package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusearchagentschedulesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusearchagentschedulesrequestDud struct { 
    


    


    

}

// Busearchagentschedulesrequest
type Busearchagentschedulesrequest struct { 
    // StartDate - Start date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - End date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // UserIds - IDs of the users for whose schedules to search
    UserIds []string `json:"userIds"`

}

// String returns a JSON representation of the model
func (o *Busearchagentschedulesrequest) String() string {
    
    
     o.UserIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Busearchagentschedulesrequest) MarshalJSON() ([]byte, error) {
    type Alias Busearchagentschedulesrequest

    if BusearchagentschedulesrequestMarshalled {
        return []byte("{}"), nil
    }
    BusearchagentschedulesrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        UserIds []string `json:"userIds"`
        *Alias
    }{

        


        


        
        UserIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

