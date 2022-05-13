package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LanguageDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Language
type Language struct { 
    


    // Name - The language name.
    Name string `json:"name"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // State
    State string `json:"state"`


    // Version
    Version string `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Language) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Language) MarshalJSON() ([]byte, error) {
    type Alias Language

    if LanguageMarshalled {
        return []byte("{}"), nil
    }
    LanguageMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DateModified time.Time `json:"dateModified"`
        
        State string `json:"state"`
        
        Version string `json:"version"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

