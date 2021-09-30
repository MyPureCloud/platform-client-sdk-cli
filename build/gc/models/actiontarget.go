package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActiontargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActiontargetDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    


    

}

// Actiontarget
type Actiontarget struct { 
    


    // Name
    Name string `json:"name"`


    // UserData - Additional user data associated with the target in key/value format.
    UserData []Keyvalue `json:"userData"`


    // SupportedMediaTypes - Supported media types of the target.
    SupportedMediaTypes []string `json:"supportedMediaTypes"`


    // State - Indicates the state of the target.
    State string `json:"state"`


    // Description - Description of the target.
    Description string `json:"description"`


    // ServiceLevel - Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
    ServiceLevel Servicelevel `json:"serviceLevel"`


    // ShortAbandonThreshold - Indicates the non-default short abandon threshold
    ShortAbandonThreshold int `json:"shortAbandonThreshold"`


    


    // CreatedDate - The date the target was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedDate - The date the target was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`

}

// String returns a JSON representation of the model
func (o *Actiontarget) String() string {
    
    
    
    
    
    
    
    
     o.UserData = []Keyvalue{{}} 
    
    
    
     o.SupportedMediaTypes = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actiontarget) MarshalJSON() ([]byte, error) {
    type Alias Actiontarget

    if ActiontargetMarshalled {
        return []byte("{}"), nil
    }
    ActiontargetMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        UserData []Keyvalue `json:"userData"`
        
        SupportedMediaTypes []string `json:"supportedMediaTypes"`
        
        State string `json:"state"`
        
        Description string `json:"description"`
        
        ServiceLevel Servicelevel `json:"serviceLevel"`
        
        ShortAbandonThreshold int `json:"shortAbandonThreshold"`
        
        
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        *Alias
    }{
        

        

        

        

        

        
        UserData: []Keyvalue{{}},
        

        

        
        SupportedMediaTypes: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

