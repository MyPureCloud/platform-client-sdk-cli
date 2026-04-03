package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradesearchunmatchedtradesbujobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradesearchunmatchedtradesbujobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Shifttradesearchunmatchedtradesbujobresponse
type Shifttradesearchunmatchedtradesbujobresponse struct { 
    


    // Status - The status of the job
    Status string `json:"status"`


    // VarType - The type of the job
    VarType string `json:"type"`


    // DownloadUrl - The URL where completed results might be available for download in case the result body for that job type is too large
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Any error information, only set if the status == 'Error'
    VarError Errorbody `json:"error"`


    // SearchUnmatchedTradesResult - Results for SearchUnmatchedTrades job type
    SearchUnmatchedTradesResult Searchunmatchedshifttradelistjobresponseitem `json:"searchUnmatchedTradesResult"`


    

}

// String returns a JSON representation of the model
func (o *Shifttradesearchunmatchedtradesbujobresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradesearchunmatchedtradesbujobresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradesearchunmatchedtradesbujobresponse

    if ShifttradesearchunmatchedtradesbujobresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradesearchunmatchedtradesbujobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        SearchUnmatchedTradesResult Searchunmatchedshifttradelistjobresponseitem `json:"searchUnmatchedTradesResult"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

