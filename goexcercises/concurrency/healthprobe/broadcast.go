package healthprobe

func Broadcast(resultsChannel <-chan Result, outs ...chan<- Result) {
	for result := range resultsChannel {
		for _, out := range outs {
			out <- result
		}
	}
	for _, out := range outs {
		close(out)
	}
}
