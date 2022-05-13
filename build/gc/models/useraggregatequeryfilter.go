package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseraggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseraggregatequeryfilterDud struct { 
    


    


    

}

// Useraggregatequeryfilter
type Useraggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Useraggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Useraggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Useraggregatequeryfilter) String() string {
    
     o.Clauses = []Useraggregatequeryclause{{}} 
     o.Predicates = []Useraggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useraggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Useraggregatequeryfilter

    if UseraggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    UseraggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Useraggregatequeryclause `json:"clauses"`
        
        Predicates []Useraggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Useraggregatequeryclause{{}},
        


        
        Predicates: []Useraggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

