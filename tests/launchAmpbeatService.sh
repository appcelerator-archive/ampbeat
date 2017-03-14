docker service create --with-registry-auth --network ampcore_infra --name ampbeat \
    --label io.amp.role="infrastructure" \
    --mode global \
    --mount type=volume,source=ampbeat,target=/containers \
    --mount type=bind,source=/var/run/docker.sock,target=/var/run/docker.sock \
    appcelerator/ampbeat
