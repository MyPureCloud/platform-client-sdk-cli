package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserobservationqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserobservationqueryfilterDud struct { 
    


    


    

}

// Userobservationqueryfilter
type Userobservationqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Userobservationqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Userobservationquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Userobservationqueryfilter) String() string {
    
     o.Clauses = []Userobservationqueryclause{{}} 
     o.Predicates = []Userobservationquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userobservationqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Userobservationqueryfilter

    if UserobservationqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    UserobservationqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Userobservationqueryclause `json:"clauses"`
        
        Predicates []Userobservationquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Userobservationqueryclause{{}},
        


        
        Predicates: []Userobservationquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

