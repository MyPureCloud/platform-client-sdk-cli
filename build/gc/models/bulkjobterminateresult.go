package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobterminateresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobterminateresultDud struct { 
    


    

}

// Bulkjobterminateresult
type Bulkjobterminateresult struct { 
    // VarError - Error details if the operation failed.
    VarError Bulkjoberror `json:"error"`


    // Entity - The result of the operation if it succeeded. For Workitem Bulk Terminate this is a summary.
    Entity Bulkjobterminateresultentity `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Bulkjobterminateresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobterminateresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobterminateresult

    if BulkjobterminateresultMarshalled {
        return []byte("{}"), nil
    }
    BulkjobterminateresultMarshalled = true

    return json.Marshal(&struct {
        
        VarError Bulkjoberror `json:"error"`
        
        Entity Bulkjobterminateresultentity `json:"entity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

