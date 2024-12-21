package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummaryaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummaryaggregatequeryclauseDud struct { 
    


    

}

// Summaryaggregatequeryclause
type Summaryaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Summaryaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Summaryaggregatequeryclause) String() string {
    
     o.Predicates = []Summaryaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summaryaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Summaryaggregatequeryclause

    if SummaryaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    SummaryaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Summaryaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Summaryaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

