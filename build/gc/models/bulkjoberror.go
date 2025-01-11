package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjoberrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjoberrorDud struct { 
    


    

}

// Bulkjoberror
type Bulkjoberror struct { 
    // Message - Error message of the bulk operation result.
    Message string `json:"message"`


    // Code - Error code of the bulk operation result.
    Code string `json:"code"`

}

// String returns a JSON representation of the model
func (o *Bulkjoberror) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjoberror) MarshalJSON() ([]byte, error) {
    type Alias Bulkjoberror

    if BulkjoberrorMarshalled {
        return []byte("{}"), nil
    }
    BulkjoberrorMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Code string `json:"code"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

