package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulingpreferencestiebreakerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulingpreferencestiebreakerDud struct { 
    


    

}

// Agentschedulingpreferencestiebreaker
type Agentschedulingpreferencestiebreaker struct { 
    // Primary - The primary ranking tiebreaker used when agents have equal preference scores
    Primary string `json:"primary"`


    // Secondary - The secondary ranking tiebreaker. Omit when primary is RandomlyAssign
    Secondary string `json:"secondary"`

}

// String returns a JSON representation of the model
func (o *Agentschedulingpreferencestiebreaker) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulingpreferencestiebreaker) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulingpreferencestiebreaker

    if AgentschedulingpreferencestiebreakerMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulingpreferencestiebreakerMarshalled = true

    return json.Marshal(&struct {
        
        Primary string `json:"primary"`
        
        Secondary string `json:"secondary"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

