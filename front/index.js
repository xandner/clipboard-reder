const { app, BrowserWindow, globalShortcut, ipcMain } = require("electron");
const net = require("net");

let shadowWindow;

function createShadowWindow() {
  shadowWindow = new BrowserWindow({
    width: 300,
    height: 200,
    frame: false,
    transparent: true,
    alwaysOnTop: true,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false,
    },
  });
  shadowWindow.loadFile("shadow.html");

  // Handle window close event
  shadowWindow.on("closed", () => {
    shadowWindow = null;
  });
}

app.whenReady().then(() => {
  globalShortcut.register("CommandOrControl+Shift+v", () => {
    let responseData = "";
    if (!shadowWindow) {
      createShadowWindow();
    } else {
      shadowWindow.show();
    }
    if (shadowWindow.isVisible()) {
      const client = net.connect({ port: 9999, host: "localhost" }, () => {
        console.log("Connected to server!");
        client.write("get_10\n");
      });
      client.on("data", (data) => {
        const string = new TextDecoder().decode(new Uint8Array(data));
        responseData = JSON.parse(string);
        console.log(responseData)
        for (const i of responseData){
          console.log(i);
        }
        shadowWindow.webContents.send("response-data", responseData);
      });
    }
  });

  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0 && !shadowWindow) {
      createShadowWindow();
    }
  });
});

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("will-quit", () => {
  globalShortcut.unregisterAll();
});
