package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GreetingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GreetingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Greeting
type Greeting struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - Greeting type
    VarType string `json:"type"`


    // OwnerType - Greeting owner type
    OwnerType string `json:"ownerType"`


    // Owner - Greeting owner
    Owner Domainentity `json:"owner"`


    // AudioFile
    AudioFile Greetingaudiofile `json:"audioFile"`


    // AudioTTS
    AudioTTS string `json:"audioTTS"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // CreatedBy
    CreatedBy string `json:"createdBy"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // ModifiedBy
    ModifiedBy string `json:"modifiedBy"`


    

}

// String returns a JSON representation of the model
func (o *Greeting) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Greeting) MarshalJSON() ([]byte, error) {
    type Alias Greeting

    if GreetingMarshalled {
        return []byte("{}"), nil
    }
    GreetingMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        OwnerType string `json:"ownerType"`
        
        Owner Domainentity `json:"owner"`
        
        AudioFile Greetingaudiofile `json:"audioFile"`
        
        AudioTTS string `json:"audioTTS"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        CreatedBy string `json:"createdBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

