package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncforecastoperationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncforecastoperationresultDud struct { 
    


    


    


    

}

// Asyncforecastoperationresult
type Asyncforecastoperationresult struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Bushorttermforecast `json:"result"`


    // Progress - Percent progress for the operation
    Progress int `json:"progress"`

}

// String returns a JSON representation of the model
func (o *Asyncforecastoperationresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncforecastoperationresult) MarshalJSON() ([]byte, error) {
    type Alias Asyncforecastoperationresult

    if AsyncforecastoperationresultMarshalled {
        return []byte("{}"), nil
    }
    AsyncforecastoperationresultMarshalled = true

    return json.Marshal(&struct { 
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Bushorttermforecast `json:"result"`
        
        Progress int `json:"progress"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

