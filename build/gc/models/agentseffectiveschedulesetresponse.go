package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentseffectiveschedulesetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentseffectiveschedulesetresponseDud struct { 
    

}

// Agentseffectiveschedulesetresponse
type Agentseffectiveschedulesetresponse struct { 
    // EffectiveBids - Bids that are effective from the startDate until the weekCount
    EffectiveBids []Agenteffectivebid `json:"effectiveBids"`

}

// String returns a JSON representation of the model
func (o *Agentseffectiveschedulesetresponse) String() string {
     o.EffectiveBids = []Agenteffectivebid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentseffectiveschedulesetresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentseffectiveschedulesetresponse

    if AgentseffectiveschedulesetresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentseffectiveschedulesetresponseMarshalled = true

    return json.Marshal(&struct {
        
        EffectiveBids []Agenteffectivebid `json:"effectiveBids"`
        *Alias
    }{

        
        EffectiveBids: []Agenteffectivebid{{}},
        

        Alias: (*Alias)(u),
    })
}

