package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablecaptionblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablecaptionblockDud struct { 
    

}

// Documentbodytablecaptionblock
type Documentbodytablecaptionblock struct { 
    // Blocks - The list of building blocks for the caption property.
    Blocks []Documentbodytablecaptionitem `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablecaptionblock) String() string {
     o.Blocks = []Documentbodytablecaptionitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablecaptionblock) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablecaptionblock

    if DocumentbodytablecaptionblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablecaptionblockMarshalled = true

    return json.Marshal(&struct {
        
        Blocks []Documentbodytablecaptionitem `json:"blocks"`
        *Alias
    }{

        
        Blocks: []Documentbodytablecaptionitem{{}},
        

        Alias: (*Alias)(u),
    })
}

