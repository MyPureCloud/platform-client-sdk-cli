package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchdownloadjobstatusresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchdownloadjobstatusresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Batchdownloadjobstatusresult
type Batchdownloadjobstatusresult struct { 
    


    // JobId - JobId returned when job was initially submitted.
    JobId string `json:"jobId"`


    // ExpectedResultCount - Number of results expected when job is completed, this includes both success and error results. This number could change as recordings are being discovered and processed.
    ExpectedResultCount int `json:"expectedResultCount"`


    // ResultCount - Current number of results available, this includes both success and error results.
    ResultCount int `json:"resultCount"`


    // ErrorCount - Current number of error results.
    ErrorCount int `json:"errorCount"`


    // Status - Current status of the job. A job is considered completed when all the submitted requests have been processed and fulfilled.
    Status string `json:"status"`


    // Results - Current set of results for the job.
    Results []Batchdownloadjobresult `json:"results"`


    

}

// String returns a JSON representation of the model
func (o *Batchdownloadjobstatusresult) String() string {
    
    
    
    
    
     o.Results = []Batchdownloadjobresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchdownloadjobstatusresult) MarshalJSON() ([]byte, error) {
    type Alias Batchdownloadjobstatusresult

    if BatchdownloadjobstatusresultMarshalled {
        return []byte("{}"), nil
    }
    BatchdownloadjobstatusresultMarshalled = true

    return json.Marshal(&struct {
        
        JobId string `json:"jobId"`
        
        ExpectedResultCount int `json:"expectedResultCount"`
        
        ResultCount int `json:"resultCount"`
        
        ErrorCount int `json:"errorCount"`
        
        Status string `json:"status"`
        
        Results []Batchdownloadjobresult `json:"results"`
        *Alias
    }{

        


        


        


        


        


        


        
        Results: []Batchdownloadjobresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

