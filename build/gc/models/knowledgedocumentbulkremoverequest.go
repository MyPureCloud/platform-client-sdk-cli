package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentbulkremoverequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentbulkremoverequestDud struct { 
    

}

// Knowledgedocumentbulkremoverequest
type Knowledgedocumentbulkremoverequest struct { 
    // Entities - List of unique identifiers referencing documents that are to be deleted
    Entities []Writableentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkremoverequest) String() string {
     o.Entities = []Writableentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentbulkremoverequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentbulkremoverequest

    if KnowledgedocumentbulkremoverequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentbulkremoverequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Writableentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Writableentity{{}},
        

        Alias: (*Alias)(u),
    })
}

