package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingstatusdetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingstatusdetailqueryclauseDud struct { 
    


    

}

// Routingstatusdetailqueryclause
type Routingstatusdetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Routingstatusdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Routingstatusdetailqueryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Routingstatusdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingstatusdetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Routingstatusdetailqueryclause

    if RoutingstatusdetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    RoutingstatusdetailqueryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Routingstatusdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Routingstatusdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

