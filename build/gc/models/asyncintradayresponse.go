package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncintradayresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncintradayresponseDud struct { 
    


    


    

}

// Asyncintradayresponse
type Asyncintradayresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Buintradayresponse `json:"result"`

}

// String returns a JSON representation of the model
func (o *Asyncintradayresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncintradayresponse) MarshalJSON() ([]byte, error) {
    type Alias Asyncintradayresponse

    if AsyncintradayresponseMarshalled {
        return []byte("{}"), nil
    }
    AsyncintradayresponseMarshalled = true

    return json.Marshal(&struct { 
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Buintradayresponse `json:"result"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

