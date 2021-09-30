package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptaggregatequeryresponseDud struct { 
    

}

// Transcriptaggregatequeryresponse
type Transcriptaggregatequeryresponse struct { 
    // Results
    Results []Transcriptaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryresponse) String() string {
    
    
     o.Results = []Transcriptaggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Transcriptaggregatequeryresponse

    if TranscriptaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    TranscriptaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Transcriptaggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Transcriptaggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

