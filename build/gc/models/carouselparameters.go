package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CarouselparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CarouselparametersDud struct { 
    

}

// Carouselparameters - Template parameters for carousel components
type Carouselparameters struct { 
    // CardParameters - A list of carousel cards with their respective template parameters
    CardParameters []Cardparameters `json:"cardParameters"`

}

// String returns a JSON representation of the model
func (o *Carouselparameters) String() string {
     o.CardParameters = []Cardparameters{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Carouselparameters) MarshalJSON() ([]byte, error) {
    type Alias Carouselparameters

    if CarouselparametersMarshalled {
        return []byte("{}"), nil
    }
    CarouselparametersMarshalled = true

    return json.Marshal(&struct {
        
        CardParameters []Cardparameters `json:"cardParameters"`
        *Alias
    }{

        
        CardParameters: []Cardparameters{{}},
        

        Alias: (*Alias)(u),
    })
}

