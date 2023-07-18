# Go Application

 ## To run the golang application generated

  + Build the go application 
  ---
      cd go
      go mod tidy
  ---
  + First run the postgress or mongodb,keycloak and jhipster if rabbitmq for connection then rabbitmq registry files.
  ---
      cd docker
      <%_ if (postgresql){  _%>
      docker-compose -f  postgresql.yml up
      <%_ } _%>
      <%_ if (mongodb){  _%>
      docker-compose -f mongodb.yml up  
      <%_ } _%>   
      <%_ if (auth){  _%>
      docker-compose -f  keycloak.yml up   
      <%_ } _%>     
	 <%_ if (eureka){  _%>
      docker-compose -f  jhipster-registry.yml up  
     <%_ } _%>   
    <%_ if (rabbitmqClient?.length||rabbitmqServer?.length){  _%>
      docker-compose -f rabbitmq.yml up
    <%_ } _%>   

  ---


  + Create a .env file in the go sub directory 
  ---
    ### .env file sample  
    - GO_MICRO_SERVICE_PORT= <port_number>
    - GO_MICRO_SERVICE_REGISTRY_URL=http://admin:admin@localhost:8761/eureka 
    - GO_MICRO_KEYCLOAK_URL=http://localhost:9080
    - GO_MICRO_CLIENT_ID=internal
    - GO_MICRO_CLIENT_SECRET=internal
    - GO_MICRO_REALM_NAME=jhipster
    + (DB_URL based on selection of Db)
    - GO_MICRO_DB_URL=postgresql://go@localhost:5433/postgres
    - GO_MICRO_MONGODB_URL: mongodb://localhost:27017
    - GO_MICRO_MESSAGE_BROKER=amqp://guest:guest@localhost:5672
---

  + Now get back to the root directory of go and start the golang service 

      go run .
  ---
  
## Build an image 
   ---
       docker build -t image_name .
   ---
### Run by passing values 
- docker run -it --network="host" -p <port_number>:<port_number> -e        db_url=postgresql://go@localhost:5433/postgres image_name 


### Run by passing 3 ENV values 

```
docker run -it --network="host" -p <port_number>:<port_number> \

-e DB_URL=postgresql://go@localhost:5433/postgres \

-e KEYCLOAK_URL=http://localhost:9080/realms/jhipster \

-e SERVICE_REGISTRY_URL=http://admin:admin@localhost:8761/eureka/apps/ \

be
```

### Code Formatting 

- Gofmt formats Go programs. It uses tabs for indentation and blanks for alignment. Alignment assumes that an editor is using a fixed-width font.

 Usage:
 ```
  gofmt [flags] [path ...]
  ```
+ Flags

  - d - Do not print reformatted sources to standard output.
	If a file's formatting is different than gofmt's, print diffs
	to standard output.

  - e - Print all (including spurious) errors.

  - l - Do not print reformatted sources to standard output.
	If a file's formatting is different from gofmt's, print its name
	to standard output.

  - r rule - Apply the rewrite rule to the source before reformatting.

  - s - Try to simplify code (after applying the rewrite rule, if any).

  - w - Do not print reformatted sources to standard output.

+ For formatting whole directory 

    ```
      gofmt  .
   
    ```

