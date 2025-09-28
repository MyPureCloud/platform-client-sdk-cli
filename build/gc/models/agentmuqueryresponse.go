package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmuqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmuqueryresponseDud struct { 
    


    


    


    


    

}

// Agentmuqueryresponse
type Agentmuqueryresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The schema of the result of the operation. Null if downloadUrl is populated, but defines the schema of the data that will be returned from the downloadUrl
    Result Agentmuscheduleresult `json:"result"`


    // DownloadUrl - The URL from which to download the result. The result will follow the schema documented by the result field
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Error details if status == 'Error'. Will always be null but the schema is documented in case of an error in an async notification
    VarError Errorbody `json:"error"`

}

// String returns a JSON representation of the model
func (o *Agentmuqueryresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmuqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentmuqueryresponse

    if AgentmuqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentmuqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Agentmuscheduleresult `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

