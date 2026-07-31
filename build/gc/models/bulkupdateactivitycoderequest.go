package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateactivitycoderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateactivitycoderequestDud struct { 
    

}

// Bulkupdateactivitycoderequest
type Bulkupdateactivitycoderequest struct { 
    // Entities - List of activity codes to update
    Entities []Bulkupdateactivitycoderequestitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateactivitycoderequest) String() string {
     o.Entities = []Bulkupdateactivitycoderequestitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateactivitycoderequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateactivitycoderequest

    if BulkupdateactivitycoderequestMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateactivitycoderequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Bulkupdateactivitycoderequestitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Bulkupdateactivitycoderequestitem{{}},
        

        Alias: (*Alias)(u),
    })
}

