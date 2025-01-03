all:
	cd server && go build -o server ./main.go && sudo docker build -t optimail-server .
	cd micro-user && go build -o server ./main.go && sudo docker build -t optimail-micro-user .
	cd mail-agent && sudo docker build -t optimail-mail-agent .
