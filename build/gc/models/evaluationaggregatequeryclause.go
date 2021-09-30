package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationaggregatequeryclauseDud struct { 
    


    

}

// Evaluationaggregatequeryclause
type Evaluationaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Evaluationaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Evaluationaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Evaluationaggregatequeryclause

    if EvaluationaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Evaluationaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Evaluationaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

