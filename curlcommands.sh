curl -d '{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}' localhost:4000/v1/movies

# need to use 'jq' for creating json variables in bash perhaps.


echo '{ "foo": 123, "bar": 456 }' | jq '.foo'

#curl -H "Content-Type: application/json" --data @body.json http://localhost:8080

