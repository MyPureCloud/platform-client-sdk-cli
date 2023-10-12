package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivityqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivityqueryclauseDud struct { 
    


    

}

// Routingactivityqueryclause
type Routingactivityqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Routingactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Routingactivityqueryclause) String() string {
    
     o.Predicates = []Routingactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivityqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Routingactivityqueryclause

    if RoutingactivityqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivityqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Routingactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Routingactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

