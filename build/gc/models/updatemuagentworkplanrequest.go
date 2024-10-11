package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemuagentworkplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemuagentworkplanrequestDud struct { 
    


    


    

}

// Updatemuagentworkplanrequest
type Updatemuagentworkplanrequest struct { 
    // UserId - The agent id for whom the work plan is updated
    UserId string `json:"userId"`


    // WorkPlanId - The current work plan ID for the agent
    WorkPlanId Valuewrapperstring `json:"workPlanId"`


    // WorkPlanOverrides - The list of work plan overrides for the agent
    WorkPlanOverrides Workplanoverridelistwrapperworkplanoverriderequest `json:"workPlanOverrides"`

}

// String returns a JSON representation of the model
func (o *Updatemuagentworkplanrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemuagentworkplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatemuagentworkplanrequest

    if UpdatemuagentworkplanrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatemuagentworkplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        WorkPlanId Valuewrapperstring `json:"workPlanId"`
        
        WorkPlanOverrides Workplanoverridelistwrapperworkplanoverriderequest `json:"workPlanOverrides"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

