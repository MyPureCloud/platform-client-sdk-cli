package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueobservationqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueobservationqueryclauseDud struct { 
    


    

}

// Queueobservationqueryclause
type Queueobservationqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Queueobservationquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Queueobservationqueryclause) String() string {
    
     o.Predicates = []Queueobservationquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueobservationqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Queueobservationqueryclause

    if QueueobservationqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    QueueobservationqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Queueobservationquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Queueobservationquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

