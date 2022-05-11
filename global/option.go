package global

var (
	Proxy     string
	ProxyMode string
)

func PrintOptions() {
	Logger.Printf(AppName+" running with args: proxy=%v, proxyMode=%v",
		Proxy, ProxyMode)
}
