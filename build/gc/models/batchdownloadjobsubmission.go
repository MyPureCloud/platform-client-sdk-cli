package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchdownloadjobsubmissionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchdownloadjobsubmissionDud struct { 
    

}

// Batchdownloadjobsubmission
type Batchdownloadjobsubmission struct { 
    // BatchDownloadRequestList - List of up to 100 items requested
    BatchDownloadRequestList []Batchdownloadrequest `json:"batchDownloadRequestList"`

}

// String returns a JSON representation of the model
func (o *Batchdownloadjobsubmission) String() string {
    
    
     o.BatchDownloadRequestList = []Batchdownloadrequest{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchdownloadjobsubmission) MarshalJSON() ([]byte, error) {
    type Alias Batchdownloadjobsubmission

    if BatchdownloadjobsubmissionMarshalled {
        return []byte("{}"), nil
    }
    BatchdownloadjobsubmissionMarshalled = true

    return json.Marshal(&struct { 
        BatchDownloadRequestList []Batchdownloadrequest `json:"batchDownloadRequestList"`
        
        *Alias
    }{
        

        
        BatchDownloadRequestList: []Batchdownloadrequest{{}},
        

        
        Alias: (*Alias)(u),
    })
}

