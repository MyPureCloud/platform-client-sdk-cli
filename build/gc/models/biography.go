package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BiographyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BiographyDud struct { 
    


    


    


    


    

}

// Biography
type Biography struct { 
    // Biography - Personal detailed description
    Biography string `json:"biography"`


    // Interests
    Interests []string `json:"interests"`


    // Hobbies
    Hobbies []string `json:"hobbies"`


    // Spouse
    Spouse string `json:"spouse"`


    // Education - User education details
    Education []Education `json:"education"`

}

// String returns a JSON representation of the model
func (o *Biography) String() string {
    
    
    
    
    
    
     o.Interests = []string{""} 
    
    
    
     o.Hobbies = []string{""} 
    
    
    
    
    
    
    
     o.Education = []Education{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Biography) MarshalJSON() ([]byte, error) {
    type Alias Biography

    if BiographyMarshalled {
        return []byte("{}"), nil
    }
    BiographyMarshalled = true

    return json.Marshal(&struct { 
        Biography string `json:"biography"`
        
        Interests []string `json:"interests"`
        
        Hobbies []string `json:"hobbies"`
        
        Spouse string `json:"spouse"`
        
        Education []Education `json:"education"`
        
        *Alias
    }{
        

        

        

        
        Interests: []string{""},
        

        

        
        Hobbies: []string{""},
        

        

        

        

        
        Education: []Education{{}},
        

        
        Alias: (*Alias)(u),
    })
}

