package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentownedmappingpreviewlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentownedmappingpreviewlistingDud struct { 
    

}

// Agentownedmappingpreviewlisting
type Agentownedmappingpreviewlisting struct { 
    // Entities
    Entities []Agentownedmappingpreview `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Agentownedmappingpreviewlisting) String() string {
    
    
     o.Entities = []Agentownedmappingpreview{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentownedmappingpreviewlisting) MarshalJSON() ([]byte, error) {
    type Alias Agentownedmappingpreviewlisting

    if AgentownedmappingpreviewlistingMarshalled {
        return []byte("{}"), nil
    }
    AgentownedmappingpreviewlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Agentownedmappingpreview `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Agentownedmappingpreview{{}},
        

        
        Alias: (*Alias)(u),
    })
}

