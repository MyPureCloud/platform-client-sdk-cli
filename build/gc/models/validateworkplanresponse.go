package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateworkplanresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateworkplanresponseDud struct { 
    


    


    

}

// Validateworkplanresponse
type Validateworkplanresponse struct { 
    // WorkPlan - The work plan reference associated with this response
    WorkPlan Workplanreference `json:"workPlan"`


    // Valid - Whether the work plan is valid or not
    Valid bool `json:"valid"`


    // Messages - Validation messages for this work plan
    Messages Validateworkplanmessages `json:"messages"`

}

// String returns a JSON representation of the model
func (o *Validateworkplanresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateworkplanresponse) MarshalJSON() ([]byte, error) {
    type Alias Validateworkplanresponse

    if ValidateworkplanresponseMarshalled {
        return []byte("{}"), nil
    }
    ValidateworkplanresponseMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlan Workplanreference `json:"workPlan"`
        
        Valid bool `json:"valid"`
        
        Messages Validateworkplanmessages `json:"messages"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

