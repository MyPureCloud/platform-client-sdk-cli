package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SdklibraryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SdklibraryDud struct { 
    


    

}

// Sdklibrary
type Sdklibrary struct { 
    // Name - The name of the SDK.
    Name string `json:"name"`


    // Version - The version of the SDK.
    Version string `json:"version"`

}

// String returns a JSON representation of the model
func (o *Sdklibrary) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sdklibrary) MarshalJSON() ([]byte, error) {
    type Alias Sdklibrary

    if SdklibraryMarshalled {
        return []byte("{}"), nil
    }
    SdklibraryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version string `json:"version"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

