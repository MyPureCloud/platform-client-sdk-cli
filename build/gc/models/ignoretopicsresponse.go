package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnoretopicsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnoretopicsresponseDud struct { 
    


    


    

}

// Ignoretopicsresponse
type Ignoretopicsresponse struct { 
    // TotalTopics - Total number of topics in current org
    TotalTopics int `json:"totalTopics"`


    // AddedTopics - Number of topics added in current request
    AddedTopics int `json:"addedTopics"`


    // UpdatedTopics - Number of topics updated in current request
    UpdatedTopics int `json:"updatedTopics"`

}

// String returns a JSON representation of the model
func (o *Ignoretopicsresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignoretopicsresponse) MarshalJSON() ([]byte, error) {
    type Alias Ignoretopicsresponse

    if IgnoretopicsresponseMarshalled {
        return []byte("{}"), nil
    }
    IgnoretopicsresponseMarshalled = true

    return json.Marshal(&struct {
        
        TotalTopics int `json:"totalTopics"`
        
        AddedTopics int `json:"addedTopics"`
        
        UpdatedTopics int `json:"updatedTopics"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

