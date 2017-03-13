# Ampbeat

Welcome to Ampbeat v0.0.1

This beat handle both docker logs and metrics in a Swarm cluster context adding meta data as stack, service name to logs/metrics.
It listen Docker containers events and for each new started container open logs and metrics streams to publish the events.

It publishes, memory, net, io, cpu metrics and all logs.


Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/appcelerator/ampbeat`

## Getting Started with Ampbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7 min



### Build

To build the binary for Ampbeat run the command below. This will generate a binary
in the same directory with the name ampbeat.

```
git clone git@github.com:appcelerator/ampbeat
cd ampbeat
make create-image
```
or directly use the docker hub image, pulling it:

```
docker pull appcelerator/ampbeat:0.0.1
(or tag latest)
```

ampbeat settings can be updated in the file ampbeat-confimage.yml, before creating the image

### Run

To run Ampbeat in a docker swarm context:

```
docker service create --with-registry-auth --network aNetwork --name ampbeat \
    --label io.amp.role="infrastructure" \
    --mode global \
    --mount type=bind,source=/var/run/docker.sock,target=/var/run/docker.sock \
    appcelerator/ampbeat
```

Where the network "aNetwork" is the same than Elasticsearch or Logstash one



### Test

To test Ampbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/ampbeat.template.json and etc/ampbeat.asciidoc

```
make update
```


### Cleanup

To clean  Ampbeat source code, run the following commands:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
