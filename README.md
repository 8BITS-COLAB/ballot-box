# Ballot box

## Descentralized ballot-box for voting

- Run `go run ./main.go voter -n registry_value` to create a new voter and keystore (descentralized).
- Run `go run ./main.go voter -a` show voter address.
- Run `go run ./main.go voter -r` show voter registry.
- Run `go run ./main.go candidate -l` to list candidates.
- Run `go run ./main.go keystore` show keystore private key.
- Run `go run ./main.go vote -c candidate_code` to vote for a candidate.
- Run `go run ./main.go vote -s` candidates and votes.
- Run `go run ./main.go vote -i` verify votes blockchain integrity.
