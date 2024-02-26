package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentversionvariationreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentversionvariationreferenceDud struct { 
    


    


    

}

// Knowledgedocumentversionvariationreference
type Knowledgedocumentversionvariationreference struct { 
    // DocumentId - The ID of the document.
    DocumentId string `json:"documentId"`


    // DocumentVariationId - The variation of the document.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the document.
    DocumentVersionId string `json:"documentVersionId"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversionvariationreference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentversionvariationreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentversionvariationreference

    if KnowledgedocumentversionvariationreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentversionvariationreferenceMarshalled = true

    return json.Marshal(&struct {
        
        DocumentId string `json:"documentId"`
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

