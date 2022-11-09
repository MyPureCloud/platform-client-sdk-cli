package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeexportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeexportjobrequestDud struct { 
    


    

}

// Knowledgeexportjobrequest
type Knowledgeexportjobrequest struct { 
    // ExportFilter - What to export.
    ExportFilter Knowledgeexportjobfilter `json:"exportFilter"`


    // FileType - File type of the document
    FileType string `json:"fileType"`

}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeexportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeexportjobrequest

    if KnowledgeexportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeexportjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ExportFilter Knowledgeexportjobfilter `json:"exportFilter"`
        
        FileType string `json:"fileType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

