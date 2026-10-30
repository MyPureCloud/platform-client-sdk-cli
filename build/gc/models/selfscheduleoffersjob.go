package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleoffersjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleoffersjobDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Selfscheduleoffersjob
type Selfscheduleoffersjob struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the job
    Status string `json:"status"`


    // DownloadUrl - URL from which to fetch results for requests. Populated if status == ‘Complete'. The downloaded data will conform to the schema defined in downloadResultTemplate
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResultTemplate - Schema template for deserializing lookup offers results returned from the downloadUrl
    DownloadResultTemplate Selfscheduleoffersresulttemplate `json:"downloadResultTemplate"`


    // VarError - Any error information, only set if the status == 'Error'
    VarError Errorbody `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Selfscheduleoffersjob) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleoffersjob) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleoffersjob

    if SelfscheduleoffersjobMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleoffersjobMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResultTemplate Selfscheduleoffersresulttemplate `json:"downloadResultTemplate"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

