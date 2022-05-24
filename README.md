# Smallbank

smallbank是H-store中用于测试分布式数据库性能的一个benchmark（https://hstore.cs.brown.edu/documentation/deployment/benchmarks/）。该benchmark也经常用于fabric的性能测试中。本项目提供了基于go语言的面向fabric的smallbank链码。

## Quick Start

~~~bash
./network.sh up -s couchdb
./network.sh createChannel -s couchdb

./network.sh deployCC -ccn smallbank -ccp ../chaincode/smallbank/go/ -ccl go -ccep "OR('Org1MSP.member','Org2MSP.member')" -cccg ../chaincode/smallbank/collections_config.json

~~~

## Other

一个调用该链码的sdk实例：https://github.com/hsj576/nisl-fabric-sdk