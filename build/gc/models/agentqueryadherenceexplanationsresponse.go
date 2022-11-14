package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentqueryadherenceexplanationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentqueryadherenceexplanationsresponseDud struct { 
    


    


    

}

// Agentqueryadherenceexplanationsresponse
type Agentqueryadherenceexplanationsresponse struct { 
    // Job - The asynchronous job handling the query
    Job Adherenceexplanationjobreference `json:"job"`


    // Result - The result of the query. May come via notification
    Result Adherenceexplanationlistingagentqueryresponse `json:"result"`


    // DownloadUrl - The URL from which to download the result. May come via notification
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Agentqueryadherenceexplanationsresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentqueryadherenceexplanationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentqueryadherenceexplanationsresponse

    if AgentqueryadherenceexplanationsresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentqueryadherenceexplanationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Adherenceexplanationjobreference `json:"job"`
        
        Result Adherenceexplanationlistingagentqueryresponse `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

