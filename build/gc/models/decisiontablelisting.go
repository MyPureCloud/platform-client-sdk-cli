package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablelistingDud struct { 
    


    


    


    

}

// Decisiontablelisting
type Decisiontablelisting struct { 
    // Entities
    Entities []Decisiontable `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Decisiontablelisting) String() string {
     o.Entities = []Decisiontable{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablelisting) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablelisting

    if DecisiontablelistingMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Decisiontable `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Decisiontable{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

