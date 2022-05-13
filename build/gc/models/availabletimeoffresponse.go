package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletimeoffresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletimeoffresponseDud struct { 
    

}

// Availabletimeoffresponse - The list of date ranges with available time off values and the current waitlist per granularity.
type Availabletimeoffresponse struct { 
    // Values
    Values []Availabletimeoffrange `json:"values"`

}

// String returns a JSON representation of the model
func (o *Availabletimeoffresponse) String() string {
     o.Values = []Availabletimeoffrange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletimeoffresponse) MarshalJSON() ([]byte, error) {
    type Alias Availabletimeoffresponse

    if AvailabletimeoffresponseMarshalled {
        return []byte("{}"), nil
    }
    AvailabletimeoffresponseMarshalled = true

    return json.Marshal(&struct {
        
        Values []Availabletimeoffrange `json:"values"`
        *Alias
    }{

        
        Values: []Availabletimeoffrange{{}},
        

        Alias: (*Alias)(u),
    })
}

