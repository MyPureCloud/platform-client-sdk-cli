package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingretentioncursorentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingretentioncursorentitylistingDud struct { 
    


    


    


    

}

// Recordingretentioncursorentitylisting
type Recordingretentioncursorentitylisting struct { 
    // Entities
    Entities []Recordingretention `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Recordingretentioncursorentitylisting) String() string {
     o.Entities = []Recordingretention{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingretentioncursorentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Recordingretentioncursorentitylisting

    if RecordingretentioncursorentitylistingMarshalled {
        return []byte("{}"), nil
    }
    RecordingretentioncursorentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Recordingretention `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Recordingretention{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

