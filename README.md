```shell
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.16.1
docker run -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.16.1


http://localhost:9200/lucene_query-tr/_doc/test
```
