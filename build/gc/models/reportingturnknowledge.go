package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgeDud struct { 
    


    


    

}

// Reportingturnknowledge
type Reportingturnknowledge struct { 
    // KnowledgeBaseId - The Knowledge Base ID that the captured knowledge data relates to.
    KnowledgeBaseId string `json:"knowledgeBaseId"`


    // Feedback - The knowledge feedback data that was captured during this reporting turn.
    Feedback Reportingturnknowledgefeedback `json:"feedback"`


    // Search - The knowledge search data that was captured during this reporting turn.
    Search Reportingturnknowledgesearch `json:"search"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledge) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledge) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledge

    if ReportingturnknowledgeMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgeMarshalled = true

    return json.Marshal(&struct {
        
        KnowledgeBaseId string `json:"knowledgeBaseId"`
        
        Feedback Reportingturnknowledgefeedback `json:"feedback"`
        
        Search Reportingturnknowledgesearch `json:"search"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

