package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaqMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaqDud struct { 
    Question string `json:"question"`


    Answer string `json:"answer"`


    SourceUri string `json:"sourceUri"`


    DocumentUrl string `json:"documentUrl"`


    DocumentDisplayName string `json:"documentDisplayName"`


    Confidence float32 `json:"confidence"`

}

// Faq
type Faq struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Faq) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faq) MarshalJSON() ([]byte, error) {
    type Alias Faq

    if FaqMarshalled {
        return []byte("{}"), nil
    }
    FaqMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

