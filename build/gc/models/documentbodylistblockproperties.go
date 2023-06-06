package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodylistblockpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodylistblockpropertiesDud struct { 
    


    

}

// Documentbodylistblockproperties
type Documentbodylistblockproperties struct { 
    // UnorderedType - The type of icon for the unordered list.
    UnorderedType string `json:"unorderedType"`


    // OrderedType - The type of icon for the ordered list.
    OrderedType string `json:"orderedType"`

}

// String returns a JSON representation of the model
func (o *Documentbodylistblockproperties) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodylistblockproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodylistblockproperties

    if DocumentbodylistblockpropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodylistblockpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        UnorderedType string `json:"unorderedType"`
        
        OrderedType string `json:"orderedType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

