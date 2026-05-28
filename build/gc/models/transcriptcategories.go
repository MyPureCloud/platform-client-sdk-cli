package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptcategoriesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptcategoriesDud struct { 
    


    

}

// Transcriptcategories
type Transcriptcategories struct { 
    // Includes - List of categories which need to be included in exact match criteria. This field is not mutually exclusive with excludes category list.
    Includes []string `json:"includes"`


    // Excludes - List of categories which need to be excluded in exact match criteria. This field is not mutually exclusive with includes category list.
    Excludes []string `json:"excludes"`

}

// String returns a JSON representation of the model
func (o *Transcriptcategories) String() string {
     o.Includes = []string{""} 
     o.Excludes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptcategories) MarshalJSON() ([]byte, error) {
    type Alias Transcriptcategories

    if TranscriptcategoriesMarshalled {
        return []byte("{}"), nil
    }
    TranscriptcategoriesMarshalled = true

    return json.Marshal(&struct {
        
        Includes []string `json:"includes"`
        
        Excludes []string `json:"excludes"`
        *Alias
    }{

        
        Includes: []string{""},
        


        
        Excludes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

