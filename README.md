# oss-go
aliyun oss tools

# build
```
go build -o oss-cli
```

# config
```
$ vim $HOME/.oss.conf
[credentialsConfig]
endpoint = "oss endpoint"
# eg: endpoint = "https://oss-cn-beijing.aliyuncs.com"
accessKeyID = "accessKeyID"
accessKeySecret = "accessKeySecret"
```

# use
```
./oss-cli put --bucket  bucket-test -f /tmp/data.tar.gz -r oss-object/data.tar.gz
```
