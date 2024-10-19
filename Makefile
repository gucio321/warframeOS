all: update

setup: update
	cp ./bashrc ~/.bashrc.WarframeOS
	echo "source ${HOME}/.bashrc.WarframeOS" >> ~/.bashrc

update:
	go install .
