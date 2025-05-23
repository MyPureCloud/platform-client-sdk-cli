package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationresponseDud struct { 
    Id string `json:"id"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    Document Knowledgedocumentreference `json:"document"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Documentvariationresponse
type Documentvariationresponse struct { 
    


    


    


    // DocumentVersion - The version of the document.
    DocumentVersion Addressableentityref `json:"documentVersion"`


    // Contexts - The context values associated with the variation.
    Contexts []Documentvariationcontext `json:"contexts"`


    


    // Priority - The priority of the variation.
    Priority int `json:"priority"`


    // Name - The name of the variation.
    Name string `json:"name"`


    // Body - The content for the variation.
    Body Documentbodyresponse `json:"body"`


    

}

// String returns a JSON representation of the model
func (o *Documentvariationresponse) String() string {
    
     o.Contexts = []Documentvariationcontext{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariationresponse) MarshalJSON() ([]byte, error) {
    type Alias Documentvariationresponse

    if DocumentvariationresponseMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationresponseMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Contexts []Documentvariationcontext `json:"contexts"`
        
        Priority int `json:"priority"`
        
        Name string `json:"name"`
        
        Body Documentbodyresponse `json:"body"`
        *Alias
    }{

        


        


        


        


        
        Contexts: []Documentvariationcontext{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

