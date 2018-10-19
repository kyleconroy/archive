package twitter

import (
	"encoding/json"
	"net"
)

type AccountCreationIP struct {
	AccountID      int64
	UserCreationIP net.IP
}

type accountCreationIPObj struct {
	IP accountCreationIP `json:"accountCreationIp"`
}

type accountCreationIP struct {
	AccountID      int64  `json:"accountId,string"`
	UserCreationIP string `json:"userCreationIp"`
}

func (ip *AccountCreationIP) UnmarshalJSON(b []byte) error {
	var wrapper accountCreationIPObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	obj := wrapper.IP

	ip.AccountID = obj.AccountID
	ip.UserCreationIP = net.ParseIP(obj.UserCreationIP)
	return nil
}
