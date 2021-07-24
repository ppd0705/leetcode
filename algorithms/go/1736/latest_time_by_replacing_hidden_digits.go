package main

func maximumTime(time string) string {
	arr := []byte(time)

	if arr[0] == '?'{
		if arr[1] >= '4' && arr[1] <= '9' {
			arr[0] = '1'
		} else {
			arr[0] = '2'
		}
	}
	if arr[1] == '?' {
		if arr[0] == '2' {
			arr[1] = '3'
		} else {
			arr[1] = '9'
		}
	}
	if arr[3] == '?' {
		arr[3] = '5'
	}
	if arr[4] == '?' {
		arr[4] = '9'
	}
	return string(arr)
}
