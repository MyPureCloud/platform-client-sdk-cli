package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuasyncschedulerunresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuasyncschedulerunresponseDud struct { 
    


    


    

}

// Buasyncschedulerunresponse
type Buasyncschedulerunresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Buschedulerun `json:"result"`

}

// String returns a JSON representation of the model
func (o *Buasyncschedulerunresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buasyncschedulerunresponse) MarshalJSON() ([]byte, error) {
    type Alias Buasyncschedulerunresponse

    if BuasyncschedulerunresponseMarshalled {
        return []byte("{}"), nil
    }
    BuasyncschedulerunresponseMarshalled = true

    return json.Marshal(&struct { 
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Buschedulerun `json:"result"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

