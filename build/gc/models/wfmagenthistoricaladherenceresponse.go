package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmagenthistoricaladherenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmagenthistoricaladherenceresponseDud struct { 
    


    


    

}

// Wfmagenthistoricaladherenceresponse
type Wfmagenthistoricaladherenceresponse struct { 
    // Job - A reference to the job that was started by the request
    Job Wfmagenthistoricaladherencejobreference `json:"job"`


    // DownloadUrls - The url list to GET the results of the agent adherence query. This field is populated only if query state is Complete
    DownloadUrls []string `json:"downloadUrls"`


    // Result - Results will come via downloadUrls, however it may come inline if small enough
    Result Wfmagenthistoricaladherenceresult `json:"result"`

}

// String returns a JSON representation of the model
func (o *Wfmagenthistoricaladherenceresponse) String() string {
    
     o.DownloadUrls = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmagenthistoricaladherenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Wfmagenthistoricaladherenceresponse

    if WfmagenthistoricaladherenceresponseMarshalled {
        return []byte("{}"), nil
    }
    WfmagenthistoricaladherenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Wfmagenthistoricaladherencejobreference `json:"job"`
        
        DownloadUrls []string `json:"downloadUrls"`
        
        Result Wfmagenthistoricaladherenceresult `json:"result"`
        *Alias
    }{

        


        
        DownloadUrls: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

