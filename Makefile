all: update

setup: update
	echo "source ${PWD}/bashrc" >> ~/.bashrc

update:
	go install .
