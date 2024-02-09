package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatelimitaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatelimitaggregatequeryclauseDud struct { 
    


    

}

// Ratelimitaggregatequeryclause
type Ratelimitaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Ratelimitaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Ratelimitaggregatequeryclause) String() string {
    
     o.Predicates = []Ratelimitaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratelimitaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Ratelimitaggregatequeryclause

    if RatelimitaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    RatelimitaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Ratelimitaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Ratelimitaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

