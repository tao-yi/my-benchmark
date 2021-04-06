benchmark:
  # -t threads
	# -c total number of http connections to keep open 
	# -d duration of the test, e.g. 2s, 2m, 2h
	wrk -t12 -c400 -d10s http://localhost:3999/

benchmark_gql:
	wrk -t12 -c400 -d10s -s ./gql.lua http://localhost:3999/

query: 
	curl -X POST \
	-H "Content-Type: application/json" \
	-d '{"query":"query{hello}"}' \
	http://localhost:3999/graphql