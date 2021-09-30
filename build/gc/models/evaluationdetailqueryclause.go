package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationdetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationdetailqueryclauseDud struct { 
    


    

}

// Evaluationdetailqueryclause
type Evaluationdetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Evaluationdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Evaluationdetailqueryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Evaluationdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationdetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Evaluationdetailqueryclause

    if EvaluationdetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationdetailqueryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Evaluationdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Evaluationdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

