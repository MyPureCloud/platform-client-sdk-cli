package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentarticleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentarticleDud struct { 
    


    Content Articlecontent `json:"content"`


    

}

// Documentarticle
type Documentarticle struct { 
    // Title - The title of the Article.
    Title string `json:"title"`


    


    // Alternatives - List of Alternative questions related to the title which helps in improving the likelihood of a match to user query.
    Alternatives []string `json:"alternatives"`

}

// String returns a JSON representation of the model
func (o *Documentarticle) String() string {
    
     o.Alternatives = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentarticle) MarshalJSON() ([]byte, error) {
    type Alias Documentarticle

    if DocumentarticleMarshalled {
        return []byte("{}"), nil
    }
    DocumentarticleMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Alternatives []string `json:"alternatives"`
        *Alias
    }{

        


        


        
        Alternatives: []string{""},
        

        Alias: (*Alias)(u),
    })
}

