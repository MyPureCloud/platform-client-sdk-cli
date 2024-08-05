package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateagentworkplanbiddingpreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateagentworkplanbiddingpreferenceDud struct { 
    


    

}

// Updateagentworkplanbiddingpreference
type Updateagentworkplanbiddingpreference struct { 
    // Submitted - Whether the preference is submitted
    Submitted bool `json:"submitted"`


    // AgentWorkPlanBidPreferences - The list of work plan bidding preferences
    AgentWorkPlanBidPreferences []Agentworkplanbiddingpreferencerequest `json:"agentWorkPlanBidPreferences"`

}

// String returns a JSON representation of the model
func (o *Updateagentworkplanbiddingpreference) String() string {
    
     o.AgentWorkPlanBidPreferences = []Agentworkplanbiddingpreferencerequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateagentworkplanbiddingpreference) MarshalJSON() ([]byte, error) {
    type Alias Updateagentworkplanbiddingpreference

    if UpdateagentworkplanbiddingpreferenceMarshalled {
        return []byte("{}"), nil
    }
    UpdateagentworkplanbiddingpreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Submitted bool `json:"submitted"`
        
        AgentWorkPlanBidPreferences []Agentworkplanbiddingpreferencerequest `json:"agentWorkPlanBidPreferences"`
        *Alias
    }{

        


        
        AgentWorkPlanBidPreferences: []Agentworkplanbiddingpreferencerequest{{}},
        

        Alias: (*Alias)(u),
    })
}

