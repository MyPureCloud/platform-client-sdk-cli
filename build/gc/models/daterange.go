package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DaterangeDud struct { 
    


    

}

// Daterange
type Daterange struct { 
    // StartDate - The inclusive start of a date range in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
    StartDate string `json:"startDate"`


    // EndDate - The inclusive end of a date range in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
    EndDate string `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Daterange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Daterange) MarshalJSON() ([]byte, error) {
    type Alias Daterange

    if DaterangeMarshalled {
        return []byte("{}"), nil
    }
    DaterangeMarshalled = true

    return json.Marshal(&struct {
        
        StartDate string `json:"startDate"`
        
        EndDate string `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

