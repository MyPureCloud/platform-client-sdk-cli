package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyaggregatequeryfilterDud struct { 
    


    


    

}

// Journeyaggregatequeryfilter
type Journeyaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Journeyaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Journeyaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Journeyaggregatequeryclause{{}} 
    
    
    
     o.Predicates = []Journeyaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Journeyaggregatequeryfilter

    if JourneyaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    JourneyaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Journeyaggregatequeryclause `json:"clauses"`
        
        Predicates []Journeyaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Journeyaggregatequeryclause{{}},
        

        

        
        Predicates: []Journeyaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

