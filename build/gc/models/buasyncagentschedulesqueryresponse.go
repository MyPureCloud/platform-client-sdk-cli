package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuasyncagentschedulesqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuasyncagentschedulesqueryresponseDud struct { 
    


    


    


    


    

}

// Buasyncagentschedulesqueryresponse
type Buasyncagentschedulesqueryresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Buagentschedulesqueryresponse `json:"result"`


    // Progress - Percent progress for the operation
    Progress int `json:"progress"`


    // DownloadUrl - The URL from which to download the result if it is too large to pass directly
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Buasyncagentschedulesqueryresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buasyncagentschedulesqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Buasyncagentschedulesqueryresponse

    if BuasyncagentschedulesqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    BuasyncagentschedulesqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Buagentschedulesqueryresponse `json:"result"`
        
        Progress int `json:"progress"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

