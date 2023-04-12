# Go Application

## prerequisites
 + JHipster
 + node
 + golang
 + docker

## To run the blueprint locally

 ### Link your blueprint globally
  - Go to the path of generator
  --- 
      cd generator-jhipster-go
  ---
  - Now link the development version og jhipster to this blueprint
  ---
   npm link generator-jhipster-go
  ---
  - Now create a new file say my-app and shift to this directory
  ---
    mkdir my-app && cd my-app
  ---
  - Now there are two ways of generating the code:
    + First using prompts

     ---
         jhipster --blueprints go
     ---  

   + Now this will generate a set of prompts and based on the answers given the code will be generated.
    + Second using jdl

      + Say reminder.jdl

         ---
             application {
                config {
                 baseName be1,
                 applicationType microservice,
                 packageName com.cmi.tic,
                 authenticationType oauth2,
                 databaseType sql,
                 prodDatabaseType postgresql,
                 devDatabaseType postgresql,
                 serviceDiscoveryType eureka,
                 serverPort 9001,
                 blueprints [go]
                 } 
              }
        ---

      + Now by using the command we can generate the files 

      ---
          jhipster jdl reminder.jdl
      ---
 ## To run the golang application generated
  + First run the postgress,keycloak and jhipster registry files.
  ---
      cd docker
      docker-compose -f  postgresql.yml up     
      docker-compose -f  keycloak.yml up     
      docker-compose -f  jhipster-registry.yml up     
  ---
  + Now get back to the root directory of go and start the golang service 
  ---
      go run .
  ---
  