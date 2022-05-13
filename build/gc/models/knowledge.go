package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeDud struct { 
    


    

}

// Knowledge
type Knowledge struct { 
    // Enabled - whether or not knowledge base is enabled
    Enabled bool `json:"enabled"`


    // KnowledgeBase - The knowledge base for messenger
    KnowledgeBase Addressableentityref `json:"knowledgeBase"`

}

// String returns a JSON representation of the model
func (o *Knowledge) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledge) MarshalJSON() ([]byte, error) {
    type Alias Knowledge

    if KnowledgeMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        KnowledgeBase Addressableentityref `json:"knowledgeBase"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

