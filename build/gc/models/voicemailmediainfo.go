package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailmediainfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailmediainfoDud struct { 
    Id string `json:"id"`


    


    


    

}

// Voicemailmediainfo
type Voicemailmediainfo struct { 
    


    // MediaFileUri
    MediaFileUri string `json:"mediaFileUri"`


    // MediaImageUri
    MediaImageUri string `json:"mediaImageUri"`


    // WaveformData
    WaveformData []float32 `json:"waveformData"`

}

// String returns a JSON representation of the model
func (o *Voicemailmediainfo) String() string {
    
    
     o.WaveformData = []float32{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailmediainfo) MarshalJSON() ([]byte, error) {
    type Alias Voicemailmediainfo

    if VoicemailmediainfoMarshalled {
        return []byte("{}"), nil
    }
    VoicemailmediainfoMarshalled = true

    return json.Marshal(&struct {
        
        MediaFileUri string `json:"mediaFileUri"`
        
        MediaImageUri string `json:"mediaImageUri"`
        
        WaveformData []float32 `json:"waveformData"`
        *Alias
    }{

        


        


        


        
        WaveformData: []float32{0.0},
        

        Alias: (*Alias)(u),
    })
}

