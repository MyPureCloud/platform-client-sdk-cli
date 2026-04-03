package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotaggregatequeryfilterDud struct { 
    


    


    

}

// Copilotaggregatequeryfilter
type Copilotaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Copilotaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Copilotaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Copilotaggregatequeryfilter) String() string {
    
     o.Clauses = []Copilotaggregatequeryclause{{}} 
     o.Predicates = []Copilotaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Copilotaggregatequeryfilter

    if CopilotaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    CopilotaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Copilotaggregatequeryclause `json:"clauses"`
        
        Predicates []Copilotaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Copilotaggregatequeryclause{{}},
        


        
        Predicates: []Copilotaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

