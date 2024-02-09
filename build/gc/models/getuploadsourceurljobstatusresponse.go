package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetuploadsourceurljobstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetuploadsourceurljobstatusresponseDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Getuploadsourceurljobstatusresponse
type Getuploadsourceurljobstatusresponse struct { 
    // Id - Id of the upload from URL job.
    Id string `json:"id"`


    // Status - Status of the upload job
    Status string `json:"status"`


    // UploadKey - Key that identifies the file in the storage including the file name
    UploadKey string `json:"uploadKey"`


    // ErrorInformation - Any error information, or null of the processing is not in failed state.
    ErrorInformation Errorbody `json:"errorInformation"`


    

}

// String returns a JSON representation of the model
func (o *Getuploadsourceurljobstatusresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getuploadsourceurljobstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Getuploadsourceurljobstatusresponse

    if GetuploadsourceurljobstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    GetuploadsourceurljobstatusresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        UploadKey string `json:"uploadKey"`
        
        ErrorInformation Errorbody `json:"errorInformation"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

