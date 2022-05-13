package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DaterangewithoptionalendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DaterangewithoptionalendDud struct { 
    


    

}

// Daterangewithoptionalend
type Daterangewithoptionalend struct { 
    // StartBusinessUnitDate - The start date for work plan rotation or an agent, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`


    // EndBusinessUnitDate - The end date for work plan rotation or an agent, interpreted in the business unit's time zone. Null denotes open ended date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`

}

// String returns a JSON representation of the model
func (o *Daterangewithoptionalend) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Daterangewithoptionalend) MarshalJSON() ([]byte, error) {
    type Alias Daterangewithoptionalend

    if DaterangewithoptionalendMarshalled {
        return []byte("{}"), nil
    }
    DaterangewithoptionalendMarshalled = true

    return json.Marshal(&struct {
        
        StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

