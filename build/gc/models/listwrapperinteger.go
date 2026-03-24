package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperintegerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperintegerDud struct { 
    

}

// Listwrapperinteger
type Listwrapperinteger struct { 
    // Values
    Values []int `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperinteger) String() string {
     o.Values = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperinteger) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperinteger

    if ListwrapperintegerMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperintegerMarshalled = true

    return json.Marshal(&struct {
        
        Values []int `json:"values"`
        *Alias
    }{

        
        Values: []int{0},
        

        Alias: (*Alias)(u),
    })
}

