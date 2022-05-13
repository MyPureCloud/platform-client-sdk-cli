package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WidgetclientconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WidgetclientconfigDud struct { 
    


    


    


    

}

// Widgetclientconfig
type Widgetclientconfig struct { 
    // V1
    V1 Widgetclientconfigv1 `json:"v1"`


    // V2
    V2 interface{} `json:"v2"`


    // V1Http
    V1Http Widgetclientconfigv1http `json:"v1-http"`


    // ThirdParty
    ThirdParty interface{} `json:"third-party"`

}

// String returns a JSON representation of the model
func (o *Widgetclientconfig) String() string {
    
     o.V2 = Interface{} 
    
     o.ThirdParty = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widgetclientconfig) MarshalJSON() ([]byte, error) {
    type Alias Widgetclientconfig

    if WidgetclientconfigMarshalled {
        return []byte("{}"), nil
    }
    WidgetclientconfigMarshalled = true

    return json.Marshal(&struct {
        
        V1 Widgetclientconfigv1 `json:"v1"`
        
        V2 interface{} `json:"v2"`
        
        V1Http Widgetclientconfigv1http `json:"v1-http"`
        
        ThirdParty interface{} `json:"third-party"`
        *Alias
    }{

        


        
        V2: Interface{},
        


        


        
        ThirdParty: Interface{},
        

        Alias: (*Alias)(u),
    })
}

