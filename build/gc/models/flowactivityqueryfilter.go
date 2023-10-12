package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowactivityqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowactivityqueryfilterDud struct { 
    


    


    

}

// Flowactivityqueryfilter
type Flowactivityqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Flowactivityqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowactivityqueryfilter) String() string {
    
     o.Clauses = []Flowactivityqueryclause{{}} 
     o.Predicates = []Flowactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowactivityqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Flowactivityqueryfilter

    if FlowactivityqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    FlowactivityqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Flowactivityqueryclause `json:"clauses"`
        
        Predicates []Flowactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Flowactivityqueryclause{{}},
        


        
        Predicates: []Flowactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

