Run the below command to build and spin up the docker container

If first time running:
```docker build -t web-server .```

Run docker:
```docker run -p 80:80 web-server```


Local build:

```
cd cmd
go run .
```

Testing

Mocks created using the `mockery` package

```
brew install mockery

// From root
mockery --all
```