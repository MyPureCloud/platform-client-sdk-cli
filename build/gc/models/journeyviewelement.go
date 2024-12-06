package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementDud struct { 
    


    


    


    


    


    

}

// Journeyviewelement - An element within a journey view
type Journeyviewelement struct { 
    // Id - The unique identifier of the element within the elements list
    Id string `json:"id"`


    // Name - The unique name of the element within the view
    Name string `json:"name"`


    // Attributes - Required attributes of the element
    Attributes Journeyviewelementattributes `json:"attributes"`


    // DisplayAttributes - Attributes that defines the visualization of the element in the journey view
    DisplayAttributes Journeyviewelementdisplayattributes `json:"displayAttributes"`


    // Filter - Any filters applied to this element
    Filter Journeyviewelementfilter `json:"filter"`


    // FollowedBy - A list of JourneyViewLink objects, listing the elements downstream of this element
    FollowedBy []Journeyviewlink `json:"followedBy"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelement) String() string {
    
    
    
    
    
     o.FollowedBy = []Journeyviewlink{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelement) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelement

    if JourneyviewelementMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Attributes Journeyviewelementattributes `json:"attributes"`
        
        DisplayAttributes Journeyviewelementdisplayattributes `json:"displayAttributes"`
        
        Filter Journeyviewelementfilter `json:"filter"`
        
        FollowedBy []Journeyviewlink `json:"followedBy"`
        *Alias
    }{

        


        


        


        


        


        
        FollowedBy: []Journeyviewlink{{}},
        

        Alias: (*Alias)(u),
    })
}

