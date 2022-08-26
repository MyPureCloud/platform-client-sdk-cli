package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentversionDud struct { 
    Id string `json:"id"`


    DatePublished time.Time `json:"datePublished"`


    Document Knowledgedocumentresponse `json:"document"`


    


    VersionNumber int `json:"versionNumber"`


    DateExpires time.Time `json:"dateExpires"`


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentversion
type Knowledgedocumentversion struct { 
    


    


    


    // RestoreFromVersionId - The globally unique identifier for the document version. If the value is provided, the document is restored to the given version. If not, it publishes the draft changes as a new version of the document.
    RestoreFromVersionId string `json:"restoreFromVersionId"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversion) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentversion) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentversion

    if KnowledgedocumentversionMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentversionMarshalled = true

    return json.Marshal(&struct {
        
        RestoreFromVersionId string `json:"restoreFromVersionId"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

