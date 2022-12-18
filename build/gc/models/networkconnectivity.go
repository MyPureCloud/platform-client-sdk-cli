package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NetworkconnectivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NetworkconnectivityDud struct { 
    


    


    


    

}

// Networkconnectivity
type Networkconnectivity struct { 
    // Carrier - The name of the network carrier.
    Carrier string `json:"carrier"`


    // BluetoothEnabled - Whether Bluetooth is enabled.
    BluetoothEnabled bool `json:"bluetoothEnabled"`


    // CellularEnabled - Whether Cellular is enabled.
    CellularEnabled bool `json:"cellularEnabled"`


    // WifiEnabled - Whether Wi-Fi is enabled.
    WifiEnabled bool `json:"wifiEnabled"`

}

// String returns a JSON representation of the model
func (o *Networkconnectivity) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Networkconnectivity) MarshalJSON() ([]byte, error) {
    type Alias Networkconnectivity

    if NetworkconnectivityMarshalled {
        return []byte("{}"), nil
    }
    NetworkconnectivityMarshalled = true

    return json.Marshal(&struct {
        
        Carrier string `json:"carrier"`
        
        BluetoothEnabled bool `json:"bluetoothEnabled"`
        
        CellularEnabled bool `json:"cellularEnabled"`
        
        WifiEnabled bool `json:"wifiEnabled"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

