package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenrecordingmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenrecordingmetadataDud struct { 
    


    


    


    


    


    


    

}

// Screenrecordingmetadata
type Screenrecordingmetadata struct { 
    // TrackId
    TrackId string `json:"trackId"`


    // MediaId
    MediaId string `json:"mediaId"`


    // ScreenId
    ScreenId string `json:"screenId"`


    // OriginX
    OriginX int `json:"originX"`


    // OriginY
    OriginY int `json:"originY"`


    // Primary
    Primary bool `json:"primary"`


    // Main
    Main bool `json:"main"`

}

// String returns a JSON representation of the model
func (o *Screenrecordingmetadata) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenrecordingmetadata) MarshalJSON() ([]byte, error) {
    type Alias Screenrecordingmetadata

    if ScreenrecordingmetadataMarshalled {
        return []byte("{}"), nil
    }
    ScreenrecordingmetadataMarshalled = true

    return json.Marshal(&struct {
        
        TrackId string `json:"trackId"`
        
        MediaId string `json:"mediaId"`
        
        ScreenId string `json:"screenId"`
        
        OriginX int `json:"originX"`
        
        OriginY int `json:"originY"`
        
        Primary bool `json:"primary"`
        
        Main bool `json:"main"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

