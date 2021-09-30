package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListitemcomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListitemcomponentDud struct { 
    


    


    


    


    


    


    

}

// Listitemcomponent - An entry in a List template.
type Listitemcomponent struct { 
    // Id - An ID assigned to this list item.
    Id string `json:"id"`


    // Rmid - An ID of the rich message instance.
    Rmid string `json:"rmid"`


    // VarType - The type of list item to render.
    VarType string `json:"type"`


    // Image - URL of an image.
    Image string `json:"image"`


    // Title - The main headline of the list item.
    Title string `json:"title"`


    // Description - Text to show in the list item description.
    Description string `json:"description"`


    // Actions - The list item actions.
    Actions Contentactions `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Listitemcomponent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listitemcomponent) MarshalJSON() ([]byte, error) {
    type Alias Listitemcomponent

    if ListitemcomponentMarshalled {
        return []byte("{}"), nil
    }
    ListitemcomponentMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Rmid string `json:"rmid"`
        
        VarType string `json:"type"`
        
        Image string `json:"image"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        Actions Contentactions `json:"actions"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

