package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListedtopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListedtopicDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Listedtopic
type Listedtopic struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Published
    Published bool `json:"published"`


    // Strictness
    Strictness string `json:"strictness"`


    // ProgramsCount
    ProgramsCount int `json:"programsCount"`


    // Tags
    Tags []string `json:"tags"`


    // Dialect
    Dialect string `json:"dialect"`


    // Participants
    Participants string `json:"participants"`


    // PhrasesCount
    PhrasesCount int `json:"phrasesCount"`


    // ModifiedBy
    ModifiedBy Addressableentityref `json:"modifiedBy"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Listedtopic) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Tags = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listedtopic) MarshalJSON() ([]byte, error) {
    type Alias Listedtopic

    if ListedtopicMarshalled {
        return []byte("{}"), nil
    }
    ListedtopicMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        Strictness string `json:"strictness"`
        
        ProgramsCount int `json:"programsCount"`
        
        Tags []string `json:"tags"`
        
        Dialect string `json:"dialect"`
        
        Participants string `json:"participants"`
        
        PhrasesCount int `json:"phrasesCount"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Tags: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

