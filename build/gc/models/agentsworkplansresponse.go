package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentsworkplansresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentsworkplansresponseDud struct { 
    


    

}

// Agentsworkplansresponse
type Agentsworkplansresponse struct { 
    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`


    // Result - The result of agent(s) work plans query
    Result Muagentsworkplansresult `json:"result"`

}

// String returns a JSON representation of the model
func (o *Agentsworkplansresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentsworkplansresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentsworkplansresponse

    if AgentsworkplansresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentsworkplansresponseMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        
        Result Muagentsworkplansresult `json:"result"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

