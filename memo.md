# TODO
goroutineを使わず、自前でタイムアウト（serial.Read）


## memo

- Run following command inside docker on work folder:  
  `swagger generate server -a serialtocsv -A serialtocsv \ --strict-additional-properties -t ./server/gen -f ./swagger/serial-csv-converter.v1.yaml`

- `sudo chmod 777 work`
- move main.go under server directory
- make handler package under server and implement
- populate handler by editing configure\_\*\*\*.go
