# Notes

- Swagger url: `<host>/swagger/index.html`
- Docker postgres image setup for local run:

  ```
  docker run -d \
  --name postgres_latest \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_USER=root_user \
  -v $(pwd):/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres
  ```
