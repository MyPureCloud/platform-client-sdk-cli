package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentvideosettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentvideosettingsDud struct { 
    


    


    


    

}

// Agentvideosettings - The settings for Agent Video
type Agentvideosettings struct { 
    // AllowCamera - whether or not agent camera is allowed
    AllowCamera bool `json:"allowCamera"`


    // AllowScreenShare - whether or not agent screen share is allowed
    AllowScreenShare bool `json:"allowScreenShare"`


    // Background - background for agent
    Background string `json:"background"`


    // BackgroundImage - background image settings for agent
    BackgroundImage Backgroundimagesettings `json:"backgroundImage"`

}

// String returns a JSON representation of the model
func (o *Agentvideosettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentvideosettings) MarshalJSON() ([]byte, error) {
    type Alias Agentvideosettings

    if AgentvideosettingsMarshalled {
        return []byte("{}"), nil
    }
    AgentvideosettingsMarshalled = true

    return json.Marshal(&struct {
        
        AllowCamera bool `json:"allowCamera"`
        
        AllowScreenShare bool `json:"allowScreenShare"`
        
        Background string `json:"background"`
        
        BackgroundImage Backgroundimagesettings `json:"backgroundImage"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

