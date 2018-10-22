package twitter

import (
	"encoding/json"
	"net"
	"time"
)

type IPAudit struct {
	AccountID int64     `json:"accountId,string"`
	CreatedAt time.Time `json:"createdAt"`
	LoginIP   net.IP    `json:"loginIp"`
}

type auditObj struct {
	Obj ipaudit `json:"ipAudit"`
}

type ipaudit IPAudit

func (ip *IPAudit) UnmarshalJSON(b []byte) error {
	var wrapper auditObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*ip = IPAudit(wrapper.Obj)
	return nil
}
