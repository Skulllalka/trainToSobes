package solutions

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const timeoutLimit = 10

type Result struct {
	msg string
	err error
}

func download(urls []string) ([]string, error) {
	var wg sync.WaitGroup
	c := make(chan Result)
	var err error
	retValue:= make([]string, 0, len(urls))

	for _, url := range urls{
		wg.Add(1)
		go func(u string){
			defer wg.Done()
			v:= fakeDownload(u)
			c<-v
		}(url)
	}

	go func(){
		wg.Wait()
		close(c)
	}()
	
	for res := range c{
		if res.err!= nil{
			err = errors.Join(err, res.err)
		}
		retValue = append(retValue, res.msg)
	}

	return retValue, err
}

func fakeDownload(url string) Result {
	r:= rand.Intn(100)

	time.Sleep(time.Duration(r)* time.Millisecond)
	if r > timeoutLimit {
		return Result{err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url))}
	}
	return Result{msg: fmt.Sprintf("downloaded data from %s \n", url)}

}

func UrlsMain() {
	msgs, err := download([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(msgs)
}
