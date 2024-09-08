package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingconversationattributesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingconversationattributesresponseDud struct { 
    


    


    


    


    

}

// Routingconversationattributesresponse
type Routingconversationattributesresponse struct { 
    // Priority - Current priority value on in-queue conversation. Range:[-25000000, 25000000]
    Priority int `json:"priority"`


    // Skills - Current routing skills on in-queue conversation
    Skills []Routingskill `json:"skills"`


    // Language - Current language on in-queue conversation
    Language Language `json:"language"`


    // Label - Current label on in-queue conversation
    Label Utilizationlabel `json:"label"`


    // ScoredAgents - Current scored agents on in-queue conversation
    ScoredAgents []Scoredagent `json:"scoredAgents"`

}

// String returns a JSON representation of the model
func (o *Routingconversationattributesresponse) String() string {
    
     o.Skills = []Routingskill{{}} 
    
    
     o.ScoredAgents = []Scoredagent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingconversationattributesresponse) MarshalJSON() ([]byte, error) {
    type Alias Routingconversationattributesresponse

    if RoutingconversationattributesresponseMarshalled {
        return []byte("{}"), nil
    }
    RoutingconversationattributesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Priority int `json:"priority"`
        
        Skills []Routingskill `json:"skills"`
        
        Language Language `json:"language"`
        
        Label Utilizationlabel `json:"label"`
        
        ScoredAgents []Scoredagent `json:"scoredAgents"`
        *Alias
    }{

        


        
        Skills: []Routingskill{{}},
        


        


        


        
        ScoredAgents: []Scoredagent{{}},
        

        Alias: (*Alias)(u),
    })
}

