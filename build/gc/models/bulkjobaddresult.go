package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobaddresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobaddresultDud struct { 
    


    

}

// Bulkjobaddresult
type Bulkjobaddresult struct { 
    // VarError - Error details if the operation failed.
    VarError Bulkjoberror `json:"error"`


    // Entity - The result of the operation if it succeeded. For Workitem Bulk Add this is a summary.
    Entity Bulkjobaddworkitemsummary `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Bulkjobaddresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobaddresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobaddresult

    if BulkjobaddresultMarshalled {
        return []byte("{}"), nil
    }
    BulkjobaddresultMarshalled = true

    return json.Marshal(&struct {
        
        VarError Bulkjoberror `json:"error"`
        
        Entity Bulkjobaddworkitemsummary `json:"entity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

