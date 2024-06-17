
// TODO: Refactor this logic. Remove messy if cases
// DOM manipulation is little tricky, each compoent affects the other

export default class KeyPressController {
    /**
     * @member currentLocation
     * @member command
     * @member mode
     * @member firstKey
     * @member pageNavTargetDiv
     */
    constructor(){
        this.currentLocation = "home.txt"
        this.command = ""
        this.mode = ""
        this.firstKey = ""
        /** 
         * * @type {HTMLElement | null }
         */
        this.pageNavTargetDiv = null
        this.list = document.getElementById('list');
        this.items = this.list.getElementsByTagName('li');
        this.currentIndex = 0;
    }
        
    //might remove this feature later
    /**
     * @param {KeyboardEvent} event 
     */
    handleNavigation(event){
        let keyPressed = event.key;
        this.pageNavTargetDiv = document.getElementById('currentText');
        console.log("Key Pressed" , event.key)
        
        if(keyPressed == "Escape" && this.this.mode == "i"){
            this.mode = ""
            let commandBox = document.getElementById("commandBox")
            if(commandBox){
                document.getElementById("commandBox").blur()
            }
            return
        }
    
        if(this.mode == "i"){
            console.log("In insert this.mode,press escape to exit")
            this.pageNavTargetDiv.innerHTML = "insert mode"
            return
        }
    
        if(keyPressed == "i"){
            this.mode = keyPressed
            event.preventDefault()
            let commandBox = document.getElementById("commandBox")
            if(commandBox){
                document.getElementById("commandBox").focus()
            }
            this.pageNavTargetDiv.innerHTML = "insert mode"
            console.log("Mode is: ",this.mode)
        }
    
        if(keyPressed == ":") {
            this.command = ""
            this.firstKey = keyPressed
            console.log("First Key",this.firstKey)
        }

        if(this.firstKey == ":"){
            this.handleNavigation(keyPressed)
        }
    }

    handlePageNavigation(){
        switch(keyPressed){
            case "Backspace":
                if(this.command.length > 0){
                    this.command = this.this.command.slice(0,-1)
                }
        
                if(this.this.command.length == 0){
                    this.firstKey = ""
                    this.command = this.currentLocation
                }
                break
            
            case "Escape":
                this.firstKey = ""
                this.command = this.currentLocation
                break
            
            case "Enter":
                this.handleCommand()
                this.firstKey = ""
                break

            default: 
                this.command += keyPressed
                break
        }
        this.pageNavTargetDiv.innerHTML = this.command
    }
    
    handleCommand(){
        let cmdName = this.command.split(":")
        if (cmdName.length < 2)  {
            console.log("Invalid Command Recieved",this.command)
            return
        }

        if(this.command.includes("gt")){
            console.log("Goto Command")
            return
        }

        switch(cmdName[1]){
            case "wq":
                alert("Save and Quit.")
                break

            case "q":
                alert("Quit")
                break

            case "gb":
                alert("go back")
                break

            default:
                alert("Oops.Seems like an invalid terminal this.command")
                break
        }
    }
}