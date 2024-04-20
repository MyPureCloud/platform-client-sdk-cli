package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramtranscriptionenginesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramtranscriptionenginesDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Programtranscriptionengines
type Programtranscriptionengines struct { 
    


    // Program - The ID of the program
    Program Baseprogramentity `json:"program"`


    // TranscriptionEngines - The program transcription engine settings
    TranscriptionEngines []Programtranscriptionengine `json:"transcriptionEngines"`


    // ModifiedBy - The user last modified the record
    ModifiedBy Addressableentityref `json:"modifiedBy"`


    // DateModified - The last modified date of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Programtranscriptionengines) String() string {
    
     o.TranscriptionEngines = []Programtranscriptionengine{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programtranscriptionengines) MarshalJSON() ([]byte, error) {
    type Alias Programtranscriptionengines

    if ProgramtranscriptionenginesMarshalled {
        return []byte("{}"), nil
    }
    ProgramtranscriptionenginesMarshalled = true

    return json.Marshal(&struct {
        
        Program Baseprogramentity `json:"program"`
        
        TranscriptionEngines []Programtranscriptionengine `json:"transcriptionEngines"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        
        TranscriptionEngines: []Programtranscriptionengine{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

