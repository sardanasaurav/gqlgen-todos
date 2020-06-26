# gqlgen-todos

####steps to get started
###### ./scripts/run.sh (starts the graphql server)
###### go run Test/main.go (starts the backend api's which gql api's will call)
###### go to url localhost:7777 (for graphql playground)
###### run your query
###sample queries
```query{
      book(id:1){
         id
         title
        author{
          id
          name
        } 
      }
  }
 ```

```
query{
    books(){
       id
       title
      author{
        id
        name
      } 
    }
}
```

```
query{
    authors(id:2){
       id
       name
       book
    }
}

```
```
query{
    authors{
       id
       name
       book
    }
}
```