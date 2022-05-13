package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptionsettingsDud struct { 
    


    


    

}

// Transcriptionsettings
type Transcriptionsettings struct { 
    // Transcription - Setting to enable/disable transcription capability
    Transcription string `json:"transcription"`


    // TranscriptionConfidenceThreshold - Configure confidence threshold. The possible values are from 1 to 100.
    TranscriptionConfidenceThreshold int `json:"transcriptionConfidenceThreshold"`


    // ContentSearchEnabled - Setting to enable/disable content search
    ContentSearchEnabled bool `json:"contentSearchEnabled"`

}

// String returns a JSON representation of the model
func (o *Transcriptionsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptionsettings) MarshalJSON() ([]byte, error) {
    type Alias Transcriptionsettings

    if TranscriptionsettingsMarshalled {
        return []byte("{}"), nil
    }
    TranscriptionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Transcription string `json:"transcription"`
        
        TranscriptionConfidenceThreshold int `json:"transcriptionConfidenceThreshold"`
        
        ContentSearchEnabled bool `json:"contentSearchEnabled"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

