<h1 align="center" style="font-weight: bold;">Project name 💻</h1>

<p align="center">
 <a href="#tech">Technologies</a> • 
 <a href="#started">Getting Started</a> • 
  <a href="#documentation">Documentation</a> 
</p>

<p align="center">
    <b>I undertook this challenge with an educational purpose to enhance my skills in Go programming language and improve my understanding of architecture and containers.</b>
</p>

<h2 id="technologies">💻 Technologies</h2>

- Golang
- Docker
- PostgreSQL

<h2 id="started">🚀 Getting started</h2>

<h3>Prerequisites</h3>
- Docker

<h3>Cloning</h3>

```bash
git clone https://github.com/vnniciusg/backend-developer-challenge.git
```

<h3>Config .env variables</h2>

Use the `.env.example` as reference to create your configuration file `.env` 

```yaml
DB_USER=YOUR_DB_USER
DB_PASS=YOUR_DB_PASS
DB_NAME=YOUR_DB_NAME
DB_PORT=5432
DB_HOST=db
DB_DRIVER=postgres
  
TEST_DB_USER=YOUR_TEST_DB_USER
TEST_DB_PASS=YOUR_TEST_DB_PASS
TEST_DB_NAME=YOUR_TEST_DB_NAME
TEST_DB_PORT=5433
TEST_DB_HOST=db_test
```


<h3>Starting</h3>

How to start your project

```bash
docker-compose up
```


<h2 id="documentation">📚Documentation</h2>

Explore the API using Swagger [here](http://localhost:8080/api/v1/docs/index.html#/)


