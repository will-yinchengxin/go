package main

import (
  _ "gin/frame/app"
  "gin/frame/route"
)

func main() {
  route.InitRouter().Run(":8070")
}
