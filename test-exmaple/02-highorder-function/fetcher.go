package fetcher

func fetchData(data func (url string) (string, error), url string) (string,error) {
	return data(url)
}
