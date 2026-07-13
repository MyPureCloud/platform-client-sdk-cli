package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedecisiontableversionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedecisiontableversionrequestDud struct { 
    

}

// Createdecisiontableversionrequest
type Createdecisiontableversionrequest struct { 
    // SourceVersion - The published, superseded, or snapshot version to create the new draft from. When not provided or null, the published version is used. Must be at least 1 when provided.
    SourceVersion int `json:"sourceVersion"`

}

// String returns a JSON representation of the model
func (o *Createdecisiontableversionrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdecisiontableversionrequest) MarshalJSON() ([]byte, error) {
    type Alias Createdecisiontableversionrequest

    if CreatedecisiontableversionrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedecisiontableversionrequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceVersion int `json:"sourceVersion"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

