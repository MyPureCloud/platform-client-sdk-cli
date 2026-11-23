package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KpiattributionaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KpiattributionaggregatequeryfilterDud struct { 
    


    


    

}

// Kpiattributionaggregatequeryfilter
type Kpiattributionaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Kpiattributionaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Kpiattributionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Kpiattributionaggregatequeryfilter) String() string {
    
     o.Clauses = []Kpiattributionaggregatequeryclause{{}} 
     o.Predicates = []Kpiattributionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Kpiattributionaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Kpiattributionaggregatequeryfilter

    if KpiattributionaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    KpiattributionaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Kpiattributionaggregatequeryclause `json:"clauses"`
        
        Predicates []Kpiattributionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Kpiattributionaggregatequeryclause{{}},
        


        
        Predicates: []Kpiattributionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

