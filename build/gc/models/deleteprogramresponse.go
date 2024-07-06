package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeleteprogramresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeleteprogramresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Deleteprogramresponse
type Deleteprogramresponse struct { 
    


    // TopicLinksJob
    TopicLinksJob Addressableentityref `json:"topicLinksJob"`


    

}

// String returns a JSON representation of the model
func (o *Deleteprogramresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deleteprogramresponse) MarshalJSON() ([]byte, error) {
    type Alias Deleteprogramresponse

    if DeleteprogramresponseMarshalled {
        return []byte("{}"), nil
    }
    DeleteprogramresponseMarshalled = true

    return json.Marshal(&struct {
        
        TopicLinksJob Addressableentityref `json:"topicLinksJob"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

