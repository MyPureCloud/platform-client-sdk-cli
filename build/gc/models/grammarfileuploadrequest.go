package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GrammarfileuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GrammarfileuploadrequestDud struct { 
    

}

// Grammarfileuploadrequest
type Grammarfileuploadrequest struct { 
    // FileType
    FileType string `json:"fileType"`

}

// String returns a JSON representation of the model
func (o *Grammarfileuploadrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Grammarfileuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Grammarfileuploadrequest

    if GrammarfileuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    GrammarfileuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileType string `json:"fileType"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

