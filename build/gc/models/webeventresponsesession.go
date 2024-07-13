package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebeventresponsesessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebeventresponsesessionDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Webeventresponsesession
type Webeventresponsesession struct { 
    


    // DurationInSeconds - Indicates how long the customer has been on the site within this session.
    DurationInSeconds int `json:"durationInSeconds"`


    // EventCount - The count of all events recorded during this session.
    EventCount int `json:"eventCount"`


    // PageviewCount - The count of all pageviews performed during this session.
    PageviewCount int `json:"pageviewCount"`


    // Referrer - The referrer of the first event in the web session.
    Referrer Referrer `json:"referrer"`


    


    // CreatedDate - Date of the session's first event, that is when the session starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Webeventresponsesession) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webeventresponsesession) MarshalJSON() ([]byte, error) {
    type Alias Webeventresponsesession

    if WebeventresponsesessionMarshalled {
        return []byte("{}"), nil
    }
    WebeventresponsesessionMarshalled = true

    return json.Marshal(&struct {
        
        DurationInSeconds int `json:"durationInSeconds"`
        
        EventCount int `json:"eventCount"`
        
        PageviewCount int `json:"pageviewCount"`
        
        Referrer Referrer `json:"referrer"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

