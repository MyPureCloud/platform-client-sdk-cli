package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekscheduleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekscheduleresponseDud struct { 
    


    

}

// Weekscheduleresponse - Response for query for week schedule for a given week in management unit
type Weekscheduleresponse struct { 
    // Result - The result of the request. The value will be null if response is large
    Result Weekschedule `json:"result"`


    // DownloadUrl - The url to fetch the result for large responses. The value is null if result contains the data
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Weekscheduleresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekscheduleresponse) MarshalJSON() ([]byte, error) {
    type Alias Weekscheduleresponse

    if WeekscheduleresponseMarshalled {
        return []byte("{}"), nil
    }
    WeekscheduleresponseMarshalled = true

    return json.Marshal(&struct { 
        Result Weekschedule `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

