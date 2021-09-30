package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Topic
type Topic struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Published
    Published bool `json:"published"`


    // Strictness
    Strictness string `json:"strictness"`


    // Programs
    Programs []Baseprogramentity `json:"programs"`


    // Tags
    Tags []string `json:"tags"`


    // Dialect
    Dialect string `json:"dialect"`


    // Participants
    Participants string `json:"participants"`


    // Phrases
    Phrases []Phrase `json:"phrases"`


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
func (o *Topic) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Programs = []Baseprogramentity{{}} 
    
    
    
     o.Tags = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
     o.Phrases = []Phrase{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topic) MarshalJSON() ([]byte, error) {
    type Alias Topic

    if TopicMarshalled {
        return []byte("{}"), nil
    }
    TopicMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        Strictness string `json:"strictness"`
        
        Programs []Baseprogramentity `json:"programs"`
        
        Tags []string `json:"tags"`
        
        Dialect string `json:"dialect"`
        
        Participants string `json:"participants"`
        
        Phrases []Phrase `json:"phrases"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        
        PublishedBy Addressableentityref `json:"publishedBy"`
        
        DatePublished time.Time `json:"datePublished"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Programs: []Baseprogramentity{{}},
        

        

        
        Tags: []string{""},
        

        

        

        

        

        

        
        Phrases: []Phrase{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

