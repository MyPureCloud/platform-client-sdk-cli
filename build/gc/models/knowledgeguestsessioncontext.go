package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestsessioncontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestsessioncontextDud struct { 
    


    

}

// Knowledgeguestsessioncontext
type Knowledgeguestsessioncontext struct { 
    // Id - The context id associated with the session.
    Id string `json:"id"`


    // Values - The list of knowledge context values associated with the session.
    Values []Entity `json:"values"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestsessioncontext) String() string {
    
     o.Values = []Entity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestsessioncontext) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestsessioncontext

    if KnowledgeguestsessioncontextMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestsessioncontextMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Values []Entity `json:"values"`
        *Alias
    }{

        


        
        Values: []Entity{{}},
        

        Alias: (*Alias)(u),
    })
}

