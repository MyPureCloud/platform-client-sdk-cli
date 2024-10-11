package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluutteranceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluutteranceDud struct { 
    Id string `json:"id"`


    


    

}

// Nluutterance
type Nluutterance struct { 
    


    // Source - The source of the utterance.
    Source string `json:"source"`


    // Segments - The list of segments that that constitute this utterance for the given intent.
    Segments []Nluutterancesegment `json:"segments"`

}

// String returns a JSON representation of the model
func (o *Nluutterance) String() string {
    
     o.Segments = []Nluutterancesegment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluutterance) MarshalJSON() ([]byte, error) {
    type Alias Nluutterance

    if NluutteranceMarshalled {
        return []byte("{}"), nil
    }
    NluutteranceMarshalled = true

    return json.Marshal(&struct {
        
        Source string `json:"source"`
        
        Segments []Nluutterancesegment `json:"segments"`
        *Alias
    }{

        


        


        
        Segments: []Nluutterancesegment{{}},
        

        Alias: (*Alias)(u),
    })
}

