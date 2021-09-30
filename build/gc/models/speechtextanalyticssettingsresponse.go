package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SpeechtextanalyticssettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SpeechtextanalyticssettingsresponseDud struct { 
    


    

}

// Speechtextanalyticssettingsresponse
type Speechtextanalyticssettingsresponse struct { 
    // DefaultProgram - Setting to choose name for the default program for topic detection
    DefaultProgram Addressableentityref `json:"defaultProgram"`


    // ExpectedDialects - Setting to choose expected dialects
    ExpectedDialects []string `json:"expectedDialects"`

}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsresponse) String() string {
    
    
    
    
    
    
     o.ExpectedDialects = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Speechtextanalyticssettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Speechtextanalyticssettingsresponse

    if SpeechtextanalyticssettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    SpeechtextanalyticssettingsresponseMarshalled = true

    return json.Marshal(&struct { 
        DefaultProgram Addressableentityref `json:"defaultProgram"`
        
        ExpectedDialects []string `json:"expectedDialects"`
        
        *Alias
    }{
        

        

        

        
        ExpectedDialects: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

