package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddworkplanrotationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddworkplanrotationrequestDud struct { 
    


    


    


    

}

// Addworkplanrotationrequest
type Addworkplanrotationrequest struct { 
    // Name - Name of this work plan rotation
    Name string `json:"name"`


    // DateRange - The date range to which this work plan rotation applies
    DateRange Daterangewithoptionalend `json:"dateRange"`


    // Agents - Agents in this work plan rotation
    Agents []Addworkplanrotationagentrequest `json:"agents"`


    // Pattern - Pattern with list of work plan IDs that rotate on a weekly basis
    Pattern Workplanpatternrequest `json:"pattern"`

}

// String returns a JSON representation of the model
func (o *Addworkplanrotationrequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Agents = []Addworkplanrotationagentrequest{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addworkplanrotationrequest) MarshalJSON() ([]byte, error) {
    type Alias Addworkplanrotationrequest

    if AddworkplanrotationrequestMarshalled {
        return []byte("{}"), nil
    }
    AddworkplanrotationrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        DateRange Daterangewithoptionalend `json:"dateRange"`
        
        Agents []Addworkplanrotationagentrequest `json:"agents"`
        
        Pattern Workplanpatternrequest `json:"pattern"`
        
        *Alias
    }{
        

        

        

        

        

        
        Agents: []Addworkplanrotationagentrequest{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

