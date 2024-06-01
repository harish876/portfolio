const DEFAULT_LOCATION = "home.txt"

let command = ""
let mode = ""
let firstKey = ""
let currentWindowLocation = handleCurrentLocation(window.location)

// TODO: Refactor this logic. Remove messy if cases
// DOM manipulation is little tricky, each compoent affects the other

document.addEventListener('keydown', function(event) {
    currentWindowLocation = handleCurrentLocation(window.location)
    let keyPressed = event.key;
    const outputDiv = document.getElementById('currentText');
    console.log("Key Pressed" , event.key)
    
    if(keyPressed == "Escape" && mode == "i"){
        mode = ""
        let commandBox = document.getElementById("commandBox")
        if(commandBox){
            document.getElementById("commandBox").blur()
        }
        console.log(currentWindowLocation)
        outputDiv.innerHTML = displayPathFormatter(currentWindowLocation.page)
        return
    }

    if(mode == "i"){
        console.log("In insert mode,press escape to exit")
        outputDiv.innerHTML = "insert mode"
        return
    }

    if(keyPressed == "i"){
        mode = keyPressed
        event.preventDefault()
        let commandBox = document.getElementById("commandBox")
        if(commandBox){
            document.getElementById("commandBox").focus()
        }
        outputDiv.innerHTML = "insert mode"
        console.log("Mode is: ",mode)
    }

    if(keyPressed == ":") {
        command = ""
        firstKey = keyPressed
        console.log("First Key",firstKey)
    }
    if(firstKey == ":"){
        switch(keyPressed){
            case "Backspace":
                if(command.length > 0){
                    command = command.slice(0,-1)
                }
        
                if(command.length == 0){
                    firstKey = ""
                    command = DEFAULT_LOCATION
                }
                break
            
            case "Escape":
                firstKey = ""
                command = displayPathFormatter(currentWindowLocation.page)
                break
            
            case "Enter":
                handleCommand(command)
                command = displayPathFormatter(currentWindowLocation.page)
                firstKey = ""
                break

            default: 
                command += keyPressed
                break
        }
        outputDiv.innerHTML = command
    }
});

/**
 * @param {string} command 
 */
function handleCommand(command){

    let cmdName = command.split(":")
    if (cmdName.length < 2)  {
        console.log("Invalid Command Recieved",command)
        return
    }

    if(command.includes("gt")){
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
            alert("Oops.Seems like an invalid terminal command")
            break
    }
}

/**
 * 
 * @param {Location} location 
 * @returns {{basePath: string , page: string }} result
 */
function handleCurrentLocation(location){
    let basePath = location.origin
    let page = ""

    switch (location.pathname){
        case "/":
            page = "home"
            break
        
        case "/readme":
            page = "readme"
            break
        
        case "/project":
            page = "project"
            break
        
        case "/blog":
            page = "blog"
            break
        
        default:
            console.log("Invalid Location. This should get errored out anyways.")
            break
    }

    return {
        basePath,
        page
    }
}

/**
 * @param {string} path 
 */
function displayPathFormatter(path){
    return `${path}.txt`
}