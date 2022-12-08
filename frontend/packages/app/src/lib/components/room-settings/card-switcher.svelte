<script lang="ts">
    import {createEventDispatcher} from "svelte"
    import { Cards } from "$lib/components/cards"
    import type { Card } from "$lib/components/cards";
    import {CardBacks} from "$lib/card-backs"

    const dispatchSwitchCardBack = createEventDispatcher()
    function switchCardBack(card: Card){
        dispatchSwitchCardBack('switchcardback', card)
    }
    
    export let open = false
    let cards: Card[] = [
        { label:"", value: CardBacks["80s"],            bgImagePath: CardBacks["80s"]            },
        { label:"", value: CardBacks.ai,                bgImagePath: CardBacks.ai                },
        { label:"", value: CardBacks.matrix,            bgImagePath: CardBacks.matrix            },
        { label:"", value: CardBacks["monogram-gray"],  bgImagePath: CardBacks["monogram-gray"]  },
        { label:"", value: CardBacks["monogram-white"], bgImagePath: CardBacks["monogram-white"] },
        { label:"", value: CardBacks.tron,              bgImagePath: CardBacks.tron              },
    ]

</script>

<dialog open={open}>
    <div class="container">
        <div class="modal">
            <Cards 
                disableHoverAnimation={true}
                cards={cards} 
                on:choose={(event) => switchCardBack(event.detail)} 
            />
        </div>
    </div>

</dialog>
  

<style>
    dialog{
        width:      100%;
        height:     50vh;
        overflow:   hidden;
        border:     none;
        
        background: none;
        z-index:    1000;
        top:        0;
        left:       0;
        margin:     0;
        padding:    0;
    }

    .container {
        width:       100%;
        height:      100%;
        display:     grid;
        place-items: center;
    }

    .modal {
        background-color: #00000024;
        border-radius:    15px;
        padding:          5rem;
        box-shadow:       0px 0px 15px #00000024;
        backdrop-filter:  blur(50px);
        display:          grid;
        place-items:      center;
        border:           1px #00000017 solid;
    }

</style>