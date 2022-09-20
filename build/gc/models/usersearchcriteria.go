package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersearchcriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersearchcriteriaDud struct { 
    


    


    


    


    


    


    


    


    

}

// Usersearchcriteria
type Usersearchcriteria struct { 
    // EndValue - The end value of the range. This field is used for range search types.
    EndValue string `json:"endValue"`


    // Values - A list of values for the search to match against
    Values []string `json:"values"`


    // StartValue - The start value of the range. This field is used for range search types.
    StartValue string `json:"startValue"`


    // Value - A value for the search to match against
    Value string `json:"value"`


    // Operator - How to apply this search criteria against other criteria
    Operator string `json:"operator"`


    // Group - Groups multiple conditions
    Group []Usersearchcriteria `json:"group"`


    // DateFormat - Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
    DateFormat string `json:"dateFormat"`


    // Fields - Field names to search against
    Fields []string `json:"fields"`


    // VarType - Search Type
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Usersearchcriteria) String() string {
    
     o.Values = []string{""} 
    
    
    
     o.Group = []Usersearchcriteria{{}} 
    
     o.Fields = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersearchcriteria) MarshalJSON() ([]byte, error) {
    type Alias Usersearchcriteria

    if UsersearchcriteriaMarshalled {
        return []byte("{}"), nil
    }
    UsersearchcriteriaMarshalled = true

    return json.Marshal(&struct {
        
        EndValue string `json:"endValue"`
        
        Values []string `json:"values"`
        
        StartValue string `json:"startValue"`
        
        Value string `json:"value"`
        
        Operator string `json:"operator"`
        
        Group []Usersearchcriteria `json:"group"`
        
        DateFormat string `json:"dateFormat"`
        
        Fields []string `json:"fields"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        
        Values: []string{""},
        


        


        


        


        
        Group: []Usersearchcriteria{{}},
        


        


        
        Fields: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

