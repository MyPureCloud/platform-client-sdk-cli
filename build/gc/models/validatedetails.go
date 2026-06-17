package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidatedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidatedetailsDud struct { 
    

}

// Validatedetails
type Validatedetails struct { 
    // Flow - The flow to validate. If you do not provide the flow ID, you must provide both the name and type.
    Flow Architectflowreference `json:"flow"`

}

// String returns a JSON representation of the model
func (o *Validatedetails) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validatedetails) MarshalJSON() ([]byte, error) {
    type Alias Validatedetails

    if ValidatedetailsMarshalled {
        return []byte("{}"), nil
    }
    ValidatedetailsMarshalled = true

    return json.Marshal(&struct {
        
        Flow Architectflowreference `json:"flow"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

