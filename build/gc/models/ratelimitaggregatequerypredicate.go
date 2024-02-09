package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatelimitaggregatequerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatelimitaggregatequerypredicateDud struct { 
    


    


    


    


    

}

// Ratelimitaggregatequerypredicate
type Ratelimitaggregatequerypredicate struct { 
    // VarType - Optional type, can usually be inferred
    VarType string `json:"type"`


    // Dimension - Left hand side for dimension predicates
    Dimension string `json:"dimension"`


    // Operator - Optional operator, default is matches
    Operator string `json:"operator"`


    // Value - Right hand side for dimension predicates
    Value string `json:"value"`


    // VarRange - Right hand side for dimension predicates
    VarRange Numericrange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Ratelimitaggregatequerypredicate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratelimitaggregatequerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Ratelimitaggregatequerypredicate

    if RatelimitaggregatequerypredicateMarshalled {
        return []byte("{}"), nil
    }
    RatelimitaggregatequerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        
        VarRange Numericrange `json:"range"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

