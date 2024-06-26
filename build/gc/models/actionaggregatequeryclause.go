package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionaggregatequeryclauseDud struct { 
    


    

}

// Actionaggregatequeryclause
type Actionaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Actionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Actionaggregatequeryclause) String() string {
    
     o.Predicates = []Actionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Actionaggregatequeryclause

    if ActionaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    ActionaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Actionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Actionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

