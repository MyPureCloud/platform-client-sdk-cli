package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppeventresponsesessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppeventresponsesessionDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Appeventresponsesession
type Appeventresponsesession struct { 
    // Id - ID of the app session.
    Id string `json:"id"`


    // DurationInSeconds - Indicates how long the customer has been in the app within this session.
    DurationInSeconds int `json:"durationInSeconds"`


    // EventCount - The count of all events recorded during this session.
    EventCount int `json:"eventCount"`


    // ScreenviewCount - The count of all screen views recorded during this session.
    ScreenviewCount int `json:"screenviewCount"`


    // Referrer - The referrer of the first event in the app session.
    Referrer Referrer `json:"referrer"`


    


    // CreatedDate - UTC timestamp of the session's first event, that is when the session starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Appeventresponsesession) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appeventresponsesession) MarshalJSON() ([]byte, error) {
    type Alias Appeventresponsesession

    if AppeventresponsesessionMarshalled {
        return []byte("{}"), nil
    }
    AppeventresponsesessionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DurationInSeconds int `json:"durationInSeconds"`
        
        EventCount int `json:"eventCount"`
        
        ScreenviewCount int `json:"screenviewCount"`
        
        Referrer Referrer `json:"referrer"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

