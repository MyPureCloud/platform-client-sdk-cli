package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatguestmediarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatguestmediarequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Webchatguestmediarequest - Object representing the guest model of a media request of a chat conversation.
type Webchatguestmediarequest struct { 
    


    // Name
    Name string `json:"name"`


    // Types - The types of media being requested.
    Types []string `json:"types"`


    // State - The state of the media request, one of PENDING|ACCEPTED|DECLINED|TIMEDOUT|CANCELLED|ERRORED.
    State string `json:"state"`


    // CommunicationId - The ID of the new media communication, if applicable.
    CommunicationId string `json:"communicationId"`


    // SecurityKey - The security information related to a media request.
    SecurityKey string `json:"securityKey"`


    

}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequest) String() string {
    
    
    
    
    
    
    
    
     o.Types = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatguestmediarequest) MarshalJSON() ([]byte, error) {
    type Alias Webchatguestmediarequest

    if WebchatguestmediarequestMarshalled {
        return []byte("{}"), nil
    }
    WebchatguestmediarequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Types []string `json:"types"`
        
        State string `json:"state"`
        
        CommunicationId string `json:"communicationId"`
        
        SecurityKey string `json:"securityKey"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Types: []string{""},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

