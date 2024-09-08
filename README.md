# Sudoku

This system includes a server component in Golang and a frontend in React. It will generate its own sudoku puzzles and has some basic play through. Still plenty to add in:

1. Allow password changes
2. Allow username and name changes
3. Finish documentation
4. Add docker

Updates to the user isn't working yet, but should be an easy enough fix. The rest other than above is working fine.

# Run the program

Upon download you have to run the following:

- npm install
- npm run build
- go build

The above will build the frontend server and compile the backend server.

From there you can run:

go-sudoku [port] [sql connection]

- Port is the port for the server to run on. Default is 3000 if no parameters
- The sql connection will be the connection string or the file it points to. By default it'll be sudoku.db

Once you do this you can goto the main page in your web browser and it will work. Still updates to do though.
