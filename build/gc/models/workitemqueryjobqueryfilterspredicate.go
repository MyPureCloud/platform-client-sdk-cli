package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueryjobqueryfilterspredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueryjobqueryfilterspredicateDud struct { 
    


    


    

}

// Workitemqueryjobqueryfilterspredicate
type Workitemqueryjobqueryfilterspredicate struct { 
    // Name - Property name.
    Name string `json:"name"`


    // Operator - Query filter predicate operator.
    Operator string `json:"operator"`


    // Values - List of values to be used in the query filter predicate.
    Values []interface{} `json:"values"`

}

// String returns a JSON representation of the model
func (o *Workitemqueryjobqueryfilterspredicate) String() string {
    
    
     o.Values = []interface{}{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueryjobqueryfilterspredicate) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueryjobqueryfilterspredicate

    if WorkitemqueryjobqueryfilterspredicateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueryjobqueryfilterspredicateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Operator string `json:"operator"`
        
        Values []interface{} `json:"values"`
        *Alias
    }{

        


        


        
        Values: []interface{}{},
        

        Alias: (*Alias)(u),
    })
}

