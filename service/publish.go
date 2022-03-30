package service

func Publish(publisher string, info string, pubSub map[string][]string) {
	channel := make(chan string)
	for _, value := range pubSub[publisher] {
		go func() {
			channel <- info
		}()
		value = <-channel
		println(value)
	}
}
