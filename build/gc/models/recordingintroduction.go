package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingintroductionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingintroductionDud struct { 
    


    

}

// Recordingintroduction
type Recordingintroduction struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the subtitle.
    Subtitle string `json:"subtitle"`

}

// String returns a JSON representation of the model
func (o *Recordingintroduction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingintroduction) MarshalJSON() ([]byte, error) {
    type Alias Recordingintroduction

    if RecordingintroductionMarshalled {
        return []byte("{}"), nil
    }
    RecordingintroductionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

