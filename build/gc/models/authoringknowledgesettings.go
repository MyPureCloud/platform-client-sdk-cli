package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthoringknowledgesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthoringknowledgesettingsDud struct { 
    


    


    


    

}

// Authoringknowledgesettings
type Authoringknowledgesettings struct { 
    // KnowledgeId - The ID of the knowledge settings to use
    KnowledgeId string `json:"knowledgeId"`


    // Description - Description of the knowledge settings
    Description string `json:"description"`


    // KnowledgeType - The type of the knowledge settings (Standard or Enhanced)
    KnowledgeType string `json:"knowledgeType"`


    // InheritFromExternal - Whether to inherit knowledge settings from externalIf disabled, query detection is always enabled on the guide.
    InheritFromExternal bool `json:"inheritFromExternal"`

}

// String returns a JSON representation of the model
func (o *Authoringknowledgesettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authoringknowledgesettings) MarshalJSON() ([]byte, error) {
    type Alias Authoringknowledgesettings

    if AuthoringknowledgesettingsMarshalled {
        return []byte("{}"), nil
    }
    AuthoringknowledgesettingsMarshalled = true

    return json.Marshal(&struct {
        
        KnowledgeId string `json:"knowledgeId"`
        
        Description string `json:"description"`
        
        KnowledgeType string `json:"knowledgeType"`
        
        InheritFromExternal bool `json:"inheritFromExternal"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

