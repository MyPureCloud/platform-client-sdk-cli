package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebeventbrowserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebeventbrowserDud struct { 
    


    


    

}

// Webeventbrowser
type Webeventbrowser struct { 
    // Family - Browser family (e.g. Chrome, Safari, Firefox).
    Family string `json:"family"`


    // Version - Browser version (e.g. 68.0.3440.84).
    Version string `json:"version"`


    // Lang - Language the browser is set to. Must conform to BCP 47.
    Lang string `json:"lang"`

}

// String returns a JSON representation of the model
func (o *Webeventbrowser) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webeventbrowser) MarshalJSON() ([]byte, error) {
    type Alias Webeventbrowser

    if WebeventbrowserMarshalled {
        return []byte("{}"), nil
    }
    WebeventbrowserMarshalled = true

    return json.Marshal(&struct {
        
        Family string `json:"family"`
        
        Version string `json:"version"`
        
        Lang string `json:"lang"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

