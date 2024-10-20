package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablecellblockwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablecellblockwithhighlightDud struct { 
    


    

}

// Documentbodytablecellblockwithhighlight
type Documentbodytablecellblockwithhighlight struct { 
    // Properties - The properties for the table cell.
    Properties Documentbodytablecellblockproperties `json:"properties"`


    // Blocks - The list of content blocks for the table.
    Blocks []Documenttablecontentblockwithhighlight `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablecellblockwithhighlight) String() string {
    
     o.Blocks = []Documenttablecontentblockwithhighlight{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablecellblockwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablecellblockwithhighlight

    if DocumentbodytablecellblockwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablecellblockwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        Properties Documentbodytablecellblockproperties `json:"properties"`
        
        Blocks []Documenttablecontentblockwithhighlight `json:"blocks"`
        *Alias
    }{

        


        
        Blocks: []Documenttablecontentblockwithhighlight{{}},
        

        Alias: (*Alias)(u),
    })
}

