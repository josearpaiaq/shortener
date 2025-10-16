## 🚀 Go Shortener App

- Implement hot realoading with the docker-compose file

  - [Hot reloading](https://docs.docker.com/compose/compose-file/compose-file-v3/#hot-reloading) with docker-compose
  - take advantage of [air](github.com/cosmtrek/air@latest) or [CompileDaemon](https://github.com/githubnemo/CompileDaemon) to watch for changes and restart the server

- Improve the landing page

  - Add [tailwindcss](https://tailwindcss.com/) or something similar to make the page look better (taking in count the landing page is just a static HTML file)
  - Add a form to enter the URL
  - Add a button to shorten the URL
  - Display the shortened URL
  - Display the original URL
  - Display the number of clicks
  - Add favicon

- Add database migrations

  - [Gorm migration](https://gorm.io/docs/migration.html)

- Add tests

  - [Test validations playground](https://github.com/go-playground/validator)

- Add user authentication
  - Add a login form
  - Add a login button
  - Display the user's information
  - Display the number of clicks
  - Add a logout button
  - Related urls to the user
