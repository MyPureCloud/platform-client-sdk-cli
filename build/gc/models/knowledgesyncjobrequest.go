package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesyncjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesyncjobrequestDud struct { 
    


    

}

// Knowledgesyncjobrequest
type Knowledgesyncjobrequest struct { 
    // UploadKey - Upload key
    UploadKey string `json:"uploadKey"`


    // SourceId - Knowledge integration source id.
    SourceId string `json:"sourceId"`

}

// String returns a JSON representation of the model
func (o *Knowledgesyncjobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesyncjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesyncjobrequest

    if KnowledgesyncjobrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesyncjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        SourceId string `json:"sourceId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

