package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocaldaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocaldaterangeDud struct { 
    


    

}

// Localdaterange
type Localdaterange struct { 
    // StartDate - The inclusive start of a date range in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // EndDate - The inclusive end of a date range in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Localdaterange) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localdaterange) MarshalJSON() ([]byte, error) {
    type Alias Localdaterange

    if LocaldaterangeMarshalled {
        return []byte("{}"), nil
    }
    LocaldaterangeMarshalled = true

    return json.Marshal(&struct { 
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

