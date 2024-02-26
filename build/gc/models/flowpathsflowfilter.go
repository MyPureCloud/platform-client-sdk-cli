package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowpathsflowfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowpathsflowfilterDud struct { 
    

}

// Flowpathsflowfilter
type Flowpathsflowfilter struct { 
    // Id - The identifier of the flow.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Flowpathsflowfilter) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowpathsflowfilter) MarshalJSON() ([]byte, error) {
    type Alias Flowpathsflowfilter

    if FlowpathsflowfilterMarshalled {
        return []byte("{}"), nil
    }
    FlowpathsflowfilterMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

