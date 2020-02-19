cd ../
docker build -t jneubaum/honestvote-peer-node:latest -f deployments/Peer.Dockerfile .
docker push jneubaum/honestvote-peer-node:latest