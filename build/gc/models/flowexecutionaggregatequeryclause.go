package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionaggregatequeryclauseDud struct { 
    


    

}

// Flowexecutionaggregatequeryclause
type Flowexecutionaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowexecutionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionaggregatequeryclause) String() string {
    
     o.Predicates = []Flowexecutionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionaggregatequeryclause

    if FlowexecutionaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Flowexecutionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Flowexecutionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

