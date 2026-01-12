package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnavailabletimestimespanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnavailabletimestimespanDud struct { 
    


    

}

// Unavailabletimestimespan
type Unavailabletimestimespan struct { 
    // StartDate - Start date of the time span. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the time span from the start date in minutes
    LengthMinutes int `json:"lengthMinutes"`

}

// String returns a JSON representation of the model
func (o *Unavailabletimestimespan) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unavailabletimestimespan) MarshalJSON() ([]byte, error) {
    type Alias Unavailabletimestimespan

    if UnavailabletimestimespanMarshalled {
        return []byte("{}"), nil
    }
    UnavailabletimestimespanMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

