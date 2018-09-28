BINARYNAME=ynab-metrics
BINARYPATH=target
BINARY = ${BINARYPATH}/${BINARYNAME}

TOKENFILE:=.accessToken
TOKEN:=$(shell cat ${TOKENFILE})

build:
	go build -o ${BINARY} -v

run:
	go build -o ${BINARY} -v ./...
	./${BINARY} --token=${TOKEN}
