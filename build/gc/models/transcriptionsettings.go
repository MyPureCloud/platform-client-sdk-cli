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


    // LowLatencyTranscriptionEnabled - Boolean flag indicating whether low latency transcription via Notification API is enabled
    LowLatencyTranscriptionEnabled bool `json:"lowLatencyTranscriptionEnabled"`


    // ContentSearchEnabled - Setting to enable/disable content search
    ContentSearchEnabled bool `json:"contentSearchEnabled"`


    // PciDssRedactionEnabled - Setting to enable/disable PCI DSS Redaction
    PciDssRedactionEnabled bool `json:"pciDssRedactionEnabled"`


    // PiiRedactionEnabled - Setting to enable/disable PII Redaction
    PiiRedactionEnabled bool `json:"piiRedactionEnabled"`

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
        
        LowLatencyTranscriptionEnabled bool `json:"lowLatencyTranscriptionEnabled"`
        
        ContentSearchEnabled bool `json:"contentSearchEnabled"`
        
        PciDssRedactionEnabled bool `json:"pciDssRedactionEnabled"`
        
        PiiRedactionEnabled bool `json:"piiRedactionEnabled"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

