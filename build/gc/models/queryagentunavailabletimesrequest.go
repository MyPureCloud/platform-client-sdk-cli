package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentunavailabletimesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentunavailabletimesrequestDud struct { 
    


    

}

// Queryagentunavailabletimesrequest
type Queryagentunavailabletimesrequest struct { 
    // StartDate - The earliest date to retrieve agent unavailability times. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The latest date to retrieve agent unavailability times. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Queryagentunavailabletimesrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentunavailabletimesrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryagentunavailabletimesrequest

    if QueryagentunavailabletimesrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryagentunavailabletimesrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

