package main

import (
    "fmt"
)

func main() {
    cmd := make(chan string)
    msg := make(chan string)
    //go ping(pings)
    fmt.Printf("*Go for Very Small Cave Adventure*\n")
    go toCave(cmd, msg, "bridge")
    var Directive string
 
    for {
        cmd <- Directive
        message := <-msg
        if message == "quit" {
            close(msg)
            break
        } else {
            fmt.Println(message)
            fmt.Scanln(&Directive)      
        }
    }  
}

func toCave(cmd <-chan string, msg chan<- string, startLocation string) {
    currentLocation := startLocation
    var inventory []string
    var message,location,item string = "","",""
    var changeLocation int = 0
    for {
        command := <-cmd
        message,location,item,changeLocation = StoryTeller(currentLocation, command, inventory)
        currentLocation = location
        fmt.Println(currentLocation)
        if currentLocation == "endgame" {             
            inventory = nil            
        }
        if changeLocation == 1 {
            command = ""
            message,location,item,changeLocation = StoryTeller(currentLocation, command, inventory)
            if item != "" {
                inventory = append(inventory,item)
                fmt.Println(item, "has been added to inventory")            
            }            
            msg <- message
        } else {
            if item != "" {
                inventory = append(inventory,item)
                fmt.Println(item, "has been added to inventory")            
            }            
            msg <- message
        }
    }
}

func StoryTeller(location string, command string,inventory []string) (string,string,string,int) {
    if command == "quit" || command == "Quit" || command == "q" {
        return "quit","quit","",0
    }
    if location == "endgame" {
        if command == "" {
            return "Troll noticed the rope in your hand and said:\n-Well, I see you have a made progress so far.\nYou:\n-Working hard.\nTroll:\n-Now it is my time to do my job .Wait a minute while I magically restore the bridge using the mana and my saw.\n\nSome time passed\n\nAnother some time passed\n\nYou went to sleep and after a while you heard the Troll waking you:\n-I am done with - now claim your victory,dude!\nYou:\n\nYou thanked the Troll and left all your inventory for him to keep(why would you need it afterwards?) - he was really happy.Then you followed the bridge to the other side of the underground river.Then go out of the cave on the other side and followed through - along the yellow brick road - to the nearest bus stop to return home...\n\n*Go for Very Small Cave Adventure*\n\nThank you for playing the Game\n\nYou earned over 9 thousand score of over 9 thousand possible.\n\nProgrammed by Nick Cheg aka giant.panda.chen and the Go Gopher\n\nThe End\n\nInput (q), (Enter) to quit.\n\n","endgame","",0
        } else {
            return "","endgame","",1
        }
    }
    if location == "bridge" {
        if command == "" {
            return "You are standing in a very small cave. There is a broken bridge in front of you and the Troll guarding it. If you input (a), (Enter) - it will take you westwards, if you input (d), (Enter) - it will take you eastwards, if you input (s), (Enter) - it will take you to the south exit. Input (g), (Enter) to talk to the Troll. Input (q), (Enter) to exit the game.","bridge","",1
        } else if (command == "g") {
            if !contains(inventory,"Google Adwords Coupon") {
                return "Troll says:\n-I am the Bridge Troll - you pay the Troll Tall\nYou nodded:\n-...\n\nYou are standing in very small cave. There is a broken bridge in front of you and the Troll guarding it. If you input (a), (Enter) - it will take you westwards, if you input (d), (Enter) - it will take you eastwards, if input (s), (Enter) - it will take you to the south exit. Input (g), (Enter) to talk to the Troll. Input (q), (Enter) to exit the game.","bridge","",0
            }
            if contains(inventory,"Google Adwords Coupon") && !contains(inventory,"dimmed flashlight") && !contains(inventory,"rope") {
                return "Troll says:\n-Good.Good.Google Adwords should do the trick, but as you can see, the bridge is broken - need a rope to fix it.\nYou said:\n-Do have the rope?\nThe Troll says:-Yeah... well... I used to... I worked on the bridge fixing, but very large Moth suddenly appeared and took it off with its paws. You know, maybe you could help us both? Find the Moth and took the rope off him. You would need a flashlight to run through the caves - otherwise Moths would not let you - mine is not good enough needs a battery replacement. You can keep the Adwords until we are done here.\nYou said:-I ain't afraid of no Month\n\n\nYou are standing in very small cave. There is a broken bridge in front of you and the Troll guarding it. If you input (a), (Enter) - it will take you westwards, if you input (d), (Enter) - it will take you eastwards, if you input (s), (Enter) - it will take you to the south exit. Input (g), (Enter) to talk to the Troll. Input (q), (Enter) to exit the game.","bridge","dimmed flashlight",0
            } else if contains(inventory,"Google Adwords Coupon") && contains(inventory,"dimmed flashlight") && !contains(inventory,"rope") {
                return "Troll says:\n-Good.Good.Google Adwords should do the trick, but as you can see, the bridge is broken - need a rope to fix it.\nYou said:\n-Do have the rope?\nThe Troll says:-Yeah... well... I used to... I worked on the bridge fixing, but very large Moth suddenly appeared and took it off with its paws. You know, maybe you could help us both? Find the Moth and took the rope off him. You would need a flashlight to run through the caves - otherwise Moths would not let you - mine is not good enough needs a battery replacement. You can keep the Adwords until we are done here.\nYou said:-I ain't afraid of no Month\n\n\nYou are standing in very small cave. There is a broken bridge in front of you and the Troll guarding it. If you input (a), (Enter) - it will take you westwards, if you input (d), (Enter) - it will take you eastwards, if you input (s), (Enter) - it will take you to the south exit. Input (g), (Enter) to talk to the Troll. Input (q), (Enter) to exit the game.","bridge","",0
            } else if contains(inventory,"Google Adwords Coupon") && contains(inventory,"rope") {
                return "","endgame","",1
            }
        } else if (command == "s") {
            return "","plateau","",1
        } else if command == "a"{
            return "","blocked passage","",1
        }  else if command == "d"{
            return "","corridor","",1
        } else if command != "" || command != "g" || command != "s" || command != "a" || command != "d" {
            return "","bridge","",1
        }
    }
    if location == "blocked passage" {
        if command == "" {
            if !contains(inventory,"flashlight") {
                return "Because it is dark here - Moths are forcing you to quit the game. Input (q), (Enter) to quit.","Dark Places","",0
            }
            if contains(inventory,"flashlight") && !contains(inventory,"scissors") {
                return "You are in small cave blocked from every end, except from one passage - the one you have came from. You hear the water dropping somewhere. Steep stone walls are around you. You noticed a scissors at the corner - probably its the Troll's - the one that he may lost somewhere overhere. Well... finders - keepers. Input (d), (Enter) to go back to the bridge cave. Input (q), (Enter) to quit.","blocked passage","scissors",0
            }
            if contains(inventory,"flashlight") && contains(inventory,"scissors") {
                return "You are in small cave blocked from every end, except from one passage - the one you have came from. You hear the water dropping somewhere. Input (d), (Enter) to go back to the bridge cave. Input (q), (Enter) to quit.","blocked passage","",0
            } 
        } else if contains(inventory,"flashlight") && command == "d" {
            return "","bridge","",1
        } else if command != "" || command != "d" {
            return "","blocked passage","",1
        }
    }
    if location == "corridor" {
        if command == "" {
            if !contains(inventory,"dimmed flashlight") && !contains(inventory,"flashlight") {
                return "Because it is dark here - Moths are forcing you to restart the game. Input (q), (Enter) to quit.","Dark Places","",0
            }
            if contains(inventory,"dimmed flashlight") && !contains(inventory,"flashlight") {
                return "In dimmed light you suddenly noticed a Battery Tree with a falling leaves - the Autumn is here... somewhere. One leaf fall near your feet - a battery for the flashlight - immediately you take it and replace the old in one in your flashlight - Input (a), (Enter) to go back to the bridge cave or (d) to go further eastwards. Input (q), (Enter) to quit.","corrdior","flashlight",0
            }
            if contains(inventory,"dimmed flashlight") && contains(inventory,"flashlight") {
                return "Just a small corridor in the cave with the battery tree growing here - input (a), (Enter) to go back to the bridge cave or (d) to go further eastwards. Input (q), (Enter) to quit.","corridor","",0
            }
        }
        if (contains(inventory,"dimmed flashlight") || contains(inventory,"flashlight")) && command == "a" {
            return "","bridge","",1
        }
        if (contains(inventory,"dimmed flashlight") || contains(inventory,"flashlight")) && command == "d" {
            return "","moth cave","",1
        } else if command != "" || command != "a" || command != "d" {
            return "","corridor","",1
        }
    }
    if location == "moth cave" {
        if command == "" {
            if !contains(inventory,"flashlight") {
                return "Because it is dark here - Moths are forcing you to restart the game. Input (q), (Enter) to quit.","Dark Places","",0
            }
            if contains(inventory,"flashlight") && !contains(inventory,"rope") {
                return "You standing in vast cave. Walls can not be seen and lost somewhere into dark. At the close range you see the lava river and small wodden fence - probably to prevent people from falling into the pit. To this fence a rope is tangled to - the other end of the rope is entangled around the giant moth. How could a moth be visible in the light? Perhaps, you can ask him in person?. Input (g), (Enter) to talk to giant moth. Input (a), (Enter) to go back to the corridor. Input (q), (Enter) to quit.","moth cave","",0
            }
            if contains(inventory,"flashlight") && contains(inventory,"rope") {
                return "You standing in vast cave. Walls can not be seen and lost somewhere into dark. At the close range you see the lava river and small wodden fence - probably to prevent people from falling into the pit. Input (a), (Enter) to go back to the corridor. Input (q), (Enter) to quit.","moth cave","",0
            }
        } else if contains(inventory,"flashlight") && !contains(inventory,"rope") && (command == "g") {
            if !contains(inventory,"scissors") {
                return "Giant Moth says:\n-I am the Giant Moth, the Behemoth\nYou agreed:\n-Ugh... okay\nThe Behemoth:\n-I am travelled across the cave when I encountered strange green man with the snake'ish thing in the hands(nodded to rope). I thought it would be funny to take it away. How could I know that the snake'ish entangles? Some green strange man'ish magic? Please help me strnager... please...\nYou said:\nI'll do what I can, but I'll keep the rope...\n\nYou standing in vast cave. Walls can not be seen and lost somewhere into dark. At the close range you see the lava river and small wodden fence - probably to prevent people from falling into the pit. To this fence a rope is tangled to - the other end of the rope is entangled around the giant moth. How could a moth be visible in the light? Perhaps, you can ask him in person?. Input (g), (Enter) to talk to giant moth. Input (a), (Enter) to go back to the corridor. Input (q), (Enter) to quit.","moth cave","",0
            }
            if contains(inventory,"scissors") {
                return "The Behemoth says:\n-I am the Giant Moth, the Behemoth. Help me please. Will you help? Please, set me free.\nYou answered:\n-I have those scissors with me.\nThe Behemoth:\n-Quick use scissors on the rope please - allow me to go, please.\nYou cut the rope with your scissors in hand and set the moth free\nThe Behemoth: Thank you stranger. Keep the rope. Shall go now need to have my revenge on the Ashen one - you know.\nThe Behemoth flew away.\n\nYou standing in vast cave. Walls can not be seen and lost somewhere into dark. At the close range you see the lava river and small wodden fence - probably to prevent people from falling into the pit. Input (a), (Enter) to go back to the corridor. Input (q), (Enter) to quit.","moth cave","rope",0                
            }
        } else if (contains(inventory,"flashlight")) && command == "a" {
            return "","corridor","",1
        } else if command != "" || command != "a" || command != "g" {
            return "","moth cave","",1
        }

    }
    if location == "plateau" {
        if command == "" {
            if !contains(inventory,"Google Adwords Coupon") {
                return "You went out of the cave to a small plateau - bright alpine grass is growing here. You can not go anywhere from here, except back to the cave as ocean waves splashing below - input (w), (Enter) to go back to the cave. You noticed a small mailbox located in the middle of the meadow. By some uknown telepathic abilities you sense that Google Adwords Coupon is inside - you open the mailbox and took the coupon. Input (q), (Enter) to quit.","plateau","Google Adwords Coupon",0
            } else {
                return "You went out of the cave to a small plateau - bright alpine grass is growing here. You can not go anywhere from here, except back to the cave as ocean waves splashing below - input (w), (Enter) to go back to the cave. There is a small mailbox located in the middle of the meadow and it is open and empty. Input (q), (Enter) to quit.","plateau","",0
            }
        }
        if command == "w" {
            return "You are standing in very small cave. There is a broken bridge in front of you and the Troll guarding it. If you input (a), (Enter) - it will take you westwards, if you input (d), (Enter) - it will take you eastwards, if you input (s), (Enter) - it will take you to the south exit. Input (g), (Enter) to talk to the Troll. Input (q), (Enter) to exit the game.","bridge","",0         
        }
        if command != "" || command != "w" {
            return "","plateau","",1
        }
    }
    return "","","",0
}

func contains(arr []string, str string) bool {
    for _, a := range arr {
        if a == str {
            return true
        }
    }
    return false
}
