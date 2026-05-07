package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeupdatetradestatejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeupdatetradestatejobresponseDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Shifttradeupdatetradestatejobresponse
type Shifttradeupdatetradestatejobresponse struct { 
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


    // UpdateTradeStateResult - Results for UpdateTradeState job type
    UpdateTradeStateResult Bulkupdateshifttradestateresultitem `json:"updateTradeStateResult"`


    

}

// String returns a JSON representation of the model
func (o *Shifttradeupdatetradestatejobresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeupdatetradestatejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeupdatetradestatejobresponse

    if ShifttradeupdatetradestatejobresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeupdatetradestatejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        UpdateTradeStateResult Bulkupdateshifttradestateresultitem `json:"updateTradeStateResult"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

