package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunitiesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunitiesrequestDud struct { 
    

}

// Bulkopportunitiesrequest
type Bulkopportunitiesrequest struct { 
    // Ids - The IDs of the opportunities
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunitiesrequest) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunitiesrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunitiesrequest

    if BulkopportunitiesrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunitiesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

