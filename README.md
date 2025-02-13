### Motivation

Unikernels seem like a really interesting way to tackle certain architectural problems effectively.  I wanted to dive in and explore them a bit (just scratching the surface, though) and see what they're all about.

### About

`gifkernel` is a proof-of-concept project—a small, single-purpose web app built for quick and easy deployment as a unikernel on AWS EC2.

### Requirements

> Note: Currently tested on Linux.

*   [Go](https://go.dev/)
      <sub>for the web server (tested on 1.23.4)</sub>
*   [Ops](https://ops.city/)
     <sub>to create and deploy the unikernel</sub>
*   [Giphy](https://developers.giphy.com/) <sub>API key</sub>
*   [Vegeta](https://github.com/tsenart/vegeta) <sub>(optional - for load testing)</sub>

### Usage

1.  Clone this repository.
2.  Open `config.json` and provide your information, especially the S3 bucket name.
3.  Open `www/index.html` and add your Giphy API key.
4.  Build:   
   ```GOOS=linux go build -o gifkernel server.go```
5. Run:<br />
   ```ops run gifkernel```
   <br /><sub>then visit http://localhost:8888.</sub>

### Test

A basic load test program is available under `cmd/bench.go`. You can test locally or against your deployed server.

### Deploy

* Ensure your ~/.aws/credentials file is setup.
* Execute the `deploy.sh` script.
