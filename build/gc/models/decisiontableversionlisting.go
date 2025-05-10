package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableversionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableversionlistingDud struct { 
    


    


    


    

}

// Decisiontableversionlisting
type Decisiontableversionlisting struct { 
    // Entities
    Entities []Decisiontableversion `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Decisiontableversionlisting) String() string {
     o.Entities = []Decisiontableversion{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableversionlisting) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableversionlisting

    if DecisiontableversionlistingMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableversionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Decisiontableversion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Decisiontableversion{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

