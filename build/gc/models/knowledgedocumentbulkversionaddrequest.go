package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentbulkversionaddrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentbulkversionaddrequestDud struct { 
    

}

// Knowledgedocumentbulkversionaddrequest
type Knowledgedocumentbulkversionaddrequest struct { 
    // Entities - List of unique identifiers referencing documents that are to be versioned
    Entities []Knowledgedocumentbulkversionaddentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkversionaddrequest) String() string {
     o.Entities = []Knowledgedocumentbulkversionaddentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentbulkversionaddrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentbulkversionaddrequest

    if KnowledgedocumentbulkversionaddrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentbulkversionaddrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgedocumentbulkversionaddentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Knowledgedocumentbulkversionaddentity{{}},
        

        Alias: (*Alias)(u),
    })
}

