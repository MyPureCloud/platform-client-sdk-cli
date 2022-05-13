package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallforwardingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallforwardingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Callforwarding
type Callforwarding struct { 
    


    // Name
    Name string `json:"name"`


    // User
    User User `json:"user"`


    // Enabled - Whether or not CallForwarding is enabled
    Enabled bool `json:"enabled"`


    // PhoneNumber - This property is deprecated. Please use the calls property
    PhoneNumber string `json:"phoneNumber"`


    // Calls - An ordered list of CallRoutes to be executed when CallForwarding is enabled
    Calls []Callroute `json:"calls"`


    // Voicemail - The type of voicemail to use with the callForwarding configuration
    Voicemail string `json:"voicemail"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    

}

// String returns a JSON representation of the model
func (o *Callforwarding) String() string {
    
    
    
    
     o.Calls = []Callroute{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callforwarding) MarshalJSON() ([]byte, error) {
    type Alias Callforwarding

    if CallforwardingMarshalled {
        return []byte("{}"), nil
    }
    CallforwardingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        User User `json:"user"`
        
        Enabled bool `json:"enabled"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        Calls []Callroute `json:"calls"`
        
        Voicemail string `json:"voicemail"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        


        


        


        
        Calls: []Callroute{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

