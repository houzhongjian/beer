package service

import (
	"github.com/houzhongjian/beer"
	"io/ioutil"
	"log"
)

func Rem(c *beer.Context)  {
	sess, _ := beer.Session().Start(c)
	beer.Session().Destroy(sess)
}

func Default(c *beer.Context) {

}

func Login(c *beer.Context) {
	log.Println(c.Param("id"))
	log.Println(c.Param("name"))
	log.Println(c.UserAgent)
}

func Detail(c *beer.Context) {
	b, err := ioutil.ReadAll(c.Body)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	log.Println(string(b))
}
