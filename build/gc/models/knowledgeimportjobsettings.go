package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeimportjobsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeimportjobsettingsDud struct { 
    


    


    


    

}

// Knowledgeimportjobsettings
type Knowledgeimportjobsettings struct { 
    // ImportAsNew - If enabled import creates a new document even if update is available.
    ImportAsNew bool `json:"importAsNew"`


    // Visible - If specified, import will override the visibility of the imported documents.
    Visible bool `json:"visible"`


    // CategoryId - If specified, import will override the category of the imported documents.
    CategoryId string `json:"categoryId"`


    // LabelIds - If specified, import will add this labels to the imported documents.
    LabelIds []string `json:"labelIds"`

}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobsettings) String() string {
    
    
    
     o.LabelIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeimportjobsettings) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeimportjobsettings

    if KnowledgeimportjobsettingsMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeimportjobsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ImportAsNew bool `json:"importAsNew"`
        
        Visible bool `json:"visible"`
        
        CategoryId string `json:"categoryId"`
        
        LabelIds []string `json:"labelIds"`
        *Alias
    }{

        


        


        


        
        LabelIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

