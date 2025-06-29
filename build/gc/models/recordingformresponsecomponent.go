package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingformresponsecomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingformresponsecomponentDud struct { 
    


    

}

// Recordingformresponsecomponent
type Recordingformresponsecomponent struct { 
    // Id - The id of the component in the original message.
    Id string `json:"id"`


    // Component - The content object capturing component response from the original message.
    Component Recordingformresponsecontent `json:"component"`

}

// String returns a JSON representation of the model
func (o *Recordingformresponsecomponent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingformresponsecomponent) MarshalJSON() ([]byte, error) {
    type Alias Recordingformresponsecomponent

    if RecordingformresponsecomponentMarshalled {
        return []byte("{}"), nil
    }
    RecordingformresponsecomponentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Component Recordingformresponsecontent `json:"component"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

