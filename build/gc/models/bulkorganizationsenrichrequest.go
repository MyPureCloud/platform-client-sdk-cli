package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkorganizationsenrichrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkorganizationsenrichrequestDud struct { 
    

}

// Bulkorganizationsenrichrequest
type Bulkorganizationsenrichrequest struct { 
    // Entities
    Entities []Externalorganizationenrichrequest `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkorganizationsenrichrequest) String() string {
     o.Entities = []Externalorganizationenrichrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkorganizationsenrichrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkorganizationsenrichrequest

    if BulkorganizationsenrichrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkorganizationsenrichrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externalorganizationenrichrequest `json:"entities"`
        *Alias
    }{

        
        Entities: []Externalorganizationenrichrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

