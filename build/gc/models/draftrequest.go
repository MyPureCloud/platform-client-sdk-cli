package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftrequestDud struct { 
    


    

}

// Draftrequest
type Draftrequest struct { 
    // Intents - Draft intent object.
    Intents []Draftintents `json:"intents"`


    // Topics - Draft topic object.
    Topics []Drafttopicrequest `json:"topics"`

}

// String returns a JSON representation of the model
func (o *Draftrequest) String() string {
     o.Intents = []Draftintents{{}} 
     o.Topics = []Drafttopicrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draftrequest) MarshalJSON() ([]byte, error) {
    type Alias Draftrequest

    if DraftrequestMarshalled {
        return []byte("{}"), nil
    }
    DraftrequestMarshalled = true

    return json.Marshal(&struct {
        
        Intents []Draftintents `json:"intents"`
        
        Topics []Drafttopicrequest `json:"topics"`
        *Alias
    }{

        
        Intents: []Draftintents{{}},
        


        
        Topics: []Drafttopicrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

