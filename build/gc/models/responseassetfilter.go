package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetfilterDud struct { 
    


    


    


    


    


    

}

// Responseassetfilter
type Responseassetfilter struct { 
    // EndValue - The end value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
    EndValue string `json:"endValue"`


    // Values - A list of values for the search to match against
    Values []string `json:"values"`


    // StartValue - The start value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
    StartValue string `json:"startValue"`


    // Fields - Field name to search against. Allowed Values: divisionId, name, contentLength, contentType, dateCreated
    Fields []string `json:"fields"`


    // Value - A value for the search to match against
    Value string `json:"value"`


    // VarType - How to apply this search criteria against other criteria
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Responseassetfilter) String() string {
    
    
    
    
    
    
     o.Values = []string{""} 
    
    
    
    
    
    
    
     o.Fields = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseassetfilter) MarshalJSON() ([]byte, error) {
    type Alias Responseassetfilter

    if ResponseassetfilterMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetfilterMarshalled = true

    return json.Marshal(&struct { 
        EndValue string `json:"endValue"`
        
        Values []string `json:"values"`
        
        StartValue string `json:"startValue"`
        
        Fields []string `json:"fields"`
        
        Value string `json:"value"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        
        Values: []string{""},
        

        

        

        

        
        Fields: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

