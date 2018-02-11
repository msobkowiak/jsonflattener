# jsonflattener
`jsonflattener` transforms nested json structure into flat key, values pars. 

The program reads json file (path to the file should be passed as command line argument), fattens the json and printout the result of a transformation into stdout. 

### Example input:
```json
{
  "first_name": "Joe",
  "last_name": "Doe",
  "age": 25,
  "hobbies": ["travel", "sport", "books"],
  "address": {
    "street": "123 Main St.",
    "city": "Berlin",
    "zip_code": "11129"
  }
}
```

### Example output:
```
first_name: "Joe"
last_name: "Doe"
age: 25
hobbies[0]: "travel"
hobbies[1]: "sport"
hobbies[2]: "books"
address.city: "Berlin"
address.zip_code: "11129"
address.street: "123 Main St."
```

# Build and run 

```
# build an executable file
go build 

# run the application
./jsonflattener path_to_the_file
```
