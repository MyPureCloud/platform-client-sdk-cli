package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulejobresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Learningmodulejobresponse - Learning module job response
type Learningmodulejobresponse struct { 
    


    // Status - The status of learning module job
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Learningmodulejobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulejobresponse

    if LearningmodulejobresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

