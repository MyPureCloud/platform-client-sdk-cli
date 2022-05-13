package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsesetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsesetDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Responseset
type Responseset struct { 
    


    // Name - The name of the ResponseSet.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Responses - Map of disposition identifiers to reactions. For example: {\"disposition.classification.callable.person\": {\"reactionType\": \"transfer\"}}.
    Responses map[string]Reaction `json:"responses"`


    // BeepDetectionEnabled - Whether to enable answering machine beep detection
    BeepDetectionEnabled bool `json:"beepDetectionEnabled"`


    

}

// String returns a JSON representation of the model
func (o *Responseset) String() string {
    
    
     o.Responses = map[string]Reaction{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseset) MarshalJSON() ([]byte, error) {
    type Alias Responseset

    if ResponsesetMarshalled {
        return []byte("{}"), nil
    }
    ResponsesetMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Responses map[string]Reaction `json:"responses"`
        
        BeepDetectionEnabled bool `json:"beepDetectionEnabled"`
        *Alias
    }{

        


        


        


        


        


        
        Responses: map[string]Reaction{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

