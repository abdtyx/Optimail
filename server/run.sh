sudo docker build -t optimail-server .
sudo docker run --rm -d --name optimail-server --network host optimail-server
