package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcetagresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcetagresponseDud struct { 
    

}

// V3sourcetagresponse
type V3sourcetagresponse struct { 
    // Name - Tag name.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *V3sourcetagresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcetagresponse) MarshalJSON() ([]byte, error) {
    type Alias V3sourcetagresponse

    if V3sourcetagresponseMarshalled {
        return []byte("{}"), nil
    }
    V3sourcetagresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

