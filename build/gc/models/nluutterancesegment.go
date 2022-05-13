package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluutterancesegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluutterancesegmentDud struct { 
    


    

}

// Nluutterancesegment
type Nluutterancesegment struct { 
    // Text - The text of the segment.
    Text string `json:"text"`


    // Entity - The entity annotation of the segment.
    Entity Namedentityannotation `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Nluutterancesegment) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluutterancesegment) MarshalJSON() ([]byte, error) {
    type Alias Nluutterancesegment

    if NluutterancesegmentMarshalled {
        return []byte("{}"), nil
    }
    NluutterancesegmentMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Entity Namedentityannotation `json:"entity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

