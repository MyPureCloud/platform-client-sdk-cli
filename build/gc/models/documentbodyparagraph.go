package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyparagraphMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyparagraphDud struct { 
    


    

}

// Documentbodyparagraph
type Documentbodyparagraph struct { 
    // Blocks - The list of blocks for the paragraph.
    Blocks []Documentcontentblock `json:"blocks"`


    // Properties - The properties for the paragraph.
    Properties Documentbodyparagraphproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Documentbodyparagraph) String() string {
     o.Blocks = []Documentcontentblock{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyparagraph) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyparagraph

    if DocumentbodyparagraphMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyparagraphMarshalled = true

    return json.Marshal(&struct {
        
        Blocks []Documentcontentblock `json:"blocks"`
        
        Properties Documentbodyparagraphproperties `json:"properties"`
        *Alias
    }{

        
        Blocks: []Documentcontentblock{{}},
        


        

        Alias: (*Alias)(u),
    })
}

