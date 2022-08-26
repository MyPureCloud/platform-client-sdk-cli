package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentvariationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentvariationDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    Document Addressableentityref `json:"document"`


    SelfUri string `json:"selfUri"`

}

// Knowledgeguestdocumentvariation
type Knowledgeguestdocumentvariation struct { 
    


    // Body - The content for the variation.
    Body Documentbody `json:"body"`


    


    


    // DocumentVersion - The version of the document.
    DocumentVersion Addressableentityref `json:"documentVersion"`


    // Contexts - The context values associated with the variation.
    Contexts []Knowledgeguestdocumentvariationcontext `json:"contexts"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentvariation) String() string {
    
    
     o.Contexts = []Knowledgeguestdocumentvariationcontext{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentvariation) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentvariation

    if KnowledgeguestdocumentvariationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentvariationMarshalled = true

    return json.Marshal(&struct {
        
        Body Documentbody `json:"body"`
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Contexts []Knowledgeguestdocumentvariationcontext `json:"contexts"`
        *Alias
    }{

        


        


        


        


        


        
        Contexts: []Knowledgeguestdocumentvariationcontext{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

