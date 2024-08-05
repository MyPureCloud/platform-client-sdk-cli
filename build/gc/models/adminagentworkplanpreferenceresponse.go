package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminagentworkplanpreferenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminagentworkplanpreferenceresponseDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Adminagentworkplanpreferenceresponse
type Adminagentworkplanpreferenceresponse struct { 
    


    // WorkPlans - The list of work plans that belong to this bid group
    WorkPlans []Workplanreference `json:"workPlans"`


    // AgentWorkPlanBidPreferences - The list of agents work plan bidding preferences
    AgentWorkPlanBidPreferences []Adminagentworkplanbiddingpreference `json:"agentWorkPlanBidPreferences"`


    

}

// String returns a JSON representation of the model
func (o *Adminagentworkplanpreferenceresponse) String() string {
     o.WorkPlans = []Workplanreference{{}} 
     o.AgentWorkPlanBidPreferences = []Adminagentworkplanbiddingpreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminagentworkplanpreferenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Adminagentworkplanpreferenceresponse

    if AdminagentworkplanpreferenceresponseMarshalled {
        return []byte("{}"), nil
    }
    AdminagentworkplanpreferenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlans []Workplanreference `json:"workPlans"`
        
        AgentWorkPlanBidPreferences []Adminagentworkplanbiddingpreference `json:"agentWorkPlanBidPreferences"`
        *Alias
    }{

        


        
        WorkPlans: []Workplanreference{{}},
        


        
        AgentWorkPlanBidPreferences: []Adminagentworkplanbiddingpreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

