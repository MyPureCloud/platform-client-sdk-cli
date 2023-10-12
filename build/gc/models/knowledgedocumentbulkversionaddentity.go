package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentbulkversionaddentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentbulkversionaddentityDud struct { 
    


    


    

}

// Knowledgedocumentbulkversionaddentity
type Knowledgedocumentbulkversionaddentity struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // RestoreFromVersionId - The globally unique identifier for the document version. If the value is provided, the document is restored to the given version.
    RestoreFromVersionId string `json:"restoreFromVersionId"`


    // RestorePrevious - Indicates if the document's previous version will be restored
    RestorePrevious bool `json:"restorePrevious"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkversionaddentity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentbulkversionaddentity) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentbulkversionaddentity

    if KnowledgedocumentbulkversionaddentityMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentbulkversionaddentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        RestoreFromVersionId string `json:"restoreFromVersionId"`
        
        RestorePrevious bool `json:"restorePrevious"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

