package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkcontactsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkcontactsrequestDud struct { 
    

}

// Bulkcontactsrequest
type Bulkcontactsrequest struct { 
    // Entities
    Entities []Externalcontact `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkcontactsrequest) String() string {
    
    
     o.Entities = []Externalcontact{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkcontactsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkcontactsrequest

    if BulkcontactsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkcontactsrequestMarshalled = true

    return json.Marshal(&struct { 
        Entities []Externalcontact `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Externalcontact{{}},
        

        
        Alias: (*Alias)(u),
    })
}

