package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    Document Knowledgedocumentreference `json:"document"`


    SelfUri string `json:"selfUri"`

}

// Documentvariation
type Documentvariation struct { 
    


    // Body - The content for the variation.
    Body Documentbody `json:"body"`


    


    


    // DocumentVersion - The version of the document.
    DocumentVersion Addressableentityref `json:"documentVersion"`


    // Contexts - The context values associated with the variation.
    Contexts []Documentvariationcontext `json:"contexts"`


    


    

}

// String returns a JSON representation of the model
func (o *Documentvariation) String() string {
    
    
     o.Contexts = []Documentvariationcontext{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariation) MarshalJSON() ([]byte, error) {
    type Alias Documentvariation

    if DocumentvariationMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationMarshalled = true

    return json.Marshal(&struct {
        
        Body Documentbody `json:"body"`
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Contexts []Documentvariationcontext `json:"contexts"`
        *Alias
    }{

        


        


        


        


        


        
        Contexts: []Documentvariationcontext{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

