package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UcuserpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UcuserpresenceDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Ucuserpresence - Presence from a given source for a user
type Ucuserpresence struct { 
    


    // Name
    Name string `json:"name"`


    // UserId - User ID of the associated Genesys Cloud user.
    UserId string `json:"userId"`


    // Source - Deprecated - The sourceID field should be used as a replacement.
    Source string `json:"source"`


    // SourceId - The registered source ID from where the presence was set
    SourceId string `json:"sourceId"`


    // PresenceDefinition
    PresenceDefinition Presencedefinition `json:"presenceDefinition"`


    // Message
    Message string `json:"message"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    

}

// String returns a JSON representation of the model
func (o *Ucuserpresence) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ucuserpresence) MarshalJSON() ([]byte, error) {
    type Alias Ucuserpresence

    if UcuserpresenceMarshalled {
        return []byte("{}"), nil
    }
    UcuserpresenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UserId string `json:"userId"`
        
        Source string `json:"source"`
        
        SourceId string `json:"sourceId"`
        
        PresenceDefinition Presencedefinition `json:"presenceDefinition"`
        
        Message string `json:"message"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

