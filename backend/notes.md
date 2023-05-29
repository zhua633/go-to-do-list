## Initialize go app in the backend folder

go mod init github.com/zhua633/go-to-do-list

go mod tidy

## Install go fiber v2 in the backend folder

go get -u github.com/gofiber/fiber/v2

### If time out then

nvim ~/.config/fish/config.fish

set -x GOPROXY “proxy.golang.org,direct”

source ~/.config/fish/config.fish

## Create client app with Vite

yarn create vite frontend -- --template react-ts

## Install dependencies

yarn add @mantine/forms @mantine/core swr @primer/octicons-react
