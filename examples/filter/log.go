package filter

import (
	"github.com/houzhongjian/beer"
	"log"
)

func Log(c *beer.Context) {
	log.Println(c.IP,c.Url,c.UserAgent)
}
