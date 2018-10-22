package twitter

import "encoding/json"

type List struct {
	URLs []string `json:"urls"`
}

type listWraper struct {
	Obj userListInfo `json:"userListInfo"`
}

type userListInfo List

func (l *List) UnmarshalJSON(b []byte) error {
	var wrapper listWraper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*l = List(wrapper.Obj)
	return nil
}
