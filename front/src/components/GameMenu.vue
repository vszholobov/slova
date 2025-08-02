<template>
  <div class="mt-6 p-4 bg-gray-100 rounded-lg">
    <div v-if="isWaitingForSecondPlayer" class="text-center">
      <p class="text-lg text-gray-600">Ожидание второго игрока...</p>
    </div>
    
    <div v-else-if="isGameActive" class="text-center">
      <SendGuessButton />
    </div>
    
    <div v-else-if="isGameFinished" class="text-center space-y-4">
      <div class="mb-4">
        <p class="text-lg font-semibold text-gray-800">
          {{ gameState.status === 'WIN' ? 'Победа!' : 'В следующий раз повезёт' }}
        </p>
        <p v-if="lastWord" class="text-gray-600">
          Последнее слово: <span class="font-semibold">{{ lastWord }}</span>
        </p>
      </div>
      
      <StartNewGameButton />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useGameStore } from '@/store/game'
import SendGuessButton from './SendGuessButton.vue'
import StartNewGameButton from './StartNewGameButton.vue'

const gameStore = useGameStore()

const gameState = computed(() => gameStore.gameState!)
const isWaitingForSecondPlayer = computed(() => gameStore.isWaitingForSecondPlayer)
const isGameActive = computed(() => gameStore.isGameActive)
const isGameFinished = computed(() => gameStore.isGameFinished)

const lastWord = computed(() => {
  const { guesses } = gameState.value
  const currentPlayer = gameStore.currentPlayer
  
  if (currentPlayer === '1') {
    return guesses.player1[guesses.player1.length - 1]
  } else {
    return guesses.player2[guesses.player2.length - 1]
  }
})
</script>