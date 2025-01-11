package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContinuousforecastsessionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContinuousforecastsessionresponseDud struct { 
    


    


    

}

// Continuousforecastsessionresponse
type Continuousforecastsessionresponse struct { 
    // SessionId - Session ID of the continuous forecast data
    SessionId string `json:"sessionId"`


    // State - State of the requested session
    State string `json:"state"`


    // Files - Link to the files containing data for requested session
    Files Sessionfiles `json:"files"`

}

// String returns a JSON representation of the model
func (o *Continuousforecastsessionresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Continuousforecastsessionresponse) MarshalJSON() ([]byte, error) {
    type Alias Continuousforecastsessionresponse

    if ContinuousforecastsessionresponseMarshalled {
        return []byte("{}"), nil
    }
    ContinuousforecastsessionresponseMarshalled = true

    return json.Marshal(&struct {
        
        SessionId string `json:"sessionId"`
        
        State string `json:"state"`
        
        Files Sessionfiles `json:"files"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

