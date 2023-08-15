package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentversionreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentversionreferenceDud struct { 
    Id string `json:"id"`


    

}

// Knowledgeguestdocumentversionreference
type Knowledgeguestdocumentversionreference struct { 
    


    // VersionId - The globally unique identifier for the version of the document.
    VersionId string `json:"versionId"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentversionreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentversionreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentversionreference

    if KnowledgeguestdocumentversionreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentversionreferenceMarshalled = true

    return json.Marshal(&struct {
        
        VersionId string `json:"versionId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

