package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Topicjob
type Topicjob struct { 
    


    // State
    State string `json:"state"`


    // Topics
    Topics []Basetopicentitiy `json:"topics"`


    // CreatedBy
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Topicjob) String() string {
    
    
    
    
    
    
    
    
     o.Topics = []Basetopicentitiy{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicjob) MarshalJSON() ([]byte, error) {
    type Alias Topicjob

    if TopicjobMarshalled {
        return []byte("{}"), nil
    }
    TopicjobMarshalled = true

    return json.Marshal(&struct { 
        
        
        State string `json:"state"`
        
        Topics []Basetopicentitiy `json:"topics"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Topics: []Basetopicentitiy{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

