package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormpageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormpageDud struct { 
    


    


    

}

// Formpage - A page in a form with title, subtitle, and components
type Formpage struct { 
    // Title - Title of the page
    Title string `json:"title"`


    // Subtitle - Subtitle of the page
    Subtitle string `json:"subtitle"`


    // PageComponents - Components on this page
    PageComponents []Formpagecomponent `json:"pageComponents"`

}

// String returns a JSON representation of the model
func (o *Formpage) String() string {
    
    
     o.PageComponents = []Formpagecomponent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formpage) MarshalJSON() ([]byte, error) {
    type Alias Formpage

    if FormpageMarshalled {
        return []byte("{}"), nil
    }
    FormpageMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        PageComponents []Formpagecomponent `json:"pageComponents"`
        *Alias
    }{

        


        


        
        PageComponents: []Formpagecomponent{{}},
        

        Alias: (*Alias)(u),
    })
}

