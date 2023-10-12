package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowactivityqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowactivityqueryclauseDud struct { 
    


    

}

// Flowactivityqueryclause
type Flowactivityqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowactivityqueryclause) String() string {
    
     o.Predicates = []Flowactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowactivityqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Flowactivityqueryclause

    if FlowactivityqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    FlowactivityqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Flowactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Flowactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

