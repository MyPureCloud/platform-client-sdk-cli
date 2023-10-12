package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeaggregatequeryclauseDud struct { 
    


    

}

// Knowledgeaggregatequeryclause
type Knowledgeaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Knowledgeaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Knowledgeaggregatequeryclause) String() string {
    
     o.Predicates = []Knowledgeaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeaggregatequeryclause

    if KnowledgeaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Knowledgeaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Knowledgeaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

