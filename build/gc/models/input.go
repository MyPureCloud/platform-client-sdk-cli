package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InputDud struct { 
    


    


    


    


    


    


    


    

}

// Input - Input component configuration
type Input struct { 
    // Id - Unique identifier for the input field
    Id string `json:"id"`


    // Title - Title of the input field
    Title string `json:"title"`


    // Subtitle - Subtitle of the input field
    Subtitle string `json:"subtitle"`


    // PlaceholderText - Placeholder text for the input
    PlaceholderText string `json:"placeholderText"`


    // IsMultipleLine - Whether the input supports multiple lines
    IsMultipleLine bool `json:"isMultipleLine"`


    // IsRequired - Whether the input is required
    IsRequired bool `json:"isRequired"`


    // KeyboardType - Type of keyboard to be shown
    KeyboardType string `json:"keyboardType"`


    // AutoCompleteType - A string value representing the keyboard and system information about the expected semantic meaning for the content that users enter
    AutoCompleteType string `json:"autoCompleteType"`

}

// String returns a JSON representation of the model
func (o *Input) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Input) MarshalJSON() ([]byte, error) {
    type Alias Input

    if InputMarshalled {
        return []byte("{}"), nil
    }
    InputMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        PlaceholderText string `json:"placeholderText"`
        
        IsMultipleLine bool `json:"isMultipleLine"`
        
        IsRequired bool `json:"isRequired"`
        
        KeyboardType string `json:"keyboardType"`
        
        AutoCompleteType string `json:"autoCompleteType"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

