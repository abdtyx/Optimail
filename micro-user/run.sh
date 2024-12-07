sudo docker build -t optimail-micro-user .
sudo docker run --rm -d --name optimail-micro-user --network host optimail-micro-user
