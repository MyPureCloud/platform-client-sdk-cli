package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListedprogramMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListedprogramDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Listedprogram
type Listedprogram struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Published
    Published bool `json:"published"`


    // TopicsCount
    TopicsCount int `json:"topicsCount"`


    // Tags
    Tags []string `json:"tags"`


    // ModifiedBy
    ModifiedBy Addressableentityref `json:"modifiedBy"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Listedprogram) String() string {
    
    
    
    
     o.Tags = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listedprogram) MarshalJSON() ([]byte, error) {
    type Alias Listedprogram

    if ListedprogramMarshalled {
        return []byte("{}"), nil
    }
    ListedprogramMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        TopicsCount int `json:"topicsCount"`
        
        Tags []string `json:"tags"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        
        Tags: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

