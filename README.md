This TruSet Revealer API server is a service that monitors the commit-reveal poll contract, and reveals votes on behalf of TruSet users.

# Environment

Go doesn't always play nicely with relative paths as imports, so to work with or execute the code in this repo you will need to clone it into your Go path. You can achieve this by modifying your Go path, but here's how I did it using the default GOPATH value (should work on MacOS or Linux):

- Check out this repo inside your GO path, under `/src/github.com/TruSet`.
  - You can achieve this by modifying your Go path, but here's how I did it (should work on MacOS or Linux)
    - `cd $GOPATH/src`
    - `mkdir -p github.com/TruSet`
    - `cd github.com/TruSet`
    - `git clone https://github.com/TruSet/TruSetAPIServer.git`
- If you want to work with the code from a workspace elsewhere, set up a symbolic link, e.g.
  - `cd <your workspace>`
  - `ln -s $GOPATH/src/github.com/TruSet/TruSetAPIServer TruSetAPIServer`
- Install dependencies:
  - `cd` into the cloned repo directory, e.g `cd <your workspace>/TruSetAPIServer`
  - `go get ./...`

# Running locally

- Deploy the app from the `TruSet/TruSet` repo and start the app with `docker-compose up revealer-api revealer-worker`
- `go run main.go`
  - Run `go run main.go -h` for usage
