package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmintegrationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmintegrationlistingDud struct { 
    

}

// Wfmintegrationlisting
type Wfmintegrationlisting struct { 
    // Entities
    Entities []Wfmintegrationresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Wfmintegrationlisting) String() string {
     o.Entities = []Wfmintegrationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmintegrationlisting) MarshalJSON() ([]byte, error) {
    type Alias Wfmintegrationlisting

    if WfmintegrationlistingMarshalled {
        return []byte("{}"), nil
    }
    WfmintegrationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Wfmintegrationresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Wfmintegrationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

