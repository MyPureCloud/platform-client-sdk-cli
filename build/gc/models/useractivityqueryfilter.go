package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityqueryfilterDud struct { 
    


    


    

}

// Useractivityqueryfilter
type Useractivityqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Useractivityqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Useractivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Useractivityqueryfilter) String() string {
    
     o.Clauses = []Useractivityqueryclause{{}} 
     o.Predicates = []Useractivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Useractivityqueryfilter

    if UseractivityqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    UseractivityqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Useractivityqueryclause `json:"clauses"`
        
        Predicates []Useractivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Useractivityqueryclause{{}},
        


        
        Predicates: []Useractivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

