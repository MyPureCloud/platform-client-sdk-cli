package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeparsejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeparsejobrequestDud struct { 
    


    

}

// Knowledgeparsejobrequest
type Knowledgeparsejobrequest struct { 
    // UploadKey - Upload key
    UploadKey string `json:"uploadKey"`


    // Hints - Hinted titles for the parser.
    Hints []string `json:"hints"`

}

// String returns a JSON representation of the model
func (o *Knowledgeparsejobrequest) String() string {
    
     o.Hints = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeparsejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeparsejobrequest

    if KnowledgeparsejobrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeparsejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        Hints []string `json:"hints"`
        *Alias
    }{

        


        
        Hints: []string{""},
        

        Alias: (*Alias)(u),
    })
}

