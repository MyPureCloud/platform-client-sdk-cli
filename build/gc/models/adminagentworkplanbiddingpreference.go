package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminagentworkplanbiddingpreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminagentworkplanbiddingpreferenceDud struct { 
    


    


    


    


    


    

}

// Adminagentworkplanbiddingpreference
type Adminagentworkplanbiddingpreference struct { 
    // Agent - The agent to whom this work plan bidding preference applies
    Agent Userreference `json:"agent"`


    // Submitted - Whether the preference is submitted
    Submitted bool `json:"submitted"`


    // AssignedWorkPlan - The work plan assigned to the agent by the bid process
    AssignedWorkPlan Workplanreference `json:"assignedWorkPlan"`


    // OverriddenWorkPlan - The work plan that overrides the assigned work plan for the agent
    OverriddenWorkPlan Workplanreference `json:"overriddenWorkPlan"`


    // OverrideReason - The reason why the assigned work plan has been overridden. This must be null without an override work plan
    OverrideReason string `json:"overrideReason"`


    // Priorities - The agent priorities for the list of work plans. The index of the priorities should match with the list of work plans that belong to bid group. It contains null if priority is not set for the work plan
    Priorities []int `json:"priorities"`

}

// String returns a JSON representation of the model
func (o *Adminagentworkplanbiddingpreference) String() string {
    
    
    
    
    
     o.Priorities = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminagentworkplanbiddingpreference) MarshalJSON() ([]byte, error) {
    type Alias Adminagentworkplanbiddingpreference

    if AdminagentworkplanbiddingpreferenceMarshalled {
        return []byte("{}"), nil
    }
    AdminagentworkplanbiddingpreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        Submitted bool `json:"submitted"`
        
        AssignedWorkPlan Workplanreference `json:"assignedWorkPlan"`
        
        OverriddenWorkPlan Workplanreference `json:"overriddenWorkPlan"`
        
        OverrideReason string `json:"overrideReason"`
        
        Priorities []int `json:"priorities"`
        *Alias
    }{

        


        


        


        


        


        
        Priorities: []int{0},
        

        Alias: (*Alias)(u),
    })
}

