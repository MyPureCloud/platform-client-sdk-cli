package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateagentunavailabletimesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateagentunavailabletimesrequestDud struct { 
    


    

}

// Validateagentunavailabletimesrequest
type Validateagentunavailabletimesrequest struct { 
    // ValidationWeekDate - The ID of the week to validate. Must correspond to the start day of week of the business unit to which the agent belongs in the format YYYY-MM-dd. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    ValidationWeekDate time.Time `json:"validationWeekDate"`


    // UnavailableTimes - Proposed changes to agent's unavailable time spans to be validated
    UnavailableTimes []Updateunavailabletime `json:"unavailableTimes"`

}

// String returns a JSON representation of the model
func (o *Validateagentunavailabletimesrequest) String() string {
    
     o.UnavailableTimes = []Updateunavailabletime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateagentunavailabletimesrequest) MarshalJSON() ([]byte, error) {
    type Alias Validateagentunavailabletimesrequest

    if ValidateagentunavailabletimesrequestMarshalled {
        return []byte("{}"), nil
    }
    ValidateagentunavailabletimesrequestMarshalled = true

    return json.Marshal(&struct {
        
        ValidationWeekDate time.Time `json:"validationWeekDate"`
        
        UnavailableTimes []Updateunavailabletime `json:"unavailableTimes"`
        *Alias
    }{

        


        
        UnavailableTimes: []Updateunavailabletime{{}},
        

        Alias: (*Alias)(u),
    })
}

