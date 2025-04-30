package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperoverridedateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperoverridedateDud struct { 
    

}

// Listwrapperoverridedate
type Listwrapperoverridedate struct { 
    // Values
    Values []Overridedate `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperoverridedate) String() string {
     o.Values = []Overridedate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperoverridedate) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperoverridedate

    if ListwrapperoverridedateMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperoverridedateMarshalled = true

    return json.Marshal(&struct {
        
        Values []Overridedate `json:"values"`
        *Alias
    }{

        
        Values: []Overridedate{{}},
        

        Alias: (*Alias)(u),
    })
}

