package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateshifttradestateresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateshifttradestateresultDud struct { 
    

}

// Bulkupdateshifttradestateresult
type Bulkupdateshifttradestateresult struct { 
    // Entities
    Entities []Bulkupdateshifttradestateresultitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresult) String() string {
    
    
     o.Entities = []Bulkupdateshifttradestateresultitem{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateshifttradestateresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateshifttradestateresult

    if BulkupdateshifttradestateresultMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateshifttradestateresultMarshalled = true

    return json.Marshal(&struct { 
        Entities []Bulkupdateshifttradestateresultitem `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Bulkupdateshifttradestateresultitem{{}},
        

        
        Alias: (*Alias)(u),
    })
}

