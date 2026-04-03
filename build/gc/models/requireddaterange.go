package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequireddaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequireddaterangeDud struct { 
    


    

}

// Requireddaterange
type Requireddaterange struct { 
    // StartDate - The start of a date range in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end of a date range in ISO-8601 format
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Requireddaterange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requireddaterange) MarshalJSON() ([]byte, error) {
    type Alias Requireddaterange

    if RequireddaterangeMarshalled {
        return []byte("{}"), nil
    }
    RequireddaterangeMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

