const { app, BrowserWindow, globalShortcut, ipcMain } = require("electron");
const  {client} = require("./modules/client");
const net=require("net");

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

    if (!shadowWindow) {
      createShadowWindow();
    } else {
      shadowWindow.show();
    }
    shadowWindow.webContents.send("display-message", "Hello, Shadow Window!");
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
