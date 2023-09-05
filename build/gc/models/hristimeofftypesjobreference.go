package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HristimeofftypesjobreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HristimeofftypesjobreferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Hristimeofftypesjobreference
type Hristimeofftypesjobreference struct { 
    


    // Status - The status of the job
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Hristimeofftypesjobreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Hristimeofftypesjobreference) MarshalJSON() ([]byte, error) {
    type Alias Hristimeofftypesjobreference

    if HristimeofftypesjobreferenceMarshalled {
        return []byte("{}"), nil
    }
    HristimeofftypesjobreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

