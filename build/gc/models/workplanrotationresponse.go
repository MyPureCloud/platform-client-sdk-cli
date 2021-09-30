package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanrotationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanrotationresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workplanrotationresponse
type Workplanrotationresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Enabled - Whether the work plan rotation is enabled for scheduling
    Enabled bool `json:"enabled"`


    // DateRange - The date range to which this work plan rotation applies
    DateRange Daterangewithoptionalend `json:"dateRange"`


    // Pattern - Pattern with ordered list of work plans that rotate on a weekly basis
    Pattern Workplanpatternresponse `json:"pattern"`


    // AgentCount - Number of agents in this work plan rotation
    AgentCount int `json:"agentCount"`


    // Agents - Agents in this work plan rotation. Populate with expand=agents for GET WorkPlanRotationsList (defaults to empty list)
    Agents []Workplanrotationagentresponse `json:"agents"`


    // Metadata - Version metadata for this work plan rotation
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Workplanrotationresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Agents = []Workplanrotationagentresponse{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanrotationresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanrotationresponse

    if WorkplanrotationresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanrotationresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        DateRange Daterangewithoptionalend `json:"dateRange"`
        
        Pattern Workplanpatternresponse `json:"pattern"`
        
        AgentCount int `json:"agentCount"`
        
        Agents []Workplanrotationagentresponse `json:"agents"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Agents: []Workplanrotationagentresponse{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

