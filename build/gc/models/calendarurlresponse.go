package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CalendarurlresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CalendarurlresponseDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Calendarurlresponse
type Calendarurlresponse struct { 
    // CalendarUrl - The calendar url for the user to subscribe with supported clients
    CalendarUrl string `json:"calendarUrl"`


    

}

// String returns a JSON representation of the model
func (o *Calendarurlresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Calendarurlresponse) MarshalJSON() ([]byte, error) {
    type Alias Calendarurlresponse

    if CalendarurlresponseMarshalled {
        return []byte("{}"), nil
    }
    CalendarurlresponseMarshalled = true

    return json.Marshal(&struct {
        
        CalendarUrl string `json:"calendarUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

