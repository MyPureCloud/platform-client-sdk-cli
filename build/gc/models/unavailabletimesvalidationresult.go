package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnavailabletimesvalidationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnavailabletimesvalidationresultDud struct { 
    


    


    

}

// Unavailabletimesvalidationresult
type Unavailabletimesvalidationresult struct { 
    // Valid - Indicates whether there are no violations in the given unavailable times
    Valid bool `json:"valid"`


    // InvalidWorkPlans - Invalid work plans that were used when validating the agents unavailable times
    InvalidWorkPlans []Workplanreference `json:"invalidWorkPlans"`


    // WorkPlanConstraintsViolationMessage - Message for set of agent unavailable times violating work plan constraints
    WorkPlanConstraintsViolationMessage Workplanconstraintsviolationmessage `json:"workPlanConstraintsViolationMessage"`

}

// String returns a JSON representation of the model
func (o *Unavailabletimesvalidationresult) String() string {
    
     o.InvalidWorkPlans = []Workplanreference{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unavailabletimesvalidationresult) MarshalJSON() ([]byte, error) {
    type Alias Unavailabletimesvalidationresult

    if UnavailabletimesvalidationresultMarshalled {
        return []byte("{}"), nil
    }
    UnavailabletimesvalidationresultMarshalled = true

    return json.Marshal(&struct {
        
        Valid bool `json:"valid"`
        
        InvalidWorkPlans []Workplanreference `json:"invalidWorkPlans"`
        
        WorkPlanConstraintsViolationMessage Workplanconstraintsviolationmessage `json:"workPlanConstraintsViolationMessage"`
        *Alias
    }{

        


        
        InvalidWorkPlans: []Workplanreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

