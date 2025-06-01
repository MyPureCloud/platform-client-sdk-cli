package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TesttopicphrasephraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TesttopicphrasephraseDud struct { 
    


    

}

// Testtopicphrasephrase
type Testtopicphrasephrase struct { 
    // Text - The phrase text
    Text string `json:"text"`


    // Strictness - The phrase strictness, default value is null
    Strictness string `json:"strictness"`

}

// String returns a JSON representation of the model
func (o *Testtopicphrasephrase) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testtopicphrasephrase) MarshalJSON() ([]byte, error) {
    type Alias Testtopicphrasephrase

    if TesttopicphrasephraseMarshalled {
        return []byte("{}"), nil
    }
    TesttopicphrasephraseMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Strictness string `json:"strictness"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

