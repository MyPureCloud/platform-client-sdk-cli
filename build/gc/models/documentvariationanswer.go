package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationanswerDud struct { 
    Id string `json:"id"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    Document Knowledgedocumentreference `json:"document"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Documentvariationanswer
type Documentvariationanswer struct { 
    


    


    


    // DocumentVersion - The version of the document.
    DocumentVersion Addressableentityref `json:"documentVersion"`


    // Contexts - The context values associated with the variation.
    Contexts []Documentvariationcontext `json:"contexts"`


    


    // Priority - The priority of the variation.
    Priority int `json:"priority"`


    // Name - The name of the variation.
    Name string `json:"name"`


    // Body - The content for the variation.
    Body Documentbodywithhighlight `json:"body"`


    

}

// String returns a JSON representation of the model
func (o *Documentvariationanswer) String() string {
    
     o.Contexts = []Documentvariationcontext{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariationanswer) MarshalJSON() ([]byte, error) {
    type Alias Documentvariationanswer

    if DocumentvariationanswerMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationanswerMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Contexts []Documentvariationcontext `json:"contexts"`
        
        Priority int `json:"priority"`
        
        Name string `json:"name"`
        
        Body Documentbodywithhighlight `json:"body"`
        *Alias
    }{

        


        


        


        


        
        Contexts: []Documentvariationcontext{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

