# configuration
A configuration importer automatically filling in config if there is an ENV set for it using struct tags

# Usage
(see example in cmd/main.go)
```
type MyConfig struct {
    Name string `env:"USER_NAME:Fred"`
    UseSSL bool `env:"USE_SSL:true"`
    Port int `env:"PORT_NO:1234"`
}


func main() {
    myConfig := MyConfig{}
    configImporter.InitConfig(&myConfig)

    fmt.Println(myConfig)
}
```

Configure the config struct with tags with the name "env".
The first field of the tag is the environmental variable that will be used if it exists.
If the environmental variable is not found, the second value will be used.

Automatically converts strings to ints and bools.



