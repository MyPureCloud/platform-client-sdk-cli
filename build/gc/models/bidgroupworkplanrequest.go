package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BidgroupworkplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BidgroupworkplanrequestDud struct { 
    


    


    SuggestedAgentCount int `json:"suggestedAgentCount"`


    AgentCountRange Agentcountrange `json:"agentCountRange"`

}

// Bidgroupworkplanrequest
type Bidgroupworkplanrequest struct { 
    // WorkPlanId - The ID of the work plan used in the bid group
    WorkPlanId string `json:"workPlanId"`


    // OverrideAgentCount - The modified agent count for this work plan
    OverrideAgentCount int `json:"overrideAgentCount"`


    


    

}

// String returns a JSON representation of the model
func (o *Bidgroupworkplanrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bidgroupworkplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Bidgroupworkplanrequest

    if BidgroupworkplanrequestMarshalled {
        return []byte("{}"), nil
    }
    BidgroupworkplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanId string `json:"workPlanId"`
        
        OverrideAgentCount int `json:"overrideAgentCount"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

