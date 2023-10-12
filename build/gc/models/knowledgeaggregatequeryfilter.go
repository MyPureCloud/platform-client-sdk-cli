package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeaggregatequeryfilterDud struct { 
    


    


    

}

// Knowledgeaggregatequeryfilter
type Knowledgeaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Knowledgeaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Knowledgeaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Knowledgeaggregatequeryfilter) String() string {
    
     o.Clauses = []Knowledgeaggregatequeryclause{{}} 
     o.Predicates = []Knowledgeaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeaggregatequeryfilter

    if KnowledgeaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Knowledgeaggregatequeryclause `json:"clauses"`
        
        Predicates []Knowledgeaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Knowledgeaggregatequeryclause{{}},
        


        
        Predicates: []Knowledgeaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

