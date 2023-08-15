package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentfeedbackresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentfeedbackresponselistingDud struct { 
    


    


    


    

}

// Knowledgedocumentfeedbackresponselisting
type Knowledgedocumentfeedbackresponselisting struct { 
    // Entities
    Entities []Knowledgedocumentfeedbackresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentfeedbackresponselisting) String() string {
     o.Entities = []Knowledgedocumentfeedbackresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentfeedbackresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentfeedbackresponselisting

    if KnowledgedocumentfeedbackresponselistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentfeedbackresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgedocumentfeedbackresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgedocumentfeedbackresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

