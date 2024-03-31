package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyparagraphwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyparagraphwithhighlightDud struct { 
    


    

}

// Documentbodyparagraphwithhighlight
type Documentbodyparagraphwithhighlight struct { 
    // Blocks - The list of blocks for the paragraph.
    Blocks []Documentcontentblockwithhighlight `json:"blocks"`


    // Properties - The properties for the paragraph.
    Properties Documentbodyparagraphproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Documentbodyparagraphwithhighlight) String() string {
     o.Blocks = []Documentcontentblockwithhighlight{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyparagraphwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyparagraphwithhighlight

    if DocumentbodyparagraphwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyparagraphwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        Blocks []Documentcontentblockwithhighlight `json:"blocks"`
        
        Properties Documentbodyparagraphproperties `json:"properties"`
        *Alias
    }{

        
        Blocks: []Documentcontentblockwithhighlight{{}},
        


        

        Alias: (*Alias)(u),
    })
}

