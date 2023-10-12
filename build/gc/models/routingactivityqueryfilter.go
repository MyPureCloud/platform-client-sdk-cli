package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivityqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivityqueryfilterDud struct { 
    


    


    

}

// Routingactivityqueryfilter
type Routingactivityqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Routingactivityqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Routingactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Routingactivityqueryfilter) String() string {
    
     o.Clauses = []Routingactivityqueryclause{{}} 
     o.Predicates = []Routingactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivityqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Routingactivityqueryfilter

    if RoutingactivityqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivityqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Routingactivityqueryclause `json:"clauses"`
        
        Predicates []Routingactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Routingactivityqueryclause{{}},
        


        
        Predicates: []Routingactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

