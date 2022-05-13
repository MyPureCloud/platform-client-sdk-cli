package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuasyncscheduleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuasyncscheduleresponseDud struct { 
    


    


    

}

// Buasyncscheduleresponse
type Buasyncscheduleresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Buschedulemetadata `json:"result"`

}

// String returns a JSON representation of the model
func (o *Buasyncscheduleresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buasyncscheduleresponse) MarshalJSON() ([]byte, error) {
    type Alias Buasyncscheduleresponse

    if BuasyncscheduleresponseMarshalled {
        return []byte("{}"), nil
    }
    BuasyncscheduleresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Buschedulemetadata `json:"result"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

