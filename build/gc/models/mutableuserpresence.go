package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MutableuserpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MutableuserpresenceDud struct { 
    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Mutableuserpresence
type Mutableuserpresence struct { 
    // Id - The user's id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Source - Deprecated - The sourceID field should be used as a replacement.
    Source string `json:"source"`


    // SourceId - Represents the ID of a registered source
    SourceId string `json:"sourceId"`


    // Primary - A boolean used to tell whether or not to set this presence source as the primary on a PATCH
    Primary bool `json:"primary"`


    // PresenceDefinition
    PresenceDefinition Presencedefinition `json:"presenceDefinition"`


    // Message
    Message string `json:"message"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    

}

// String returns a JSON representation of the model
func (o *Mutableuserpresence) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mutableuserpresence) MarshalJSON() ([]byte, error) {
    type Alias Mutableuserpresence

    if MutableuserpresenceMarshalled {
        return []byte("{}"), nil
    }
    MutableuserpresenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Source string `json:"source"`
        
        SourceId string `json:"sourceId"`
        
        Primary bool `json:"primary"`
        
        PresenceDefinition Presencedefinition `json:"presenceDefinition"`
        
        Message string `json:"message"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

