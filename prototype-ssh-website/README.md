# SSH Portfolio Prototype

This is a prototype for an SSH-based portfolio website.

## SSH Host Keys

The project includes a set of SSH host keys in the `.ssh/` folder:

- `.ssh/term_info_ed25519`: The private host key.
- `.ssh/term_info_ed25519.pub`: The public host key.

### Purpose

These keys are used by the SSH server to identify itself to clients. In SSH, when a client connects to a server, the server presents its public key. If the client has never connected to this server before, it will prompt the user to trust the key. If it has connected before, it verifies that the key matches the one it has stored in `~/.ssh/known_hosts`.

Using a fixed host key (instead of generating a new one every time the server starts) ensures that users don't get "Host key verification failed" errors every time you restart your server.

### Changing Host Keys

If you want to use your own keys, you can generate a new ed25519 key pair:

```bash
# From the prototype-ssh-website directory
ssh-keygen -t ed25519 -f .ssh/term_info_ed25519 -N ""
```

This will overwrite the existing keys with a new pair.

## Testing the Website

### Running with Docker

You can build and run the project using Docker with the following commands:

```bash
# Build the docker image
docker build -t ssh-portfolio .

# Stop and remove any existing container with the same name
docker stop my-portfolio || true && docker rm my-portfolio || true

# Run the container
docker run -d -p 2222:2222 --name my-portfolio ssh-portfolio
```

### Running with Go

If you have Go installed, you can run the server directly:

```bash
go run main.go
```

The server will start listening on `localhost:2222`.

### Connecting to the Server

Once the server is running, you can connect to it using any SSH client:

```bash
ssh localhost -p 2222
```

**Note:** If you have previously connected to a different server on `localhost:2222` (or if you changed the host keys), you might need to clear the old key from your `known_hosts` file:

```bash
ssh-keygen -f "$HOME/.ssh/known_hosts" -R "[localhost]:2222"
```

docker build -t ssh-portfolio . ; docker stop my-portfolio && docker rm my-portfolio ; docker run -d -p 2222:2222 --name my-portfolio ssh-portfolio ; ssh-keygen -f "/home/nixos/.ssh/known_hosts" -R "[localhost]:2222" ; ssh localhost -p 2222
