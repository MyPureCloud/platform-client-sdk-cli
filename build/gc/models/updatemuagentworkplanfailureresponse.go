package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemuagentworkplanfailureresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemuagentworkplanfailureresponseDud struct { 
    


    


    

}

// Updatemuagentworkplanfailureresponse
type Updatemuagentworkplanfailureresponse struct { 
    // User - The user for whom the work plan update has failed
    User Userreference `json:"user"`


    // Failure - The work plan update failure reason
    Failure string `json:"failure"`


    // NotFoundWorkPlanId - The id of the work plan that has not been found
    NotFoundWorkPlanId string `json:"notFoundWorkPlanId"`

}

// String returns a JSON representation of the model
func (o *Updatemuagentworkplanfailureresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemuagentworkplanfailureresponse) MarshalJSON() ([]byte, error) {
    type Alias Updatemuagentworkplanfailureresponse

    if UpdatemuagentworkplanfailureresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdatemuagentworkplanfailureresponseMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Failure string `json:"failure"`
        
        NotFoundWorkPlanId string `json:"notFoundWorkPlanId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

