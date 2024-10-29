package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentvariationresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentvariationresponselistingDud struct { 
    


    


    


    

}

// Documentvariationresponselisting
type Documentvariationresponselisting struct { 
    // Entities
    Entities []Documentvariationresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Documentvariationresponselisting) String() string {
     o.Entities = []Documentvariationresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentvariationresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Documentvariationresponselisting

    if DocumentvariationresponselistingMarshalled {
        return []byte("{}"), nil
    }
    DocumentvariationresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Documentvariationresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Documentvariationresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

