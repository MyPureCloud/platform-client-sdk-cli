package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingstatusresponseDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    ErrorDetails []Schedulingprocessingerror `json:"errorDetails"`


    SchedulingResultUri string `json:"schedulingResultUri"`


    PercentComplete int `json:"percentComplete"`

}

// Schedulingstatusresponse
type Schedulingstatusresponse struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Schedulingstatusresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Schedulingstatusresponse

    if SchedulingstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    SchedulingstatusresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

