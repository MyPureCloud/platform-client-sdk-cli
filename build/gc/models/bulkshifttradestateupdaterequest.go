package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkshifttradestateupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkshifttradestateupdaterequestDud struct { 
    

}

// Bulkshifttradestateupdaterequest
type Bulkshifttradestateupdaterequest struct { 
    // Entities - The shift trades to update
    Entities []Bulkupdateshifttradestaterequestitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkshifttradestateupdaterequest) String() string {
     o.Entities = []Bulkupdateshifttradestaterequestitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkshifttradestateupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkshifttradestateupdaterequest

    if BulkshifttradestateupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    BulkshifttradestateupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Bulkupdateshifttradestaterequestitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Bulkupdateshifttradestaterequestitem{{}},
        

        Alias: (*Alias)(u),
    })
}

