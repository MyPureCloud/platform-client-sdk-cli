package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Program
type Program struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Published
    Published bool `json:"published"`


    // Topics
    Topics []Basetopicentitiy `json:"topics"`


    // Tags
    Tags []string `json:"tags"`


    // ModifiedBy
    ModifiedBy Addressableentityref `json:"modifiedBy"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // PublishedBy
    PublishedBy Addressableentityref `json:"publishedBy"`


    // DatePublished - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DatePublished time.Time `json:"datePublished"`


    

}

// String returns a JSON representation of the model
func (o *Program) String() string {
    
    
    
     o.Topics = []Basetopicentitiy{{}} 
     o.Tags = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Program) MarshalJSON() ([]byte, error) {
    type Alias Program

    if ProgramMarshalled {
        return []byte("{}"), nil
    }
    ProgramMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        Topics []Basetopicentitiy `json:"topics"`
        
        Tags []string `json:"tags"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        
        PublishedBy Addressableentityref `json:"publishedBy"`
        
        DatePublished time.Time `json:"datePublished"`
        *Alias
    }{

        


        


        


        


        
        Topics: []Basetopicentitiy{{}},
        


        
        Tags: []string{""},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

