package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VideosettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VideosettingsDud struct { 
    


    


    

}

// Videosettings - The settings to enable video chat
type Videosettings struct { 
    // Enabled - whether or not video is enabled
    Enabled bool `json:"enabled"`


    // Agent - Video Settings for agent
    Agent Agentvideosettings `json:"agent"`


    // User - Video Settings for user
    User Uservideosettings `json:"user"`

}

// String returns a JSON representation of the model
func (o *Videosettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Videosettings) MarshalJSON() ([]byte, error) {
    type Alias Videosettings

    if VideosettingsMarshalled {
        return []byte("{}"), nil
    }
    VideosettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Agent Agentvideosettings `json:"agent"`
        
        User Uservideosettings `json:"user"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

