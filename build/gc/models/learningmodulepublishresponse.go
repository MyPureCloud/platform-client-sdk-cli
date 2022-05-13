package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepublishresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepublishresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Learningmodulepublishresponse - Learning module publish response
type Learningmodulepublishresponse struct { 
    


    // Version - The version of published learning module
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepublishresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepublishresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepublishresponse

    if LearningmodulepublishresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepublishresponseMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

