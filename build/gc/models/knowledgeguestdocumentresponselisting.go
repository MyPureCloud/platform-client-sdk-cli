package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentresponselistingDud struct { 
    


    


    


    

}

// Knowledgeguestdocumentresponselisting
type Knowledgeguestdocumentresponselisting struct { 
    // Entities
    Entities []Knowledgeguestdocumentresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentresponselisting) String() string {
     o.Entities = []Knowledgeguestdocumentresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentresponselisting

    if KnowledgeguestdocumentresponselistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgeguestdocumentresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgeguestdocumentresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

