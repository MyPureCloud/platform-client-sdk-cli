package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobterminateresultentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobterminateresultentityDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Bulkjobterminateresultentity
type Bulkjobterminateresultentity struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Bulkjobterminateresultentity) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobterminateresultentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobterminateresultentity

    if BulkjobterminateresultentityMarshalled {
        return []byte("{}"), nil
    }
    BulkjobterminateresultentityMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

