package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyaggregatequeryclauseDud struct { 
    


    

}

// Journeyaggregatequeryclause
type Journeyaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Journeyaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryclause) String() string {
    
     o.Predicates = []Journeyaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Journeyaggregatequeryclause

    if JourneyaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    JourneyaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Journeyaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Journeyaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

