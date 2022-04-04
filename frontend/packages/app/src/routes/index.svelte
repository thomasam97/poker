<script context="module" lang="ts">


    // export const ssr = false;
	export const prerender = false;
	// export async function load({ page, fetch, session, stuff }) {
	export async function load({ url, params, props, fetch }) {
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