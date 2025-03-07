package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresentedknowledgedocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresentedknowledgedocumentDud struct { 
    


    


    


    

}

// Presentedknowledgedocument
type Presentedknowledgedocument struct { 
    // DocumentId - The ID of the document.
    DocumentId string `json:"documentId"`


    // DocumentVariationId - The variation of the document.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the document.
    DocumentVersionId string `json:"documentVersionId"`


    // SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.
    SurfacingMethod string `json:"surfacingMethod"`

}

// String returns a JSON representation of the model
func (o *Presentedknowledgedocument) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presentedknowledgedocument) MarshalJSON() ([]byte, error) {
    type Alias Presentedknowledgedocument

    if PresentedknowledgedocumentMarshalled {
        return []byte("{}"), nil
    }
    PresentedknowledgedocumentMarshalled = true

    return json.Marshal(&struct {
        
        DocumentId string `json:"documentId"`
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        
        SurfacingMethod string `json:"surfacingMethod"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

