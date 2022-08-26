package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeexportjobfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeexportjobfilterDud struct { 
    


    

}

// Knowledgeexportjobfilter
type Knowledgeexportjobfilter struct { 
    // DocumentsFilter - Filters for narrowing down which documents to export.
    DocumentsFilter Knowledgeexportjobdocumentsfilter `json:"documentsFilter"`


    // VersionFilter - Specifies what version should be exported.
    VersionFilter string `json:"versionFilter"`

}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeexportjobfilter) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeexportjobfilter

    if KnowledgeexportjobfilterMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeexportjobfilterMarshalled = true

    return json.Marshal(&struct {
        
        DocumentsFilter Knowledgeexportjobdocumentsfilter `json:"documentsFilter"`
        
        VersionFilter string `json:"versionFilter"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

