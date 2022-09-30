var cut_index = 0
if task.headers[0].Number.Uint64() > 2000
	task = nil
for i, header := range task.headers {
	if header.Number > 2000 {
		cut_index = i
		break
	}
}
if cut_index > 0
	task.headers = task.headers[:(len(task.headers)-cut_index)]
