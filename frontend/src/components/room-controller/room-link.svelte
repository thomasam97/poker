<script lang="ts">
    export let link = ""

    let showTooltip = false
    function onRoomLinkClick(event: MouseEvent){
        event.preventDefault()
        navigator.clipboard.writeText(link);
        showTooltip = true
        setTimeout(() => { 
            showTooltip = false
        },2_000)
    }

</script>

<a 
    class="tooltip"
    class:show={showTooltip}
    href={link} 
    on:click={onRoomLinkClick}
    data-tooltip="Link has been copied"
>
    Copy Room Link
</a>

<style>
    a {
        display:      grid;
        place-content: center;
    }

    .tooltip {
        position: relative;
        /* border-bottom: 1px dotted white; */
    }

/* Tooltip box */
.tooltip:before {
    content:          attr(data-tooltip); 
    position:         absolute;
    width:            100px;
    background-color: rgb(109, 109, 109);
    color:            var(--color-red);
    text-align:        center;
    padding:           10px;
    line-height:       1.2;
    border-radius:     6px;
    z-index:           1;
    opacity:           0;
    transition:        opacity .6s;
    top:               40px;
    left:              50%;
    margin-left:       -60px;
    font-size:         0.75em;
}

/* Tooltip arrow */
.tooltip:after {
    content:      "";
    position:     absolute;
    top:          30px;
    left:         50%;
    margin-left:  -5px;
    border-width: 5px;
    border-style: solid;
    opacity:      0;
    transition:   opacity .6s;
    border-color: transparent transparent rgb(109, 109, 109) transparent;
    /* border-color: #062B45 transparent transparent transparent; */
}

.tooltip::before,
.tooltip::after {
    pointer-events: none;
}

.tooltip.show::before, 
.tooltip.show::after {
    opacity: 1;
    pointer-events: unset;
}
</style>