module github.com/statefb/serial-to-text-api

go 1.15

replace local.packages/config => ./config

replace local.packages/converter => ./converter

replace local.packages/data => ./data

replace local.packages/handler => ./handler

replace local.packages/mymiddleware => ./mymiddleware

replace local.packages/sender => ./sender

replace local.packages/states => ./states

replace local.packages/gen/models => ./gen/models

replace local.packages/gen/restapi => ./gen/restapi

replace local.packages/gen/restapi/serialtocsv => ./gen/restapi/serialtocsv

replace local.packages/gen/restapi/serialtocsv/common => ./gen/restapi/serialtocsv/common

require (
	github.com/go-openapi/loads v0.20.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/jlaffaye/ftp v0.0.0-20201112195030-9aae4d151126 // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	local.packages/config v0.0.0-00010101000000-000000000000 // indirect
	local.packages/converter v0.0.0-00010101000000-000000000000 // indirect
	local.packages/data v0.0.0-00010101000000-000000000000 // indirect
	local.packages/gen/models v0.0.0-00010101000000-000000000000 // indirect
	local.packages/gen/restapi v0.0.0-00010101000000-000000000000
	local.packages/gen/restapi/serialtocsv v0.0.0-00010101000000-000000000000
	local.packages/gen/restapi/serialtocsv/common v0.0.0-00010101000000-000000000000 // indirect
	local.packages/handler v0.0.0-00010101000000-000000000000 // indirect
	local.packages/mymiddleware v0.0.0-00010101000000-000000000000 // indirect
	local.packages/sender v0.0.0-00010101000000-000000000000 // indirect
	local.packages/states v0.0.0-00010101000000-000000000000 // indirect
)
