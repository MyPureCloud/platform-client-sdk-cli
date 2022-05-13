package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerytimeofflimitvaluesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerytimeofflimitvaluesresponseDud struct { 
    

}

// Querytimeofflimitvaluesresponse - The list of date ranges with the time off limit, allocated and waitlisted values.
type Querytimeofflimitvaluesresponse struct { 
    // Values
    Values []Timeofflimitvaluerange `json:"values"`

}

// String returns a JSON representation of the model
func (o *Querytimeofflimitvaluesresponse) String() string {
     o.Values = []Timeofflimitvaluerange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querytimeofflimitvaluesresponse) MarshalJSON() ([]byte, error) {
    type Alias Querytimeofflimitvaluesresponse

    if QuerytimeofflimitvaluesresponseMarshalled {
        return []byte("{}"), nil
    }
    QuerytimeofflimitvaluesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Values []Timeofflimitvaluerange `json:"values"`
        *Alias
    }{

        
        Values: []Timeofflimitvaluerange{{}},
        

        Alias: (*Alias)(u),
    })
}

