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
      docker-compose -f  postgresql.yml up
      docker-compose -f mongodb.yml up     
      docker-compose -f  keycloak.yml up     
      docker-compose -f  jhipster-registry.yml up   
      docker-compose -f rabbitmq.yml up
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
    

