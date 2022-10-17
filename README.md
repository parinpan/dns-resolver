# DNS Resolver by Fachrin
A web app to query DNS records of a hostname along with its nameservers. Check it out live demo version here: https://resolver.fachr.in


## Program Flow
![DNSResolverFlow](https://user-images.githubusercontent.com/14908455/196105029-739ddfd5-c23e-443f-a31c-171243b24e8a.png)

## How to Build & Run (Docker Required)
If you don't have Docker installed on your machine, please do install first. Check out: https://docs.docker.com/get-docker/

### Build
Execute following command in your terminal
```bash
make build
```

### Run
Execute following command in your terminal
```bash
make run
```
Run command will build and run the image at the same time. A docker container will run on: http://localhost:9999, open it on your browser to interact with the app


## Preview
A preview of the web app querying DNS records (ANY) of **tesla.com** returning records for the hostname and its nameservers

![Screenshot 2022-10-17 at 08 41 04](https://user-images.githubusercontent.com/14908455/196106475-3f70803e-770d-4cbd-94e6-e15f2b8f1231.png)

Check out: https://resolver.fachr.in :)