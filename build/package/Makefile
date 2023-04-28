BINARYNAME=ynab-metrics
BINARYPATH=target
BINARY = ${BINARYPATH}/${BINARYNAME}
CFGFILE=config.json

TOKENFILE:=.accessToken
TOKEN:=$(shell cat ${TOKENFILE})

build:
	go build -o ${BINARY} -v

run:
	./${BINARY} --config ${CFGFILE}
