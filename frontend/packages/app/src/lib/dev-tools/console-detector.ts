
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
	
	window.addEventListener('resize', handleResize);
	hasEventListener = true
	handleResize()
}


let wasDevToolsOpen = false
function handleResize(){
	const widthThreshold = globalThis.outerWidth - globalThis.innerWidth > threshold;
	const heightThreshold = globalThis.outerHeight - globalThis.innerHeight > threshold;
	// const orientation = widthThreshold ? 'vertical' : 'horizontal';
	const isDevToolsOpen = widthThreshold || heightThreshold

	if( wasDevToolsOpen === isDevToolsOpen){ return }

	emitEvent(isDevToolsOpen)
	wasDevToolsOpen = isDevToolsOpen

}

function emitEvent(isOpen: boolean) {
	const eventKey = "devtoolschange"
	const event = new globalThis.CustomEvent(eventKey, {detail: {isOpen}} )
	globalThis.dispatchEvent(event);
}