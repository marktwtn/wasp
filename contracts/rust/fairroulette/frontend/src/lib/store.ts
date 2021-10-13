import { derived, get, Readable, Writable, writable } from 'svelte/store';
import { BettingStep, calculateRoundLengthLeft } from './../lib/app';
import type { IRound } from './models/IRound';
import type { Buffer, IKeyPair } from './wasp_client';
import { Base58 } from './wasp_client/crypto/base58';

const RESET_ROUND: IRound = {
    active: false,
    logs: [],
    players: [],
    betSelection: undefined,
    betAmount: undefined,
    betPlaced: false,
    winningNumber: undefined,
    startedAt: undefined,
    number: undefined,
}

export const seed: Writable<Buffer> = writable()
export const seedString: Readable<string> = derived(seed, $seed => Base58.encode($seed))
export const keyPair: Writable<IKeyPair> = writable()
export const address: Writable<string> = writable()
export const addressIndex: Writable<number> = writable(0)
export const balance: Writable<bigint> = writable(0n)

export const round: Writable<IRound> = writable(RESET_ROUND)

export const timestamp: Writable<number> = writable()
export const timeToFinished: Readable<number> = derived(timestamp, $timestamp => $timestamp ? calculateRoundLengthLeft($timestamp) : 0)

export const placingBet: Writable<boolean> = writable(false)
export const showBettingSystem: Writable<boolean> = writable(false)
export const bettingStep: Writable<BettingStep> = writable(1)

export const showWinningNumber: Writable<boolean> = writable(false)

export const firstTimeRequestingFunds: Writable<boolean> = writable(true)
export const requestingFunds: Writable<boolean> = writable(false)

export const isAWinnerPlayer: Writable<boolean> = writable(false)

export const addressesHistory: Writable<string[]> = writable([])

export function resetRound(): void {
    round.set({ ...RESET_ROUND, winningNumber: get(round)?.winningNumber, players: [], logs: get(round)?.logs })
}

export function showWinnerAnimation(): void {
    isAWinnerPlayer.set(true)
    setTimeout(() => { isAWinnerPlayer.set(false) }, 20000)
}

export function resetBettingSystem(): void {
    showBettingSystem.set(false)
    bettingStep.set(BettingStep.NumberChoice)
    round.update(($round) => ({ ...$round, betSelection: undefined, betAmount: undefined }));
}