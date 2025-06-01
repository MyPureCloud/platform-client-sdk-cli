package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptsfiltersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptsfiltersDud struct { 
    


    


    


    


    

}

// Transcriptsfilters
type Transcriptsfilters struct { 
    // MediaType - The media type of the transcripts, default value is all 
    MediaType string `json:"mediaType"`


    // StartTimeMs - start time to filter by, default value is 7 days into the past
    StartTimeMs int `json:"startTimeMs"`


    // EndTimeMs - end time to filter by, default value is current time
    EndTimeMs int `json:"endTimeMs"`


    // Queues - list of queues ids to filter by
    Queues []string `json:"queues"`


    // Flows - list of flows ids to filter by
    Flows []string `json:"flows"`

}

// String returns a JSON representation of the model
func (o *Transcriptsfilters) String() string {
    
    
    
     o.Queues = []string{""} 
     o.Flows = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptsfilters) MarshalJSON() ([]byte, error) {
    type Alias Transcriptsfilters

    if TranscriptsfiltersMarshalled {
        return []byte("{}"), nil
    }
    TranscriptsfiltersMarshalled = true

    return json.Marshal(&struct {
        
        MediaType string `json:"mediaType"`
        
        StartTimeMs int `json:"startTimeMs"`
        
        EndTimeMs int `json:"endTimeMs"`
        
        Queues []string `json:"queues"`
        
        Flows []string `json:"flows"`
        *Alias
    }{

        


        


        


        
        Queues: []string{""},
        


        
        Flows: []string{""},
        

        Alias: (*Alias)(u),
    })
}

