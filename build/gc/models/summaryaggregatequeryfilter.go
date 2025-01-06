package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummaryaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummaryaggregatequeryfilterDud struct { 
    


    


    

}

// Summaryaggregatequeryfilter
type Summaryaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Summaryaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Summaryaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Summaryaggregatequeryfilter) String() string {
    
     o.Clauses = []Summaryaggregatequeryclause{{}} 
     o.Predicates = []Summaryaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summaryaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Summaryaggregatequeryfilter

    if SummaryaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    SummaryaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Summaryaggregatequeryclause `json:"clauses"`
        
        Predicates []Summaryaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Summaryaggregatequeryclause{{}},
        


        
        Predicates: []Summaryaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

