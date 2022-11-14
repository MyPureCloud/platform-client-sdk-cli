package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricalshrinkageresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricalshrinkageresponseDud struct { 
    


    


    


    

}

// Wfmhistoricalshrinkageresponse
type Wfmhistoricalshrinkageresponse struct { 
    // OperationId - The operationId for which to listen
    OperationId string `json:"operationId"`


    // DownloadUrls - The url list to GET the results of the Historical Shrinkage query. This field is populated only if query state is Complete
    DownloadUrls []string `json:"downloadUrls"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Historicalshrinkageresultlisting `json:"downloadResult"`


    // State - The state of the shrinkage query
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricalshrinkageresponse) String() string {
    
     o.DownloadUrls = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricalshrinkageresponse) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricalshrinkageresponse

    if WfmhistoricalshrinkageresponseMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricalshrinkageresponseMarshalled = true

    return json.Marshal(&struct {
        
        OperationId string `json:"operationId"`
        
        DownloadUrls []string `json:"downloadUrls"`
        
        DownloadResult Historicalshrinkageresultlisting `json:"downloadResult"`
        
        State string `json:"state"`
        *Alias
    }{

        


        
        DownloadUrls: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

