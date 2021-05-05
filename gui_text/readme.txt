export FYNE_FONT=/System/Library/Fonts/AquaKana.ttc

$Env:FYNE_FONT='C:\Windows\Fonts\YuGothM.ttc'
go run .

go build -ldflags="-H windowsgui"
go build -ldflags="-H darwingui"

go get fyne.io/fyne/cmd/fyne
fyne package -icon icon.png