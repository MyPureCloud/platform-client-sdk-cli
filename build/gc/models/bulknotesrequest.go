package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulknotesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulknotesrequestDud struct { 
    

}

// Bulknotesrequest
type Bulknotesrequest struct { 
    // Entities
    Entities []Note `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulknotesrequest) String() string {
    
    
     o.Entities = []Note{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulknotesrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulknotesrequest

    if BulknotesrequestMarshalled {
        return []byte("{}"), nil
    }
    BulknotesrequestMarshalled = true

    return json.Marshal(&struct { 
        Entities []Note `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Note{{}},
        

        
        Alias: (*Alias)(u),
    })
}

