package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GreetingaudiofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GreetingaudiofileDud struct { 
    


    


    

}

// Greetingaudiofile
type Greetingaudiofile struct { 
    // DurationMilliseconds
    DurationMilliseconds int `json:"durationMilliseconds"`


    // SizeBytes
    SizeBytes int `json:"sizeBytes"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Greetingaudiofile) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Greetingaudiofile) MarshalJSON() ([]byte, error) {
    type Alias Greetingaudiofile

    if GreetingaudiofileMarshalled {
        return []byte("{}"), nil
    }
    GreetingaudiofileMarshalled = true

    return json.Marshal(&struct { 
        DurationMilliseconds int `json:"durationMilliseconds"`
        
        SizeBytes int `json:"sizeBytes"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

