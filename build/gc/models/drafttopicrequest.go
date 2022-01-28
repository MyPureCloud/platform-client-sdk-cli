package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DrafttopicrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DrafttopicrequestDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Drafttopicrequest
type Drafttopicrequest struct { 
    // Id - Id for a topic.
    Id string `json:"id"`


    // Name - Name/Label for a topic.
    Name string `json:"name"`


    // Phrases - The phrases that are extracted for a topic.
    Phrases []string `json:"phrases"`


    

}

// String returns a JSON representation of the model
func (o *Drafttopicrequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Phrases = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Drafttopicrequest) MarshalJSON() ([]byte, error) {
    type Alias Drafttopicrequest

    if DrafttopicrequestMarshalled {
        return []byte("{}"), nil
    }
    DrafttopicrequestMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Phrases []string `json:"phrases"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Phrases: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

