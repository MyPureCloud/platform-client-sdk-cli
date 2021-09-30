package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopybuforecastrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopybuforecastrequestDud struct { 
    


    

}

// Copybuforecastrequest
type Copybuforecastrequest struct { 
    // Description - The description for the forecast
    Description string `json:"description"`


    // WeekDate - The start date of the new forecast to create from the existing forecast. Must correspond to the start day of week for the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`

}

// String returns a JSON representation of the model
func (o *Copybuforecastrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copybuforecastrequest) MarshalJSON() ([]byte, error) {
    type Alias Copybuforecastrequest

    if CopybuforecastrequestMarshalled {
        return []byte("{}"), nil
    }
    CopybuforecastrequestMarshalled = true

    return json.Marshal(&struct { 
        Description string `json:"description"`
        
        WeekDate time.Time `json:"weekDate"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

