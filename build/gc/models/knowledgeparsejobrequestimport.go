package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeparsejobrequestimportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeparsejobrequestimportDud struct { 
    


    

}

// Knowledgeparsejobrequestimport
type Knowledgeparsejobrequestimport struct { 
    // Edits - Override of the result of the parse.
    Edits []Knowledgeparserecord `json:"edits"`


    // Excludes - Excluded results.
    Excludes []string `json:"excludes"`

}

// String returns a JSON representation of the model
func (o *Knowledgeparsejobrequestimport) String() string {
     o.Edits = []Knowledgeparserecord{{}} 
     o.Excludes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeparsejobrequestimport) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeparsejobrequestimport

    if KnowledgeparsejobrequestimportMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeparsejobrequestimportMarshalled = true

    return json.Marshal(&struct {
        
        Edits []Knowledgeparserecord `json:"edits"`
        
        Excludes []string `json:"excludes"`
        *Alias
    }{

        
        Edits: []Knowledgeparserecord{{}},
        


        
        Excludes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

