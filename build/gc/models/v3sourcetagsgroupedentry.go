package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcetagsgroupedentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcetagsgroupedentryDud struct { 
    


    

}

// V3sourcetagsgroupedentry
type V3sourcetagsgroupedentry struct { 
    // Source - Knowledge source reference for this entry.
    Source V3sourceref `json:"source"`


    // Tags - Distinct tags for this source.
    Tags []V3sourcetagresponse `json:"tags"`

}

// String returns a JSON representation of the model
func (o *V3sourcetagsgroupedentry) String() string {
    
     o.Tags = []V3sourcetagresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcetagsgroupedentry) MarshalJSON() ([]byte, error) {
    type Alias V3sourcetagsgroupedentry

    if V3sourcetagsgroupedentryMarshalled {
        return []byte("{}"), nil
    }
    V3sourcetagsgroupedentryMarshalled = true

    return json.Marshal(&struct {
        
        Source V3sourceref `json:"source"`
        
        Tags []V3sourcetagresponse `json:"tags"`
        *Alias
    }{

        


        
        Tags: []V3sourcetagresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

