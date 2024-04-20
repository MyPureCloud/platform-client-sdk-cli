package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperintervalDud struct { 
    

}

// Listwrapperinterval
type Listwrapperinterval struct { 
    // Values - Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperinterval) String() string {
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperinterval) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperinterval

    if ListwrapperintervalMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperintervalMarshalled = true

    return json.Marshal(&struct {
        
        Values []string `json:"values"`
        *Alias
    }{

        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

