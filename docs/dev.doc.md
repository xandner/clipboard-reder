## Developer docs
### Connect to server:
Server is running on port 9999. you can connect to it with route : `ws://localhost:9999/clipboard`. when you connected to this route, 10 records of last stored data will send for you.

if you want to get data and search in stored record, send this format to route:
```json
{
"on":"search",
"param":"text you want to search"
}
```
if you want to set data in clipboard, must send record id like:
```json
{
"on":"set",
"param":"record id"
}
```
