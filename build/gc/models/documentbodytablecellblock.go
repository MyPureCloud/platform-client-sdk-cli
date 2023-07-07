package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablecellblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablecellblockDud struct { 
    


    

}

// Documentbodytablecellblock
type Documentbodytablecellblock struct { 
    // Blocks - The list of content blocks for the table.
    Blocks []Documenttablecontentblock `json:"blocks"`


    // Properties - The properties for the table cell.
    Properties Documentbodytablecellblockproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablecellblock) String() string {
     o.Blocks = []Documenttablecontentblock{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablecellblock) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablecellblock

    if DocumentbodytablecellblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablecellblockMarshalled = true

    return json.Marshal(&struct {
        
        Blocks []Documenttablecontentblock `json:"blocks"`
        
        Properties Documentbodytablecellblockproperties `json:"properties"`
        *Alias
    }{

        
        Blocks: []Documenttablecontentblock{{}},
        


        

        Alias: (*Alias)(u),
    })
}

