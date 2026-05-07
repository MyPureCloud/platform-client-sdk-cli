package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradequeryagenttradesjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradequeryagenttradesjobresponseDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Shifttradequeryagenttradesjobresponse
type Shifttradequeryagenttradesjobresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the job
    Status string `json:"status"`


    // VarType - The type of the job
    VarType string `json:"type"`


    // DownloadUrl - The URL where completed results might be available for download in case the result body for that job type is too large
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Any error information, only set if the status == 'Error'
    VarError Errorbody `json:"error"`


    // QueryAgentTradesResult - Results for QueryAgentTrades job type
    QueryAgentTradesResult Shifttradelistjobresponse `json:"queryAgentTradesResult"`


    

}

// String returns a JSON representation of the model
func (o *Shifttradequeryagenttradesjobresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradequeryagenttradesjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradequeryagenttradesjobresponse

    if ShifttradequeryagenttradesjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradequeryagenttradesjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        QueryAgentTradesResult Shifttradelistjobresponse `json:"queryAgentTradesResult"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

