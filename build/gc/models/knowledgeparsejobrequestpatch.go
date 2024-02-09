package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeparsejobrequestpatchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeparsejobrequestpatchDud struct { 
    

}

// Knowledgeparsejobrequestpatch
type Knowledgeparsejobrequestpatch struct { 
    // Hints - Hinted titles for the parser.
    Hints []string `json:"hints"`

}

// String returns a JSON representation of the model
func (o *Knowledgeparsejobrequestpatch) String() string {
     o.Hints = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeparsejobrequestpatch) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeparsejobrequestpatch

    if KnowledgeparsejobrequestpatchMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeparsejobrequestpatchMarshalled = true

    return json.Marshal(&struct {
        
        Hints []string `json:"hints"`
        *Alias
    }{

        
        Hints: []string{""},
        

        Alias: (*Alias)(u),
    })
}

