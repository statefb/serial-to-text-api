![redoc](https://github.com/statefb/serial-to-text-api/workflows/redoc/badge.svg?branch=master)
![release](https://github.com/statefb/serial-to-text-api/workflows/release/badge.svg)

# Serial to text API

Convert serial streaming into text format such as csv, json with key data.

## Purpose and Motivation

TODO

## Features

WIP

## Install
You can choose from one of the following ways:
* Download from [release](https://github.com/statefb/serial-to-text-api/releases)
* `$git clone` & `$go build`

## Usage
### Setup configuration
* Place [config.yml](https://github.com/statefb/serial-to-text-api/blob/master/server/config.yml) on the same directory as executable.
* For detail, see [Configuration](#Configuration)

### Run server
Port number can be set arbitrarily.  
`$./serial-to-text --host=0.0.0.0 --port=3000`

## Example
Following is the python snippet:
```snippet.py
import requests

url = 'http://serial-to-text:3000/'

res = requests.get(url + 'status')
print(res)  # status ("waiting", "collecting", etc) can be checked

# start collecting data
res = requests.put(url + 'start', body=json.dump({
    "lotId": "lot123",
    "bagId": "bag456"
}))

# get current collected signal
res = requests.get(url + 'data')
print(res)
# output example:
"""
[
    {
        "timestamp": "2021-01-01T10:00:01",
        "value": "S  T  1.523kg"
    },
    {
        "timestamp": "2021-01-01T10:00:11",
        "value": "S  T  2.341kg"
    }
]
"""

# stop collecting data
res = requests.put(url + 'stop')

# send collected data with key information.
res = requests.put(url + 'send')
# output json file example:
"""
{
  "key": {
    "lotId": "lot123",
    "bagId": "bag456"
  },
  "data": [
    {
        "timestamp": "2021-01-01T10:00:01",
        "value": "S  T  1.523kg"
    },
    {
        "timestamp": "2021-01-01T10:00:11",
        "value": "S  T  2.341kg"
    }
  ]
}

# NOTE: run reset before next lot.
res = requests.put(url + 'reset')
"""

```

## Configuration
* serial
    * mock: If true, use dummy serial port for testing. Default is false.
    * maxRecordSize: Max number of serial signal to keep per one lot. Excess will be ignored.
    * bufferSize: Buffer size to read signal per one trial.
    * endSignature: Termination signal. In many case it can be newline code. See the reference manual of device which you will use.
    * name: Serial port name to connect.
    * baud: Baud rate.
    * databits: Length of data bits.
    * parity: Parity. Should be one of "None", "Odd" or "Even".
    * stopbits: Length of Stopbit. Should be one of "1", "2" or "15"(1 and half).
* sender
    * method: Method to send signal. Should be "ftp". Of course, ftp is not recommended if place the server on the internet.
    * ftp
        * path: Path to save.
        * uri: URI for ftp server in the form of "<address>:<port>".
        * name: User name.
        * password: Password to login.

## Other References
* [API Reference](https://statefb.github.io/serial-to-text-api/)