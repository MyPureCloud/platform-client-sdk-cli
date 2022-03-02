package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscripttopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscripttopicDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    TopicPhrase string `json:"topicPhrase"`


    TranscriptPhrase string `json:"transcriptPhrase"`


    Confidence int `json:"confidence"`


    StartTimeMilliseconds int `json:"startTimeMilliseconds"`


    

}

// Transcripttopic
type Transcripttopic struct { 
    


    


    


    


    


    


    // Duration
    Duration Topicduration `json:"duration"`

}

// String returns a JSON representation of the model
func (o *Transcripttopic) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcripttopic) MarshalJSON() ([]byte, error) {
    type Alias Transcripttopic

    if TranscripttopicMarshalled {
        return []byte("{}"), nil
    }
    TranscripttopicMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        Duration Topicduration `json:"duration"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

