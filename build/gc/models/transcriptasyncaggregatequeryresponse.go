package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptasyncaggregatequeryresponseDud struct { 
    


    

}

// Transcriptasyncaggregatequeryresponse
type Transcriptasyncaggregatequeryresponse struct { 
    // Results
    Results []Transcriptaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Transcriptasyncaggregatequeryresponse) String() string {
     o.Results = []Transcriptaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Transcriptasyncaggregatequeryresponse

    if TranscriptasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    TranscriptasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Transcriptaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Transcriptaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

