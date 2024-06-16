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


    // SourceId - Knowledge integration source id.
    SourceId string `json:"sourceId"`


    // JsonFileVersion - Requested version of the exported json file. Available versions are 2 and 3, default is 2
    JsonFileVersion int `json:"jsonFileVersion"`

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
        
        SourceId string `json:"sourceId"`
        
        JsonFileVersion int `json:"jsonFileVersion"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

