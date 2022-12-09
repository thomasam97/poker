const KEY_CARD_BACK = "cardBack"

export class Stash {

    /**
     * Stores the given card back in local storage
     * @param cardBack 
     */
    public storeCardBack(cardBack: string) {
        window.localStorage.setItem(KEY_CARD_BACK, cardBack)
    }

    /**
     * Retrieves the stored card back from local storage
     * @returns 
     */
    public CardBack(): string {
        const storedCardBack = window.localStorage.getItem(KEY_CARD_BACK)
        if (!storedCardBack){
            return ""
        }

        return storedCardBack
    }
}