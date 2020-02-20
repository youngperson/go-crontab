# 本地测试使用

## docker安装单机版etcd
docker run -d  \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=/etcd-data \
  --name etcd quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://0.0.0.0:2380

## docker安装单机版MongoDB
docker pull mongo
docker run --name mongodb-standalone -v /data/mongo:/data/db -p 27017:27017 -d mongo
