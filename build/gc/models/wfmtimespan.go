package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmtimespanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmtimespanDud struct { 
    


    

}

// Wfmtimespan
type Wfmtimespan struct { 
    // StartDate - Start date of the time span. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the time span from the start date in minutes
    LengthMinutes int `json:"lengthMinutes"`

}

// String returns a JSON representation of the model
func (o *Wfmtimespan) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmtimespan) MarshalJSON() ([]byte, error) {
    type Alias Wfmtimespan

    if WfmtimespanMarshalled {
        return []byte("{}"), nil
    }
    WfmtimespanMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

