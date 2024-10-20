package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodylistDud struct { 
    


    

}

// Documentbodylist
type Documentbodylist struct { 
    // Properties - Properties for the UnorderedList or OrderedList.
    Properties Documentbodylistblockproperties `json:"properties"`


    // Blocks - The list of items for an OrderedList or an UnorderedList.
    Blocks []Documentbodylistblock `json:"blocks"`

}

// String returns a JSON representation of the model
func (o *Documentbodylist) String() string {
    
     o.Blocks = []Documentbodylistblock{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodylist) MarshalJSON() ([]byte, error) {
    type Alias Documentbodylist

    if DocumentbodylistMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodylistMarshalled = true

    return json.Marshal(&struct {
        
        Properties Documentbodylistblockproperties `json:"properties"`
        
        Blocks []Documentbodylistblock `json:"blocks"`
        *Alias
    }{

        


        
        Blocks: []Documentbodylistblock{{}},
        

        Alias: (*Alias)(u),
    })
}

