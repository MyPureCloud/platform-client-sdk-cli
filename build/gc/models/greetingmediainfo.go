package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GreetingmediainfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GreetingmediainfoDud struct { 
    Id string `json:"id"`


    


    

}

// Greetingmediainfo
type Greetingmediainfo struct { 
    


    // MediaFileUri
    MediaFileUri string `json:"mediaFileUri"`


    // MediaImageUri
    MediaImageUri string `json:"mediaImageUri"`

}

// String returns a JSON representation of the model
func (o *Greetingmediainfo) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Greetingmediainfo) MarshalJSON() ([]byte, error) {
    type Alias Greetingmediainfo

    if GreetingmediainfoMarshalled {
        return []byte("{}"), nil
    }
    GreetingmediainfoMarshalled = true

    return json.Marshal(&struct { 
        
        
        MediaFileUri string `json:"mediaFileUri"`
        
        MediaImageUri string `json:"mediaImageUri"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

