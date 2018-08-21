package config

import (
    "os"
    "fmt"
    "log"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type SignalServerRequestStringConfig struct {
    ListAll string
    RegMe   string
}

type SignalServerConfig struct {
    ListenPort int
}

type AddressInfo struct {
    IP   string
    Port int
}

type TotalConfig struct {
    Nodes      []*AddressInfo
    Signals    []*AddressInfo
    SSConfig      SignalServerConfig
    SSCommand     SignalServerRequestStringConfig
}

func NewTotalConfig() *TotalConfig {
    return &TotalConfig{Nodes : make([]*AddressInfo, 0, 2)}
}

func NewAddressInfo(ip string, port int) *AddressInfo {
    return &AddressInfo { IP:ip, Port:port}
}

func (self *TotalConfig) BuildConnectedNodesList() string {
    sb := strings.Builder{}
    for _, addInfo := range self.Nodes {
        sb.WriteString(fmt.Sprintf("%s:%d\n", addInfo.IP, addInfo.Port))
    }

    return sb.String()
}

func (self *TotalConfig) BuildListAllGetRequestList() []string {
    out := make([]string, 0, len(self.Signals))
    for _, addrInfo := range self.Signals {
        out = append(out, fmt.Sprintf("http://%s:%d/%s", addrInfo.IP, addrInfo.Port, self.SSCommand.ListAll))
    }
    
    return out
}
func (self *TotalConfig) BuildListPort() string {
    return fmt.Sprintf(":%d", self.SSConfig.ListenPort)
}

func (self *TotalConfig) BuildListAllURL() string {
    return fmt.Sprintf("/%s", self.SSCommand.ListAll)
}

func (self *TotalConfig) BuildRegMeURL() string {
    return fmt.Sprintf("/%s", self.SSCommand.RegMe)
}

func (self *TotalConfig) ReFileWithHardcodedValues() {
    self.SSConfig  = SignalServerConfig{ListenPort:30001}
    self.SSCommand = SignalServerRequestStringConfig{RegMe : "regme", ListAll : "listall"}

    self.Nodes      = make([]*AddressInfo, 0, 2)
    self.Nodes      = append(self.Nodes, NewAddressInfo("127.0.0.1", 30001))
    self.Nodes      = append(self.Nodes, NewAddressInfo("I am bad IP", 666))

    self.Signals    = make([]*AddressInfo, 0, 2)
    self.Signals    = append(self.Signals, NewAddressInfo("127.0.0.1", 30001))
    self.Signals    = append(self.Signals, NewAddressInfo("I am bad IP", 666))
    
    bdata, err := json.MarshalIndent(self, "  ", "\t")
    if err != nil {
        log.Fatal(err.Error())
    }

    file, err := os.OpenFile("config/config.cfg", os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer file.Close()

    _, err = file.Write(bdata)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func (self *TotalConfig) LoadConfig() {
    file, err := os.OpenFile("config/config.cfg", os.O_RDONLY,0666)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err.Error())
    }

    if err := json.Unmarshal(bytes, self); err != nil {
        log.Fatal(err.Error())
    }
}