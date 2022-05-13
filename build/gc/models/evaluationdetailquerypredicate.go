package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationdetailquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationdetailquerypredicateDud struct { 
    


    


    


    


    


    

}

// Evaluationdetailquerypredicate
type Evaluationdetailquerypredicate struct { 
    // VarType - Optional type, can usually be inferred
    VarType string `json:"type"`


    // Dimension - Left hand side for dimension predicates
    Dimension string `json:"dimension"`


    // Metric - Left hand side for metric predicates
    Metric string `json:"metric"`


    // Operator - Optional operator, default is matches
    Operator string `json:"operator"`


    // Value - Right hand side for dimension or metric predicates
    Value string `json:"value"`


    // VarRange - Right hand side for dimension or metric predicates
    VarRange Numericrange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Evaluationdetailquerypredicate) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationdetailquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Evaluationdetailquerypredicate

    if EvaluationdetailquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    EvaluationdetailquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        Metric string `json:"metric"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        
        VarRange Numericrange `json:"range"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

