package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanpatternresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanpatternresponseDud struct { 
    

}

// Workplanpatternresponse
type Workplanpatternresponse struct { 
    // WorkPlans - List of work plans in order of rotation on a weekly basis
    WorkPlans []Workplanreference `json:"workPlans"`

}

// String returns a JSON representation of the model
func (o *Workplanpatternresponse) String() string {
     o.WorkPlans = []Workplanreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanpatternresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanpatternresponse

    if WorkplanpatternresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanpatternresponseMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlans []Workplanreference `json:"workPlans"`
        *Alias
    }{

        
        WorkPlans: []Workplanreference{{}},
        

        Alias: (*Alias)(u),
    })
}

