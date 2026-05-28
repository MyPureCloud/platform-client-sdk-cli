package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpportunityenrollmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpportunityenrollmentDud struct { 
    Id string `json:"id"`

}

// Opportunityenrollment
type Opportunityenrollment struct { 
    

}

// String returns a JSON representation of the model
func (o *Opportunityenrollment) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opportunityenrollment) MarshalJSON() ([]byte, error) {
    type Alias Opportunityenrollment

    if OpportunityenrollmentMarshalled {
        return []byte("{}"), nil
    }
    OpportunityenrollmentMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

