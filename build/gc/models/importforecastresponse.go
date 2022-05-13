package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportforecastresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportforecastresponseDud struct { 
    


    


    

}

// Importforecastresponse
type Importforecastresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation. Always null, result will come via notification
    Result Bushorttermforecast `json:"result"`

}

// String returns a JSON representation of the model
func (o *Importforecastresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importforecastresponse) MarshalJSON() ([]byte, error) {
    type Alias Importforecastresponse

    if ImportforecastresponseMarshalled {
        return []byte("{}"), nil
    }
    ImportforecastresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Bushorttermforecast `json:"result"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

