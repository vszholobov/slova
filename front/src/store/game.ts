import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { GameState } from '@/types'

export const useGameStore = defineStore('game', () => {
  const gameState = ref<GameState | null>(null)
  const isConnected = ref(false)
  const sessionId = ref<string | null>(null)

  const currentPlayer = computed(() => {
    if (!gameState.value) return null
    return gameState.value.you
  })

  const isCurrentPlayer = computed(() => (playerNum: string) => {
    return currentPlayer.value === playerNum
  })

  const isGameActive = computed(() => {
    return gameState.value?.status === 'ACTIVE'
  })

  const isWaitingForSecondPlayer = computed(() => {
    return gameState.value?.status === 'NEW'
  })

  const isGameFinished = computed(() => {
    return gameState.value?.status === 'WIN' || gameState.value?.status === 'SKIP'
  })

  const hasVote = computed(() => {
    return gameState.value?.vote !== null
  })

  const setGameState = (state: GameState) => {
    gameState.value = state
  }

  const setConnectionStatus = (connected: boolean) => {
    isConnected.value = connected
  }

  const setSessionId = (id: string) => {
    sessionId.value = id
  }

  const reset = () => {
    gameState.value = null
    isConnected.value = false
    sessionId.value = null
  }

  return {
    gameState,
    isConnected,
    sessionId,
    currentPlayer,
    isCurrentPlayer,
    isGameActive,
    isWaitingForSecondPlayer,
    isGameFinished,
    hasVote,
    setGameState,
    setConnectionStatus,
    setSessionId,
    reset
  }
}) 