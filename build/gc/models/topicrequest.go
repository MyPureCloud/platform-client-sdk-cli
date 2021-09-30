package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Topicrequest
type Topicrequest struct { 
    // Name - The topic name
    Name string `json:"name"`


    // Description - The topic description
    Description string `json:"description"`


    // Strictness - The topic strictness, default value is 72
    Strictness string `json:"strictness"`


    // ProgramIds - The ids of programs associated to the topic
    ProgramIds []string `json:"programIds"`


    // Tags - The topic tags
    Tags []string `json:"tags"`


    // Dialect - The topic dialect
    Dialect string `json:"dialect"`


    // Participants - The topic participants, default value is All
    Participants string `json:"participants"`


    // Phrases - The topic phrases
    Phrases []Phrase `json:"phrases"`

}

// String returns a JSON representation of the model
func (o *Topicrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ProgramIds = []string{""} 
    
    
    
     o.Tags = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
     o.Phrases = []Phrase{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicrequest) MarshalJSON() ([]byte, error) {
    type Alias Topicrequest

    if TopicrequestMarshalled {
        return []byte("{}"), nil
    }
    TopicrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Strictness string `json:"strictness"`
        
        ProgramIds []string `json:"programIds"`
        
        Tags []string `json:"tags"`
        
        Dialect string `json:"dialect"`
        
        Participants string `json:"participants"`
        
        Phrases []Phrase `json:"phrases"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        ProgramIds: []string{""},
        

        

        
        Tags: []string{""},
        

        

        

        

        

        

        
        Phrases: []Phrase{{}},
        

        
        Alias: (*Alias)(u),
    })
}

