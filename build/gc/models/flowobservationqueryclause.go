package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowobservationqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowobservationqueryclauseDud struct { 
    


    

}

// Flowobservationqueryclause
type Flowobservationqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowobservationquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowobservationqueryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Flowobservationquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowobservationqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Flowobservationqueryclause

    if FlowobservationqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    FlowobservationqueryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Flowobservationquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Flowobservationquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

