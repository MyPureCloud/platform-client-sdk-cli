package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodylistblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodylistblockDud struct { 
    


    


    

}

// Documentbodylistblock
type Documentbodylistblock struct { 
    // VarType - The type of the list block.
    VarType string `json:"type"`


    // Properties - The properties for the list block.
    Properties Documentbodylistitemproperties `json:"properties"`


    // Blocks - The list of items for an OrderedList or an UnorderedList.
    Blocks []Documentlistcontentblock `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodylistblock) String() string {
    
    
     o.Blocks = []Documentlistcontentblock{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodylistblock) MarshalJSON() ([]byte, error) {
    type Alias Documentbodylistblock

    if DocumentbodylistblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodylistblockMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Properties Documentbodylistitemproperties `json:"properties"`
        
        Blocks []Documentlistcontentblock `json:"blocks"`
        *Alias
    }{

        


        


        
        Blocks: []Documentlistcontentblock{{}},
        

        Alias: (*Alias)(u),
    })
}

