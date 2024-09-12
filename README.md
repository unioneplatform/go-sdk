Mostly generated code from the [proto](https://github.com/unioneplatform/proto) repository.

The main branch can be broken. Make sure to use tagged releases.

To use the default client(s), you can try something like:

```go
ctx := context.Background()
client, _ := iam.NewClient(ctx)
defer client.Close()
out, err := client.WhoAmI(ctx, &info.GetInfoRequest{})
log.Println(out, err)
...
```
