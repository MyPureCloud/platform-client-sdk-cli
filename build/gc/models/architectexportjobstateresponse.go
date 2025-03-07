package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectexportjobstateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectexportjobstateresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Architectexportjobstateresponse
type Architectexportjobstateresponse struct { 
    


    // Status - Status of the Architect Export Job
    Status string `json:"status"`


    // Command - The command executed by the Architect Job
    Command string `json:"command"`


    // DownloadUrl - The signed URL for downloading exported Architect data. If more than one flow was exported as part of the job, the URL provides a zipped folder containing all flows.
    DownloadUrl string `json:"downloadUrl"`


    // Messages - Warnings and Errors messages of the Architect Job
    Messages []Architectjobmessage `json:"messages"`


    

}

// String returns a JSON representation of the model
func (o *Architectexportjobstateresponse) String() string {
    
    
    
     o.Messages = []Architectjobmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectexportjobstateresponse) MarshalJSON() ([]byte, error) {
    type Alias Architectexportjobstateresponse

    if ArchitectexportjobstateresponseMarshalled {
        return []byte("{}"), nil
    }
    ArchitectexportjobstateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Command string `json:"command"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        Messages []Architectjobmessage `json:"messages"`
        *Alias
    }{

        


        


        


        


        
        Messages: []Architectjobmessage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

