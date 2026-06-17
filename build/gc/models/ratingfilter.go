package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatingfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatingfilterDud struct { 
    


    


    


    

}

// Ratingfilter
type Ratingfilter struct { 
    // Operator - The comparison operator for review rating filtering.
    Operator string `json:"operator"`


    // From - The lower bound for the Between operator
    From int `json:"from"`


    // To - The upper bound for the Between operator
    To int `json:"to"`


    // Values - One or more rating values to filter by
    Values []int `json:"values"`

}

// String returns a JSON representation of the model
func (o *Ratingfilter) String() string {
    
    
    
     o.Values = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratingfilter) MarshalJSON() ([]byte, error) {
    type Alias Ratingfilter

    if RatingfilterMarshalled {
        return []byte("{}"), nil
    }
    RatingfilterMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        From int `json:"from"`
        
        To int `json:"to"`
        
        Values []int `json:"values"`
        *Alias
    }{

        


        


        


        
        Values: []int{0},
        

        Alias: (*Alias)(u),
    })
}

