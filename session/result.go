package session

import "github.com/houzhongjian/beer/utils"

type SessionResult struct {
	val string
}

func (res SessionResult) Int() (int, error) {
	return utils.ParseInt(res.val)
}

func (res SessionResult) String() string {
	return res.val
}
