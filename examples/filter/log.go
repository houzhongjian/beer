package filter

import (
	"github.com/houzhongjian/beer"
	"github.com/houzhongjian/beer/log"
)

func Log(c *beer.Context) {
	log.Info(c.IP,c.Url,c.UserAgent)
}
