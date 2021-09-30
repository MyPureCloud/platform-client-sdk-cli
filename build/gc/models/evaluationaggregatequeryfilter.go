package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationaggregatequeryfilterDud struct { 
    


    


    

}

// Evaluationaggregatequeryfilter
type Evaluationaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Evaluationaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Evaluationaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Evaluationaggregatequeryclause{{}} 
    
    
    
     o.Predicates = []Evaluationaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Evaluationaggregatequeryfilter

    if EvaluationaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    EvaluationaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Evaluationaggregatequeryclause `json:"clauses"`
        
        Predicates []Evaluationaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Evaluationaggregatequeryclause{{}},
        

        

        
        Predicates: []Evaluationaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

