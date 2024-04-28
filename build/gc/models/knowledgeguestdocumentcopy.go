package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentcopyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentcopyDud struct { 
    


    


    


    


    


    SessionId string `json:"sessionId"`


    Application Knowledgeguestsearchclientapplication `json:"application"`

}

// Knowledgeguestdocumentcopy
type Knowledgeguestdocumentcopy struct { 
    // DocumentVariationId - The variation of the document whose content was copied.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the document whose content was copied.
    DocumentVersionId string `json:"documentVersionId"`


    // SearchId - The search that surfaced the document whose content was copied.
    SearchId string `json:"searchId"`


    // QueryType - The type of the query that surfaced the document.
    QueryType string `json:"queryType"`


    // SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.
    SurfacingMethod string `json:"surfacingMethod"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentcopy) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentcopy) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentcopy

    if KnowledgeguestdocumentcopyMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentcopyMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        
        SearchId string `json:"searchId"`
        
        QueryType string `json:"queryType"`
        
        SurfacingMethod string `json:"surfacingMethod"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

