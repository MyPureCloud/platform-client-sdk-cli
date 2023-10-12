package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResolutionaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResolutionaggregatequeryfilterDud struct { 
    


    


    

}

// Resolutionaggregatequeryfilter
type Resolutionaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Resolutionaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Resolutionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Resolutionaggregatequeryfilter) String() string {
    
     o.Clauses = []Resolutionaggregatequeryclause{{}} 
     o.Predicates = []Resolutionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resolutionaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Resolutionaggregatequeryfilter

    if ResolutionaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    ResolutionaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Resolutionaggregatequeryclause `json:"clauses"`
        
        Predicates []Resolutionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Resolutionaggregatequeryclause{{}},
        


        
        Predicates: []Resolutionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

