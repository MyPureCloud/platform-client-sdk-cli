package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesettingsDud struct { 
    


    


    

}

// Knowledgesettings
type Knowledgesettings struct { 
    // KnowledgeId - The ID of the knowledge settings to use
    KnowledgeId string `json:"knowledgeId"`


    // Description - Description of the knowledge settings
    Description string `json:"description"`


    // KnowledgeType - The type of the knowledge settings (Standard or Enhanced)
    KnowledgeType string `json:"knowledgeType"`

}

// String returns a JSON representation of the model
func (o *Knowledgesettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesettings) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesettings

    if KnowledgesettingsMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesettingsMarshalled = true

    return json.Marshal(&struct {
        
        KnowledgeId string `json:"knowledgeId"`
        
        Description string `json:"description"`
        
        KnowledgeType string `json:"knowledgeType"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

