package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateshifttradestateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateshifttradestateresponseDud struct { 
    


    


    

}

// Bulkupdateshifttradestateresponse
type Bulkupdateshifttradestateresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Bulkupdateshifttradestateresult `json:"result"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateshifttradestateresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateshifttradestateresponse

    if BulkupdateshifttradestateresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateshifttradestateresponseMarshalled = true

    return json.Marshal(&struct { 
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Bulkupdateshifttradestateresult `json:"result"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

