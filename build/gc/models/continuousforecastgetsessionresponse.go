package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContinuousforecastgetsessionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContinuousforecastgetsessionresponseDud struct { 
    


    


    


    

}

// Continuousforecastgetsessionresponse
type Continuousforecastgetsessionresponse struct { 
    // SessionId - Latest session ID of the business unit
    SessionId string `json:"sessionId"`


    // LastSuccessfulSessionId - Last successful session ID of the business unit
    LastSuccessfulSessionId string `json:"lastSuccessfulSessionId"`


    // State - State of the latest session
    State string `json:"state"`


    // ErrorCode - Failed session error code
    ErrorCode string `json:"errorCode"`

}

// String returns a JSON representation of the model
func (o *Continuousforecastgetsessionresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Continuousforecastgetsessionresponse) MarshalJSON() ([]byte, error) {
    type Alias Continuousforecastgetsessionresponse

    if ContinuousforecastgetsessionresponseMarshalled {
        return []byte("{}"), nil
    }
    ContinuousforecastgetsessionresponseMarshalled = true

    return json.Marshal(&struct {
        
        SessionId string `json:"sessionId"`
        
        LastSuccessfulSessionId string `json:"lastSuccessfulSessionId"`
        
        State string `json:"state"`
        
        ErrorCode string `json:"errorCode"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

