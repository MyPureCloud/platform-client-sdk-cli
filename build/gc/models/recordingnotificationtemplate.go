package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingnotificationtemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingnotificationtemplateDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Recordingnotificationtemplate
type Recordingnotificationtemplate struct { 
    


    // Language - Template language.
    Language string `json:"language"`


    // Header - The template header.
    Header Recordingtemplateheader `json:"header"`


    // Body - The template body.
    Body Recordingtemplatebody `json:"body"`


    // Buttons - Template buttons
    Buttons []Recordingtemplatebutton `json:"buttons"`


    // Footer - The template footer.
    Footer Recordingtemplatefooter `json:"footer"`

}

// String returns a JSON representation of the model
func (o *Recordingnotificationtemplate) String() string {
    
    
    
     o.Buttons = []Recordingtemplatebutton{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingnotificationtemplate) MarshalJSON() ([]byte, error) {
    type Alias Recordingnotificationtemplate

    if RecordingnotificationtemplateMarshalled {
        return []byte("{}"), nil
    }
    RecordingnotificationtemplateMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        
        Header Recordingtemplateheader `json:"header"`
        
        Body Recordingtemplatebody `json:"body"`
        
        Buttons []Recordingtemplatebutton `json:"buttons"`
        
        Footer Recordingtemplatefooter `json:"footer"`
        *Alias
    }{

        


        


        


        


        
        Buttons: []Recordingtemplatebutton{{}},
        


        

        Alias: (*Alias)(u),
    })
}

