package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AudiostateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AudiostateDud struct { 
    


    

}

// Audiostate
type Audiostate struct { 
    // CanHear - Indicates that this communication's audio allows its participant to hear others.
    CanHear bool `json:"canHear"`


    // CanSpeak - Indicates that this communication's audio allows others to hear this participant.
    CanSpeak bool `json:"canSpeak"`

}

// String returns a JSON representation of the model
func (o *Audiostate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Audiostate) MarshalJSON() ([]byte, error) {
    type Alias Audiostate

    if AudiostateMarshalled {
        return []byte("{}"), nil
    }
    AudiostateMarshalled = true

    return json.Marshal(&struct {
        
        CanHear bool `json:"canHear"`
        
        CanSpeak bool `json:"canSpeak"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

