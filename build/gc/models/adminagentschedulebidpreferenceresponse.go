package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminagentschedulebidpreferenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminagentschedulebidpreferenceresponseDud struct { 
    


    

}

// Adminagentschedulebidpreferenceresponse
type Adminagentschedulebidpreferenceresponse struct { 
    // Result - The agents' schedule set preferences
    Result Adminagentschedulesetpreferences `json:"result"`


    // DownloadUrl - URL to retrieve results when the response contains a large dataset. If provided, the downloaded data will follow the same schema as the result.
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Adminagentschedulebidpreferenceresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminagentschedulebidpreferenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Adminagentschedulebidpreferenceresponse

    if AdminagentschedulebidpreferenceresponseMarshalled {
        return []byte("{}"), nil
    }
    AdminagentschedulebidpreferenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Result Adminagentschedulesetpreferences `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

