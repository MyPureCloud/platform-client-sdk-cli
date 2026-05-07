package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeaddtradejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeaddtradejobresponseDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Shifttradeaddtradejobresponse
type Shifttradeaddtradejobresponse struct { 
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


    // AddTradeResult - Results for AddTrade job type
    AddTradeResult Shifttraderesponseitem `json:"addTradeResult"`


    

}

// String returns a JSON representation of the model
func (o *Shifttradeaddtradejobresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeaddtradejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeaddtradejobresponse

    if ShifttradeaddtradejobresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeaddtradejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        AddTradeResult Shifttraderesponseitem `json:"addTradeResult"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

