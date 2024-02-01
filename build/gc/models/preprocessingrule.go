package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PreprocessingruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PreprocessingruleDud struct { 
    


    


    


    

}

// Preprocessingrule
type Preprocessingrule struct { 
    // Find - The regular expression to which file lines are to be matched.
    Find string `json:"find"`


    // ReplaceWith - The string to be substituted for each match.
    ReplaceWith string `json:"replaceWith"`


    // Global - Replaces all matching substrings in every line.
    Global bool `json:"global"`


    // IgnoreCase - Enables case-insensitive matching
    IgnoreCase bool `json:"ignoreCase"`

}

// String returns a JSON representation of the model
func (o *Preprocessingrule) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Preprocessingrule) MarshalJSON() ([]byte, error) {
    type Alias Preprocessingrule

    if PreprocessingruleMarshalled {
        return []byte("{}"), nil
    }
    PreprocessingruleMarshalled = true

    return json.Marshal(&struct {
        
        Find string `json:"find"`
        
        ReplaceWith string `json:"replaceWith"`
        
        Global bool `json:"global"`
        
        IgnoreCase bool `json:"ignoreCase"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

