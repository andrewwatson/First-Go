package main

import (

    "io/ioutil"
    "fmt"
    "encoding/json"
    // "io"
    // "strings"
    // "log"
)


type Policy struct {
    Type string
    Body []byte
}

type Message struct {
    user, id string
}

func main() {
    filename := "policies.json"
    jsonRaw, readErr := ioutil.ReadFile(filename)

    if readErr != nil {
        fmt.Println(readErr)
    }


    policyMap := map[string]string {}
    var f interface{}
    decodeErr := json.Unmarshal(jsonRaw, &f)

    if decodeErr != nil {
        fmt.Println(decodeErr)
    }

    fmt.Println(f)

    m := f.(map[string]interface{})

    fmt.Println(len(m))

    for k, v := range m {

        fmt.Println("running...")

        switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case int:
            fmt.Println(k, "is int", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
                fmt.Println(i, u)
            }
        case map[string]interface{}:
            fmt.Println(k, "is a map")
            for i, u := range vv {
                fmt.Println("1", i, "2", u)
                policyMap[i] = "x"
            }

        default:
            fmt.Println(k, "is of a type I don't know how to handle", vv)
        }
    }

    // dec := json.NewDecoder(strings.NewReader(string(jsonRaw)))

    // for {

    //     var m Message
    //     if err := dec.Decode(&m); err == io.EOF {
    //         break
    //     } else if err != nil {
    //         log.Fatal(err)
    //     }
    //     fmt.Printf("%s: %s\n", m.user, m.id)
    //     policyMap[m.id] = m.user
    // }
}

