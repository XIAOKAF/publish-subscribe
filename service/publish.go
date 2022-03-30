package service

func Publish(publisher string, info string, pubSub map[string][]string) {
	getInfo := make(map[string]string)
	channel := make(chan string)
	for _, value := range pubSub[publisher] {
		go func() {
			channel <- info
		}()
		getInfo[value] = <-channel
		println(publisher + "给" + value + "发来了一条信息：" + getInfo[value])
	}
}
