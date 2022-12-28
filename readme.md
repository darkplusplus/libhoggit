# libhoggit

Provides an interface to https://atwar.online/.

```
client := libhoggit.Client("gaw")
state, err := client.GetState()
if err != nil {
    fmt.Printf("%v", err)
    panic(err)
}

for _, obj := range state.Objects {
    fmt.Printf("%d\t%s\t%s//%s//%s\n",
        obj.LastSeen,
        obj.ID,
        obj.Group,
        obj.Type,
        obj.Platform)
}
```