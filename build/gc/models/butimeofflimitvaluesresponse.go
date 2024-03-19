package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitvaluesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitvaluesresponseDud struct { 
    

}

// Butimeofflimitvaluesresponse
type Butimeofflimitvaluesresponse struct { 
    // Values
    Values []Butimeofflimitvaluerange `json:"values"`

}

// String returns a JSON representation of the model
func (o *Butimeofflimitvaluesresponse) String() string {
     o.Values = []Butimeofflimitvaluerange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitvaluesresponse) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitvaluesresponse

    if ButimeofflimitvaluesresponseMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitvaluesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Values []Butimeofflimitvaluerange `json:"values"`
        *Alias
    }{

        
        Values: []Butimeofflimitvaluerange{{}},
        

        Alias: (*Alias)(u),
    })
}

