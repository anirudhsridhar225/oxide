# Oxide

A fully online, in-the-browser text editor.

## Dependencies

- `go`
- `air` (`go install github.com/air-verse/air@latest`)
- `swag` (`go install github.com/swaggo/swag/cmd/swag@latest`)

Note: make sure your go installation and the bin folder are added to PATH.
    For example in linux, that would be `~/go/bin`

## Tech Stack

    - frontend
        - htmx
        - vite
        - less

    - backend
        - golang
        - chi
        - gorm
        - swag

## Why I'm doing this

just because i can honestly. no real motivation past that.

## Contributing

A `CONTRIBUTING.md` file is still a WIP.

## Usage

1. Clone the project:
    `git clone https://github.com/anirudhsridhar225/oxide.git`

2. Cd into the frontend and backend folders and install all dependencies:
    `cd frontend && pnpm install`
    `cd backend && go install`

3. Cd into the frontend and start the vite server, cd into the backend and start the go server:
    `cd frontend && pnpm dev`
    `cd backend && go mod clean && swag init --parseDependency --parseInternal && air`

4. The frontend is served at `localhost:3000` while the backend is served at `localhost:8000`.

Note: you can access the swagger documentation at `localhost:8000/swagger/index.html`

## TODOs

- [ ] write a golang lsp serving endpoint
- [ ] write a rust/cpp based backend microservice to generate syntax highlighting through wasm via tree-sitter
- [ ] integrate wasm into the frontend
- [ ] design the frontend for the editor interface
- [ ] write a golang endpoint to create and execute docker containers that mount a `/tmp` filesystem for users to run their code in
- [ ] optimise spinning up the docker containers so i dont have to make a 1000 dockerfiles
- [ ] write a frontend/backend microservice to manage the terminal emulation for the code editor
- [ ] research and implement debugging
- [ ] research the implementation of a plugin ecosystem (i really love vim)
- [ ] make it look nice (optional)
- [ ] have different colorscheme options etc. for the frontend
- [ ] make user sessions air-tight and enable github based SSO
- [ ] make it so that there is a way for people to import their repositories and then code and push/pull/pr their changes
        - [ ] maybe think about dockerizing that as well
- [ ] clean the code up

## Finally

I hope this project, whenever I finish it, provides a good UX to everyone.

Made with :heart: by andy
