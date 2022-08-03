# Ballot box

## Descentralized ballot-box for voting

- Run `go run ./main.go setup` to create a new ballot box with a basic configuration.
- Run `go run ./main.go voter -n registry_value -k secret_key` to create a new voter and keystore (descentralized).
- Run `go run ./main.go voter -a -k secret_key -p private_key` show voter address.
- Run `go run ./main.go voter -r -k secret_key -p private_key` show voter registry.
- Run `go run ./main.go candidate -l` to list candidates.
- Run `go run ./main.go vote -c candidate_code -k secret_key -p private_key` to vote for a candidate.
- Run `go run ./main.go vote -s` candidates and votes.
- Run `go run ./agent/agent.go` worker to verify votes blockchain integrity.
- Run `go run ./main.go network -l port` to listen on a port.
  > All peers will be connected to the same network and between them will be able to communicate.
- Open you browser and go to http://localhost:port to access the home page.
- Run `go run ./main.go config -s KEY=VALUE ` to set a configuration value.
- Run `go run ./main.go config -g KEY ` to get a configuration value.
- Run `go install` to substitute `go run ./main.go` to `ballot-box` command.
