package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentvariationanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentvariationanswerDud struct { 
    Id string `json:"id"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    Document Addressableentityref `json:"document"`


    


    SelfUri string `json:"selfUri"`

}

// Knowledgeguestdocumentvariationanswer
type Knowledgeguestdocumentvariationanswer struct { 
    


    


    


    // DocumentVersion - The version of the document.
    DocumentVersion Addressableentityref `json:"documentVersion"`


    // Contexts - The context values associated with the variation.
    Contexts []Knowledgeguestdocumentvariationcontext `json:"contexts"`


    


    // Body - The content for the variation.
    Body Documentbodywithhighlight `json:"body"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentvariationanswer) String() string {
    
     o.Contexts = []Knowledgeguestdocumentvariationcontext{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentvariationanswer) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentvariationanswer

    if KnowledgeguestdocumentvariationanswerMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentvariationanswerMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Contexts []Knowledgeguestdocumentvariationcontext `json:"contexts"`
        
        Body Documentbodywithhighlight `json:"body"`
        *Alias
    }{

        


        


        


        


        
        Contexts: []Knowledgeguestdocumentvariationcontext{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

