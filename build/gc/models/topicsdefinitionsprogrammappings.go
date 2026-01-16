package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicsdefinitionsprogrammappingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicsdefinitionsprogrammappingsDud struct { 
    


    


    


    


    

}

// Topicsdefinitionsprogrammappings
type Topicsdefinitionsprogrammappings struct { 
    // Program
    Program Baseprogramentity `json:"program"`


    // Queues
    Queues []Addressableentityref `json:"queues"`


    // Flows
    Flows []Addressableentityref `json:"flows"`


    // ModifiedBy
    ModifiedBy Addressableentityref `json:"modifiedBy"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Topicsdefinitionsprogrammappings) String() string {
    
     o.Queues = []Addressableentityref{{}} 
     o.Flows = []Addressableentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicsdefinitionsprogrammappings) MarshalJSON() ([]byte, error) {
    type Alias Topicsdefinitionsprogrammappings

    if TopicsdefinitionsprogrammappingsMarshalled {
        return []byte("{}"), nil
    }
    TopicsdefinitionsprogrammappingsMarshalled = true

    return json.Marshal(&struct {
        
        Program Baseprogramentity `json:"program"`
        
        Queues []Addressableentityref `json:"queues"`
        
        Flows []Addressableentityref `json:"flows"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        
        Queues: []Addressableentityref{{}},
        


        
        Flows: []Addressableentityref{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

