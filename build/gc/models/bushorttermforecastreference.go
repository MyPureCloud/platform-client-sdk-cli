package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BushorttermforecastreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BushorttermforecastreferenceDud struct { 
    


    


    Description string `json:"description"`


    SelfUri string `json:"selfUri"`

}

// Bushorttermforecastreference
type Bushorttermforecastreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    


    

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bushorttermforecastreference) MarshalJSON() ([]byte, error) {
    type Alias Bushorttermforecastreference

    if BushorttermforecastreferenceMarshalled {
        return []byte("{}"), nil
    }
    BushorttermforecastreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

