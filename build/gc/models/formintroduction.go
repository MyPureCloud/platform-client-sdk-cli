package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormintroductionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormintroductionDud struct { 
    


    


    


    

}

// Formintroduction - Form introduction section with title, subtitle, image, and button text
type Formintroduction struct { 
    // Title - Title of the introduction
    Title string `json:"title"`


    // ImageUrl - URL of the image to display
    ImageUrl string `json:"imageUrl"`


    // Subtitle - Subtitle of the introduction
    Subtitle string `json:"subtitle"`


    // ButtonText - Text for the start button
    ButtonText string `json:"buttonText"`

}

// String returns a JSON representation of the model
func (o *Formintroduction) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formintroduction) MarshalJSON() ([]byte, error) {
    type Alias Formintroduction

    if FormintroductionMarshalled {
        return []byte("{}"), nil
    }
    FormintroductionMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        ImageUrl string `json:"imageUrl"`
        
        Subtitle string `json:"subtitle"`
        
        ButtonText string `json:"buttonText"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

