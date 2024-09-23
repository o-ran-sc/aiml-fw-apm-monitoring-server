package subscribe

type subscribeItem struct {
	done chan (bool)
}

type subscribeMap map[string]subscribeItem

type SubscribeInfo struct {
	Host string `json:"host"`
	Name string `json:"name"`
}
