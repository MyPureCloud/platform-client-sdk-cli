package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementfilterrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementfilterrangeDud struct { 
    


    


    


    


    


    

}

// Journeyviewelementfilterrange - the range of attribute values to filter on. At least one comparator must be defined
type Journeyviewelementfilterrange struct { 
    // Lt - comparator: less than
    Lt Journeyviewelementfilterrangedata `json:"lt"`


    // Lte - comparator: less than or equal
    Lte Journeyviewelementfilterrangedata `json:"lte"`


    // Gt - comparator: greater than
    Gt Journeyviewelementfilterrangedata `json:"gt"`


    // Gte - comparator: greater than or equal
    Gte Journeyviewelementfilterrangedata `json:"gte"`


    // Eq - comparator: is equal to
    Eq Journeyviewelementfilterrangedata `json:"eq"`


    // Neq - comparator: is not equal to
    Neq Journeyviewelementfilterrangedata `json:"neq"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilterrange) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementfilterrange) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementfilterrange

    if JourneyviewelementfilterrangeMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementfilterrangeMarshalled = true

    return json.Marshal(&struct {
        
        Lt Journeyviewelementfilterrangedata `json:"lt"`
        
        Lte Journeyviewelementfilterrangedata `json:"lte"`
        
        Gt Journeyviewelementfilterrangedata `json:"gt"`
        
        Gte Journeyviewelementfilterrangedata `json:"gte"`
        
        Eq Journeyviewelementfilterrangedata `json:"eq"`
        
        Neq Journeyviewelementfilterrangedata `json:"neq"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

