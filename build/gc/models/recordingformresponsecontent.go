package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingformresponsecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingformresponsecontentDud struct { 
    


    

}

// Recordingformresponsecontent
type Recordingformresponsecontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // ButtonResponse - Button response content.
    ButtonResponse Buttonresponse `json:"buttonResponse"`

}

// String returns a JSON representation of the model
func (o *Recordingformresponsecontent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingformresponsecontent) MarshalJSON() ([]byte, error) {
    type Alias Recordingformresponsecontent

    if RecordingformresponsecontentMarshalled {
        return []byte("{}"), nil
    }
    RecordingformresponsecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        ButtonResponse Buttonresponse `json:"buttonResponse"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

