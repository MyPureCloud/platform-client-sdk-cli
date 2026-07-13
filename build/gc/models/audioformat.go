package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AudioformatMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AudioformatDud struct { 
    Channels int `json:"channels"`


    BitsPerSample int `json:"bitsPerSample"`


    SampleRate int `json:"sampleRate"`


    Encoding string `json:"encoding"`

}

// Audioformat
type Audioformat struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Audioformat) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Audioformat) MarshalJSON() ([]byte, error) {
    type Alias Audioformat

    if AudioformatMarshalled {
        return []byte("{}"), nil
    }
    AudioformatMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

