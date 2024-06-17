import KeyPressController from "./keypress.js";

document.addEventListener('DOMContentLoaded', () => {
    console.log("is this called only once")
    const keyPressController  = new KeyPressController()
    document.addEventListener('keydown',keyPressController.handleNavigation);
}, false);