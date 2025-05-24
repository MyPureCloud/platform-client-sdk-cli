package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuheadcountforecastresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuheadcountforecastresponseDud struct { 
    


    

}

// Buheadcountforecastresponse
type Buheadcountforecastresponse struct { 
    // Result - The headcount forecast, null when downloadUrl is provided
    Result Buheadcountforecastbuplanninggroupheadcountforecastresult `json:"result"`


    // DownloadUrl - Download URL.  Null unless the response is too large to pass directly through the api
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Buheadcountforecastresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buheadcountforecastresponse) MarshalJSON() ([]byte, error) {
    type Alias Buheadcountforecastresponse

    if BuheadcountforecastresponseMarshalled {
        return []byte("{}"), nil
    }
    BuheadcountforecastresponseMarshalled = true

    return json.Marshal(&struct {
        
        Result Buheadcountforecastbuplanninggroupheadcountforecastresult `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

