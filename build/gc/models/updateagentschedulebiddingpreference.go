package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateagentschedulebiddingpreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateagentschedulebiddingpreferenceDud struct { 
    


    

}

// Updateagentschedulebiddingpreference
type Updateagentschedulebiddingpreference struct { 
    // Submitted - Whether the preference is submitted
    Submitted bool `json:"submitted"`


    // AgentScheduleBidPreferences - The schedule bidding preferences
    AgentScheduleBidPreferences []Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferences"`

}

// String returns a JSON representation of the model
func (o *Updateagentschedulebiddingpreference) String() string {
    
     o.AgentScheduleBidPreferences = []Agentschedulebiddingpreferencepriority{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateagentschedulebiddingpreference) MarshalJSON() ([]byte, error) {
    type Alias Updateagentschedulebiddingpreference

    if UpdateagentschedulebiddingpreferenceMarshalled {
        return []byte("{}"), nil
    }
    UpdateagentschedulebiddingpreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Submitted bool `json:"submitted"`
        
        AgentScheduleBidPreferences []Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferences"`
        *Alias
    }{

        


        
        AgentScheduleBidPreferences: []Agentschedulebiddingpreferencepriority{{}},
        

        Alias: (*Alias)(u),
    })
}

