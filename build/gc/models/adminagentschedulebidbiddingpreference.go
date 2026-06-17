package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminagentschedulebidbiddingpreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminagentschedulebidbiddingpreferenceDud struct { 
    


    


    


    


    


    


    

}

// Adminagentschedulebidbiddingpreference
type Adminagentschedulebidbiddingpreference struct { 
    // Agent - The agent to whom this schedule bid preference applies
    Agent Userreference `json:"agent"`


    // Submitted - Indicates whether the preference has been submitted
    Submitted bool `json:"submitted"`


    // AssignedScheduleSetId - The schedule set assigned to the agent by the bid process. This will be set after bid is processed
    AssignedScheduleSetId string `json:"assignedScheduleSetId"`


    // OverriddenScheduleSetId - The schedule set that overrides the assigned schedule set for the agent
    OverriddenScheduleSetId string `json:"overriddenScheduleSetId"`


    // OverrideReason - The reason the assigned schedule set has been overridden. This must be null if no override schedule is set
    OverrideReason string `json:"overrideReason"`


    // AgentScheduleBidPreferencePriorities - The agent schedule set preferences
    AgentScheduleBidPreferencePriorities []Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferencePriorities"`


    // EndDate - The end date of this scheduling set preference relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Adminagentschedulebidbiddingpreference) String() string {
    
    
    
    
    
     o.AgentScheduleBidPreferencePriorities = []Agentschedulebiddingpreferencepriority{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminagentschedulebidbiddingpreference) MarshalJSON() ([]byte, error) {
    type Alias Adminagentschedulebidbiddingpreference

    if AdminagentschedulebidbiddingpreferenceMarshalled {
        return []byte("{}"), nil
    }
    AdminagentschedulebidbiddingpreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        Submitted bool `json:"submitted"`
        
        AssignedScheduleSetId string `json:"assignedScheduleSetId"`
        
        OverriddenScheduleSetId string `json:"overriddenScheduleSetId"`
        
        OverrideReason string `json:"overrideReason"`
        
        AgentScheduleBidPreferencePriorities []Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferencePriorities"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        


        


        


        


        
        AgentScheduleBidPreferencePriorities: []Agentschedulebiddingpreferencepriority{{}},
        


        

        Alias: (*Alias)(u),
    })
}

