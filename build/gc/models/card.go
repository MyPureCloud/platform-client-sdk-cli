package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CardDud struct { 
    


    


    


    


    

}

// Card
type Card struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Description - Text to show in the description.
    Description string `json:"description"`


    // Url - URL of an image.
    Url string `json:"url"`


    // DefaultAction - The default action to be taken.
    DefaultAction Cardaction `json:"defaultAction"`


    // Actions - List of possible action objects.
    Actions []Cardaction `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Card) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Actions = []Cardaction{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Card) MarshalJSON() ([]byte, error) {
    type Alias Card

    if CardMarshalled {
        return []byte("{}"), nil
    }
    CardMarshalled = true

    return json.Marshal(&struct { 
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Url string `json:"url"`
        
        DefaultAction Cardaction `json:"defaultAction"`
        
        Actions []Cardaction `json:"actions"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Actions: []Cardaction{{}},
        

        
        Alias: (*Alias)(u),
    })
}

