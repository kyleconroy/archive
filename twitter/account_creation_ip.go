package twitter

import (
	"encoding/json"
	"net"
)

type AccountCreationIP struct {
	AccountID      int64  `json:"accountId,string"`
	UserCreationIP net.IP `json:"userCreationIp"`
}

type accountCreationIPObj struct {
	IP accountCreationIP `json:"accountCreationIp"`
}

type accountCreationIP AccountCreationIP

func (ip *AccountCreationIP) UnmarshalJSON(b []byte) error {
	var wrapper accountCreationIPObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*ip = AccountCreationIP(wrapper.IP)
	return nil
}
