curl -d '{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}' localhost:4000/v1/movies

# need to use 'jq' for creating json variables in bash perhaps.


echo '{ "foo": 123, "bar": 456 }' | jq '.foo'

#curl -H "Content-Type: application/json" --data @text.json http://localhost:8080

curl -H "Content-Type: application/json" --data @text.json localhost:4000/v1/movies

