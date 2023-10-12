package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentbulkupdateentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentbulkupdateentityDud struct { 
    


    


    

}

// Knowledgedocumentbulkupdateentity
type Knowledgedocumentbulkupdateentity struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // CategoryId - The category associated with the document.
    CategoryId string `json:"categoryId"`


    // LabelIds - The ids of labels associated with the document.
    LabelIds []string `json:"labelIds"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkupdateentity) String() string {
    
    
     o.LabelIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentbulkupdateentity) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentbulkupdateentity

    if KnowledgedocumentbulkupdateentityMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentbulkupdateentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        CategoryId string `json:"categoryId"`
        
        LabelIds []string `json:"labelIds"`
        *Alias
    }{

        


        


        
        LabelIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

