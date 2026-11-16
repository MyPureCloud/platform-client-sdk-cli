package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptoutlinesegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptoutlinesegmentDud struct { 
    


    


    


    

}

// Transcriptoutlinesegment
type Transcriptoutlinesegment struct { 
    // Header - The header of the outline segment.
    Header string `json:"header"`


    // Description - The description of the outline segment.
    Description string `json:"description"`


    // FirstPhraseStartTimeMs - The start time of the first phrase of the segment in milliseconds.
    FirstPhraseStartTimeMs int `json:"firstPhraseStartTimeMs"`


    // LastPhraseStartTimeMs - The start time of the last phrase of the segment in milliseconds.
    LastPhraseStartTimeMs int `json:"lastPhraseStartTimeMs"`

}

// String returns a JSON representation of the model
func (o *Transcriptoutlinesegment) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptoutlinesegment) MarshalJSON() ([]byte, error) {
    type Alias Transcriptoutlinesegment

    if TranscriptoutlinesegmentMarshalled {
        return []byte("{}"), nil
    }
    TranscriptoutlinesegmentMarshalled = true

    return json.Marshal(&struct {
        
        Header string `json:"header"`
        
        Description string `json:"description"`
        
        FirstPhraseStartTimeMs int `json:"firstPhraseStartTimeMs"`
        
        LastPhraseStartTimeMs int `json:"lastPhraseStartTimeMs"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

