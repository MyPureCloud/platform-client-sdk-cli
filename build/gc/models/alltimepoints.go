package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlltimepointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlltimepointsDud struct { 
    User Userreference `json:"user"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`


    AllTimePoints int `json:"allTimePoints"`

}

// Alltimepoints
type Alltimepoints struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Alltimepoints) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alltimepoints) MarshalJSON() ([]byte, error) {
    type Alias Alltimepoints

    if AlltimepointsMarshalled {
        return []byte("{}"), nil
    }
    AlltimepointsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

