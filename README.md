# Multitenant Hosting Platform

### What it does?
1. Hosts apps either locally or on cloud
2. For local: spins up go routines as servers for multiple apps.
3. For cloud: Deploys and hosts on netlify

### How to run locally?

export NETLIFY_TOKEN="Personal access token"
export INDEX_FILE_PATH="path to idex file that returns hello world"
`go run main.go`

This will run the platform server

### How to create Apps and host them?

#### Using cli
1. Get the cli code from https://github.com/goyalbhumika/multitenant-hosting-cli
2. Build the cli `go build -o mycli`
3. Deploy using './mycli create-app --name crazy-app-107 --deploy_type cloud' for cloud
4. Deploy using './mycli create-app --name crazy-app-107 --deploy_type local' for local

##### To access local apps
Output of cli would be: `App created successfully: {Name:app-2 Port:9002 DNS:app-2.gravityfalls42.hitchhiker}`
1. Add `app-2.gravityfalls42.hitchhiker` to /etc/hosts
2. Access using curl http://app-2.gravityfalls42.hitchhiker:9002
3. The above command returns "Hello world app-2"

##### To access cloud
1. Deployment should be done on netlify with a DNS provided by netlify.
2. To create custom DNS, use duck DNS and add to netlify custom domain.
