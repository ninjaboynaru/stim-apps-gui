
build:
	go build -o ./stim-apps-gui

run:
	./stim-apps-gui

buildrun:
	go build -o ./stim-apps-gui && ./stim-apps-gui

fynedemo:
	go run fyne.io/fyne/v2/cmd/fyne_demo@latest