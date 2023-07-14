BINARYNAME=ynab-metrics
BINARYPATH=target
BINARY = ${BINARYPATH}/${BINARYNAME}
CFGFILE=config.json

build:
	go build -o ${BINARY} -v

run:
	./${BINARY} --config ${CFGFILE}
