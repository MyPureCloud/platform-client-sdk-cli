package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PerformancepredictionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PerformancepredictionresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Performancepredictionresponse
type Performancepredictionresponse struct { 
    


    // WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // ScheduleId - The ID of the schedule this performance prediction is associated with
    ScheduleId string `json:"scheduleId"`


    // DownloadUrl - The url to GET the results of the performance prediction. This field is populated only if query state is 'Complete'
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Performancepredictionoutputs `json:"downloadResult"`


    // State - The state of the performance prediction
    State string `json:"state"`


    

}

// String returns a JSON representation of the model
func (o *Performancepredictionresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Performancepredictionresponse) MarshalJSON() ([]byte, error) {
    type Alias Performancepredictionresponse

    if PerformancepredictionresponseMarshalled {
        return []byte("{}"), nil
    }
    PerformancepredictionresponseMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        ScheduleId string `json:"scheduleId"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult Performancepredictionoutputs `json:"downloadResult"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

