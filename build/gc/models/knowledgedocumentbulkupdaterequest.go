package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentbulkupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentbulkupdaterequestDud struct { 
    

}

// Knowledgedocumentbulkupdaterequest
type Knowledgedocumentbulkupdaterequest struct { 
    // Entities - List of unique identifiers referencing documents that are to be updated
    Entities []Knowledgedocumentbulkupdateentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkupdaterequest) String() string {
     o.Entities = []Knowledgedocumentbulkupdateentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentbulkupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentbulkupdaterequest

    if KnowledgedocumentbulkupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentbulkupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgedocumentbulkupdateentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Knowledgedocumentbulkupdateentity{{}},
        

        Alias: (*Alias)(u),
    })
}

