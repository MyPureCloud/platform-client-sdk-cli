package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanjobexceptionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanjobexceptionDud struct { 
    


    

}

// Activityplanjobexception
type Activityplanjobexception struct { 
    // ExceptionType - The type of error
    ExceptionType string `json:"exceptionType"`


    // Occurrences - The occurrences in which this error occurred
    Occurrences []Activityplanoccurrencereference `json:"occurrences"`

}

// String returns a JSON representation of the model
func (o *Activityplanjobexception) String() string {
    
     o.Occurrences = []Activityplanoccurrencereference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanjobexception) MarshalJSON() ([]byte, error) {
    type Alias Activityplanjobexception

    if ActivityplanjobexceptionMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanjobexceptionMarshalled = true

    return json.Marshal(&struct {
        
        ExceptionType string `json:"exceptionType"`
        
        Occurrences []Activityplanoccurrencereference `json:"occurrences"`
        *Alias
    }{

        


        
        Occurrences: []Activityplanoccurrencereference{{}},
        

        Alias: (*Alias)(u),
    })
}

