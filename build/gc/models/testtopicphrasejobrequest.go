package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TesttopicphrasejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TesttopicphrasejobrequestDud struct { 
    


    

}

// Testtopicphrasejobrequest
type Testtopicphrasejobrequest struct { 
    // Topic - The topic to test
    Topic Testtopicphrasetopic `json:"topic"`


    // TranscriptsFilters - The transcripts filters
    TranscriptsFilters Transcriptsfilters `json:"transcriptsFilters"`

}

// String returns a JSON representation of the model
func (o *Testtopicphrasejobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testtopicphrasejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Testtopicphrasejobrequest

    if TesttopicphrasejobrequestMarshalled {
        return []byte("{}"), nil
    }
    TesttopicphrasejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Topic Testtopicphrasetopic `json:"topic"`
        
        TranscriptsFilters Transcriptsfilters `json:"transcriptsFilters"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

