package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyresponseDud struct { 
    

}

// Documentbodyresponse
type Documentbodyresponse struct { 
    // Blocks - The list of building blocks for the document body.
    Blocks []Documentbodyblock `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodyresponse) String() string {
     o.Blocks = []Documentbodyblock{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyresponse) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyresponse

    if DocumentbodyresponseMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyresponseMarshalled = true

    return json.Marshal(&struct {
        
        Blocks []Documentbodyblock `json:"blocks"`
        *Alias
    }{

        
        Blocks: []Documentbodyblock{{}},
        

        Alias: (*Alias)(u),
    })
}

