package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PrefixlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PrefixlistingDud struct { 
    


    


    


    

}

// Prefixlisting
type Prefixlisting struct { 
    // Entities
    Entities []Prefixlistingitem `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Prefixlisting) String() string {
     o.Entities = []Prefixlistingitem{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Prefixlisting) MarshalJSON() ([]byte, error) {
    type Alias Prefixlisting

    if PrefixlistingMarshalled {
        return []byte("{}"), nil
    }
    PrefixlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Prefixlistingitem `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Prefixlistingitem{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

