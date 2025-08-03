package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermrequirementsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermrequirementsresponseDud struct { 
    


    


    


    

}

// Longtermrequirementsresponse
type Longtermrequirementsresponse struct { 
    // Status - Status of the long term forecast
    Status string `json:"status"`


    // ErrorCode - Error code when status is Failed
    ErrorCode string `json:"errorCode"`


    // LongTermRequirements - For schema documentation only, always null, schema for staffing forecast result at downloadUrl
    LongTermRequirements Longtermrequirements `json:"longTermRequirements"`


    // DownloadUrl - Download URL for the staffing forecast result
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Longtermrequirementsresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermrequirementsresponse) MarshalJSON() ([]byte, error) {
    type Alias Longtermrequirementsresponse

    if LongtermrequirementsresponseMarshalled {
        return []byte("{}"), nil
    }
    LongtermrequirementsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        ErrorCode string `json:"errorCode"`
        
        LongTermRequirements Longtermrequirements `json:"longTermRequirements"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

