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
    // SessionId - The ID of the latest session, regardless of the session's status
    SessionId string `json:"sessionId"`


    // LastSuccessfulSessionId - The ID of the last session that has a state of Complete
    LastSuccessfulSessionId string `json:"lastSuccessfulSessionId"`


    // State - The state of the latest session
    State string `json:"state"`


    // ErrorCode - The error code if the latest session has a state of Error
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

