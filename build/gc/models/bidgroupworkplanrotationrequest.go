package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BidgroupworkplanrotationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BidgroupworkplanrotationrequestDud struct { 
    


    

}

// Bidgroupworkplanrotationrequest
type Bidgroupworkplanrotationrequest struct { 
    // WorkPlanRotationId - The ID of the work plan rotation used in the bid group
    WorkPlanRotationId string `json:"workPlanRotationId"`


    // AgentCount - The count of agents that can be assigned to this work plan rotation
    AgentCount int `json:"agentCount"`

}

// String returns a JSON representation of the model
func (o *Bidgroupworkplanrotationrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bidgroupworkplanrotationrequest) MarshalJSON() ([]byte, error) {
    type Alias Bidgroupworkplanrotationrequest

    if BidgroupworkplanrotationrequestMarshalled {
        return []byte("{}"), nil
    }
    BidgroupworkplanrotationrequestMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanRotationId string `json:"workPlanRotationId"`
        
        AgentCount int `json:"agentCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

