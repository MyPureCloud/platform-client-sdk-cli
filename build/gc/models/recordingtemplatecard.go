package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingtemplatecardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingtemplatecardDud struct { 
    


    


    

}

// Recordingtemplatecard
type Recordingtemplatecard struct { 
    // Header - The template header.
    Header Recordingtemplateheader `json:"header"`


    // Body - The template body.
    Body Recordingtemplatebody `json:"body"`


    // Buttons - Template buttons.
    Buttons []Recordingtemplatebutton `json:"buttons"`

}

// String returns a JSON representation of the model
func (o *Recordingtemplatecard) String() string {
    
    
     o.Buttons = []Recordingtemplatebutton{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingtemplatecard) MarshalJSON() ([]byte, error) {
    type Alias Recordingtemplatecard

    if RecordingtemplatecardMarshalled {
        return []byte("{}"), nil
    }
    RecordingtemplatecardMarshalled = true

    return json.Marshal(&struct {
        
        Header Recordingtemplateheader `json:"header"`
        
        Body Recordingtemplatebody `json:"body"`
        
        Buttons []Recordingtemplatebutton `json:"buttons"`
        *Alias
    }{

        


        


        
        Buttons: []Recordingtemplatebutton{{}},
        

        Alias: (*Alias)(u),
    })
}

