package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentvariationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentvariationcontextDud struct { 
    


    

}

// Knowledgeguestdocumentvariationcontext
type Knowledgeguestdocumentvariationcontext struct { 
    // Context - The knowledge context associated with the variation.
    Context Addressableentityref `json:"context"`


    // Values - The list of knowledge context values associated with the variation.
    Values []Addressableentityref `json:"values"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentvariationcontext) String() string {
    
     o.Values = []Addressableentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentvariationcontext) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentvariationcontext

    if KnowledgeguestdocumentvariationcontextMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentvariationcontextMarshalled = true

    return json.Marshal(&struct {
        
        Context Addressableentityref `json:"context"`
        
        Values []Addressableentityref `json:"values"`
        *Alias
    }{

        


        
        Values: []Addressableentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

