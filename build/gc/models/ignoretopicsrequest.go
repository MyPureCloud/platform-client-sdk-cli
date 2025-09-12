package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnoretopicsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnoretopicsrequestDud struct { 
    

}

// Ignoretopicsrequest
type Ignoretopicsrequest struct { 
    // Topics - List of topics to be ignored
    Topics []Ignoretopic `json:"topics"`

}

// String returns a JSON representation of the model
func (o *Ignoretopicsrequest) String() string {
     o.Topics = []Ignoretopic{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignoretopicsrequest) MarshalJSON() ([]byte, error) {
    type Alias Ignoretopicsrequest

    if IgnoretopicsrequestMarshalled {
        return []byte("{}"), nil
    }
    IgnoretopicsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Topics []Ignoretopic `json:"topics"`
        *Alias
    }{

        
        Topics: []Ignoretopic{{}},
        

        Alias: (*Alias)(u),
    })
}

