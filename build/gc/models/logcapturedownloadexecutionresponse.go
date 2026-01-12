package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogcapturedownloadexecutionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogcapturedownloadexecutionresponseDud struct { 
    


    User Addressableentityref `json:"user"`


    State string `json:"state"`


    DateStart time.Time `json:"dateStart"`


    FileUrl string `json:"fileUrl"`


    SelfUri string `json:"selfUri"`

}

// Logcapturedownloadexecutionresponse
type Logcapturedownloadexecutionresponse struct { 
    // Id - Id of file download job.
    Id string `json:"id"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Logcapturedownloadexecutionresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logcapturedownloadexecutionresponse) MarshalJSON() ([]byte, error) {
    type Alias Logcapturedownloadexecutionresponse

    if LogcapturedownloadexecutionresponseMarshalled {
        return []byte("{}"), nil
    }
    LogcapturedownloadexecutionresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

