### Some followed documentations:
    File structure:
    https://www.youtube.com/watch?v=OFud4iPuAH8
    https://github.com/golang-standards/project-layout/tree/master

    Hot reload:
    https://courses.devopsdirective.com/docker-beginner-to-pro/lessons/06-building-container-images/03-api-golang-dockerfile

    Migrations:
    https://github.com/golang-migrate/migrate




## Project purpose
    Made this project to:
    Evaluate the pros and cons of developing an Web API with Golang, in comparison with other stacks like NestJS(Typescrypt) and Laravel(PHP)
    Get more familiar with golang
    Having a template for golang Web Servers



# Key features i want to analyse
    conventional project file structure
    request validation
    authentication
    available ORM's, query builders
    Dockerfile for development and production
    Livereload with docker
    debugging
    performance vs other frameworks / languages
    unit tests
    speed for developing slightly complex features
    conventional programing paradigms and philosophy


# Progress
    OK - development dockerfile with livereload
    OK - simple working api
    OK - validation
    OK - config file, env
    OK - service layer pattern
    OK - OpenAPI configuration
    OK - some db integration    
    OK - debug in container
    OK - sqlc
    OK - vulnerability scan
    OK - migrations
    OK - traces - very unsatisfied
    OK - metrics creation
    production build / ci/cd pipeline
    implement logs, metrics monitoring
    unit tests - suffering
    OK - authentication
    Skipped - authorization


# Setup Keycloak
    - add the following line to your etc/hosts this file is at C:\Windows\System32\drivers\etc if you're using windows
    127.0.0.1 mykeycloak

    - Go to http://localhost:8082/ and login with user:admin password:admin
    - Create "app" realm and change to that realm
    - Create the client appclient with client authentication set to ON and the valid redirect URIs: "http://localhost:8082/*", otherwise you get a invalid_redirect_url on the Oauth2 Login Flow
    - On appclient, go to credentials and copy the Client Secret
    - Update KEYCLOAK_CLIENT_SECRET variable on dock-compose
    - create a user: "test" and then the password "test", password don't need to be temporary
    - After you successfuly login on the http://localhost:8082/api/v1/auth2/login route, you should get and rawIDToken on the body, that with the prefix "Bearer " will be your authentication token, you can paste it on the Authorization login box on swagger

    - explanation: I had difficulties working with diferent hostnames for keycloak, for example, using http://localhost:8082 on the browser and http://keycloak:8080 inside docker, i had to keep the same hostname inside and outside docker: http://mykeycloak:8080.
    The introspector method failed with different hostnames, probably a security issue
    i couldn't change the keycloak port in development mode, so i had to keep it running in 8080:8080
    That was really painful to setup 