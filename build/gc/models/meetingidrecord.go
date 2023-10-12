package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MeetingidrecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MeetingidrecordDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Meetingidrecord
type Meetingidrecord struct { 
    


    // Ephemeral - Boolean flag for ephemeral status of the created record
    Ephemeral bool `json:"ephemeral"`


    // ConferenceId - The conferenceId used to generate a meetingId
    ConferenceId string `json:"conferenceId"`


    // DateExpired - Instant at which the meetingId record will no longer be considered active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpired time.Time `json:"dateExpired"`


    

}

// String returns a JSON representation of the model
func (o *Meetingidrecord) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Meetingidrecord) MarshalJSON() ([]byte, error) {
    type Alias Meetingidrecord

    if MeetingidrecordMarshalled {
        return []byte("{}"), nil
    }
    MeetingidrecordMarshalled = true

    return json.Marshal(&struct {
        
        Ephemeral bool `json:"ephemeral"`
        
        ConferenceId string `json:"conferenceId"`
        
        DateExpired time.Time `json:"dateExpired"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

