package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateworkplanrotationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateworkplanrotationrequestDud struct { 
    


    


    


    


    


    

}

// Updateworkplanrotationrequest
type Updateworkplanrotationrequest struct { 
    // Name - Name of this work plan rotation
    Name string `json:"name"`


    // Enabled - Whether the work plan rotation is enabled for scheduling
    Enabled bool `json:"enabled"`


    // DateRange - The date range to which this work plan rotation applies
    DateRange Daterangewithoptionalend `json:"dateRange"`


    // Agents - Agents in this work plan rotation
    Agents []Updateworkplanrotationagentrequest `json:"agents"`


    // Pattern - Pattern with list of work plan IDs that rotate on a weekly basis
    Pattern Workplanpatternrequest `json:"pattern"`


    // Metadata - Version metadata for this work plan rotation
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateworkplanrotationrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Agents = []Updateworkplanrotationagentrequest{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateworkplanrotationrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateworkplanrotationrequest

    if UpdateworkplanrotationrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateworkplanrotationrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        DateRange Daterangewithoptionalend `json:"dateRange"`
        
        Agents []Updateworkplanrotationagentrequest `json:"agents"`
        
        Pattern Workplanpatternrequest `json:"pattern"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Agents: []Updateworkplanrotationagentrequest{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

