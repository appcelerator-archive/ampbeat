docker service create --network ampcore_infra --name kibana \
--replicas 1 \
--label io.amp.role="infrastructure" \
-p 5601:5601 \
-e ELASTICSEARCH_URL=http://elasticsearch:9200 \
blacktop/kibana:5.1
