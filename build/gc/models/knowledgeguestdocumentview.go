package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentviewDud struct { 
    


    


    


    

}

// Knowledgeguestdocumentview
type Knowledgeguestdocumentview struct { 
    // DocumentVariationId - The variation of the viewed document.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the viewed document.
    DocumentVersionId string `json:"documentVersionId"`


    // SearchId - The search that surfaced the viewed document.
    SearchId string `json:"searchId"`


    // QueryType - The type of the query that surfaced the document.
    QueryType string `json:"queryType"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentview) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentview) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentview

    if KnowledgeguestdocumentviewMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentviewMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        
        SearchId string `json:"searchId"`
        
        QueryType string `json:"queryType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

