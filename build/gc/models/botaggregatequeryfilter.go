package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotaggregatequeryfilterDud struct { 
    


    


    

}

// Botaggregatequeryfilter
type Botaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Botaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Botaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Botaggregatequeryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Botaggregatequeryclause{{}} 
    
    
    
     o.Predicates = []Botaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Botaggregatequeryfilter

    if BotaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    BotaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Botaggregatequeryclause `json:"clauses"`
        
        Predicates []Botaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Botaggregatequeryclause{{}},
        

        

        
        Predicates: []Botaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

