package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodylistwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodylistwithhighlightDud struct { 
    


    

}

// Documentbodylistwithhighlight
type Documentbodylistwithhighlight struct { 
    // Properties - Properties for the UnorderedList or OrderedList.
    Properties Documentbodylistblockproperties `json:"properties"`


    // Blocks - The list of items for an OrderedList or an UnorderedList.
    Blocks []Documentbodylistblockwithhighlight `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodylistwithhighlight) String() string {
    
     o.Blocks = []Documentbodylistblockwithhighlight{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodylistwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodylistwithhighlight

    if DocumentbodylistwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodylistwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        Properties Documentbodylistblockproperties `json:"properties"`
        
        Blocks []Documentbodylistblockwithhighlight `json:"blocks"`
        *Alias
    }{

        


        
        Blocks: []Documentbodylistblockwithhighlight{{}},
        

        Alias: (*Alias)(u),
    })
}

