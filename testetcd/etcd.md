```shell
# 创建网络
docker network create --driver bridge --subnet=10.2.36.0/16 --gateway=10.2.1.1 etcdnet
```

```shell
# etcd-node1
docker run -d \
-p 8479:2379 \
-p 8381:2380 \
--name etcd-node1 \
--network=etcdnet \
--ip 10.2.36.1 \
k8s.gcr.io/etcd:3.5.3-0 \
etcd \
--name node1 \
--advertise-client-urls http://10.2.36.1:2379 \
--initial-advertise-peer-urls http://10.2.36.1:2380 \
--listen-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 \
--initial-cluster-token etcd-cluster \
--initial-cluster "node1=http://10.2.36.1:2380,node2=http://10.2.36.2:2380,node3=http://10.2.36.3:2380" \
--initial-cluster-state new
```

```shell
# etcd-node2
docker run -d \
-p 8579:2379 \
-p 8382:2380 \
--name etcd-node2 \
--network=etcdnet \
--ip 10.2.36.2 \
k8s.gcr.io/etcd:3.5.3-0 \
etcd \
--name node2 \
--advertise-client-urls http://10.2.36.2:2379 \
--initial-advertise-peer-urls http://10.2.36.2:2380 \
--listen-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 \
--initial-cluster-token etcd-cluster \
--initial-cluster "node1=http://10.2.36.1:2380,node2=http://10.2.36.2:2380,node3=http://10.2.36.3:2380" \
--initial-cluster-state new
```

```shell
# etcd-node3
docker run -d \
-p 8679:2379 \
-p 8383:2380 \
--name etcd-node3 \
--network=etcdnet \
--ip 10.2.36.3 \
k8s.gcr.io/etcd:3.5.3-0 \
etcd \
--name node3 \
--advertise-client-urls http://10.2.36.3:2379 \
--initial-advertise-peer-urls http://10.2.36.3:2380 \
--listen-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 \
--initial-cluster-token etcd-cluster \
--initial-cluster "node1=http://10.2.36.1:2380,node2=http://10.2.36.2:2380,node3=http://10.2.36.3:2380" \
--initial-cluster-state new
```

```shell
# etcd 检查集群状态
etcdctl --write-out=table --endpoints=localhost:8479,localhost:8579,localhost:8679  endpoint health
          +----------------+--------+-------------+-------+
          |    ENDPOINT    | HEALTH |    TOOK     | ERROR |
          +----------------+--------+-------------+-------+
          | localhost:8479 |   true | 10.527416ms |       |
          | localhost:8679 |   true | 10.807167ms |       |
          | localhost:8579 |   true | 10.669792ms |       |
          +----------------+--------+-------------+-------+

# 查看集群节点初试配置
etcdctl --write-out=table --endpoints=localhost:8479,localhost:8579,localhost:8679 endpoint status
```

```shell
# 集群存读键值对
etcdctl --endpoints=localhost:8479,localhost:8579,localhost:8679 put foo "Hello World!"
etcdctl --endpoints=localhost:8479,localhost:8579,localhost:8679 get foo
etcdctl --endpoints=localhost:8479,localhost:8579,localhost:8679 --write-out="json" get foo
```

