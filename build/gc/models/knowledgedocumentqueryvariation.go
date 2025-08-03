package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentqueryvariationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentqueryvariationDud struct { 
    Id string `json:"id"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    Document Knowledgedocumentreference `json:"document"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentqueryvariation
type Knowledgedocumentqueryvariation struct { 
    


    


    


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


    // Chunks - The chunk blocks associated with the variation.
    Chunks []Documentvariationquerychunkblock `json:"chunks"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentqueryvariation) String() string {
    
     o.Contexts = []Documentvariationcontext{{}} 
    
    
    
     o.Chunks = []Documentvariationquerychunkblock{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentqueryvariation) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentqueryvariation

    if KnowledgedocumentqueryvariationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentqueryvariationMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Contexts []Documentvariationcontext `json:"contexts"`
        
        Priority int `json:"priority"`
        
        Name string `json:"name"`
        
        Body Documentbodyresponse `json:"body"`
        
        Chunks []Documentvariationquerychunkblock `json:"chunks"`
        *Alias
    }{

        


        


        


        


        
        Contexts: []Documentvariationcontext{{}},
        


        


        


        


        


        
        Chunks: []Documentvariationquerychunkblock{{}},
        


        

        Alias: (*Alias)(u),
    })
}

