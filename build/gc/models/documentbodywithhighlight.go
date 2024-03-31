package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodywithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodywithhighlightDud struct { 
    

}

// Documentbodywithhighlight
type Documentbodywithhighlight struct { 
    // Blocks - The list of building blocks for the document body.
    Blocks []Documentbodyblockwithhighlight `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodywithhighlight) String() string {
     o.Blocks = []Documentbodyblockwithhighlight{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodywithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodywithhighlight

    if DocumentbodywithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodywithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        Blocks []Documentbodyblockwithhighlight `json:"blocks"`
        *Alias
    }{

        
        Blocks: []Documentbodyblockwithhighlight{{}},
        

        Alias: (*Alias)(u),
    })
}

