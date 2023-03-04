/* eslint-disable no-useless-escape */

export function initConsoleMessage(){
	globalThis.addEventListener("devtoolschange", handleDevToolsChange)
}

function handleDevToolsChange(event:unknown){
	const ev = event as CustomEvent<{isOpen:boolean}>

	if(!ev.detail.isOpen){ return }

	printConsoleMessage()
	globalThis.removeEventListener("devtoolschange",handleDevToolsChange)
}

function printConsoleMessage(){
	// console.clear()


	const banner =[
		"  ____                               __    ____                            ____                __       __  __     ",
		" /\\  _`\\                  __        /\\ \\__/\\  _`\\   __                    /\\  _`\\             /\\ \\     /\\ \\/\\ \\     ",
		" \\ \\,\\L\\_\\  _____   _ __ /\\_\\    ___\\ \\ ,_\\ \\ \\L\\_\\/\\_\\    ___     ____   \\ \\ \\L\\_\\    ___ ___\\ \\ \\____\\ \\ \\_\\ \\    ",
		"  \\/_\\__ \\ /\\ '__`\\/\\`'__\\/\\ \\ /' _ `\\ \\ \\/\\ \\  _\\L\\/\\ \\ /' _ `\\  /',__\\   \\ \\ \\L_L  /' __` __`\\ \\ '__`\\\\ \\  _  \\   ",
		"    /\\ \\L\\ \\ \\ \\L\\ \\ \\ \\/ \\ \\ \\/\\ \\/\\ \\ \\ \\_\\ \\ \\L\\ \\ \\ \\/\\ \\/\\ \\/\\__, `\\   \\ \\ \\/, \\/\\ \\/\\ \\/\\ \\ \\ \\L\\ \\\\ \\ \\ \\ \\  ",
		"    \\ `\\____\\ \\ ,__/\\ \\_\\  \\ \\_\\ \\_\\ \\_\\ \\__\\\\ \\____/\\ \\_\\ \\_\\ \\_\\/\\____/    \\ \\____/\\ \\_\\ \\_\\ \\_\\ \\_,__/ \\ \\_\\ \\_\\ ",
		"	 \\/_____/\\ \\ \\/  \\/_/   \\/_/\\/_/\\/_/\\/__/ \\/___/  \\/_/\\/_/\\/_/\\/___/      \\/___/  \\/_/\\/_/\\/_/\\/___/   \\/_/\\/_/ ",
		"			  \\ \\_\\                                                                                                 ",
		"			   \\/_/                                                                                                 ",
	].join("\n")

	console.log(`%c${banner}`, " color: #EA5768")
	// console.log("%cHi! Are you interested in working on apps like this? Do you want to work on apps like this?\\nCheck out our vacancies:", "color: #EA5768; font-size:30px");
	console.log("%cHi! Are you interested in working on apps like this?\nCheck out our vacancies at:", "color: #EA5768; font-size:30px");
	console.log("%chttps://sprinteins.com/jobs", "background: #EA5768; color:white; font-size:30px; padding: 0.5rem 1rem; text-transform: uppercase;")

	
}