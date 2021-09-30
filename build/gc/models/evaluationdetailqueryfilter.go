package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationdetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationdetailqueryfilterDud struct { 
    


    


    

}

// Evaluationdetailqueryfilter
type Evaluationdetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Evaluationdetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Evaluationdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Evaluationdetailqueryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Evaluationdetailqueryclause{{}} 
    
    
    
     o.Predicates = []Evaluationdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationdetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Evaluationdetailqueryfilter

    if EvaluationdetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    EvaluationdetailqueryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Evaluationdetailqueryclause `json:"clauses"`
        
        Predicates []Evaluationdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Evaluationdetailqueryclause{{}},
        

        

        
        Predicates: []Evaluationdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

