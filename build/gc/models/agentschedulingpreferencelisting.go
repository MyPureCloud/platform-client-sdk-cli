package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulingpreferencelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulingpreferencelistingDud struct { 
    

}

// Agentschedulingpreferencelisting
type Agentschedulingpreferencelisting struct { 
    // Entities
    Entities []Agentschedulingpreference `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Agentschedulingpreferencelisting) String() string {
     o.Entities = []Agentschedulingpreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulingpreferencelisting) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulingpreferencelisting

    if AgentschedulingpreferencelistingMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulingpreferencelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Agentschedulingpreference `json:"entities"`
        *Alias
    }{

        
        Entities: []Agentschedulingpreference{{}},
        

        Alias: (*Alias)(u),
    })
}

