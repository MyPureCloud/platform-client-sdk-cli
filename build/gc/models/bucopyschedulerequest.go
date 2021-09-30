package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BucopyschedulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BucopyschedulerequestDud struct { 
    


    

}

// Bucopyschedulerequest
type Bucopyschedulerequest struct { 
    // Description - The description for the new schedule
    Description string `json:"description"`


    // WeekDate - The start weekDate for the new copy of the schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`

}

// String returns a JSON representation of the model
func (o *Bucopyschedulerequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bucopyschedulerequest) MarshalJSON() ([]byte, error) {
    type Alias Bucopyschedulerequest

    if BucopyschedulerequestMarshalled {
        return []byte("{}"), nil
    }
    BucopyschedulerequestMarshalled = true

    return json.Marshal(&struct { 
        Description string `json:"description"`
        
        WeekDate time.Time `json:"weekDate"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

