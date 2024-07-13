const net = require("net");

class TCPClient {
  constructor(host, port) {
    this.host = host;
    this.port = port;
    this.client = new net.Socket();
    this.connected = false;

    this.client.on("error", (err) => {
      console.error("Socket error:", err.message);
      this.client.destroy();
    });

    this.client.on("close", () => {
      console.log("Connection closed");
      this.connected = false;
    });
  }

  connect() {
    return new Promise((resolve, reject) => {
      if (this.connected) {
        resolve();
        return;
      }

      this.client.connect(this.port, this.host, () => {
        console.log("Connected to server!");
        this.connected = true;
        resolve();
      });

      this.client.once("error", (err) => {
        console.error("Connection error:", err.message);
        this.connected = false;
        reject(err);
      });
    });
  }

  send(message) {
    return new Promise((resolve, reject) => {
      if (!this.connected) {
        reject(new Error("Client not connected"));
        return;
      }

      this.client.write(message,(err) => {
        console.log("Sending:", message);
        if (err) {
          console.error("Error sending data:", err.message);
          reject(err);
        }
        console.log("Sent:", message);

        // Wait for a response from the server
        this.client.on("data", (data) => {
          console.log("Received:", data.toString());
          resolve(data.toString());
        });
      });
    });
  }

  disconnect() {
    if (this.connected) {
      this.client.end();
    }
  }
}
(()=>{
    n=new TCPClient('localhost', 9999);
    n.connect().then(()=>{
        n.send('Hello, Server!').then(()=>{
            n.disconnect();
        });
    }).catch((err)=>{
        console.error(err);
    });
})()

module.exports = TCPClient;
