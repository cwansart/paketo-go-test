# Paketo Go test

This is a small sample app that tests how to build an OCI container using Paketo.io.

## Build the container

To build the container, you need to install Docker and Paketo first. Check their corresponding documentations for that.

To build the _paketo-go-test_ app run the following command:

```
$ pack build paketo-go-test --path . --builder paketobuildpacks/builder:base
```

If you see the message _Successfully built image paketo-go-test_ then you're ready to start the newly build container.

## Run the container

You can simply run it by the given name with Docker:

```
$ docker run --rm -p 8080:8080 paketo-go-test
```