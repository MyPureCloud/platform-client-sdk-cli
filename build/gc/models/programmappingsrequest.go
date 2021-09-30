package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgrammappingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgrammappingsrequestDud struct { 
    


    

}

// Programmappingsrequest
type Programmappingsrequest struct { 
    // QueueIds - The program queues
    QueueIds []string `json:"queueIds"`


    // FlowIds - The program flows
    FlowIds []string `json:"flowIds"`

}

// String returns a JSON representation of the model
func (o *Programmappingsrequest) String() string {
    
    
     o.QueueIds = []string{""} 
    
    
    
     o.FlowIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programmappingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Programmappingsrequest

    if ProgrammappingsrequestMarshalled {
        return []byte("{}"), nil
    }
    ProgrammappingsrequestMarshalled = true

    return json.Marshal(&struct { 
        QueueIds []string `json:"queueIds"`
        
        FlowIds []string `json:"flowIds"`
        
        *Alias
    }{
        

        
        QueueIds: []string{""},
        

        

        
        FlowIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

