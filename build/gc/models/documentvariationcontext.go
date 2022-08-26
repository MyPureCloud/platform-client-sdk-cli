package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationcontextDud struct { 
    


    

}

// Documentvariationcontext
type Documentvariationcontext struct { 
    // Context - The knowledge context associated with the variation.
    Context Knowledgecontextreference `json:"context"`


    // Values - The list of knowledge context values associated with the variation.
    Values []Knowledgecontextvaluereference `json:"values"`

}

// String returns a JSON representation of the model
func (o *Documentvariationcontext) String() string {
    
     o.Values = []Knowledgecontextvaluereference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariationcontext) MarshalJSON() ([]byte, error) {
    type Alias Documentvariationcontext

    if DocumentvariationcontextMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationcontextMarshalled = true

    return json.Marshal(&struct {
        
        Context Knowledgecontextreference `json:"context"`
        
        Values []Knowledgecontextvaluereference `json:"values"`
        *Alias
    }{

        


        
        Values: []Knowledgecontextvaluereference{{}},
        

        Alias: (*Alias)(u),
    })
}

