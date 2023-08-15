package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentversionreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentversionreferenceDud struct { 
    Id string `json:"id"`


    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentversionreference
type Knowledgedocumentversionreference struct { 
    


    


    // VersionId - The globally unique identifier for the version of the document.
    VersionId string `json:"versionId"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversionreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentversionreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentversionreference

    if KnowledgedocumentversionreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentversionreferenceMarshalled = true

    return json.Marshal(&struct {
        
        VersionId string `json:"versionId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

