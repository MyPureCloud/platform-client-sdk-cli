package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanbiddingpreferenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanbiddingpreferenceresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentworkplanbiddingpreferenceresponse
type Agentworkplanbiddingpreferenceresponse struct { 
    


    // Submitted - Whether the preference is submitted
    Submitted bool `json:"submitted"`


    // AssignedWorkPlan - The work plan assigned to the agent by the bid process
    AssignedWorkPlan Workplanreference `json:"assignedWorkPlan"`


    // OverriddenWorkPlan - The work plan that overrides the assigned work plan for the agent
    OverriddenWorkPlan Workplanreference `json:"overriddenWorkPlan"`


    // OverrideReason - The reason why the assigned work plan has been overridden. This must be null without an override work plan
    OverrideReason string `json:"overrideReason"`


    // AgentWorkPlanBidPreferences - The list of work plan bidding preferences
    AgentWorkPlanBidPreferences []Agentworkplanbiddingpreference `json:"agentWorkPlanBidPreferences"`


    

}

// String returns a JSON representation of the model
func (o *Agentworkplanbiddingpreferenceresponse) String() string {
    
    
    
    
     o.AgentWorkPlanBidPreferences = []Agentworkplanbiddingpreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanbiddingpreferenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanbiddingpreferenceresponse

    if AgentworkplanbiddingpreferenceresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanbiddingpreferenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Submitted bool `json:"submitted"`
        
        AssignedWorkPlan Workplanreference `json:"assignedWorkPlan"`
        
        OverriddenWorkPlan Workplanreference `json:"overriddenWorkPlan"`
        
        OverrideReason string `json:"overrideReason"`
        
        AgentWorkPlanBidPreferences []Agentworkplanbiddingpreference `json:"agentWorkPlanBidPreferences"`
        *Alias
    }{

        


        


        


        


        


        
        AgentWorkPlanBidPreferences: []Agentworkplanbiddingpreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

