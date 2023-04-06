
/**
 * Inspired by:
 * https://github.com/sindresorhus/devtools-detect
 * By Sindre Sorhus
 * 
 */

const threshold = 170;

export function initConsoleDetection(){
	ensureEventListener()
}

let hasEventListener = false
function ensureEventListener(){
	if(hasEventListener){ return }

	globalThis.window.addEventListener('resize', handleResize);
	hasEventListener = true
	handleResize()
}


function handleResize(){
	const widthThreshold = globalThis.outerWidth - globalThis.innerWidth > threshold;
	const heightThreshold = globalThis.outerHeight - globalThis.innerHeight > threshold;
	const isDevToolsOpen = widthThreshold || heightThreshold

	emitEvent(isDevToolsOpen)

}

function emitEvent(isOpen: boolean) {
	const eventKey = "devtoolschange"
	const event = new globalThis.CustomEvent(eventKey, {detail: {isOpen}} )
	globalThis.dispatchEvent(event);
}