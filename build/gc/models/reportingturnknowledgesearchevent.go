package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgesearcheventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgesearcheventDud struct { 
    


    


    


    


    

}

// Reportingturnknowledgesearchevent
type Reportingturnknowledgesearchevent struct { 
    // SearchId - The ID of this knowledge search.
    SearchId string `json:"searchId"`


    // KnowledgeBaseId - The Knowledge Base ID that the captured knowledge data relates to.
    KnowledgeBaseId string `json:"knowledgeBaseId"`


    // Documents - The list of search documents that the feedback applies to.
    Documents []Reportingturnknowledgedocument `json:"documents"`


    // SearchQuery - The search query that was used to search the Knowledge Base documents for a matching question.
    SearchQuery string `json:"searchQuery"`


    // AnswerDocumentId - The document ID of the search answer.
    AnswerDocumentId string `json:"answerDocumentId"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgesearchevent) String() string {
    
    
     o.Documents = []Reportingturnknowledgedocument{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgesearchevent) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgesearchevent

    if ReportingturnknowledgesearcheventMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgesearcheventMarshalled = true

    return json.Marshal(&struct {
        
        SearchId string `json:"searchId"`
        
        KnowledgeBaseId string `json:"knowledgeBaseId"`
        
        Documents []Reportingturnknowledgedocument `json:"documents"`
        
        SearchQuery string `json:"searchQuery"`
        
        AnswerDocumentId string `json:"answerDocumentId"`
        *Alias
    }{

        


        


        
        Documents: []Reportingturnknowledgedocument{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

