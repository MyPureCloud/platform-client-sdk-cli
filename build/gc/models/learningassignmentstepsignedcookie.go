package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentstepsignedcookieMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentstepsignedcookieDud struct { 
    Url string `json:"url"`


    CookieValues map[string]string `json:"cookieValues"`

}

// Learningassignmentstepsignedcookie
type Learningassignmentstepsignedcookie struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Learningassignmentstepsignedcookie) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentstepsignedcookie) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentstepsignedcookie

    if LearningassignmentstepsignedcookieMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentstepsignedcookieMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

