package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PromptassetuploadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PromptassetuploadDud struct { 
    Url string `json:"url"`


    Headers map[string]string `json:"headers"`

}

// Promptassetupload
type Promptassetupload struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Promptassetupload) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Promptassetupload) MarshalJSON() ([]byte, error) {
    type Alias Promptassetupload

    if PromptassetuploadMarshalled {
        return []byte("{}"), nil
    }
    PromptassetuploadMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

