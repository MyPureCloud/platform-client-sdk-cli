package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingstatusdetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingstatusdetailqueryfilterDud struct { 
    


    


    

}

// Routingstatusdetailqueryfilter
type Routingstatusdetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Routingstatusdetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Routingstatusdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Routingstatusdetailqueryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Routingstatusdetailqueryclause{{}} 
    
    
    
     o.Predicates = []Routingstatusdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingstatusdetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Routingstatusdetailqueryfilter

    if RoutingstatusdetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    RoutingstatusdetailqueryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Routingstatusdetailqueryclause `json:"clauses"`
        
        Predicates []Routingstatusdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Routingstatusdetailqueryclause{{}},
        

        

        
        Predicates: []Routingstatusdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

