<script context="module" lang="ts">
import Room from "./room/[roomid]/[playername]/room.svelte";


    // export const ssr = false;
	export const prerender = false;
	// export async function load({ page, fetch, session, stuff }) {
	export async function load({ url, params, props, fetch }) {
        console.debug('[DEBUG] ', {params: params.roomid} )
        // return {
        //     status: 200,
        // }
        const roomID = params.roomid
        if(!roomID){
            return {
                status: 302,
                redirect: `/room/`
            };
        }

        return {
            status: 302,
            redirect: `/room/${params.roomid}`
        };

    }
</script>

<h1>/</h1>