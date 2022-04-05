/// <reference types="cypress" />

import { API } from "api"

describe('Feature: Playthrough', () => {
    const api = new API("ws://localhost:7788/api")
    const playerName = "John"
    const roomName = "not-important."+Math.random().toString(36);
    const playerCardPlayerNameSelector = makeTestId("player-card__player-name")
    const roomNameSelector = makeTestId("room-name")
    const playerCardBackSelector = makeTestId("player-card__back")

    beforeEach(() => {
        cy.visit('http://localhost:60370')
        const adminName = "admin"
        api.register(roomName, adminName, () => {})
    })

    context("Given a user joins as a player", () => {
        describe("When (s)he joins a room", () => { 
            beforeEach(() => { 
                cy.get(roomNameSelector)
                    .clear()
                    .type(roomName)

                const playerNameSelector = makeTestId("player-name")
                cy.get(playerNameSelector).type(playerName)
                
                const buttonJoinSelector = makeTestId('button-join')
                cy.get(buttonJoinSelector).click()
            })

            it("Then (s)he will see her/his name above a card", () => { 
                const playerCardPlayerNameSelector = makeTestId("player-card__player-name")
                cy.get(playerCardPlayerNameSelector)
                    .should('contain', playerName)
            })

            describe("When (s)he chooses a card and the room is revealed", () => { 
                const adminsChosenCard = "3"
                beforeEach(() => { 
                    cy.get(playerCardPlayerNameSelector)
                        .should('contain', playerName)
                        .then(() => api.startRoom() )
                        .then(() => api.choose(adminsChosenCard) )

                    const cardSelector = makeTestId("card-3")
                    cy.get(cardSelector)
                        .click()
                        .then(() => api.revealRoom())    
            
                })

                it("Then (s)he sees her/his chosen card", () => { 
                    cy.get(playerCardBackSelector).should('be.visible')
                    cy.get(playerCardPlayerNameSelector).should('contain', playerName)
                })
                
                it("And (s)he sees the chosen cards of the others", () => { 
                    cy.get(playerCardBackSelector)
                        .should('be.visible')
                        .should('have.length',2)
                    cy.get(playerCardPlayerNameSelector)
                        .should('have.length', 2)
                })
            })

        }) 
    }) 
}) 
       

function makeTestId(testId){
    return `[test-id="${testId}"]`
}