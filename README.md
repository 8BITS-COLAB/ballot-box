# Ballot box

## Descentralized ballot-box for voting

- Run `go run ./main.go voter -n registry_value -k secret_key` to create a new voter and keystore (descentralized).
- Run `go run ./main.go voter -a -k secret_key` show voter address.
- Run `go run ./main.go voter -r -k secret_key` show voter registry.
- Run `go run ./main.go candidate -l` to list candidates.
- Run `go run ./main.go keystore` show keystore private key.
- Run `go run ./main.go vote -c candidate_code -k secret_key` to vote for a candidate.
- Run `go run ./main.go vote -s` candidates and votes.
- Run `go run ./agent/agent.go` work to verify votes blockchain integrity.
