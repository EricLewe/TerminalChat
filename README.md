# TerminalChat

These projects may be viewed as a an attempt to catch as many concept as possible related to web developement.
The application itself is a chat terminal inspired by state of the art chat apps.

### Requirements
  - Go
  - Docker engine (and docker compose)
### Installation

Clone repo and cd into TerminalChat all projects are vendored and in a docker container so lets use docker.

Run the following commands:
```sh
$ sudo docker-compose build
$ sudo docker-compose start server
$ sudo docker-compose run client
```


### Usage of program
 - Once logged in you can type
 - Use up and down arrowkeys or scrollwheel on mouse. To scroll your conversations.
 - Type join 1 to join conversation 1 (your name must be on the conversation title) so eventually use another numbers.
 - type !weather to see the weather near your location (i put my faith in the peer struct for the global ip)
