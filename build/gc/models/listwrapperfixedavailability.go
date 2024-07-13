package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperfixedavailabilityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperfixedavailabilityDud struct { 
    

}

// Listwrapperfixedavailability
type Listwrapperfixedavailability struct { 
    // Values
    Values []Fixedavailability `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperfixedavailability) String() string {
     o.Values = []Fixedavailability{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperfixedavailability) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperfixedavailability

    if ListwrapperfixedavailabilityMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperfixedavailabilityMarshalled = true

    return json.Marshal(&struct {
        
        Values []Fixedavailability `json:"values"`
        *Alias
    }{

        
        Values: []Fixedavailability{{}},
        

        Alias: (*Alias)(u),
    })
}

