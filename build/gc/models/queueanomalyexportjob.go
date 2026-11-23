package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueanomalyexportjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueanomalyexportjobDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Queueanomalyexportjob
type Queueanomalyexportjob struct { 
    


    // Status - The status of the export job
    Status string `json:"status"`


    // DownloadUrl - The URL from which to download the completed export. Populated when status is Complete
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - Anomalies always come through downloadUrl, the schema included here is just for documentation
    DownloadTemplate Queueanomalyexportjobtemplate `json:"downloadTemplate"`


    

}

// String returns a JSON representation of the model
func (o *Queueanomalyexportjob) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueanomalyexportjob) MarshalJSON() ([]byte, error) {
    type Alias Queueanomalyexportjob

    if QueueanomalyexportjobMarshalled {
        return []byte("{}"), nil
    }
    QueueanomalyexportjobMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Queueanomalyexportjobtemplate `json:"downloadTemplate"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

