package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryrequestpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryrequestpredicateDud struct { 
    


    

}

// Developmentactivityaggregatequeryrequestpredicate
type Developmentactivityaggregatequeryrequestpredicate struct { 
    // Dimension - Each predicates specifies a dimension.
    Dimension string `json:"dimension"`


    // Value - Corresponding value for dimensions in predicates. If the dimension is type, Valid Values: Informational (deprecated), AssessedContent (deprecated), Assessment (deprecated), Coaching, External, Native
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestpredicate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryrequestpredicate) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryrequestpredicate

    if DevelopmentactivityaggregatequeryrequestpredicateMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryrequestpredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

