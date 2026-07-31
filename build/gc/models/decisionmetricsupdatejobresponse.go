package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsupdatejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsupdatejobresponseDud struct { 
    


    


    


    


    

}

// Decisionmetricsupdatejobresponse
type Decisionmetricsupdatejobresponse struct { 
    // UploadKey - The S3 key for the uploaded decision metrics file
    UploadKey string `json:"uploadKey"`


    // Job - The update job
    Job Decisionmetricsjobreference `json:"job"`


    // Status - The status of the update job
    Status string `json:"status"`


    // Metadata - The metadata of the update job
    Metadata Wfmentitymetadata `json:"metadata"`


    // UpdateErrors - Errors occurred during update process, which will be non empty when status is `Error`
    UpdateErrors []Decisionmetricsupdateerror `json:"updateErrors"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsupdatejobresponse) String() string {
    
    
    
    
     o.UpdateErrors = []Decisionmetricsupdateerror{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsupdatejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsupdatejobresponse

    if DecisionmetricsupdatejobresponseMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsupdatejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        Job Decisionmetricsjobreference `json:"job"`
        
        Status string `json:"status"`
        
        Metadata Wfmentitymetadata `json:"metadata"`
        
        UpdateErrors []Decisionmetricsupdateerror `json:"updateErrors"`
        *Alias
    }{

        


        


        


        


        
        UpdateErrors: []Decisionmetricsupdateerror{{}},
        

        Alias: (*Alias)(u),
    })
}

