package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulebiddingpreferenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulebiddingpreferenceresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentschedulebiddingpreferenceresponse
type Agentschedulebiddingpreferenceresponse struct { 
    


    // Submitted - Whether the preference is submitted
    Submitted bool `json:"submitted"`


    // AssignedScheduleSetId - The schedule set assigned to the agent by the bid process. Will be set after bid is processed
    AssignedScheduleSetId string `json:"assignedScheduleSetId"`


    // OverriddenScheduleSetId - The schedule set that overrides the assigned schedule set for the agent
    OverriddenScheduleSetId string `json:"overriddenScheduleSetId"`


    // OverrideReason - The reason why the assigned schedule set has been overridden. This must be null without an override schedule set
    OverrideReason string `json:"overrideReason"`


    // AgentScheduleBidPreferences - The schedule bidding preferences
    AgentScheduleBidPreferences []Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferences"`


    

}

// String returns a JSON representation of the model
func (o *Agentschedulebiddingpreferenceresponse) String() string {
    
    
    
    
     o.AgentScheduleBidPreferences = []Agentschedulebiddingpreferencepriority{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulebiddingpreferenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulebiddingpreferenceresponse

    if AgentschedulebiddingpreferenceresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulebiddingpreferenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Submitted bool `json:"submitted"`
        
        AssignedScheduleSetId string `json:"assignedScheduleSetId"`
        
        OverriddenScheduleSetId string `json:"overriddenScheduleSetId"`
        
        OverrideReason string `json:"overrideReason"`
        
        AgentScheduleBidPreferences []Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferences"`
        *Alias
    }{

        


        


        


        


        


        
        AgentScheduleBidPreferences: []Agentschedulebiddingpreferencepriority{{}},
        


        

        Alias: (*Alias)(u),
    })
}

