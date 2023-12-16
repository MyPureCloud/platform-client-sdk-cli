package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PerformancepredictionrecalculationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PerformancepredictionrecalculationresponseDud struct { 
    


    


    


    

}

// Performancepredictionrecalculationresponse
type Performancepredictionrecalculationresponse struct { 
    // OperationId - The operationId for which to listen
    OperationId string `json:"operationId"`


    // DownloadUrl - The url to GET the results of the performance prediction. This field is populated only if query state is 'Complete'
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Performancepredictionoutputs `json:"downloadResult"`


    // State - The state of the performance prediction
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Performancepredictionrecalculationresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Performancepredictionrecalculationresponse) MarshalJSON() ([]byte, error) {
    type Alias Performancepredictionrecalculationresponse

    if PerformancepredictionrecalculationresponseMarshalled {
        return []byte("{}"), nil
    }
    PerformancepredictionrecalculationresponseMarshalled = true

    return json.Marshal(&struct {
        
        OperationId string `json:"operationId"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult Performancepredictionoutputs `json:"downloadResult"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

