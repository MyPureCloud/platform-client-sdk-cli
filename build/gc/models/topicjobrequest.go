package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicjobrequestDud struct { 
    

}

// Topicjobrequest
type Topicjobrequest struct { 
    // TopicIds - The ids of the topics used for this job
    TopicIds []string `json:"topicIds"`

}

// String returns a JSON representation of the model
func (o *Topicjobrequest) String() string {
     o.TopicIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Topicjobrequest

    if TopicjobrequestMarshalled {
        return []byte("{}"), nil
    }
    TopicjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        TopicIds []string `json:"topicIds"`
        *Alias
    }{

        
        TopicIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

