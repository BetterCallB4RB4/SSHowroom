docker rm -f sshowroom
docker build -t sshowroom:0.0.1 .
docker run -d -p 2222:2222  --name sshowroom sshowroom:0.0.1
ssh-keygen -R "[localhost]:2222"
ssh localhost -p 2222



docker rm -f sshowroom ; docker build -t sshowroom:0.0.1 . ; docker run -d -p 2222:2222 --rm --name sshowroom sshowroom:0.0.1 ; ssh-keygen -R "[localhost]:2222" ; ssh -o StrictHostKeyChecking=accept-new localhost -p 2222

