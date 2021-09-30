package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserobservationqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserobservationqueryclauseDud struct { 
    


    

}

// Userobservationqueryclause
type Userobservationqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Userobservationquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Userobservationqueryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Userobservationquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userobservationqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Userobservationqueryclause

    if UserobservationqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    UserobservationqueryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Userobservationquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Userobservationquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

