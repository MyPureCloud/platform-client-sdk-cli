package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicenowsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicenowsettingsDud struct { 
    


    


    


    

}

// Servicenowsettings
type Servicenowsettings struct { 
    // KnowledgeBaseIds - Filter source by knowledge base ids.
    KnowledgeBaseIds []string `json:"knowledgeBaseIds"`


    // Language - Filter source by language.
    Language string `json:"language"`


    // Categories - Filter source by categories.
    Categories []string `json:"categories"`


    // BaseUrl - The base URL to resources.
    BaseUrl string `json:"baseUrl"`

}

// String returns a JSON representation of the model
func (o *Servicenowsettings) String() string {
     o.KnowledgeBaseIds = []string{""} 
    
     o.Categories = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicenowsettings) MarshalJSON() ([]byte, error) {
    type Alias Servicenowsettings

    if ServicenowsettingsMarshalled {
        return []byte("{}"), nil
    }
    ServicenowsettingsMarshalled = true

    return json.Marshal(&struct {
        
        KnowledgeBaseIds []string `json:"knowledgeBaseIds"`
        
        Language string `json:"language"`
        
        Categories []string `json:"categories"`
        
        BaseUrl string `json:"baseUrl"`
        *Alias
    }{

        
        KnowledgeBaseIds: []string{""},
        


        


        
        Categories: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

