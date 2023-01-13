package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkresponseDud struct { 
    


    


    

}

// Wfmhistoricaladherencebulkresponse
type Wfmhistoricaladherencebulkresponse struct { 
    // Job - A reference to the job that was started by the request
    Job Wfmhistoricaladherencebulkjobreference `json:"job"`


    // DownloadUrls - The uri list to GET the results of the Historical Adherence query. This field is populated only if query state is Complete
    DownloadUrls []string `json:"downloadUrls"`


    // DownloadResult - Results will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Wfmhistoricaladherencebulkresult `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkresponse) String() string {
    
     o.DownloadUrls = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkresponse) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkresponse

    if WfmhistoricaladherencebulkresponseMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Wfmhistoricaladherencebulkjobreference `json:"job"`
        
        DownloadUrls []string `json:"downloadUrls"`
        
        DownloadResult Wfmhistoricaladherencebulkresult `json:"downloadResult"`
        *Alias
    }{

        


        
        DownloadUrls: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

