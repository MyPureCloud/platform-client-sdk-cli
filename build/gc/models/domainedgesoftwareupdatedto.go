package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainedgesoftwareupdatedtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainedgesoftwareupdatedtoDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Domainedgesoftwareupdatedto
type Domainedgesoftwareupdatedto struct { 
    // Version - Version
    Version Domainedgesoftwareversiondto `json:"version"`


    // MaxDownloadRate
    MaxDownloadRate int `json:"maxDownloadRate"`


    // DownloadStartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DownloadStartTime time.Time `json:"downloadStartTime"`


    // ExecuteStartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExecuteStartTime time.Time `json:"executeStartTime"`


    // ExecuteStopTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExecuteStopTime time.Time `json:"executeStopTime"`


    // ExecuteOnIdle
    ExecuteOnIdle bool `json:"executeOnIdle"`


    // Status
    Status string `json:"status"`


    // EdgeUri
    EdgeUri string `json:"edgeUri"`


    // CallDrainingWaitTimeSeconds
    CallDrainingWaitTimeSeconds int `json:"callDrainingWaitTimeSeconds"`


    // Current
    Current bool `json:"current"`

}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareupdatedto) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainedgesoftwareupdatedto) MarshalJSON() ([]byte, error) {
    type Alias Domainedgesoftwareupdatedto

    if DomainedgesoftwareupdatedtoMarshalled {
        return []byte("{}"), nil
    }
    DomainedgesoftwareupdatedtoMarshalled = true

    return json.Marshal(&struct { 
        Version Domainedgesoftwareversiondto `json:"version"`
        
        MaxDownloadRate int `json:"maxDownloadRate"`
        
        DownloadStartTime time.Time `json:"downloadStartTime"`
        
        ExecuteStartTime time.Time `json:"executeStartTime"`
        
        ExecuteStopTime time.Time `json:"executeStopTime"`
        
        ExecuteOnIdle bool `json:"executeOnIdle"`
        
        Status string `json:"status"`
        
        EdgeUri string `json:"edgeUri"`
        
        CallDrainingWaitTimeSeconds int `json:"callDrainingWaitTimeSeconds"`
        
        Current bool `json:"current"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

