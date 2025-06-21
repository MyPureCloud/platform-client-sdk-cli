package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneratedguidecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneratedguidecontentDud struct { 
    Instruction string `json:"instruction"`


    Variables []Variable `json:"variables"`


    Resources Guideversionresources `json:"resources"`

}

// Generatedguidecontent
type Generatedguidecontent struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Generatedguidecontent) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generatedguidecontent) MarshalJSON() ([]byte, error) {
    type Alias Generatedguidecontent

    if GeneratedguidecontentMarshalled {
        return []byte("{}"), nil
    }
    GeneratedguidecontentMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

