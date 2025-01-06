package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterlabelfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterlabelfilterDud struct { 
    

}

// Supportcenterlabelfilter
type Supportcenterlabelfilter struct { 
    // Labels - Labels to filter by.
    Labels []Addressableentityref `json:"labels"`

}

// String returns a JSON representation of the model
func (o *Supportcenterlabelfilter) String() string {
     o.Labels = []Addressableentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterlabelfilter) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterlabelfilter

    if SupportcenterlabelfilterMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterlabelfilterMarshalled = true

    return json.Marshal(&struct {
        
        Labels []Addressableentityref `json:"labels"`
        *Alias
    }{

        
        Labels: []Addressableentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

