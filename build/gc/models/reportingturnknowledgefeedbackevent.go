package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgefeedbackeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgefeedbackeventDud struct { 
    


    


    


    


    


    

}

// Reportingturnknowledgefeedbackevent
type Reportingturnknowledgefeedbackevent struct { 
    // SearchId - The ID of this knowledge search.
    SearchId string `json:"searchId"`


    // KnowledgeBaseId - The Knowledge Base ID that the captured knowledge data relates to.
    KnowledgeBaseId string `json:"knowledgeBaseId"`


    // Documents - The list of search documents that the feedback applies to.
    Documents []Reportingturnknowledgedocument `json:"documents"`


    // FeedbackRating - The feedback rating for the search (1.0 - 5.0). 1 = Negative, 5 = Positive.
    FeedbackRating int `json:"feedbackRating"`


    // DocumentVariationId - The variation of the document.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the document.
    DocumentVersionId string `json:"documentVersionId"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgefeedbackevent) String() string {
    
    
     o.Documents = []Reportingturnknowledgedocument{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgefeedbackevent) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgefeedbackevent

    if ReportingturnknowledgefeedbackeventMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgefeedbackeventMarshalled = true

    return json.Marshal(&struct {
        
        SearchId string `json:"searchId"`
        
        KnowledgeBaseId string `json:"knowledgeBaseId"`
        
        Documents []Reportingturnknowledgedocument `json:"documents"`
        
        FeedbackRating int `json:"feedbackRating"`
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        *Alias
    }{

        


        


        
        Documents: []Reportingturnknowledgedocument{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

