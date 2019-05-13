## hungry-honeybadger

**Run Projekt**

**1. Install go & and configure GOPATH**

[GO](https://golang.org/doc/install)


**Clone Projekt**

with shh:

`git clone git@gitlab.mi.hdm-stuttgart.de:mwa/ss19/hungry-honeybadger.git`

with https:

`git clone https://gitlab.mi.hdm-stuttgart.de/mwa/ss19/hungry-honeybadger.git`


**2. Install dependencies**

run in root folder
`go get -d ./api && cd web && npm install `

# Development 

**Dev with Docker Fronted**

run in web folder
`npm run devFull`

starts backend and DB und serve frontend via nginx 

**Frontend**

*Start frontend with dev Server*

run in web folder
`npm run dev`

*start frontend with nginx as dev Server within docker*

run in web folder
`npm run dockerRun`

# Production 

**1. Install Docker**

[Docker](https://docs.docker.com/install/)

**Frontend**

*start prod build*

run in web folder
`npm run prod`
 