package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrappershiftstartvarianceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrappershiftstartvarianceDud struct { 
    

}

// Listwrappershiftstartvariance
type Listwrappershiftstartvariance struct { 
    // Values
    Values []Shiftstartvariance `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrappershiftstartvariance) String() string {
    
    
     o.Values = []Shiftstartvariance{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrappershiftstartvariance) MarshalJSON() ([]byte, error) {
    type Alias Listwrappershiftstartvariance

    if ListwrappershiftstartvarianceMarshalled {
        return []byte("{}"), nil
    }
    ListwrappershiftstartvarianceMarshalled = true

    return json.Marshal(&struct { 
        Values []Shiftstartvariance `json:"values"`
        
        *Alias
    }{
        

        
        Values: []Shiftstartvariance{{}},
        

        
        Alias: (*Alias)(u),
    })
}

