package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodylistblockwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodylistblockwithhighlightDud struct { 
    


    


    

}

// Documentbodylistblockwithhighlight
type Documentbodylistblockwithhighlight struct { 
    // VarType - The type of the list block.
    VarType string `json:"type"`


    // Properties - The properties for the list block.
    Properties Documentbodylistitemproperties `json:"properties"`


    // Blocks - The list of items for an OrderedList or an UnorderedList.
    Blocks []Documentlistcontentblockwithhighlight `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodylistblockwithhighlight) String() string {
    
    
     o.Blocks = []Documentlistcontentblockwithhighlight{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodylistblockwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodylistblockwithhighlight

    if DocumentbodylistblockwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodylistblockwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Properties Documentbodylistitemproperties `json:"properties"`
        
        Blocks []Documentlistcontentblockwithhighlight `json:"blocks"`
        *Alias
    }{

        


        


        
        Blocks: []Documentlistcontentblockwithhighlight{{}},
        

        Alias: (*Alias)(u),
    })
}

