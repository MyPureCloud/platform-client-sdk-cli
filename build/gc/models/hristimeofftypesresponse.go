package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HristimeofftypesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HristimeofftypesresponseDud struct { 
    


    

}

// Hristimeofftypesresponse
type Hristimeofftypesresponse struct { 
    // Job - The asynchronous job handling the query
    Job Hristimeofftypesjobreference `json:"job"`


    // Entities - List of time off types. It is available only via notification
    Entities []Hristimeofftyperesponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Hristimeofftypesresponse) String() string {
    
     o.Entities = []Hristimeofftyperesponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Hristimeofftypesresponse) MarshalJSON() ([]byte, error) {
    type Alias Hristimeofftypesresponse

    if HristimeofftypesresponseMarshalled {
        return []byte("{}"), nil
    }
    HristimeofftypesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Hristimeofftypesjobreference `json:"job"`
        
        Entities []Hristimeofftyperesponse `json:"entities"`
        *Alias
    }{

        


        
        Entities: []Hristimeofftyperesponse{{}},
        

        Alias: (*Alias)(u),
    })
}

