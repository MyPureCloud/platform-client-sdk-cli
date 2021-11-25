package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftrequestDud struct { 
    


    

}

// Draftrequest
type Draftrequest struct { 
    // Intents - Draft intent object.
    Intents []Draftintents `json:"intents"`


    // Topic
    Topic []Drafttopics `json:"topic"`

}

// String returns a JSON representation of the model
func (o *Draftrequest) String() string {
    
    
     o.Intents = []Draftintents{{}} 
    
    
    
     o.Topic = []Drafttopics{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draftrequest) MarshalJSON() ([]byte, error) {
    type Alias Draftrequest

    if DraftrequestMarshalled {
        return []byte("{}"), nil
    }
    DraftrequestMarshalled = true

    return json.Marshal(&struct { 
        Intents []Draftintents `json:"intents"`
        
        Topic []Drafttopics `json:"topic"`
        
        *Alias
    }{
        

        
        Intents: []Draftintents{{}},
        

        

        
        Topic: []Drafttopics{{}},
        

        
        Alias: (*Alias)(u),
    })
}

