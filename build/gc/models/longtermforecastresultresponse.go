package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermforecastresultresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermforecastresultresponseDud struct { 
    


    

}

// Longtermforecastresultresponse
type Longtermforecastresultresponse struct { 
    // Result - The result of the operation.  Populated whenever the result is small enough to pass through the api directly
    Result Longtermforecastresult `json:"result"`


    // DownloadUrl - The download url to fetch the result.  Only populated if the result is too large to pass through the api directly
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Longtermforecastresultresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermforecastresultresponse) MarshalJSON() ([]byte, error) {
    type Alias Longtermforecastresultresponse

    if LongtermforecastresultresponseMarshalled {
        return []byte("{}"), nil
    }
    LongtermforecastresultresponseMarshalled = true

    return json.Marshal(&struct { 
        Result Longtermforecastresult `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

