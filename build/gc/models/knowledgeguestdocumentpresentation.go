package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentpresentationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentpresentationDud struct { 
    


    


    


    


    SessionId string `json:"sessionId"`


    Application Knowledgeguestsearchclientapplication `json:"application"`

}

// Knowledgeguestdocumentpresentation
type Knowledgeguestdocumentpresentation struct { 
    // Documents - The presented documents
    Documents []Presentedknowledgedocument `json:"documents"`


    // SearchId - The search that surfaced the documents that were presented.
    SearchId string `json:"searchId"`


    // QueryType - The type of the query that surfaced the documents.
    QueryType string `json:"queryType"`


    // SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.Generative: A generated answer in a snippet was shown.
    SurfacingMethod string `json:"surfacingMethod"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentpresentation) String() string {
     o.Documents = []Presentedknowledgedocument{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentpresentation) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentpresentation

    if KnowledgeguestdocumentpresentationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentpresentationMarshalled = true

    return json.Marshal(&struct {
        
        Documents []Presentedknowledgedocument `json:"documents"`
        
        SearchId string `json:"searchId"`
        
        QueryType string `json:"queryType"`
        
        SurfacingMethod string `json:"surfacingMethod"`
        *Alias
    }{

        
        Documents: []Presentedknowledgedocument{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

