package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneratemeetingidrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneratemeetingidrequestDud struct { 
    


    


    

}

// Generatemeetingidrequest
type Generatemeetingidrequest struct { 
    // ConferenceId - The conferenceId for which to generate a meetingId
    ConferenceId string `json:"conferenceId"`


    // Ephemeral - Boolean flag for ephemeral status of the created record
    Ephemeral bool `json:"ephemeral"`


    // ExpireTimeDays - Number of days the meetingId record will remain active
    ExpireTimeDays int `json:"expireTimeDays"`

}

// String returns a JSON representation of the model
func (o *Generatemeetingidrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generatemeetingidrequest) MarshalJSON() ([]byte, error) {
    type Alias Generatemeetingidrequest

    if GeneratemeetingidrequestMarshalled {
        return []byte("{}"), nil
    }
    GeneratemeetingidrequestMarshalled = true

    return json.Marshal(&struct {
        
        ConferenceId string `json:"conferenceId"`
        
        Ephemeral bool `json:"ephemeral"`
        
        ExpireTimeDays int `json:"expireTimeDays"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

