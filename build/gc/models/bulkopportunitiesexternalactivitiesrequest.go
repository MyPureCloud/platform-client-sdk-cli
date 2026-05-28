package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunitiesexternalactivitiesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunitiesexternalactivitiesrequestDud struct { 
    

}

// Bulkopportunitiesexternalactivitiesrequest
type Bulkopportunitiesexternalactivitiesrequest struct { 
    // Ids - The IDs of the external activities
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunitiesexternalactivitiesrequest) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunitiesexternalactivitiesrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunitiesexternalactivitiesrequest

    if BulkopportunitiesexternalactivitiesrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunitiesexternalactivitiesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

